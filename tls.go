package twilio

import (
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"errors"
	"net/http"
	"os"

	"github.com/kevinburke/rest/v2/restclient"
)

// distrustedSPKIHashes contains hex-encoded SHA-256 hashes of the
// SubjectPublicKeyInfo of intermediate certificates that we don't consider
// trustworthy. This is the same list Go distrusts as of Go 1.27.3 (see
// CVE-2026-111943 and golang/go#82345); we apply it here so users on older Go
// releases are protected too.
var distrustedSPKIHashes = map[string]struct{}{
	"ebed6475c2c2d80fcb199f3fe8dd6e6b3ce60b77eb15d4e42c3dec0c6f5b5105": {},
	"49dd353dc2b919a3770645117d0e60c16115e7d1bc259eefe734b00ea434fe5f": {},
	"e0254c7fcdc265e2f0f8a428dabb75ab4ae039b4872d62f5ee57f0877d5a8a41": {},
	"5c0660a2b164910d8ff589c18e14552e9e0ec237dcefe2cb736d41a4811f81a0": {},
	"43bffa19d6ef651451f6841e019931e96f6f2eaeff28ab4b6f1ce4576de2717b": {},
	"63ecefee431d49523bc1136c4e72d0206a477086caf606cd5cf025686c109cf8": {},
	"ec4ffa1ad0ac88081314b7aaffef6a91f97cd57291bb5ba686c5415f5eb4f6b6": {},
	"92c063d67e8f14db71d288e11266ffdaf69cbf115f7bab14f3bc2f68d25e0e30": {},
}

var errDistrustedIntermediate = errors.New("twilio: certificate chain contains a distrusted intermediate certificate")

func containsDistrustedSPKIHash(chain []*x509.Certificate) bool {
	for _, c := range chain {
		h := sha256.Sum256(c.RawSubjectPublicKeyInfo)
		if _, ok := distrustedSPKIHashes[hex.EncodeToString(h[:])]; ok {
			return true
		}
	}
	return false
}

// verifyConnection rejects a connection if every verified chain contains a
// distrusted certificate. Like Go's own filtering, a connection is allowed if
// at least one verified chain is clean.
func verifyConnection(cs tls.ConnectionState) error {
	if len(cs.VerifiedChains) == 0 {
		// Verification was skipped by the caller (InsecureSkipVerify).
		return nil
	}
	for _, chain := range cs.VerifiedChains {
		if !containsDistrustedSPKIHash(chain) {
			return nil
		}
	}
	return errDistrustedIntermediate
}

// newDefaultTransport returns a copy of http.DefaultTransport that refuses
// chains containing distrusted intermediates, wrapped so that
// DEBUG_HTTP_TRAFFIC=true continues to work.
func newDefaultTransport() *restclient.Transport {
	var base *http.Transport
	if dt, ok := http.DefaultTransport.(*http.Transport); ok {
		base = dt.Clone()
	} else {
		base = &http.Transport{Proxy: http.ProxyFromEnvironment, ForceAttemptHTTP2: true}
	}
	if base.TLSClientConfig == nil {
		base.TLSClientConfig = &tls.Config{}
	}
	prev := base.TLSClientConfig.VerifyConnection
	base.TLSClientConfig.VerifyConnection = func(cs tls.ConnectionState) error {
		if err := verifyConnection(cs); err != nil {
			return err
		}
		if prev != nil {
			return prev(cs)
		}
		return nil
	}
	return &restclient.Transport{
		RoundTripper: base,
		Debug:        os.Getenv("DEBUG_HTTP_TRAFFIC") == "true",
		Output:       os.Stderr,
	}
}
