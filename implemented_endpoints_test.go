package twilio

// Hermetic path-construction tests for the call, conference, and message
// sub-resource endpoints. The OpenAPI tests validate the response structs; these
// verify each method targets the correct URL and HTTP method. lastPath is
// defined in sip_test.go.

import (
	"context"
	"net/url"
	"testing"
)

// lastMethod returns the HTTP method of the most recent request to the server.
func lastMethod(s *Server) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.URLs) == 0 {
		return ""
	}
	return s.Methods[len(s.Methods)-1]
}

func TestImplementedEndpointPaths(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	empty := url.Values{}

	cases := []struct {
		name       string
		call       func(c *Client) error
		wantMethod string
		wantPath   string
	}{
		{
			"MessageUpdate",
			func(c *Client) error { _, err := c.Messages.Update(ctx, "MM123", empty); return err },
			"POST", "/2010-04-01/Accounts/AC123/Messages/MM123.json",
		},
		{
			"MessageRedact",
			func(c *Client) error { _, err := c.Messages.Redact(ctx, "MM123"); return err },
			"POST", "/2010-04-01/Accounts/AC123/Messages/MM123.json",
		},
		{
			"MessageCreateFeedback",
			func(c *Client) error { _, err := c.Messages.CreateFeedback(ctx, "MM123", "confirmed"); return err },
			"POST", "/2010-04-01/Accounts/AC123/Messages/MM123/Feedback.json",
		},
		{
			"MediaDelete",
			func(c *Client) error { return c.Media.Delete(ctx, "MM123", "ME123") },
			"DELETE", "/2010-04-01/Accounts/AC123/Messages/MM123/Media/ME123.json",
		},
		{
			"CallDelete",
			func(c *Client) error { return c.Calls.Delete(ctx, "CA123") },
			"DELETE", "/2010-04-01/Accounts/AC123/Calls/CA123.json",
		},
		{
			"CallGetCallRecordings",
			func(c *Client) error { _, err := c.Calls.GetCallRecordings(ctx, "CA123", empty); return err },
			"GET", "/2010-04-01/Accounts/AC123/Calls/CA123/Recordings.json",
		},
		{
			"CallCreateRecording",
			func(c *Client) error { _, err := c.Calls.CreateRecording(ctx, "CA123", empty); return err },
			"POST", "/2010-04-01/Accounts/AC123/Calls/CA123/Recordings.json",
		},
		{
			"CallGetRecording",
			func(c *Client) error { _, err := c.Calls.GetRecording(ctx, "CA123", "RE123"); return err },
			"GET", "/2010-04-01/Accounts/AC123/Calls/CA123/Recordings/RE123.json",
		},
		{
			"CallUpdateRecording",
			func(c *Client) error { _, err := c.Calls.UpdateRecording(ctx, "CA123", "RE123", empty); return err },
			"POST", "/2010-04-01/Accounts/AC123/Calls/CA123/Recordings/RE123.json",
		},
		{
			"CallDeleteRecording",
			func(c *Client) error { return c.Calls.DeleteRecording(ctx, "CA123", "RE123") },
			"DELETE", "/2010-04-01/Accounts/AC123/Calls/CA123/Recordings/RE123.json",
		},
		{
			"CallGetEvents",
			func(c *Client) error { _, err := c.Calls.GetEvents(ctx, "CA123", empty); return err },
			"GET", "/2010-04-01/Accounts/AC123/Calls/CA123/Events.json",
		},
		{
			"CallGetNotifications",
			func(c *Client) error { _, err := c.Calls.GetNotifications(ctx, "CA123", empty); return err },
			"GET", "/2010-04-01/Accounts/AC123/Calls/CA123/Notifications.json",
		},
		{
			"CallGetNotification",
			func(c *Client) error { _, err := c.Calls.GetNotification(ctx, "CA123", "NO123"); return err },
			"GET", "/2010-04-01/Accounts/AC123/Calls/CA123/Notifications/NO123.json",
		},
		{
			"CallCreatePayment",
			func(c *Client) error { _, err := c.Calls.CreatePayment(ctx, "CA123", empty); return err },
			"POST", "/2010-04-01/Accounts/AC123/Calls/CA123/Payments.json",
		},
		{
			"CallUpdatePayment",
			func(c *Client) error { _, err := c.Calls.UpdatePayment(ctx, "CA123", "PK123", empty); return err },
			"POST", "/2010-04-01/Accounts/AC123/Calls/CA123/Payments/PK123.json",
		},
		{
			"ConferenceUpdate",
			func(c *Client) error { _, err := c.Conferences.Update(ctx, "CF123", empty); return err },
			"POST", "/2010-04-01/Accounts/AC123/Conferences/CF123.json",
		},
		{
			"ParticipantCreate",
			func(c *Client) error { _, err := c.Conferences.Participants("CF123").Create(ctx, empty); return err },
			"POST", "/2010-04-01/Accounts/AC123/Conferences/CF123/Participants.json",
		},
		{
			"ParticipantGet",
			func(c *Client) error { _, err := c.Conferences.Participants("CF123").Get(ctx, "CA999"); return err },
			"GET", "/2010-04-01/Accounts/AC123/Conferences/CF123/Participants/CA999.json",
		},
		{
			"ParticipantUpdate",
			func(c *Client) error {
				_, err := c.Conferences.Participants("CF123").Update(ctx, "CA999", empty)
				return err
			},
			"POST", "/2010-04-01/Accounts/AC123/Conferences/CF123/Participants/CA999.json",
		},
		{
			"ParticipantDelete",
			func(c *Client) error { return c.Conferences.Participants("CF123").Delete(ctx, "CA999") },
			"DELETE", "/2010-04-01/Accounts/AC123/Conferences/CF123/Participants/CA999.json",
		},
		{
			"ParticipantGetPage",
			func(c *Client) error { _, err := c.Conferences.Participants("CF123").GetPage(ctx, empty); return err },
			"GET", "/2010-04-01/Accounts/AC123/Conferences/CF123/Participants.json",
		},
		{
			"ConferenceGetRecordings",
			func(c *Client) error { _, err := c.Conferences.GetRecordings(ctx, "CF123", empty); return err },
			"GET", "/2010-04-01/Accounts/AC123/Conferences/CF123/Recordings.json",
		},
		{
			"ConferenceGetRecording",
			func(c *Client) error { _, err := c.Conferences.GetRecording(ctx, "CF123", "RE123"); return err },
			"GET", "/2010-04-01/Accounts/AC123/Conferences/CF123/Recordings/RE123.json",
		},
		{
			"ConferenceUpdateRecording",
			func(c *Client) error {
				_, err := c.Conferences.UpdateRecording(ctx, "CF123", "RE123", empty)
				return err
			},
			"POST", "/2010-04-01/Accounts/AC123/Conferences/CF123/Recordings/RE123.json",
		},
		{
			"ConferenceDeleteRecording",
			func(c *Client) error { return c.Conferences.DeleteRecording(ctx, "CF123", "RE123") },
			"DELETE", "/2010-04-01/Accounts/AC123/Conferences/CF123/Recordings/RE123.json",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			client, s := getServer([]byte(`{}`))
			defer s.Close()
			if err := tc.call(client); err != nil {
				t.Fatalf("%s: unexpected error: %v", tc.name, err)
			}
			if got := lastMethod(s); got != tc.wantMethod {
				t.Errorf("%s: method = %s, want %s", tc.name, got, tc.wantMethod)
			}
			if got := lastPath(s); got != tc.wantPath {
				t.Errorf("%s: path = %s, want %s", tc.name, got, tc.wantPath)
			}
		})
	}
}
