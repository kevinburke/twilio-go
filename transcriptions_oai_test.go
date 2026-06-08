package twilio

// Transcriptions conformance checks against twilio-oai (account-level
// transcriptions). Shared machinery lives in oai_test.go.

import (
	"reflect"
	"strings"
	"testing"
)

var transcriptionSchemas = map[string]reflect.Type{
	"api.v2010.account.transcription": reflect.TypeFor[Transcription](),
}

// transcriptionPath matches the account-level Transcriptions resource only. The
// spec also has Transcriptions sub-paths under Recordings and Calls; those use
// different schemas and are classified by the recordings and calls tests.
func transcriptionPath(p string) bool {
	if strings.Contains(p, "/Recordings/") || strings.Contains(p, "/Calls/") {
		return false
	}
	return strings.Contains(p, "/Transcriptions.json") || strings.Contains(p, "/Transcriptions/")
}

// transcriptionOperations classifies every account-level Transcriptions
// operation. All are supported by TranscriptionService.
var transcriptionOperations = map[string]string{
	"GET Transcriptions.json":          "supported", // TranscriptionService.GetPage
	"GET Transcriptions/{Sid}.json":    "supported", // Get
	"DELETE Transcriptions/{Sid}.json": "supported", // Delete
}

func TestTranscriptionsOpenAPISchemaCoverage(t *testing.T) {
	assertSchemaCoverage(t, loadSpec(t), transcriptionSchemas)
}

func TestTranscriptionsOpenAPIExampleDecode(t *testing.T) {
	spec := loadSpec(t)
	if n := decodeExamples(t, spec, transcriptionPath, transcriptionSchemas); n == 0 {
		t.Fatal("no Transcriptions examples were decoded; the spec layout may have changed")
	}
}

func TestTranscriptionsOpenAPIEndpoints(t *testing.T) {
	assertOperationsClassified(t, loadSpec(t), transcriptionPath, transcriptionOperations)
}
