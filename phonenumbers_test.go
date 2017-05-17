package twilio

import (
	"context"
	"testing"
	"time"

	"github.com/kevinburke/rest"
)

func TestGetNumberPage(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Parallel()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	numbers, err := envClient.IncomingNumbers.GetPage(ctx, WithPageSize(1000))
	if err != nil {
		t.Fatal(err)
	}
	if len(numbers.IncomingPhoneNumbers) == 0 {
		t.Error("expected to get a list of phone numbers, got back 0")
	}
}

func TestBuyNumber(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Parallel()
	_, err := envClient.IncomingNumbers.BuyNumber(context.Background(), "+1foobar")
	if err == nil {
		t.Fatal("expected to get an error, got nil")
	}
	rerr, ok := err.(*rest.Error)
	if !ok {
		t.Fatal("couldn't cast err to a rest.Error")
	}
	expected := "+1foobar is not a valid number"
	if rerr.Title != expected {
		t.Errorf("expected Title to be %s, got %s", expected, rerr.Title)
	}
	if rerr.StatusCode != 400 {
		t.Errorf("expected StatusCode to be 400, got %d", rerr.StatusCode)
	}
}
