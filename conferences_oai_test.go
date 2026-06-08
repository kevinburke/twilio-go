package twilio

// Conferences conformance checks against twilio-oai: the conference itself, its
// participants, and its recordings. Shared machinery lives in oai_test.go.

import (
	"reflect"
	"strings"
	"testing"
)

// conferenceSchemas covers the three representations returned under
// /Conferences. The conference_recording schema is identical to the
// account-level recording schema, so it is modeled by the same Recording
// struct.
var conferenceSchemas = map[string]reflect.Type{
	"api.v2010.account.conference":                      reflect.TypeFor[Conference](),
	"api.v2010.account.conference.participant":          reflect.TypeFor[Participant](),
	"api.v2010.account.conference.conference_recording": reflect.TypeFor[Recording](),
}

func conferencePath(p string) bool {
	return strings.Contains(p, "/Conferences.json") || strings.Contains(p, "/Conferences/")
}

// conferenceOperations classifies every Conferences operation. The library
// supports the full conference surface: listing/fetching/updating conferences,
// the participant lifecycle, and conference recordings.
var conferenceOperations = map[string]string{
	"GET Conferences.json":        "supported", // ConferenceService.GetPage
	"GET Conferences/{Sid}.json":  "supported", // Get
	"POST Conferences/{Sid}.json": "supported", // Update

	"GET Conferences/{ConferenceSid}/Participants.json":              "supported", // Participants(sid).GetPage
	"POST Conferences/{ConferenceSid}/Participants.json":             "supported", // Participants(sid).Create
	"GET Conferences/{ConferenceSid}/Participants/{CallSid}.json":    "supported", // Participants(sid).Get
	"POST Conferences/{ConferenceSid}/Participants/{CallSid}.json":   "supported", // Participants(sid).Update
	"DELETE Conferences/{ConferenceSid}/Participants/{CallSid}.json": "supported", // Participants(sid).Delete

	"GET Conferences/{ConferenceSid}/Recordings.json":          "supported", // GetRecordings
	"GET Conferences/{ConferenceSid}/Recordings/{Sid}.json":    "supported", // GetRecording
	"POST Conferences/{ConferenceSid}/Recordings/{Sid}.json":   "supported", // UpdateRecording
	"DELETE Conferences/{ConferenceSid}/Recordings/{Sid}.json": "supported", // DeleteRecording
}

func TestConferencesOpenAPISchemaCoverage(t *testing.T) {
	assertSchemaCoverage(t, loadSpec(t), conferenceSchemas)
}

func TestConferencesOpenAPIExampleDecode(t *testing.T) {
	spec := loadSpec(t)
	if n := decodeExamples(t, spec, conferencePath, conferenceSchemas); n == 0 {
		t.Fatal("no Conferences examples were decoded; the spec layout may have changed")
	}
}

func TestConferencesOpenAPIEndpoints(t *testing.T) {
	assertOperationsClassified(t, loadSpec(t), conferencePath, conferenceOperations)
}
