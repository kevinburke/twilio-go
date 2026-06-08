package twilio

// SIP conformance checks against twilio-oai. The shared machinery (spec
// loading, schema coverage, example decoding) lives in oai_test.go.

import (
	"reflect"
	"strings"
	"testing"
)

// sipSchemas maps a component schema name in twilio_api_v2010.yaml to the Go
// struct in this package that models it. Every SIP schema with at least one
// property is listed; the three Auth mapping schemas share SIPAuthMapping.
var sipSchemas = map[string]reflect.Type{
	"api.v2010.account.sip.sip_credential_list":                                                                       reflect.TypeFor[SIPCredentialList](),
	"api.v2010.account.sip.sip_credential_list.sip_credential":                                                        reflect.TypeFor[SIPCredential](),
	"api.v2010.account.sip.sip_domain":                                                                                reflect.TypeFor[SIPDomain](),
	"api.v2010.account.sip.sip_domain.sip_credential_list_mapping":                                                    reflect.TypeFor[SIPDomainCredentialListMapping](),
	"api.v2010.account.sip.sip_domain.sip_ip_access_control_list_mapping":                                             reflect.TypeFor[SIPDomainIPAccessControlListMapping](),
	"api.v2010.account.sip.sip_ip_access_control_list":                                                                reflect.TypeFor[SIPIPAccessControlList](),
	"api.v2010.account.sip.sip_ip_access_control_list.sip_ip_address":                                                 reflect.TypeFor[SIPIPAddress](),
	"api.v2010.account.sip.sip_domain.sip_auth.sip_auth_calls.sip_auth_calls_credential_list_mapping":                 reflect.TypeFor[SIPAuthMapping](),
	"api.v2010.account.sip.sip_domain.sip_auth.sip_auth_calls.sip_auth_calls_ip_access_control_list_mapping":          reflect.TypeFor[SIPAuthMapping](),
	"api.v2010.account.sip.sip_domain.sip_auth.sip_auth_registrations.sip_auth_registrations_credential_list_mapping": reflect.TypeFor[SIPAuthMapping](),
}

func sipPath(path string) bool { return strings.Contains(path, "/SIP/") }

// TestSIPOpenAPISchemaCoverage verifies that every property in each SIP schema
// from twilio-oai is modeled by a json tag in the corresponding Go struct.
func TestSIPOpenAPISchemaCoverage(t *testing.T) {
	assertSchemaCoverage(t, loadSpec(t), sipSchemas)
}

// TestSIPOpenAPIExampleDecode decodes every SIP instance example in the spec
// into the matching Go struct with unknown fields disallowed.
func TestSIPOpenAPIExampleDecode(t *testing.T) {
	spec := loadSpec(t)
	if n := decodeExamples(t, spec, sipPath, sipSchemas); n == 0 {
		t.Fatal("no SIP instance examples were decoded; the spec layout may have changed")
	}
}
