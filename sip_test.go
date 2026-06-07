package twilio

import (
	"context"
	"net/url"
	"testing"
)

// Response fixtures below are taken verbatim from the official OpenAPI
// definitions (github.com/twilio/twilio-oai); see sip_oai_test.go for the
// spec-conformance checks that validate the structs against the full spec.

var sipDomainInstance = []byte(`{"account_sid": "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "api_version": "2010-04-01", "auth_type": "IP_ACL", "date_created": "Mon, 20 Jul 2015 17:27:10 +0000", "date_updated": "Mon, 20 Jul 2015 17:27:10 +0000", "domain_name": "dunder-mifflin-scranton.sip.twilio.com", "friendly_name": "Scranton Office", "sip_registration": true, "sid": "SDaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "subresource_uris": {"credential_list_mappings": "/2010-04-01/Accounts/ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa/SIP/Domains/SDaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa/CredentialListMappings.json"}, "uri": "/2010-04-01/Accounts/ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa/SIP/Domains/SDaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.json", "voice_fallback_method": "POST", "voice_fallback_url": null, "voice_method": "POST", "voice_status_callback_method": "POST", "voice_status_callback_url": null, "voice_url": "https://dundermifflin.example.com/twilio/app.php", "emergency_calling_enabled": true, "secure": true, "byoc_trunk_sid": "BYaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "emergency_caller_sid": "PNaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}`)

var sipCredentialInstance = []byte(`{"account_sid": "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "credential_list_sid": "CLaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "date_created": "Wed, 19 Aug 2015 19:48:45 +0000", "date_updated": "Wed, 19 Aug 2015 19:48:45 +0000", "sid": "CRaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "uri": "/2010-04-01/Accounts/ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa/SIP/CredentialLists/CLaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa/Credentials/CRaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.json", "username": "1440013725.28"}`)

var sipAuthRegistrationMappingInstance = []byte(`{"account_sid": "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "date_created": "Thu, 30 Jul 2015 20:00:00 +0000", "date_updated": "Thu, 30 Jul 2015 20:00:00 +0000", "friendly_name": "friendly_name", "sid": "CLaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}`)

var sipIPAddressInstance = []byte(`{"account_sid": "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "date_created": "Mon, 20 Jul 2015 17:27:10 +0000", "date_updated": "Mon, 20 Jul 2015 17:27:10 +0000", "friendly_name": "office", "ip_access_control_list_sid": "ALaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "ip_address": "192.168.1.1", "cidr_prefix_length": 32, "sid": "IPaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "uri": "/2010-04-01/Accounts/ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa/SIP/IpAccessControlLists/ALaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa/IpAddresses/IPaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.json"}`)

func lastPath(s *Server) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.URLs) == 0 {
		return ""
	}
	return s.URLs[len(s.URLs)-1].Path
}

func TestSIPDomainCreate(t *testing.T) {
	t.Parallel()
	client, s := getServer(sipDomainInstance)
	defer s.Close()
	data := url.Values{
		"DomainName":   {"dunder-mifflin-scranton.sip.twilio.com"},
		"FriendlyName": {"Scranton Office"},
	}
	domain, err := client.SIP.Domains.Create(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := lastPath(s), "/2010-04-01/Accounts/AC123/SIP/Domains.json"; got != want {
		t.Errorf("expected request path %s, got %s", want, got)
	}
	if domain.DomainName != "dunder-mifflin-scranton.sip.twilio.com" {
		t.Errorf("unexpected DomainName %q", domain.DomainName)
	}
	if !domain.SIPRegistration {
		t.Error("expected SIPRegistration to be true")
	}
	if !domain.Secure {
		t.Error("expected Secure to be true")
	}
	if domain.VoiceURL != "https://dundermifflin.example.com/twilio/app.php" {
		t.Errorf("unexpected VoiceURL %q", domain.VoiceURL)
	}
}

func TestSIPCredentialNestedPath(t *testing.T) {
	t.Parallel()
	client, s := getServer(sipCredentialInstance)
	defer s.Close()
	clSid := "CLaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	crSid := "CRaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	cred, err := client.SIP.CredentialLists.Credentials(clSid).Get(context.Background(), crSid)
	if err != nil {
		t.Fatal(err)
	}
	want := "/2010-04-01/Accounts/AC123/SIP/CredentialLists/" + clSid + "/Credentials/" + crSid + ".json"
	if got := lastPath(s); got != want {
		t.Errorf("expected request path %s, got %s", want, got)
	}
	if cred.Username != "1440013725.28" {
		t.Errorf("unexpected Username %q", cred.Username)
	}
	if cred.CredentialListSid != clSid {
		t.Errorf("unexpected CredentialListSid %q", cred.CredentialListSid)
	}
}

func TestSIPAuthRegistrationMappingCreate(t *testing.T) {
	t.Parallel()
	client, s := getServer(sipAuthRegistrationMappingInstance)
	defer s.Close()
	domainSid := "SDaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	data := url.Values{"CredentialListSid": {"CLaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}}
	mapping, err := client.SIP.Domains.AuthRegistrationsCredentialListMappings(domainSid).Create(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}
	want := "/2010-04-01/Accounts/AC123/SIP/Domains/" + domainSid + "/Auth/Registrations/CredentialListMappings.json"
	if got := lastPath(s); got != want {
		t.Errorf("expected request path %s, got %s", want, got)
	}
	if mapping.Sid != "CLaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" {
		t.Errorf("unexpected Sid %q", mapping.Sid)
	}
}

func TestSIPIPAddressNestedPath(t *testing.T) {
	t.Parallel()
	client, s := getServer(sipIPAddressInstance)
	defer s.Close()
	alSid := "ALaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	data := url.Values{
		"FriendlyName": {"office"},
		"IpAddress":    {"192.168.1.1"},
	}
	addr, err := client.SIP.IPAccessControlLists.IPAddresses(alSid).Create(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}
	want := "/2010-04-01/Accounts/AC123/SIP/IpAccessControlLists/" + alSid + "/IpAddresses.json"
	if got := lastPath(s); got != want {
		t.Errorf("expected request path %s, got %s", want, got)
	}
	if addr.IPAddress != "192.168.1.1" {
		t.Errorf("unexpected IPAddress %q", addr.IPAddress)
	}
	if addr.CIDRPrefixLength != 32 {
		t.Errorf("expected CIDRPrefixLength 32, got %d", addr.CIDRPrefixLength)
	}
}
