package twilio

import (
	"context"
	"testing"
)

func TestGetCallSummary(t *testing.T) {
	t.Parallel()
	client, s := getServer(insightsCallSummaryResponse)
	t.Cleanup(s.Close)

	sid := "CA04917eab5c194f4c86207384933c0c41"
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
