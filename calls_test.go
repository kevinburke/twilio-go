package twilio

import (
	"context"
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
		sid := "CAa98f7bbc9bc4980a44b128ca4884ca73"
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
	t.Run("GetRange", func(t *testing.T) {
		t.Parallel()
		// call history in this range:
		// 5 CA14b8432 2016-10-27 23:27:08 UTC
		// 4 CAa5eba99 2016-10-27 23:26:38 UTC
		// 3 CA5757109 2016-10-27 23:26:07 UTC (date created is 23:25)
		// 2 CA6d27370 2016-10-27 02:34:21 UTC
		// 1 CA47b862c 2016-10-27 02:34:07 UTC
		//
		// We want to filter for calls 2 and 3. The first page is completely
		// discarded, the second page has 2 results
		data := url.Values{}
		data.Set("PageSize", "2")
		// use fixed zone to avoid daylight savings issues.
		nyc := time.FixedZone("America/New_York", 60*60*-4)
		start := time.Date(2016, 10, 26, 22, 34, 15, 00, nyc)
		end := time.Date(2016, 10, 27, 19, 26, 10, 00, nyc)
		iter := envClient.Calls.GetCallsInRange(start, end, data)
		count := 0
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		for {
			count++
			page, err := iter.Next(ctx)
			if err == NoMoreResults {
				break
			}
			if err != nil {
				t.Fatal(err)
			}
			if len(page.Calls) != 2 {
				t.Fatalf("expected 2 calls in first result set, got %d", len(page.Calls))
			}
			// the 19:25 call on 10-27
			if page.Calls[0].Sid != "CA5757109d6dcbc4bebf6847f5dd45191e" {
				t.Errorf("wrong sid for first call: want %q got %q", "CA575710", page.Calls[0].Sid[:8])
			}
			if page.Calls[1].Sid != "CA6d27370cbbfb605521fe8800bb73f2d2" {
				t.Errorf("wrong sid: got %q", page.Calls[1].Sid)
			}
		}
		if count != 2 {
			t.Errorf("wrong count, expected exactly 2 calls to Next(), got %d", count)
		}
	})
}
