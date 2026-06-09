package twilio

// Addresses conformance checks against twilio-oai. Shared machinery lives in
// oai_test.go.

import (
	"reflect"
	"strings"
	"testing"
)

var addressSchemas = map[string]reflect.Type{
	"api.v2010.account.address":                        reflect.TypeFor[Address](),
	"api.v2010.account.address.dependent_phone_number": reflect.TypeFor[DependentPhoneNumber](),
}

// addressPath matches the Addresses resource and its DependentPhoneNumbers
// sub-resource.
func addressPath(path string) bool {
	return strings.Contains(path, "/Addresses.json") || strings.Contains(path, "/Addresses/")
}

// addressOperations classifies every Addresses operation in the spec. See
// assertOperationsClassified for the contract. Keep this in sync with the spec;
// an unclassified operation fails the test.
var addressOperations = map[string]string{
	"POST Addresses.json":                                   "supported", // AddressService.Create
	"GET Addresses.json":                                    "supported", // AddressService.GetPage
	"GET Addresses/{Sid}.json":                              "supported", // AddressService.Get
	"POST Addresses/{Sid}.json":                             "supported", // AddressService.Update
	"DELETE Addresses/{Sid}.json":                           "supported", // AddressService.Delete
	"GET Addresses/{AddressSid}/DependentPhoneNumbers.json": "supported", // AddressService.GetDependentPhoneNumbers
}

func TestAddressesOpenAPISchemaCoverage(t *testing.T) {
	assertSchemaCoverage(t, loadSpec(t), addressSchemas)
}

func TestAddressesOpenAPIExampleDecode(t *testing.T) {
	spec := loadSpec(t)
	if n := decodeExamples(t, spec, addressPath, addressSchemas); n == 0 {
		t.Fatal("no Addresses instance examples were decoded; the spec layout may have changed")
	}
}

func TestAddressesOpenAPIEndpoints(t *testing.T) {
	assertOperationsClassified(t, loadSpec(t), addressPath, addressOperations)
}
