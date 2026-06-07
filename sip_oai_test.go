package twilio

// This file checks the SIP structs in this package against Twilio's official
// OpenAPI definitions (github.com/twilio/twilio-oai). It is the only
// spec-conformance test in the library and is intentionally scoped to the SIP
// resources.
//
// The test reads the checked-out spec on disk. Point it at a twilio-oai
// checkout with the TWILIO_OAI_DIR environment variable; otherwise it looks in
// $HOME/src/github.com/twilio/twilio-oai. If the spec is not present (for
// example in CI) the test skips rather than fails, so it never blocks a build
// that lacks the spec. To run it:
//
//	TWILIO_OAI_DIR=/path/to/twilio-oai go test -run TestSIPOpenAPI ./...

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"sort"
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

// openAPISpec is the subset of the OpenAPI document we need to read.
type openAPISpec struct {
	Components struct {
		Schemas map[string]struct {
			Properties map[string]json.RawMessage `json:"properties"`
		} `json:"schemas"`
	} `json:"components"`
	// The value is a map whose keys are HTTP methods ("get", "post", ...)
	// alongside non-operation keys ("servers", "description", "x-twilio"), so
	// it is kept raw and decoded per-method below.
	Paths map[string]map[string]json.RawMessage `json:"paths"`
}

// oaiOperation is a single path operation (one HTTP method).
type oaiOperation struct {
	Responses map[string]struct {
		Content struct {
			JSON struct {
				Schema   json.RawMessage `json:"schema"`
				Examples map[string]struct {
					Value json.RawMessage `json:"value"`
				} `json:"examples"`
			} `json:"application/json"`
		} `json:"content"`
	} `json:"responses"`
}

// httpMethods are the path-object keys that denote operations.
var httpMethods = map[string]bool{
	"get": true, "post": true, "put": true, "delete": true,
	"patch": true, "options": true, "head": true,
}

func loadSIPSpec(t *testing.T) *openAPISpec {
	t.Helper()
	dir := os.Getenv("TWILIO_OAI_DIR")
	if dir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			t.Skipf("twilio-oai spec not available: %v", err)
		}
		dir = filepath.Join(home, "src", "github.com", "twilio", "twilio-oai")
	}
	path := filepath.Join(dir, "spec", "json", "twilio_api_v2010.json")
	b, err := os.ReadFile(path)
	if err != nil {
		t.Skipf("twilio-oai spec not found at %s (set TWILIO_OAI_DIR to enable this test): %v", path, err)
	}
	spec := new(openAPISpec)
	if err := json.Unmarshal(b, spec); err != nil {
		t.Fatalf("could not parse %s: %v", path, err)
	}
	return spec
}

// jsonTags returns the set of json field names exposed by the struct type,
// recursing into embedded structs.
func jsonTags(typ reflect.Type) map[string]bool {
	tags := make(map[string]bool)
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if f.Anonymous && f.Type.Kind() == reflect.Struct {
			for k := range jsonTags(f.Type) {
				tags[k] = true
			}
			continue
		}
		tag := f.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}
		tags[strings.Split(tag, ",")[0]] = true
	}
	return tags
}

// TestSIPOpenAPISchemaCoverage verifies that every property in each SIP schema
// from twilio-oai is modeled by a json tag in the corresponding Go struct. A
// failure means the spec gained a field this library does not yet expose.
func TestSIPOpenAPISchemaCoverage(t *testing.T) {
	spec := loadSIPSpec(t)
	for name, typ := range sipSchemas {
		schema, ok := spec.Components.Schemas[name]
		if !ok {
			t.Errorf("schema %q not found in spec; the spec may have renamed or removed it", name)
			continue
		}
		tags := jsonTags(typ)
		var missing []string
		for prop := range schema.Properties {
			if !tags[prop] {
				missing = append(missing, prop)
			}
		}
		if len(missing) > 0 {
			sort.Strings(missing)
			t.Errorf("%s (%s) is missing struct fields for spec properties: %s",
				typ.Name(), name, strings.Join(missing, ", "))
		}
	}
}

// TestSIPOpenAPIExampleDecode decodes every instance-resource example in the
// spec into the matching Go struct with unknown fields disallowed. This catches
// both unmodeled fields and type mismatches (for example a field the spec types
// as an integer that we typed as a string).
func TestSIPOpenAPIExampleDecode(t *testing.T) {
	spec := loadSIPSpec(t)

	// Build a lookup from "#/components/schemas/<name>" to Go type so we can
	// match a response's schema $ref to one of our structs.
	byRef := make(map[string]reflect.Type, len(sipSchemas))
	for name, typ := range sipSchemas {
		byRef["#/components/schemas/"+name] = typ
	}

	decoded := 0
	for path, methods := range spec.Paths {
		if !strings.Contains(path, "/SIP/") {
			continue
		}
		for method, rawOp := range methods {
			if !httpMethods[method] {
				continue
			}
			var op oaiOperation
			if err := json.Unmarshal(rawOp, &op); err != nil {
				t.Errorf("%s %s: could not parse operation: %v", strings.ToUpper(method), path, err)
				continue
			}
			for code, resp := range op.Responses {
				// Only instance responses reference a schema directly; list
				// responses use an inline object whose paging wrapper is not
				// modeled by the resource struct.
				var ref struct {
					Ref string `json:"$ref"`
				}
				if err := json.Unmarshal(resp.Content.JSON.Schema, &ref); err != nil {
					continue
				}
				typ, ok := byRef[ref.Ref]
				if !ok {
					continue
				}
				for exName, ex := range resp.Content.JSON.Examples {
					dec := json.NewDecoder(strings.NewReader(string(ex.Value)))
					dec.DisallowUnknownFields()
					v := reflect.New(typ).Interface()
					if err := dec.Decode(v); err != nil {
						t.Errorf("%s %s %s example %q: decoding into %s failed: %v",
							strings.ToUpper(method), path, code, exName, typ.Name(), err)
						continue
					}
					decoded++
				}
			}
		}
	}
	if decoded == 0 {
		t.Fatal("no SIP instance examples were decoded; the spec layout may have changed")
	}
	t.Logf("decoded %d SIP instance examples against this package's structs", decoded)
}
