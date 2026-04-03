// Package token generates valid tokens for Twilio Client SDKs.

// Create them on your server and pass them to a client to help verify a
// client's identity, and to grant access to features in client API's.
//
// For more information, please see the Twilio docs:
// https://www.twilio.com/docs/api/rest/access-tokens
package token

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"
)

const jwtContentType = "twilio-fpa;v=1"

// AccessToken holds properties that can generate a signature to talk to the
// Messaging/Voice/Video API's.
type AccessToken struct {
	NotBefore  time.Time
	accountSid string
	apiKey     string
	apiSecret  []byte
	identity   string
	ttl        time.Duration
	grants     map[string]any
	mu         sync.Mutex
}

// New generates a new AccessToken with the specified properties. identity is
// a unique ID for a particular user.
//
// To generate an apiKey and apiSecret follow this instructions:
// https://www.twilio.com/docs/api/rest/access-tokens#jwt-format
func New(accountSid, apiKey, apiSecret, identity string, ttl time.Duration) *AccessToken {
	return &AccessToken{
		accountSid: accountSid,
		apiKey:     apiKey,
		apiSecret:  []byte(apiSecret),
		identity:   identity,
		ttl:        ttl,
		grants:     make(map[string]any),
	}
}

func (t *AccessToken) AddGrant(grant Grant) {
	t.mu.Lock()
	t.grants["identity"] = t.identity
	t.grants[grant.Key()] = grant.ToPayload()
	t.mu.Unlock()
}

type stdToken struct {
	Audience  string `json:"aud,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	ID        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`
}

type token struct {
	Grants map[string]any `json:"grants"`
	*stdToken
}

const header = `{"alg":"HS256","cty":"` + jwtContentType + `","typ":"JWT"}`

var headerb64 []byte

func init() {
	headerb64 = make([]byte, base64.URLEncoding.EncodedLen(len(header)))
	base64.URLEncoding.Encode(headerb64, []byte(header))
	headerb64 = bytes.TrimRight(headerb64, "=")
}

func (t *AccessToken) JWT() (string, error) {
	now := time.Now().UTC()

	stdClaims := &stdToken{
		ID:        fmt.Sprintf("%s-%d", t.apiKey, now.Unix()),
		ExpiresAt: now.Add(t.ttl).Unix(),
		Issuer:    t.apiKey,
		IssuedAt:  now.Unix(),
		Subject:   t.accountSid,
	}
	if !t.NotBefore.IsZero() {
		stdClaims.NotBefore = t.NotBefore.UTC().Unix()
	}

	t.mu.Lock()
	defer t.mu.Unlock()
	claims := token{
		t.grants,
		stdClaims,
	}
	// marshal header
	data, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	datab64 := make([]byte, base64.URLEncoding.EncodedLen(len(data)))
	base64.URLEncoding.Encode(datab64, data)
	datab64 = bytes.TrimRight(datab64, "=")
	parts := append(headerb64, '.')
	parts = append(parts, datab64...)
	hasher := hmac.New(sha256.New, t.apiSecret)
	hasher.Write(parts)

	seg := string(parts) + "." + base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return strings.TrimRight(seg, "="), nil
}
