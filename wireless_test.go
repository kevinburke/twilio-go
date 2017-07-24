package twilio

import (
	"context"
	"testing"
)

func TestGetSim(t *testing.T) {
	t.Parallel()
	client, server := getServer(simGetResponse)
	defer server.Close()
	sim, err := client.Wireless.Sims.Get(context.Background(), "DEe10f758e920e43318ad80677505fcf90")
	if err != nil {
		t.Fatal(err)
	}
	if sim.Sid != "DEe10f758e920e43318ad80677505fcf90" {
		t.Errorf("sim: got sid %q, want %q", sim.Sid, "DEe10f758e920e43318ad80677505fcf90")
	}
	if sim.Status != StatusActive {
		t.Errorf("sim: got status %q, want %q", sim.Status, StatusActive)
	}
}
