package twilio

import (
	"net/url"
	"testing"
)

func TestLookupPhoneNumbersGet(t *testing.T) {
	t.Parallel()
	client, s := getServer(phoneLookupResponse)
	defer s.Close()
	ctx := t.Context()

	data := url.Values{}
	data.Add("Type", "carrier")
	data.Add("Type", "caller-name")
	lookup, err := client.Lookup.LookupPhoneNumbers.Get(ctx, "4157012312", data)
	if err != nil {
		t.Fatal(err)
	}
	if lookup.PhoneNumber != "+14157012312" {
		t.Errorf("expected PhoneNumber to be %s, got %s", "4157012312", lookup.PhoneNumber)
	}
	if lookup.CallerName.CallerName != "CCSF" {
		t.Errorf("expected CallerName to be %s, got %s", "CCSF", lookup.CallerName.CallerName)
	}
	if lookup.Carrier.Type != "landline" {
		t.Errorf("expected Carrier.Type to be %s, got %s", "landline", lookup.Carrier.Type)
	}
}
