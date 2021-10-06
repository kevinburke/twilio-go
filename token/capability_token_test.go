package token

import (
	"testing"
	"time"
)

const (
	WORKSPACE_SID = "WS1234567890"
	WORKER_SID    = "WK1234567890"
)

func TestCapabilityJWT(t *testing.T) {
	t.Parallel()

	opts := CapabilityTokenOptions{
		Ttl:                            time.Duration(1000 * time.Second),
		AccountSid:                     ACC_SID,
		ApiKey:                         API_KEY,
		ApiSecret:                      API_SECRET,
		WorkspaceSid:                   WORKSPACE_SID,
		WorkerSid:                      WORKER_SID,
		Version:                        "v1",
		DefaultWorkerCapabilities:      true,
		DefaultEventBridgeCapabilities: true,
	}

	opts.Policies = append(opts.Policies, BuildWorkspacePolicy(opts, "", nil))
	opts.Policies = append(opts.Policies, BuildWorkspacePolicy(opts, "", []string{"**"}))
	opts.Policies = append(opts.Policies, BuildWorkspacePolicy(opts, "POST", []string{"Activities"}))
	opts.Policies = append(opts.Policies, BuildWorkspacePolicy(opts, "POST", []string{"Workers", opts.WorkerSid, "Reservations", "**"}))

	capTkn := NewCapabilityToken(opts)
	jwtString, err := capTkn.JWT()
	if err != nil {
		t.Error("Unexpected error when generating the capability token", err)
	}

	t.Log(jwtString)

	if jwtString == "" {
		t.Error("token returned is empty")
	}
}
