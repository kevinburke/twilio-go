package twilio

import (
	"context"
	"testing"
)

func TestGetVideoRecording(t *testing.T) {
	t.Parallel()
	client, server := getServer(videoRecordingResponse)
	defer server.Close()
	recording, err := client.Video.VideoRecordings.Get(context.Background(), "RT63868a235fc1cf3987e6a2b67346273f")
	if err != nil {
		t.Fatal(err)
	}
	if recording.Sid != "RT63868a235fc1cf3987e6a2b67346273f" {
		t.Errorf("recording: got sid %q, want %q", recording.Sid, "RT63868a235fc1cf3987e6a2b67346273f")
	}
	if recording.Status != StatusProcessing {
		t.Errorf("recording: got status %q, want %q", recording.Status, StatusProcessing)
	}
	for key, value := range recording.GroupingSids {
		if key != "room_sid" {
			t.Errorf("recording.GroupungSids: got key %q, want %q", key, "room_sid")
		}
		if value != "RM58f1e8f2b1c6b88ca90a012a4be0c279" {
			t.Errorf("recording.GroupungSids: got value %q, want %q", value, "RM58f1e8f2b1c6b88ca90a012a4be0c279")
		}
	}
}
