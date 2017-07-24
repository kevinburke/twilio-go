package twilioclient

import (
	"net/url"
	"strings"
	"sync"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Capability struct {
	accountSid   string
	authToken    []byte
	capabilities []string

	incomingClientName       string
	shouldBuildIncomingScope bool

	shouldBuildOutgoingScope bool
	outgoingParams           map[string]string
	appSid                   string
	mu                       sync.Mutex
}

func NewCapability(sid, token string) *Capability {
	return &Capability{
		accountSid: sid,
		authToken:  []byte(token),
	}
}

// Registers this client to accept incoming calls by the given `clientName`.
// If your app TwiML <Dial>s `clientName`, this client will receive the call.
func (c *Capability) AllowClientIncoming(clientName string) {
	c.shouldBuildIncomingScope = true
	c.incomingClientName = clientName
}

// Allows this client to call your application with id `appSid` (See https://www.twilio.com/user/account/apps).
// When the call connects, Twilio will call your voiceUrl REST endpoint.
// The `appParams` argument will get passed through to your voiceUrl REST endpoint as GET or POST parameters.
func (c *Capability) AllowClientOutgoing(appSid string, appParams map[string]string) {
	c.shouldBuildOutgoingScope = true
	c.appSid = appSid
	c.outgoingParams = appParams
}

func (c *Capability) AllowEventStream(filters map[string]string) {
	params := map[string]string{
		"path": "/2010-04-01/Events",
	}
	if len(filters) > 0 {
		params["params"] = url.QueryEscape(generateParamString(filters))
	}
	c.addCapability("stream", "subscribe", params)
}

type customClaim struct {
	*jwt.StandardClaims
	Scope string
}

// Generate the twilio capability token. Deliver this token to your
// JS/iOS/Android Twilio client.
func (c *Capability) GenerateToken(ttl time.Duration) (string, error) {
	c.doBuildIncomingScope()
	c.doBuildOutgoingScope()
	now := time.Now().UTC()
	cc := &customClaim{
		Scope: strings.Join(c.capabilities, " "),
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: now.Add(ttl).Unix(),
			Issuer:    c.accountSid,
			IssuedAt:  now.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cc)
	return token.SignedString([]byte(c.authToken))
}

func (c *Capability) doBuildOutgoingScope() {
	if c.shouldBuildOutgoingScope {
		values := map[string]string{}
		values["appSid"] = c.appSid
		if c.incomingClientName != "" {
			values["clientName"] = c.incomingClientName
		}

		if c.outgoingParams != nil {
			values["appParams"] = generateParamString(c.outgoingParams)
		}

		c.addCapability("client", "outgoing", values)
	}
}

func (c *Capability) doBuildIncomingScope() {
	if c.shouldBuildIncomingScope {
		values := map[string]string{}
		values["clientName"] = c.incomingClientName
		c.addCapability("client", "incoming", values)
	}
}

func (c *Capability) addCapability(service, privilege string, params map[string]string) {
	c.capabilities = append(c.capabilities, scopeUriFor(service, privilege, params))
}

func scopeUriFor(service, privilege string, params map[string]string) string {
	scopeUri := "scope:" + service + ":" + privilege
	if len(params) > 0 {
		scopeUri += "?" + generateParamString(params)
	}
	return scopeUri
}

func generateParamString(params map[string]string) string {
	values := make(url.Values)
	for key, val := range params {
		values.Add(key, val)
	}
	return values.Encode()
}
