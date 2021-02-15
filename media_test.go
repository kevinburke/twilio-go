package twilio

import (
	"context"
	"errors"
	"os"
	"strings"
	"testing"
	"time"
)

func TestGetURL(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Parallel()
	t.Skip("broken because of twilio archiving old messages")
	sid := os.Getenv("TWILIO_ACCOUNT_SID")
	c := NewClient(sid, os.Getenv("TWILIO_AUTH_TOKEN"), nil)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	// These are tied to Kevin's account, sorry I don't have a better way to do
	// this.
	u, err := c.Media.GetURL(ctx, "MM89a8c4a6891c53054e9cd604922bfb61", "ME4f366233682e811f63f73220bc07fc34")
	if err != nil {
		t.Fatal(err)
	}
	if u == nil {
		t.Fatal(errors.New("got nil url"))
	}
	str := u.String()
	if !strings.HasPrefix(str, "https://s3-external-1.amazonaws.com/media.twiliocdn.com/"+sid) {
		t.Errorf("wrong url: %s", str)
	}
}

func TestGetImage(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Parallel()
	t.Skip("broken because of twilio archiving old messages")
	sid := os.Getenv("TWILIO_ACCOUNT_SID")
	c := NewClient(sid, os.Getenv("TWILIO_AUTH_TOKEN"), nil)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	// These are tied to Kevin's account, sorry I don't have a better way to do
	// this.
	i, err := c.Media.GetImage(ctx, "MM89a8c4a6891c53054e9cd604922bfb61", "ME4f366233682e811f63f73220bc07fc34")
	if err != nil {
		t.Fatal(err)
	}
	bounds := i.Bounds()
	if bounds.Max.X < 50 || bounds.Max.Y < 50 {
		t.Errorf("Invalid picture bounds: %v", bounds)
	}
}
