package twilio

// Phone-number conformance checks against twilio-oai: incoming (owned) numbers,
// available numbers, and outgoing caller IDs. Shared machinery lives in
// oai_test.go.

import (
	"reflect"
	"strings"
	"testing"
)

var incomingNumberSchemas = map[string]reflect.Type{
	"api.v2010.account.incoming_phone_number": reflect.TypeFor[IncomingPhoneNumber](),
}

// availableNumberSchemas maps the three number-type schemas the library exposes
// (Local, Mobile, TollFree) to AvailableNumber. The seven available-number type
// schemas in the spec are structurally identical; the library models them with
// one struct.
var availableNumberSchemas = map[string]reflect.Type{
	"api.v2010.account.available_phone_number_country.available_phone_number_local":     reflect.TypeFor[AvailableNumber](),
	"api.v2010.account.available_phone_number_country.available_phone_number_mobile":    reflect.TypeFor[AvailableNumber](),
	"api.v2010.account.available_phone_number_country.available_phone_number_toll_free": reflect.TypeFor[AvailableNumber](),
}

var outgoingCallerIDSchemas = map[string]reflect.Type{
	"api.v2010.account.outgoing_caller_id": reflect.TypeFor[OutgoingCallerID](),
}

func incomingNumberPath(p string) bool {
	return strings.Contains(p, "/IncomingPhoneNumbers.json") || strings.Contains(p, "/IncomingPhoneNumbers/")
}

func availableNumberPath(p string) bool {
	return strings.Contains(p, "/AvailablePhoneNumbers.json") || strings.Contains(p, "/AvailablePhoneNumbers/")
}

func outgoingCallerIDPath(p string) bool {
	return strings.Contains(p, "/OutgoingCallerIds.json") || strings.Contains(p, "/OutgoingCallerIds/")
}

// incomingNumberOperations classifies every IncomingPhoneNumbers operation. The
// library supports the base list/create, typed Local/TollFree purchasing, and
// the instance fetch/update/release; it does not support the typed list views,
// Mobile purchasing, or the AssignedAddOns sub-resources.
var incomingNumberOperations = map[string]string{
	"GET IncomingPhoneNumbers.json":           "supported", // IncomingNumberService.GetPage
	"POST IncomingPhoneNumbers.json":          "supported", // Create / BuyNumber
	"GET IncomingPhoneNumbers/Local.json":     "unimplemented: list local owned numbers",
	"POST IncomingPhoneNumbers/Local.json":    "supported", // Local.Create
	"GET IncomingPhoneNumbers/Mobile.json":    "unimplemented: list mobile owned numbers",
	"POST IncomingPhoneNumbers/Mobile.json":   "unimplemented: purchase mobile number (no Mobile service)",
	"GET IncomingPhoneNumbers/TollFree.json":  "unimplemented: list toll-free owned numbers",
	"POST IncomingPhoneNumbers/TollFree.json": "supported", // TollFree.Create
	"GET IncomingPhoneNumbers/{Sid}.json":     "supported", // Get
	"POST IncomingPhoneNumbers/{Sid}.json":    "supported", // Update
	"DELETE IncomingPhoneNumbers/{Sid}.json":  "supported", // Release

	"GET IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns.json":                                     "unimplemented: assigned add-ons list",
	"POST IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns.json":                                    "unimplemented: assign add-on",
	"GET IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns/{Sid}.json":                               "unimplemented: assigned add-on fetch",
	"DELETE IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns/{Sid}.json":                            "unimplemented: remove assigned add-on",
	"GET IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns/{AssignedAddOnSid}/Extensions.json":       "unimplemented: assigned add-on extensions list",
	"GET IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns/{AssignedAddOnSid}/Extensions/{Sid}.json": "unimplemented: assigned add-on extension fetch",
}

