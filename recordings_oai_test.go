package twilio

// Recordings conformance checks against twilio-oai (account-level recordings).
// Shared machinery lives in oai_test.go.

import (
	"reflect"
	"strings"
	"testing"
)

var recordingSchemas = map[string]reflect.Type{
	"api.v2010.account.recording": reflect.TypeFor[Recording](),
}

// recordingPath matches the account-level Recordings resource and its
// sub-resources, excluding the call- and conference-scoped recordings, which
// are distinct resources with their own schemas.
func recordingPath(p string) bool {
	if strings.Contains(p, "/Calls/") || strings.Contains(p, "/Conferences/") {
		return false
	}
	return strings.Contains(p, "/Recordings.json") || strings.Contains(p, "/Recordings/")
}

// recordingOperations classifies every account-level Recordings operation. The
// library supports listing, fetching, and deleting recordings, plus listing a
// recording's transcriptions; the add-on result sub-resources and addressing an
// individual transcription under a recording are not implemented.
var recordingOperations = map[string]string{
	"GET Recordings.json":          "supported", // RecordingService.GetPage
	"GET Recordings/{Sid}.json":    "supported", // Get
	"DELETE Recordings/{Sid}.json": "supported", // Delete

	"GET Recordings/{RecordingSid}/Transcriptions.json":          "supported", // GetTranscriptions
	"GET Recordings/{RecordingSid}/Transcriptions/{Sid}.json":    "unimplemented: recording-scoped transcription fetch",
	"DELETE Recordings/{RecordingSid}/Transcriptions/{Sid}.json": "unimplemented: recording-scoped transcription delete",

	"GET Recordings/{ReferenceSid}/AddOnResults.json":                                             "unimplemented: add-on results list",
	"GET Recordings/{ReferenceSid}/AddOnResults/{Sid}.json":                                       "unimplemented: add-on result fetch",
	"DELETE Recordings/{ReferenceSid}/AddOnResults/{Sid}.json":                                    "unimplemented: add-on result delete",
	"GET Recordings/{ReferenceSid}/AddOnResults/{AddOnResultSid}/Payloads.json":                   "unimplemented: add-on result payloads list",
	"GET Recordings/{ReferenceSid}/AddOnResults/{AddOnResultSid}/Payloads/{Sid}.json":             "unimplemented: add-on result payload fetch",
	"DELETE Recordings/{ReferenceSid}/AddOnResults/{AddOnResultSid}/Payloads/{Sid}.json":          "unimplemented: add-on result payload delete",
	"GET Recordings/{ReferenceSid}/AddOnResults/{AddOnResultSid}/Payloads/{PayloadSid}/Data.json": "unimplemented: add-on result payload data",
}

func TestRecordingsOpenAPISchemaCoverage(t *testing.T) {
	assertSchemaCoverage(t, loadSpec(t), recordingSchemas)
}

func TestRecordingsOpenAPIExampleDecode(t *testing.T) {
	spec := loadSpec(t)
	if n := decodeExamples(t, spec, recordingPath, recordingSchemas); n == 0 {
		t.Fatal("no Recordings examples were decoded; the spec layout may have changed")
	}
}

func TestRecordingsOpenAPIEndpoints(t *testing.T) {
	assertOperationsClassified(t, loadSpec(t), recordingPath, recordingOperations)
}
