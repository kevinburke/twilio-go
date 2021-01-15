package twilio

import (
	"context"
	"testing"
)

func TestGetSimCertificates(t *testing.T) {
	t.Parallel()
	client, server := getServer(simCertsGetResponse)
	defer server.Close()
	certs, err := client.Wireless.TrustOnBoard.GetSimCertificates(context.Background(), "DEe10f758e920e43318ad80677505fcf90")
	if err != nil {
		t.Fatal(err)
	}

	if len(certs) != 2 {
		t.Fatalf("certs: got length of %d, expected %d", len(certs), 2)
	}
	if certs[0].Sid != "WCf618acb530d15e36a3335d5d27b12ff7" {
		t.Errorf("certs[0]: got sid %q, want %q", certs[0].Sid, "WCf618acb530d15e36a3335d5d27b12ff7")
	}
	if certs[0].SimSid != "DEe10f758e920e43318ad80677505fcf90" {
		t.Errorf("certs[0]: got sim_sid %q, want %q", certs[0].SimSid, "DEe10f758e920e43318ad80677505fcf90")
	}

	if certs[1].Sid != "7d65de8867740536a678e49eafced195" {
		t.Errorf("certs[1]: got sid %q, want %q", certs[1].Sid, "7d65de8867740536a678e49eafced195")
	}
	if certs[1].SimSid != "DEe10f758e920e43318ad80677505fcf90" {
		t.Errorf("certs[1]: got sim_sid %q, want %q", certs[1].SimSid, "DEe10f758e920e43318ad80677505fcf90")
	}
}
