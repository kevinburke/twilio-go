package twilio

import (
	"context"
	"net/url"
	"testing"
)

func TestGetAddress(t *testing.T) {
	t.Parallel()
	client, server := getServer(addressInstance)
	defer server.Close()
	addr, err := client.Addresses.Get(context.Background(), "ADaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	if err != nil {
		t.Fatal(err)
	}
	if addr.FriendlyName != "Main Office" {
		t.Errorf("got bad friendly name: %s", addr.FriendlyName)
	}
	if addr.City != "SF" || addr.Region != "CA" || addr.IsoCountry != "US" {
		t.Errorf("got bad address fields: %+v", addr)
	}
	if addr.StreetSecondary != "Suite 300" {
		t.Errorf("got bad street_secondary: %s", addr.StreetSecondary)
	}
}

func TestCreateAddress(t *testing.T) {
	t.Parallel()
	client, server := getServer(addressInstance)
	defer server.Close()
	data := url.Values{}
	data.Set("CustomerName", "name")
	data.Set("Street", "4th")
	data.Set("City", "SF")
	data.Set("Region", "CA")
	data.Set("PostalCode", "94019")
	data.Set("IsoCountry", "US")
	addr, err := client.Addresses.Create(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}
	if addr.Sid != "ADaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" {
		t.Errorf("got bad sid: %s", addr.Sid)
	}
	if want := "/2010-04-01/Accounts/AC123/Addresses.json"; server.URLs[0].Path != want {
		t.Errorf("got path %s, want %s", server.URLs[0].Path, want)
	}
}

func TestGetDependentPhoneNumbers(t *testing.T) {
	t.Parallel()
	client, server := getServer(dependentPhoneNumbersPage)
	defer server.Close()
	page, err := client.Addresses.GetDependentPhoneNumbers(context.Background(),
		"ADaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", url.Values{})
	if err != nil {
		t.Fatal(err)
	}
	if len(page.DependentPhoneNumbers) != 1 {
		t.Fatalf("got %d dependent phone numbers, want 1", len(page.DependentPhoneNumbers))
	}
	num := page.DependentPhoneNumbers[0]
	if num.PhoneNumber.Friendly() != "+1 925-271-7005" {
		t.Errorf("got bad phone number: %s", num.PhoneNumber.Friendly())
	}
	if num.Capabilities == nil || !num.Capabilities.SMS {
		t.Errorf("expected SMS capability, got %+v", num.Capabilities)
	}
	want := "/2010-04-01/Accounts/AC123/Addresses/ADaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa/DependentPhoneNumbers.json"
	if server.URLs[0].Path != want {
		t.Errorf("got path %s, want %s", server.URLs[0].Path, want)
	}
}
