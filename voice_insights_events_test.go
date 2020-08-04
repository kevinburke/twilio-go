package twilio

import (
	"context"
	"net/url"
	"testing"
)

func TestGetCallEventsPage(t *testing.T) {
	t.Parallel()
	client, s := getServer(insightsCallEventsResponse)
	defer s.Close()
	sid := "NO00ed1fb4aa449be2434d54ec8e411abc"
	page, err := client.Insights.VoiceInsights(sid).Events.GetPage(context.Background(), url.Values{})
	if err != nil {
		t.Fatal(err)
	}
	if len(page.Events) != 1 {
		t.Errorf("expected 1 event, got %d", len(page.Events))
	}
	if page.Meta.Key != "events" {
		t.Errorf("expected Key to be 'events', got %s", page.Meta.Key)
	}
	if page.Events[0].Group != "connection" {
		t.Errorf("expected Group to be connection, got %s", page.Events[0].Group)
	}
	if page.Events[0].Name != "completed" {
		t.Errorf("expected Name to be 'completed', got %s", page.Events[0].Name)
	}
	if page.Events[0].Edge != "client_edge" {
		t.Errorf("expected Edge to be 'client_edge', got %s", page.Events[0].Edge)
	}
}
