package twilio

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/kevinburke/rest/resterror"
)

// invalid status here on purpose to check we use a different one.
var notFoundResp = []byte("{\"code\": 20404, \"message\": \"The requested resource /2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Calls/unknown.json was not found\", \"more_info\": \"https://www.twilio.com/docs/errors/20404\", \"status\": 428}")

var errorServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(404)
	w.Write(notFoundResp)
}))

func Test404Error(t *testing.T) {
	t.Parallel()
	client := NewClient("", "", nil)
	client.Base = errorServer.URL
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()
	sid := "unknown"
	_, err := client.Calls.Get(ctx, sid)
	if err == nil {
		t.Fatal("expected non-nil error, got nil")
	}
	rerr, ok := err.(*resterror.Error)
	if !ok {
		t.Fatalf("expected to convert err %v to resterror.Error, couldn't", err)
	}
	if !strings.Contains(rerr.Title, "The requested resource /2010-04-01") {
		t.Errorf("expected Title to contain 'The requested resource', got %s", rerr.Title)
	}
	if rerr.ID != "20404" {
		t.Errorf("expected ID to be 20404, got %s", rerr.ID)
	}
	if rerr.Type != "https://www.twilio.com/docs/errors/20404" {
		t.Errorf("expected Type to be a Twilio URL, got %s", rerr.Type)
	}
	if rerr.Status != 404 {
		t.Errorf("expected StatusCode to be 404, got %d", rerr.Status)
	}
}

func TestContext(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Parallel()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	_, err := envClient.Calls.Get(ctx, "unknown")
	if err == nil {
		t.Fatal("expected Err to be non-nil, got nil")
	}
	// I wish it had context.DeadlineExceeded, doesn't seem to be the case.
	ok := strings.Contains(err.Error(), "deadline exceeded") || strings.Contains(err.Error(), "canceled")
	if !ok {
		t.Errorf("bad error message: %v", err)
	}
}

func TestCancelStopsRequest(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	sid := "CAa98f7bbc9bc4980a44b128ca4884ca73"
	go func() {
		time.Sleep(30 * time.Millisecond)
		cancel()
	}()
	_, err := envClient.Calls.Get(ctx, sid)
	if err == nil {
		t.Fatal("expected Err to be non-nil, got nil")
	}
	if !strings.Contains(err.Error(), "canceled") {
		t.Errorf("bad error message: %#v", err)
	}
}

func TestOnBehalfOf(t *testing.T) {
	t.Parallel()
	want := "/2010-04-01/Accounts/AC345/Calls/CA123.json"
	called := false
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != want {
			t.Errorf("expected Path to be %s, got %s", want, r.URL.Path)
		}
		called = true
		w.WriteHeader(200)
		w.Write([]byte("{}"))
	}))
	defer s.Close()
	c := NewClient("AC123", "456bef", nil)
	c.Base = s.URL
	c.RequestOnBehalfOf("AC345")
	c.Calls.Get(context.Background(), "CA123")
	if called != true {
		t.Errorf("expected called to be true, got false")
	}
}
