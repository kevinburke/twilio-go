package twilio

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/hex"
	"math/big"
	"testing"
	"time"
)

func testCert(t *testing.T, cn string) *x509.Certificate {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	tmpl := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: cn},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(time.Hour),
	}
	der, err := x509.CreateCertificate(rand.Reader, tmpl, tmpl, &key.PublicKey, key)
	if err != nil {
		t.Fatal(err)
	}
	c, err := x509.ParseCertificate(der)
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func TestVerifyConnectionDistrust(t *testing.T) {
	leaf := testCert(t, "leaf")
	bad := testCert(t, "bad-intermediate")
	good := testCert(t, "good-intermediate")
	h := sha256.Sum256(bad.RawSubjectPublicKeyInfo)
	key := hex.EncodeToString(h[:])
	distrustedSPKIHashes[key] = struct{}{}
	defer delete(distrustedSPKIHashes, key)

	tests := []struct {
		name    string
		chains  [][]*x509.Certificate
		wantErr bool
	}{
		{"no verified chains", nil, false},
		{"clean chain", [][]*x509.Certificate{{leaf, good}}, false},
		{"distrusted chain", [][]*x509.Certificate{{leaf, bad}}, true},
		{"one clean alternative", [][]*x509.Certificate{{leaf, bad}, {leaf, good}}, false},
	}
	for _, tt := range tests {
		err := verifyConnection(tls.ConnectionState{VerifiedChains: tt.chains})
		if (err != nil) != tt.wantErr {
			t.Errorf("%s: got err %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestDefaultTransportHasVerifyConnection(t *testing.T) {
	tr := newDefaultTransport()
	if tr == nil || tr.RoundTripper == nil {
		t.Fatal("nil transport")
	}
}
