package twilio

import (
	"context"
	"testing"
)

func TestGetCallSummary(t *testing.T) {
	t.Parallel()
	client, s := getServer(insightsCallSummaryResponse)
	defer s.Close()
	sid := "NO00ed1fb4aa449be2434d54ec8e411abc"
	summary, err := client.Insights.VoiceInsights(sid).Summary.Get(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if summary.CallSid != sid {
		t.Errorf("expected Sid to be %s, got %s", sid, summary.CallSid)
	}
	if summary.CallType != "carrier" {
		t.Errorf("expected CallType to be %s, got %s", "carrier", summary.CallType)
	}
}