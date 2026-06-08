package twilio

import (
	"context"
	"encoding/json"
	"net/url"
	"testing"
	"time"
)

func TestCalls(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	t.Cleanup(func() { cancel() })
	t.Run("Get", func(t *testing.T) {
		t.Parallel()
		// Self-seeding: fetch a recent call and round-trip it, rather than
		// hardcoding a SID that rots or pins the test to specific account data.
		page, err := envClient.Calls.GetPage(ctx, url.Values{"PageSize": []string{"1"}})
		if err != nil {
			t.Fatal(err)
		}
		if len(page.Calls) == 0 {
			t.Skip("no calls in account to fetch")
		}
		sid := page.Calls[0].Sid
		call, err := envClient.Calls.Get(ctx, sid)
		if err != nil {
			t.Fatal(err)
		}
		if call.Sid != sid {
			t.Errorf("expected Sid to equal %s, got %s", sid, call.Sid)
		}
	})
	t.Run("GetRecordings", func(t *testing.T) {
		t.Parallel()
		sid := "CA14365760c10f73392c5440bdfb70c212"
		recordings, err := envClient.Calls.GetRecordings(ctx, sid, nil)
		if err != nil {
			t.Fatal(err)
		}
		if l := len(recordings.Recordings); l != 1 {
			t.Fatalf("expected 1 recording, got %d", l)
		}
		rsid := "REd04242a0544234abba080942e0535505"
		if r := recordings.Recordings[0].Sid; r != rsid {
			t.Errorf("expected recording sid to be %s, got %s", rsid, r)
		}
		if recordings.NextPageURI.Valid {
			t.Errorf("expected next page uri to be invalid, got %v", recordings.NextPageURI)
		}
	})
	t.Run("Post", func(t *testing.T) {
		t.Parallel()
		client, server := getServer(makeCallResponse)
		defer server.Close()
		u, _ := url.Parse("https://kev.inburke.com/zombo/zombocom.mp3")
		call, err := client.Calls.MakeCall("+19253920364", "+19252717005", u)
		if err != nil {
			t.Fatal(err)
		}
		if call.To != "+19252717005" {
			t.Errorf("Wrong To phone number: %s", call.To)
		}
		if call.Status != StatusQueued {
			t.Errorf("Wrong status: %s", call.Status)
		}
	})
	// A former GetRange subtest pinned GetCallsInRange to a fixed 2016 call
	// window in a since-archived account; it was removed as no longer useful.
	// The date-range iteration logic is covered by the hermetic message-range
	// tests. See V3.md if GetCallsInRange warrants a data-independent live test.
}

// TestCallAnsweredByUnmarshal guards against a regression where a non-null
// answered_by string failed to decode the entire Call. The API returns
// answered_by as null or as a string ("human", "machine", and the more
// specific AMD values).
func TestCallAnsweredByUnmarshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		json  string
		valid bool
		want  AnsweredBy
	}{
		{`{"answered_by":"human"}`, true, AnsweredByHuman},
		{`{"answered_by":"machine"}`, true, AnsweredByMachine},
		{`{"answered_by":"machine_end_beep"}`, true, AnsweredBy("machine_end_beep")},
		{`{"answered_by":null}`, false, ""},
		{`{}`, false, ""},
	}
	for _, tc := range cases {
		var call Call
		if err := json.Unmarshal([]byte(tc.json), &call); err != nil {
			t.Errorf("%s: unexpected error: %v", tc.json, err)
			continue
		}
		if call.AnsweredBy.Valid != tc.valid {
			t.Errorf("%s: Valid = %v, want %v", tc.json, call.AnsweredBy.Valid, tc.valid)
		}
		if call.AnsweredBy.AnsweredBy != tc.want {
			t.Errorf("%s: AnsweredBy = %q, want %q", tc.json, call.AnsweredBy.AnsweredBy, tc.want)
		}
	}
}
