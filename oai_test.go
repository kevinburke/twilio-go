package twilio

// This file holds shared machinery for the OpenAPI conformance tests, which
// check this package's response structs against Twilio's official OpenAPI
// definitions (github.com/twilio/twilio-oai). Per-resource tests live in
// sip_oai_test.go, messages_oai_test.go, and calls_oai_test.go.
//
// The tests read a checked-out spec on disk. Point them at a twilio-oai
// checkout with the TWILIO_OAI_DIR environment variable; otherwise they look in
// $HOME/src/github.com/twilio/twilio-oai. If the spec is not present (for
// example in CI) the tests skip rather than fail, so they never block a build
// that lacks the spec. To run them:
//
//	TWILIO_OAI_DIR=/path/to/twilio-oai go test -run OpenAPI ./...
//
// Each resource gets up to three checks:
//
//   - schema coverage: every property in a spec schema is modeled by a json tag
//     in the corresponding struct, and simple primitive schema types match the
//     Go field type unless explicitly excepted (catches missing attributes and
//     many type mismatches).
//   - example decode: every instance example in the spec decodes into the
//     struct with unknown fields disallowed (catches missing attributes and
//     incorrect types).
//   - endpoint classification: every operation (method + path) in the spec for
//     the resource is classified as supported or unimplemented, and the test
//     fails if the spec adds an operation we have not classified (catches the
//     spec gaining endpoints, and documents which ones we do not implement).

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	gotypes "github.com/kevinburke/go-types"
)

// accountPrefix is the portion of every core-API path before the resource. It
// is stripped to form the short keys used in the endpoint-classification maps.
const accountPrefix = "/2010-04-01/Accounts/{AccountSid}/"

// openAPISpec is the subset of the OpenAPI document we need to read.
type openAPISpec struct {
	Components struct {
		Schemas map[string]openAPISchema `json:"schemas"`
	} `json:"components"`
	// The value is a map whose keys are HTTP methods ("get", "post", ...)
	// alongside non-operation keys ("servers", "description", "x-twilio"), so
	// it is kept raw and decoded per-method below.
	Paths map[string]map[string]json.RawMessage `json:"paths"`
}

type openAPISchema struct {
	Type       string                     `json:"type"`
	Properties map[string]json.RawMessage `json:"properties"`
}

type openAPIProperty struct {
	Type   string `json:"type"`
	Ref    string `json:"$ref"`
	Format string `json:"format"`
	Items  struct {
		Type string `json:"type"`
		Ref  string `json:"$ref"`
	} `json:"items"`
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

// schemaTypeExceptions documents places where the pinned OpenAPI schema type
// does not match Twilio's response examples or observed API behavior. Keep this
// list small: each entry disables one schema-vs-Go type assertion.
var schemaTypeExceptions = map[string]string{
	"api.v2010.account.call.call_recording.price": "twilio-oai marks this as number, but its response examples and the API return string prices",
	"api.v2010.account.transcription.price":       "twilio-oai marks this as number, but its response examples and the API return string prices",

	"api.v2010.account.available_phone_number_country.available_phone_number_local.latitude":      "twilio-oai marks this as number, but its response examples return string coordinates",
	"api.v2010.account.available_phone_number_country.available_phone_number_local.longitude":     "twilio-oai marks this as number, but its response examples return string coordinates",
	"api.v2010.account.available_phone_number_country.available_phone_number_mobile.latitude":     "twilio-oai marks this as number, but its response examples return string coordinates",
	"api.v2010.account.available_phone_number_country.available_phone_number_mobile.longitude":    "twilio-oai marks this as number, but its response examples return string coordinates",
	"api.v2010.account.available_phone_number_country.available_phone_number_toll_free.latitude":  "twilio-oai marks this as number, but its response examples return string coordinates",
	"api.v2010.account.available_phone_number_country.available_phone_number_toll_free.longitude": "twilio-oai marks this as number, but its response examples return string coordinates",
}

func loadSpec(t *testing.T) *openAPISpec {
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

// jsonFields returns the struct fields keyed by their json tag names, recursing
// into embedded structs.
func jsonFields(typ reflect.Type) map[string]reflect.StructField {
	fields := make(map[string]reflect.StructField)
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if f.Anonymous && f.Type.Kind() == reflect.Struct {
			maps.Copy(fields, jsonFields(f.Type))
			continue
		}
		tag := f.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}
		tag = strings.Split(tag, ",")[0]
		if tag == "" {
			continue
		}
		fields[tag] = f
	}
	return fields
}

