package twilio

import (
	"context"
	"net/url"
	"testing"
)

func TestGetCallMetrics(t *testing.T) {
	t.Parallel()
	client, s := getServer(insightsCallMetricsResponse)
	defer s.Close()
	sid := "NO00ed1fb4aa449be2434d54ec8e411abc"
	page, err := client.Insights.VoiceInsights(sid).Metrics.GetPage(context.Background(), url.Values{})
	if err != nil {
		t.Fatal(err)
	}
	if len(page.Metrics) != 8 {
		t.Errorf("expected 8 metrics, got %d", len(page.Metrics))
	}
	if page.Meta.Key != "metrics" {
		t.Errorf("expected Key to be 'metrics', got %s", page.Meta.Key)
	}
	if page.Metrics[0].CallSid != sid {
		t.Errorf("expected CallSid to be %s, got %s", sid, page.Metrics[0].CallSid)
	}
	if page.Metrics[0].Edge != "carrier_edge" {
		t.Errorf("expected Edge to be 'carrier_edge', got %s", page.Metrics[0].Edge)
	}
	if page.Metrics[0].CarrierEdge == nil {
		t.Error("expected Carrier Edge metrics to be available")
	}
}