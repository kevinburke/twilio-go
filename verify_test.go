package twilio

import (
	"context"
	"net/url"
	"testing"
)

func TestVerifyPhoneNumbersCreate(t *testing.T) {
	t.Parallel()
	client, s := getServer(verifyResponse)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	data := url.Values{}
	data.Add("to", "+14159373912")
	data.Add("channel", "sms")
	verify, err := client.Verify.Verifications.Create(ctx, "VA9e0bd45bfa7d9b9e7dca86cf94c7d4f8", data)
	if err != nil {
		t.Fatal(err)
	}
	if verify.To != "+14159373912" {
		t.Errorf("expected To to be %s, got %s", "+14159373912", verify.To)
	}
	if verify.Valid {
		t.Errorf("expected Valid to be %t, got %t", false, true)
	}
	if verify.Channel != "sms" {
		t.Errorf("expected Channel to be %s, got %s", "sms", verify.Channel)
	}
	if verify.Lookup.Carrier.Type != "mobile" {
		t.Errorf("expected Lookup.Carrier to be %s, got %s", "mobile", verify.Lookup.Carrier.Type)
	}
}

func TestVerifyPhoneNumbersGet(t *testing.T) {
	t.Parallel()
	client, s := getServer(verifyResponse)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	verify, err := client.Verify.Verifications.Get(ctx, "VA9e0bd45bfa7d9b9e7dca86cf94c7d4f8", "+14159373912")
	if err != nil {
		t.Fatal(err)
	}
	if verify.To != "+14159373912" {
		t.Errorf("expected To to be %s, got %s", "+14159373912", verify.To)
	}
	if verify.Valid {
		t.Errorf("expected Valid to be %t, got %t", false, true)
	}
	if verify.Channel != "sms" {
		t.Errorf("expected Channel to be %s, got %s", "sms", verify.Channel)
	}
	if verify.Lookup.Carrier.Type != "mobile" {
		t.Errorf("expected Lookup.Carrier to be %s, got %s", "mobile", verify.Lookup.Carrier.Type)
	}
}

func TestVerifyPhoneNumbersCheck(t *testing.T) {
	t.Parallel()
	client, s := getServer(verifyCheckResponse)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	data := url.Values{}
	data.Add("code", "1234")
	data.Add("to", "+14159373912")
	verify, err := client.Verify.Verifications.Check(ctx, "VA9e0bd45bfa7d9b9e7dca86cf94c7d4f8", data)
	if err != nil {
		t.Fatal(err)
	}
	if verify.To != "+14159373912" {
		t.Errorf("expected To to be %s, got %s", "+14159373912", verify.To)
	}
	if !verify.Valid {
		t.Errorf("expected Valid to be %t, got %t", true, false)
	}
	if verify.Status != "approved" {
		t.Errorf("expected Status to be %s, got %s", "approved", verify.Status)
	}
	if verify.Channel != "sms" {
		t.Errorf("expected Channel to be %s, got %s", "sms", verify.Channel)
	}
}