func schemaTypeForProperty(spec *openAPISpec, prop openAPIProperty) string {
	if prop.Type != "" {
		return prop.Type
	}
	if prop.Ref == "" {
		return ""
	}
	name := strings.TrimPrefix(prop.Ref, "#/components/schemas/")
	schema := spec.Components.Schemas[name]
	return schema.Type
}

func schemaTypeLabel(spec *openAPISpec, prop openAPIProperty) string {
	if prop.Ref != "" {
		name := strings.TrimPrefix(prop.Ref, "#/components/schemas/")
		refType := spec.Components.Schemas[name].Type
		if refType != "" {
			return prop.Ref + " (" + refType + ")"
		}
		return prop.Ref
	}
	if prop.Type == "array" && prop.Items.Ref != "" {
		return "array of " + prop.Items.Ref
	}
	if prop.Type == "array" && prop.Items.Type != "" {
		return "array of " + prop.Items.Type
	}
	return prop.Type
}

func deref(typ reflect.Type) reflect.Type {
	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	return typ
}

func schemaTypeCompatible(spec *openAPISpec, fieldType reflect.Type, prop openAPIProperty) bool {
	typ := deref(fieldType)
	schemaType := schemaTypeForProperty(spec, prop)
	switch schemaType {
	case "":
		// Some schemas intentionally omit a primitive type for free-form JSON
		// objects such as encryption_details and call event payloads.
		return true
	case "string":
		return typ.Kind() == reflect.String ||
			typ == reflect.TypeFor[TwilioTime]() ||
			typ == reflect.TypeFor[TwilioDuration]() ||
			typ == reflect.TypeFor[TwilioDurationMS]() ||
			typ == reflect.TypeFor[Segments]() ||
			typ == reflect.TypeFor[NumMedia]() ||
			typ == reflect.TypeFor[NullAnsweredBy]() ||
			typ == reflect.TypeFor[gotypes.NullString]()
	case "integer":
		switch typ.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return true
		default:
			return false
		}
	case "number":
		switch typ.Kind() {
		case reflect.Float32, reflect.Float64:
			return true
		default:
			return false
		}
	case "boolean":
		return typ.Kind() == reflect.Bool
	case "object":
		return typ.Kind() == reflect.Map ||
			typ.Kind() == reflect.Struct ||
			typ == reflect.TypeFor[json.RawMessage]()
	case "array":
		return typ.Kind() == reflect.Slice || typ.Kind() == reflect.Array
	default:
		return true
	}
}

// assertSchemaCoverage verifies that every property in each named spec schema
// is modeled by a json tag in the mapped Go struct. A failure means the spec
// has a field this library does not yet expose on that struct.
func assertSchemaCoverage(t *testing.T, spec *openAPISpec, schemas map[string]reflect.Type) {
	t.Helper()
	for name, typ := range schemas {
		schema, ok := spec.Components.Schemas[name]
		if !ok {
			t.Errorf("schema %q not found in spec; the spec may have renamed or removed it", name)
			continue
		}
		fields := jsonFields(typ)
		var missing []string
		for prop, raw := range schema.Properties {
			field, ok := fields[prop]
			if !ok {
				missing = append(missing, prop)
				continue
			}

			var def openAPIProperty
			if err := json.Unmarshal(raw, &def); err != nil {
				t.Errorf("%s (%s) property %q could not be decoded: %v", typ.Name(), name, prop, err)
				continue
			}
			exceptionKey := name + "." + prop
			compatible := schemaTypeCompatible(spec, field.Type, def)
			if !compatible {
				if _, ok := schemaTypeExceptions[exceptionKey]; ok {
					continue
				}
				t.Errorf("%s.%s (%s) has Go field type %s, incompatible with spec type %s",
					typ.Name(), field.Name, exceptionKey, field.Type, schemaTypeLabel(spec, def))
				continue
			}
			if reason, ok := schemaTypeExceptions[exceptionKey]; ok {
				t.Errorf("stale schema type exception %s (%s): Go field type %s is now compatible with spec type %s",
					exceptionKey, reason, field.Type, schemaTypeLabel(spec, def))
			}
		}
		if len(missing) > 0 {
			sort.Strings(missing)
			t.Errorf("%s (%s) is missing struct fields for spec properties: %s",
				typ.Name(), name, strings.Join(missing, ", "))
		}
	}
}

