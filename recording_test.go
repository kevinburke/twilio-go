package twilio

import (
	"context"
	"path/filepath"
	"testing"
	"time"
)

var recordingInstance = []byte(`{
	"account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"api_version": "2010-04-01",
	"call_sid": "CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"conference_sid": "CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"channels": 1,
	"date_created": "Fri, 14 Oct 2016 21:56:34 +0000",
	"date_updated": "Fri, 14 Oct 2016 21:56:38 +0000",
	"start_time": "Fri, 14 Oct 2016 21:56:34 +0000",
	"price": "0.04",
	"price_unit": "USD",
	"duration": "4",
	"sid": "REd04242a0544234abba080942e0535505",
	"source": "StartConferenceRecordingAPI",
	"status": "completed",
	"error_code": null,
	"uri": "/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
	"subresource_uris": {
	  "add_on_results": "/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/AddOnResults.json",
	  "transcriptions": "/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Transcriptions.json"
	},
	"encryption_details": {
	  "encryption_public_key_sid": "CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	  "encryption_cek": "OV4h6zrsxMIW7h0Zfqwfn6TI2GCNl54KALlg8wn8YB8KYZhXt6HlgvBWAmQTlfYVeLWydMiCewY0YkDDT1xmNe5huEo9vjuKBS5OmYK4CZkSx1NVv3XOGrZHpd2Pl/5WJHVhUK//AUO87uh5qnUP2E0KoLh1nyCLeGcEkXU0RfpPn/6nxjof/n6m6OzZOyeIRK4Oed5+rEtjqFDfqT0EVKjs6JAxv+f0DCc1xYRHl2yV8bahUPVKs+bHYdy4PVszFKa76M/Uae4jFA9Lv233JqWcxj+K2UoghuGhAFbV/JQIIswY2CBYI8JlVSifSqNEl9vvsTJ8bkVMm3MKbG2P7Q==",
	  "encryption_iv": "8I2hhNIYNTrwxfHk"
	}
  }`)

func TestGetRecording(t *testing.T) {
	t.Parallel()
	client, s := getServer(recordingInstance)
	defer s.Close()

	sid := "REd04242a0544234abba080942e0535505"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var fullPath string

	fn := client.FullPath
	client.FullPath = func(pathPart string) string {
		fullPath = fn(pathPart)
		return fullPath
	}

	recording, err := client.Recordings.Get(ctx, sid)
	if err != nil {
		t.Fatal(err)
	}
	if fullPath != "/2010-04-01/Accounts/AC123/Recordings/REd04242a0544234abba080942e0535505.json" {
		t.Errorf("expected fullPath to equal %s, got %s", "/2010-04-01/Accounts/AC123/Recordings/REd04242a0544234abba080942e0535505.json", fullPath)
	}
	if recording.Sid != sid {
		t.Errorf("expected Sid to equal %s, got %s", sid, recording.Sid)
	}
}

func TestRecordingPageIter(t *testing.T) {
	t.Parallel()
	client, s := getServer(recordingInstance)
	defer s.Close()

	client.Recordings.SubType = filepath.Join(conferencePathPart, "CF1234")
	it := client.Recordings.GetPageIterator(nil)
	if it.p.pathPart != "Conferences/CF1234/Recordings" {
		t.Errorf("expected pathPart to equal Conferences/CF1234/Recordings, got %s", it.p.pathPart)
	}

	client.Recordings.SubType = filepath.Join(callsPathPart, "SID1234")
	it = client.Recordings.GetPageIterator(nil)
	if it.p.pathPart != "Calls/SID1234/Recordings" {
		t.Errorf("expected pathPart to equal Calls/SID1234/Recordings, got %s", it.p.pathPart)
	}
}
