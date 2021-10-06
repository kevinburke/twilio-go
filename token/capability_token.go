package token

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"strings"
	"sync"
	"time"
)

type CapabilityToken struct {
	accountSid   string
	apiKey       string
	apiSecret    []byte
	workspaceSid string
	workerSid    string
	version      string
	ttl          time.Duration
	policies     []map[string]interface{}
	mu           sync.Mutex
}

type cToken struct {
	*stdToken
	Policies     []map[string]interface{} `json:"policies"`
	Version      string                   `json:"version"`
	FriendlyName string                   `json:"friendly_name"`
	AccountSid   string                   `json:"account_sid"`
	WorkspaceSid string                   `json:"workspace_sid"`
	Channel      string                   `json:"channel"`
	WorkerSid    string                   `json:"worker_sid"`
}

type Policy struct {
	URL    string
	Method string
}

func (p Policy) ToMap() map[string]interface{} {
	var mapPolicy = make(map[string]interface{})

	mapPolicy["url"] = p.URL
	mapPolicy["method"] = p.Method
	mapPolicy["allow"] = true

	return mapPolicy
}

type CapabilityTokenOptions struct {
	Ttl                            time.Duration
	AccountSid                     string
	ApiKey                         string
	ApiSecret                      string
	Policies                       []Policy
	WorkspaceSid                   string
	WorkerSid                      string
	Version                        string
	DefaultWorkerCapabilities      bool
	DefaultEventBridgeCapabilities bool
}

const EVENT_URL_BASE = "https://event-bridge.twilio.com/v1/wschannels"
const TASKROUTER_BASE_URL = "https://taskrouter.twilio.com"

const WORKSPACES = "Workspaces"
const WORKERS = "Workers"
const ACTIVITIES = "Activities"
const RESERVATIONS = "Reservations"
const TASKS = "Tasks"

func DefaultWorkerPolicies(workspaceSid, workerSid, version string) []Policy {
	return []Policy{
		// Activities
		{
			URL:    strings.Join([]string{TASKROUTER_BASE_URL, version, workspaceSid, ACTIVITIES}, "/"),
			Method: "GET",
		},
		// Tasks
		{
			URL:    strings.Join([]string{TASKROUTER_BASE_URL, version, WORKSPACES, workspaceSid, TASKS, "**"}, "/"),
			Method: "GET",
		},
		// Reservations
		{
			URL:    strings.Join([]string{TASKROUTER_BASE_URL, version, WORKSPACES, workspaceSid, WORKERS, workerSid, RESERVATIONS, "**"}, "/"),
			Method: "GET",
		},
		// Worker Fetch
		{
			URL:    strings.Join([]string{TASKROUTER_BASE_URL, version, WORKSPACES, workspaceSid, WORKERS, workerSid}, "/"),
			Method: "GET",
		},
	}
}

func DefaultEventBridgePolicies(accountSid, channelId string) []Policy {
	url := strings.Join([]string{EVENT_URL_BASE, accountSid, channelId}, "/")
	return []Policy{
		{
			URL:    url,
			Method: "GET",
		},
		{
			URL:    url,
			Method: "POST",
		},
	}
}

func BuildWorkspacePolicy(opts CapabilityTokenOptions, method string, resources []string) Policy {
	if method == "" {
		method = "GET"
	}

	fullResourceArray := []string{TASKROUTER_BASE_URL, opts.Version, WORKSPACES, opts.WorkspaceSid}
	fullResourceArray = append(fullResourceArray, resources...)

	url := strings.Join(fullResourceArray, "/")

	return Policy{
		URL:    url,
		Method: method,
	}
}

// NewCapabilityToken generates a new CapabilityToken with the provided CapabilityTokenOptions.
//
// To see more visit: https://www.twilio.com/docs/taskrouter/js-sdk/workspace/constructing-jwts
func NewCapabilityToken(opts CapabilityTokenOptions) *CapabilityToken {
	policies := opts.Policies

	if opts.DefaultWorkerCapabilities {
		policies = append(policies, DefaultWorkerPolicies(opts.WorkspaceSid, opts.WorkerSid, opts.Version)...)
	}

	if opts.DefaultEventBridgeCapabilities {
		policies = append(policies, DefaultEventBridgePolicies(opts.AccountSid, opts.WorkerSid)...)
	}

	var policyMaps = make([]map[string]interface{}, 0)

	for _, policy := range policies {
		if policy.ToMap() != nil {
			policyMaps = append(policyMaps, policy.ToMap())
		}
	}

	return &CapabilityToken{
		accountSid:   opts.AccountSid,
		apiKey:       opts.ApiKey,
		apiSecret:    []byte(opts.ApiSecret),
		policies:     policyMaps,
		workspaceSid: opts.WorkspaceSid,
		workerSid:    opts.WorkerSid,
		ttl:          opts.Ttl,
		version:      opts.Version,
	}
}

func (t *CapabilityToken) JWT() (string, error) {
	now := time.Now().UTC()

	stdClaims := &stdToken{
		ExpiresAt: now.Add(t.ttl).Unix(),
		Issuer:    t.accountSid,
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	claims := cToken{
		stdClaims,
		t.policies,
		t.version,
		t.workerSid,
		t.accountSid,
		t.workspaceSid,
		t.workspaceSid,
		t.workerSid,
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
