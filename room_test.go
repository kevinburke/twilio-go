package twilio

import (
	"context"
	"testing"
)

func TestGetRoom(t *testing.T) {
	t.Parallel()
	client, server := getServer(roomResponse)
	defer server.Close()
	room, err := client.Video.Rooms.Get(context.Background(), "RMca86cf94c7d4f89e0bd45bfa7d9b9e7d")
	if err != nil {
		t.Fatal(err)
	}
	if room.Sid != "RMca86cf94c7d4f89e0bd45bfa7d9b9e7d" {
		t.Errorf("room: got sid %q, want %q", room.Sid, "RMca86cf94c7d4f89e0bd45bfa7d9b9e7d")
	}
	if room.Status != StatusInProgress {
		t.Errorf("room: got status %q, want %q", room.Status, StatusInProgress)
	}
	if room.Type != RoomTypePeerToPeer {
		t.Errorf("room: got type %q, want %q", room.Type, RoomTypePeerToPeer)
	}
}
