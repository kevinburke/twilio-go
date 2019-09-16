package twilio

import (
	"context"
	"testing"
)

func TestFetchMessagingResource(t *testing.T) {
	t.Parallel()
	client, server := getServer(messagingResourceResponse)
	defer server.Close()
	msr, err := client.Resource.ServiceResourceService.Fetch(context.Background(), "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	if err != nil {
		t.Fatal(err)
	}
	if msr.Sid != "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" {
		t.Errorf("messaging resource: got sid %q, want %q", msr.Sid, "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	}
}

func TestFetchPhoneNumberResource(t *testing.T) {
	t.Parallel()
	client, server := getServer(phoneNumberResourceResponse)
	defer server.Close()
	pnr, err := client.Resource.ServiceResourceService.PhoneNumber.Fetch(context.Background(), "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", "PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	if err != nil {
		t.Fatal(err)
	}
	if pnr.ServiceSid != "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" {
		t.Errorf("messaging resource: got sid %q, want %q", pnr.ServiceSid, "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	}
	if pnr.Sid != "PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" {
		t.Errorf("phone number resource: got sid %q, want %q", pnr.Sid, "PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	}
}

func TestFetchAlphaSenderResource(t *testing.T) {
	t.Parallel()
	client, server := getServer(alphaSenderResourceResponse)
	defer server.Close()
	asr, err := client.Resource.ServiceResourceService.AlphaSender.Fetch(context.Background(), "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", "AIXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	if err != nil {
		t.Fatal(err)
	}
	if asr.ServiceSid != "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" {
		t.Errorf("messaging resource: got sid %q, want %q", asr.ServiceSid, "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	}
	if asr.Sid != "AIXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" {
		t.Errorf("alpha sender resource: got sid %q, want %q", asr.Sid, "AIXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	}
}

func TestFetchShortCodeResource(t *testing.T) {
	t.Parallel()
	client, server := getServer(shortCodeResourceResponse)
	defer server.Close()
	asr, err := client.Resource.ServiceResourceService.ShortCode.Fetch(context.Background(), "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", "SCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	if err != nil {
		t.Fatal(err)
	}
	if asr.ServiceSid != "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" {
		t.Errorf("messaging resource: got sid %q, want %q", asr.ServiceSid, "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	}
	if asr.Sid != "SCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" {
		t.Errorf("short code resource: got sid %q, want %q", asr.Sid, "SCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	}
}


var messagingResourceResponse = []byte(`
{
  "account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "sid": "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "date_created": "2015-07-30T20:12:31Z",
  "date_updated": "2015-07-30T20:12:33Z",
  "friendly_name": "My Service!",
  "inbound_request_url": "https://www.example.com/",
  "inbound_method": "POST",
  "fallback_url": "https://www.example.com",
  "fallback_method": "GET",
  "status_callback": "https://www.example.com",
  "sticky_sender": true,
  "smart_encoding": false,
  "mms_converter": true,
  "fallback_to_long_code": true,
  "scan_message_content": "inherit",
  "area_code_geomatch": true,
  "validity_period": 600,
  "synchronous_validation": true,
  "links": {
    "phone_numbers": "https://messaging.twilio.com/v1/Services/MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers",
    "short_codes": "https://messaging.twilio.com/v1/Services/MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ShortCodes",
    "alpha_senders": "https://messaging.twilio.com/v1/Services/MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/AlphaSenders",
    "messages": "https://messaging.twilio.com/v1/Services/MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages",
    "broadcasts": "https://messaging.twilio.com/v1/Services/MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Broadcasts"
  },
  "url": "https://messaging.twilio.com/v1/Services/MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}
`)

var phoneNumberResourceResponse = []byte(`
{
  "sid": "PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "service_sid": "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "date_created": "2015-07-30T20:12:31Z",
  "date_updated": "2015-07-30T20:12:33Z",
  "phone_number": "+987654321",
  "country_code": "US",
  "capabilities": [],
  "url": "https://messaging.twilio.com/v1/Services/MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}
`)

var alphaSenderResourceResponse = []byte(`
{
  "sid": "AIXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "service_sid": "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "date_created": "2015-07-30T20:12:31Z",
  "date_updated": "2015-07-30T20:12:33Z",
  "alpha_sender": "My company",
  "capabilities": [],
  "url": "https://messaging.twilio.com/v1/Services/MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/AlphaSenders/AIXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}
`)

var shortCodeResourceResponse = []byte(`
{
  "sid": "SCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "service_sid": "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "date_created": "2015-07-30T20:12:31Z",
  "date_updated": "2015-07-30T20:12:33Z",
  "short_code": "12345",
  "country_code": "US",
  "capabilities": [],
  "url": "https://messaging.twilio.com/v1/Services/MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ShortCodes/SCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}
`)