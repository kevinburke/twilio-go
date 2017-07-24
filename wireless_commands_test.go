package twilio

import (
	"context"
	"net/url"
	"testing"
)

func TestCreateCommand(t *testing.T) {
	t.Parallel()
	client, server := getServer(cmdCreateResponse)
	defer server.Close()
	cmd, err := client.Wireless.Commands.Create(context.Background(), url.Values{
		"Sim":     []string{"DEe10f758e920e43318ad80677505fcf90"},
		"Command": []string{"twilio-go testing!"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if cmd.Direction != Direction("to_sim") {
		t.Errorf("bad direction: %q", cmd.Direction)
	}
	if cmd.Command != "twilio-go testing!" {
		t.Errorf("bad command text: %q", cmd.Command)
	}
	if cmd.SimSid != "DEe10f758e920e43318ad80677505fcf90" {
		t.Errorf("bad sim sid: %q", cmd.SimSid)
	}
}