// availableNumberOperations classifies every AvailablePhoneNumbers operation.
// The library exposes the country list and Local/Mobile/TollFree searches.
var availableNumberOperations = map[string]string{
	"GET AvailablePhoneNumbers.json":                                "supported", // SupportedCountries.Get
	"GET AvailablePhoneNumbers/{CountryCode}.json":                  "unimplemented: country resource",
	"GET AvailablePhoneNumbers/{CountryCode}/Local.json":            "supported", // AvailableNumbers.Local
	"GET AvailablePhoneNumbers/{CountryCode}/Mobile.json":           "supported", // AvailableNumbers.Mobile
	"GET AvailablePhoneNumbers/{CountryCode}/TollFree.json":         "supported", // AvailableNumbers.TollFree
	"GET AvailablePhoneNumbers/{CountryCode}/MachineToMachine.json": "unimplemented: machine-to-machine search",
	"GET AvailablePhoneNumbers/{CountryCode}/National.json":         "unimplemented: national search",
	"GET AvailablePhoneNumbers/{CountryCode}/SharedCost.json":       "unimplemented: shared-cost search",
	"GET AvailablePhoneNumbers/{CountryCode}/Voip.json":             "unimplemented: voip search",
}

// outgoingCallerIDOperations classifies every OutgoingCallerIds operation. All
// are supported by OutgoingCallerIDService.
var outgoingCallerIDOperations = map[string]string{
	"GET OutgoingCallerIds.json":          "supported", // GetPage
	"POST OutgoingCallerIds.json":         "supported", // Create (caller-id validation)
	"GET OutgoingCallerIds/{Sid}.json":    "supported", // Get
	"POST OutgoingCallerIds/{Sid}.json":   "supported", // Update
	"DELETE OutgoingCallerIds/{Sid}.json": "supported", // Delete
}

func TestIncomingNumbersOpenAPISchemaCoverage(t *testing.T) {
	assertSchemaCoverage(t, loadSpec(t), incomingNumberSchemas)
}

func TestIncomingNumbersOpenAPIExampleDecode(t *testing.T) {
	spec := loadSpec(t)
	if n := decodeExamples(t, spec, incomingNumberPath, incomingNumberSchemas); n == 0 {
		t.Fatal("no IncomingPhoneNumbers examples were decoded; the spec layout may have changed")
	}
}

func TestIncomingNumbersOpenAPIEndpoints(t *testing.T) {
	assertOperationsClassified(t, loadSpec(t), incomingNumberPath, incomingNumberOperations)
}

func TestAvailableNumbersOpenAPISchemaCoverage(t *testing.T) {
	assertSchemaCoverage(t, loadSpec(t), availableNumberSchemas)
}

func TestAvailableNumbersOpenAPIExampleDecode(t *testing.T) {
	spec := loadSpec(t)
	if n := decodeExamples(t, spec, availableNumberPath, availableNumberSchemas); n == 0 {
		t.Fatal("no AvailablePhoneNumbers examples were decoded; the spec layout may have changed")
	}
}

func TestAvailableNumbersOpenAPIEndpoints(t *testing.T) {
	assertOperationsClassified(t, loadSpec(t), availableNumberPath, availableNumberOperations)
}

func TestOutgoingCallerIDsOpenAPISchemaCoverage(t *testing.T) {
	assertSchemaCoverage(t, loadSpec(t), outgoingCallerIDSchemas)
}

func TestOutgoingCallerIDsOpenAPIExampleDecode(t *testing.T) {
	spec := loadSpec(t)
	if n := decodeExamples(t, spec, outgoingCallerIDPath, outgoingCallerIDSchemas); n == 0 {
		t.Fatal("no OutgoingCallerIds examples were decoded; the spec layout may have changed")
	}
}

func TestOutgoingCallerIDsOpenAPIEndpoints(t *testing.T) {
	assertOperationsClassified(t, loadSpec(t), outgoingCallerIDPath, outgoingCallerIDOperations)
}
