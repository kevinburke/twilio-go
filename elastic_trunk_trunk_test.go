package twilio

import (
	"context"
	"net/url"
	"testing"
)

func TestFetchTrunkResource(t *testing.T) {
	t.Parallel()
	client, server := getServer(trunkResourceResponse)
	defer server.Close()
	etr, err := client.ElasticTrunks.Trunks.Get(context.Background(), "TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	if err != nil {
		t.Fatal(err)
	}
	if etr.Sid != "TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" {
		t.Errorf("messaging resource: got sid %q, want %q", etr.Sid, "TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	}
}

func TestCreateTrunkResource(t *testing.T) {
	t.Parallel()
	client, server := getServer(trunkResourceResponse)
	defer server.Close()

	values := url.Values{}
	values.Add("friendly_name", "friendly_name")

	etr, err := client.ElasticTrunks.Trunks.Create(context.Background(), values)
	if err != nil {
		t.Fatal(err)
	}
	if etr.FriendlyName != "friendly_name" {
		t.Errorf("messaging resource: got sid %q, want %q", etr.Sid, "friendly_name")
	}
}

func TestUpdateTrunkResource(t *testing.T) {
	t.Parallel()
	client, server := getServer(trunkResourceResponse)
	defer server.Close()

	values := url.Values{}
	values.Add("friendly_name", "friendly_name")

	etr, err := client.ElasticTrunks.Trunks.Update(context.Background(), "TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", values)
	if err != nil {
		t.Fatal(err)
	}
	if etr.FriendlyName != "friendly_name" {
		t.Errorf("messaging resource: got sid %q, want %q", etr.Sid, "friendly_name")
	}
}

func TestDeleteTrunkResource(t *testing.T) {
	t.Parallel()
	client, server := getServer(trunkResourceResponse)
	defer server.Close()

	err := client.ElasticTrunks.Trunks.Delete(context.Background(), "TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	if err != nil {
		t.Fatal(err)
	}
}

var trunkResourceResponse = []byte(`
{
  "sid": "TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "domain_name": "test.pstn.twilio.com",
  "disaster_recovery_method": "POST",
  "disaster_recovery_url": "http://disaster-recovery.com",
  "friendly_name": "friendly_name",
  "secure": false,
  "cnam_lookup_enabled": false,
  "recording": {
    "mode": "do-not-record",
    "trim": "do-not-trim"
  },
  "transfer_mode": "disable-all",
  "auth_type": "",
  "auth_type_set": [],
  "date_created": "2015-01-02T11:23:45Z",
  "date_updated": "2015-01-02T11:23:45Z",
  "url": "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "links": {
    "origination_urls": "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/OriginationUrls",
    "credential_lists": "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/CredentialLists",
    "ip_access_control_lists": "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAccessControlLists",
    "phone_numbers": "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers"
  }
}
`)
