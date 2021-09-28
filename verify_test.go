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
	data.Add("to", "+14155551234")
	data.Add("channel", "sms")
	verify, err := client.Verify.Verifications.Create(ctx, "VA9e0bd45bfa7d9b9e7dca86cf94c7d4f8", data)
	if err != nil {
		t.Fatal(err)
	}
	if verify.To != "+14155551234" {
		t.Errorf("expected To to be %s, got %s", "+14155551234", verify.To)
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

	verify, err := client.Verify.Verifications.Get(ctx, "VA9e0bd45bfa7d9b9e7dca86cf94c7d4f8", "+14155551234")
	if err != nil {
		t.Fatal(err)
	}
	if verify.To != "+14155551234" {
		t.Errorf("expected To to be %s, got %s", "+14155551234", verify.To)
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
	data.Add("to", "+14155551234")
	verify, err := client.Verify.Verifications.Check(ctx, "VA9e0bd45bfa7d9b9e7dca86cf94c7d4f8", data)
	if err != nil {
		t.Fatal(err)
	}
	if verify.To != "+14155551234" {
		t.Errorf("expected To to be %s, got %s", "+14155551234", verify.To)
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

func TestVerifyAccessTokenCreate(t *testing.T) {
	t.Parallel()
	client, s := getServer(verifyAccessTokenResponse)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	data := url.Values{}
	data.Add("Identity", "foo")
	data.Add("FactorType", "push")
	verify, err := client.Verify.AccessTokens.Create(ctx, "VA9e0bd45bfa7d9b9e7dca86cf94c7d4f8", data)
	if err != nil {
		t.Fatal(err)
	}
	if verify.Token != "token.stub" {
		t.Errorf("expected Token to be %s, got %s", "token.stub", verify.Token)
	}
}

func TestVerifyChallengeCreate(t *testing.T) {
	t.Parallel()
	client, s := getServer(verifyChallengeResponse)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	data := url.Values{}
	data.Add("FactorSid", "YFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	data.Add("FactorType", "push")
	data.Add("Details.Message", "Hi! Mr. John Doe, would you like to sign up?")
	data.Add("Details.Fields", "{\"label\": \"Action\", \"value\": \"Sign up in portal\"}")
	data.Add("Details.Fields", "{\"label\": \"Location\", \"value\": \"California\"}")
	challenge, err := client.Verify.Challenges.Create(ctx, "VA9e0bd45bfa7d9b9e7dca86cf94c7d4f8", "ff483d1ff591898a9942916050d2ca3f", data)
	if err != nil {
		t.Fatal(err)
	}
	if challenge.FactorSid != "YFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" {
		t.Errorf("expected FactorSid to be %s, got %s", "YFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", challenge.FactorSid)
	}
	if challenge.FactorType != "push" {
		t.Errorf("expected FactorType to be %s, got %s", "push", challenge.FactorType)
	}
}

func TestVerifyChallengeCheck(t *testing.T) {
	t.Parallel()
	client, s := getServer(verifyResponse)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	challenge, err := client.Verify.Challenges.Get(ctx, "VA9e0bd45bfa7d9b9e7dca86cf94c7d4f8", "ff483d1ff591898a9942916050d2ca3f", "YCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	if err != nil {
		t.Fatal(err)
	}
	if challenge.Sid != "YCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" {
		t.Errorf("expected Sid to be %s, got %s", "YCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", challenge.Sid)
	}
}

