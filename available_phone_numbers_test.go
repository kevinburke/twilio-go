package twilio

import (
	"context"
	"net/url"
	"testing"
)

func TestSearchAvailablePhoneNumbers(t *testing.T) {
	t.Parallel()
	client, server := getServer(availablePhoneNumbers)
	defer server.Close()

	data := url.Values{}
	data.Set("Contains", "571*******")
	data.Set("InRegion", "VA")
	data.Set("SmsEnabled", "true")
	data.Set("VoiceEnabled", "true")
	data.Set("PageSize", "1")

	res, err := client.AvailableNumbers.Local.GetPage(context.Background(), "US", data)
	if err != nil {
		t.Fatal(err)
	}

	if len(res.Numbers) != 1 {
		t.Errorf("expected 1 number, got %d", len(res.Numbers))
	}

	if res.Numbers[0].FriendlyName != "(571) 200-0596" {
		t.Errorf("unexpected friendly name: %s", res.Numbers[0].FriendlyName)
	}

	if res.Numbers[0].PhoneNumber != "+15712000596" {
		t.Errorf("unexpected phone number: %s", res.Numbers[0].PhoneNumber)
	}
}

func TestSupportedCountries(t *testing.T) {
	t.Parallel()
	client, server := getServer(supportedCountries)
	defer server.Close()

	res, err := client.AvailableNumbers.SupportedCountries.Get(context.Background(), false)
	if err != nil {
		t.Fatal(err)
	}

	if len(res.Countries) != 2 {
		t.Errorf("expected 2 countries, got %d", len(res.Countries))
	}

	if res.Countries[0].Country != "South Africa" {
		t.Errorf("unexpected country: %s", res.Countries[0].Country)
	}

	if res.Countries[1].Country != "Peru" {
		t.Errorf("unexpected country: %s", res.Countries[0].Country)
	}
}