// decodeExamples decodes every response example in the spec (for paths matching
// pathFilter) into the struct mapped from the response schema, with unknown
// fields disallowed. This catches both unmodeled fields and type mismatches. It
// returns the number of resource representations decoded; a caller that expects
// coverage should fail if this is zero.
//
// It handles two response shapes:
//
//   - instance responses, whose schema is a direct $ref to a resource schema:
//     the whole example is decoded into the struct.
//   - list responses, whose schema is an inline object with an array property
//     whose items $ref a resource schema: each element of that array in the
//     example is decoded into the struct. The paging wrapper around the array
//     is not modeled by the resource struct, so it is not decoded.
func decodeExamples(t *testing.T, spec *openAPISpec, pathFilter func(string) bool, schemas map[string]reflect.Type) int {
	t.Helper()
	byRef := make(map[string]reflect.Type, len(schemas))
	for name, typ := range schemas {
		byRef["#/components/schemas/"+name] = typ
	}

	// decode reports an error if raw (a single resource representation) does not
	// decode into typ with unknown fields disallowed.
	decode := func(label string, typ reflect.Type, raw json.RawMessage) bool {
		dec := json.NewDecoder(strings.NewReader(string(raw)))
		dec.DisallowUnknownFields()
		v := reflect.New(typ).Interface()
		if err := dec.Decode(v); err != nil {
			t.Errorf("%s: decoding into %s failed: %v", label, typ.Name(), err)
			return false
		}
		return true
	}

	decoded := 0
	for path, methods := range spec.Paths {
		if !pathFilter(path) {
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
				var schema struct {
					Ref        string `json:"$ref"`
					Properties map[string]struct {
						Type  string `json:"type"`
						Items struct {
							Ref string `json:"$ref"`
						} `json:"items"`
					} `json:"properties"`
				}
				if err := json.Unmarshal(resp.Content.JSON.Schema, &schema); err != nil {
					continue
				}
				label := strings.ToUpper(method) + " " + strings.TrimPrefix(path, accountPrefix) + " " + code

				// Instance response: schema is a direct $ref.
				if typ, ok := byRef[schema.Ref]; ok {
					for exName, ex := range resp.Content.JSON.Examples {
						if decode(label+" example "+exName, typ, ex.Value) {
							decoded++
						}
					}
					continue
				}

				// List response: find the array property whose items $ref a
				// known schema, then decode each element of that array.
				for prop, def := range schema.Properties {
					typ, ok := byRef[def.Items.Ref]
					if def.Type != "array" || !ok {
						continue
					}
					for exName, ex := range resp.Content.JSON.Examples {
						var wrapper map[string]json.RawMessage
						if err := json.Unmarshal(ex.Value, &wrapper); err != nil {
							continue
						}
						var items []json.RawMessage
						if err := json.Unmarshal(wrapper[prop], &items); err != nil {
							continue
						}
						for i, item := range items {
							if decode(fmt.Sprintf("%s example %s [%s #%d]", label, exName, prop, i), typ, item) {
								decoded++
							}
						}
					}
				}
			}
		}
	}
	return decoded
}

// assertOperationsClassified enumerates every operation (method + short path)
// in the spec for paths matching pathFilter and checks it against classified, a
// map keyed by "<METHOD> <short path>" (for example "GET Messages/{Sid}.json",
// where the short path has the /2010-04-01/Accounts/{AccountSid}/ prefix
// stripped). Each value is "supported" or a string beginning with
// "unimplemented" plus a short note.
//
// The test fails if the spec contains an operation absent from classified (the
// spec gained an endpoint we have not triaged) or if classified contains an
// entry absent from the spec (a stale entry to remove). Operations marked
// unimplemented are logged, not failed, so the map documents the known gap
// without turning the suite red.
func assertOperationsClassified(t *testing.T, spec *openAPISpec, pathFilter func(string) bool, classified map[string]string) {
	t.Helper()
	seen := make(map[string]bool, len(classified))
	var unimplemented []string
	supported := 0
	for path, methods := range spec.Paths {
		if !pathFilter(path) {
			continue
		}
		short := strings.TrimPrefix(path, accountPrefix)
		for method := range methods {
			if !httpMethods[method] {
				continue
			}
			key := strings.ToUpper(method) + " " + short
			seen[key] = true
			status, ok := classified[key]
			if !ok {
				t.Errorf("unclassified endpoint %q: add it to the classification map as supported or unimplemented", key)
				continue
			}
			if strings.HasPrefix(status, "unimplemented") {
				unimplemented = append(unimplemented, key+" ("+status+")")
			} else {
				supported++
			}
		}
	}
	for key := range classified {
		if !seen[key] {
			t.Errorf("stale classification entry %q: not present in the spec, remove it", key)
		}
	}
	sort.Strings(unimplemented)
	t.Logf("%d operations supported, %d unimplemented", supported, len(unimplemented))
	for _, u := range unimplemented {
		t.Logf("  unimplemented: %s", u)
	}
}
