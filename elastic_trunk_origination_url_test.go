package twilio

import (
	"context"
	"net/url"
	"testing"
)

func TestFetchOriginationUrlResource(t *testing.T) {
	t.Parallel()
	client, server := getServer(trunkResourceResponse)
	defer server.Close()
	etr, err := client.ElasticTrunks.OriginationUrls.Get(context.Background(), "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	if err != nil {
		t.Fatal(err)
	}
	if etr.Sid != "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" {
		t.Errorf("messaging resource: got sid %q, want %q", etr.Sid, "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	}
}

func TestCreateOriginationUrlResource(t *testing.T) {
	t.Parallel()
	client, server := getServer(trunkResourceResponse)
	defer server.Close()

	values := url.Values{}
	values.Add("friendly_name", "friendly_name")

	etr, err := client.ElasticTrunks.OriginationUrls.Create(context.Background(), values)
	if err != nil {
		t.Fatal(err)
	}
	if etr.FriendlyName != "friendly_name" {
		t.Errorf("messaging resource: got sid %q, want %q", etr.Sid, "friendly_name")
	}
}

func TestUpdateOriginationUrlResource(t *testing.T) {
	t.Parallel()
	client, server := getServer(trunkResourceResponse)
	defer server.Close()

	values := url.Values{}
	values.Add("friendly_name", "friendly_name")

	etr, err := client.ElasticTrunks.OriginationUrls.Update(context.Background(), "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", values)
	if err != nil {
		t.Fatal(err)
	}
	if etr.FriendlyName != "friendly_name" {
		t.Errorf("messaging resource: got sid %q, want %q", etr.Sid, "friendly_name")
	}
}

func TestDeleteOriginationUrlResource(t *testing.T) {
	t.Parallel()
	client, server := getServer(trunkResourceResponse)
	defer server.Close()

	err := client.ElasticTrunks.OriginationUrls.Delete(context.Background(), "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	if err != nil {
		t.Fatal(err)
	}
}

var originationUrlResourceResponse = []byte(`
{
  "weight": 1,
  "date_updated": "2018-05-07T20:50:58Z",
  "enabled": true,
  "friendly_name": "friendly_name",
  "account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "priority": 1,
  "sip_url": "sip://sip-box.com:1234",
  "sid": "OUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "date_created": "2018-05-07T20:50:58Z",
  "trunk_sid": "TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "url": "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/OriginationUrls/OUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}
`)
