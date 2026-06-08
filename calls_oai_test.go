package twilio

// Calls conformance checks against twilio-oai. Shared machinery lives in
// oai_test.go.

import (
	"reflect"
	"strings"
	"testing"
)

var callSchemas = map[string]reflect.Type{
	"api.v2010.account.call":                            reflect.TypeFor[Call](),
	"api.v2010.account.call.call_recording":             reflect.TypeFor[CallRecording](),
	"api.v2010.account.call.call_event":                 reflect.TypeFor[CallEventLog](),
	"api.v2010.account.call.call_notification":          reflect.TypeFor[CallNotification](),
	"api.v2010.account.call.call_notification-instance": reflect.TypeFor[CallNotification](),
	"api.v2010.account.call.payments":                   reflect.TypeFor[Payment](),
}

// callPath matches the Calls resource and its sub-resources, excluding the
// unrelated SIP Domain "Auth/Calls" paths.
func callPath(path string) bool {
	if strings.Contains(path, "/SIP/") {
		return false
	}
	return strings.Contains(path, "/Calls.json") || strings.Contains(path, "/Calls/")
}

// callOperations classifies every Calls operation in the spec. See
// assertOperationsClassified for the contract. The library implements the call
// lifecycle plus events, notifications, payments, and call recordings; siprec,
// streams, realtime transcriptions, and user-defined messages are not yet
// modeled.
var callOperations = map[string]string{
	"POST Calls.json":         "supported", // CallService.Create / MakeCall
	"GET Calls.json":          "supported", // CallService.GetPage
	"GET Calls/{Sid}.json":    "supported", // CallService.Get
	"POST Calls/{Sid}.json":   "supported", // CallService.Update / Cancel / Hangup / Redirect
	"DELETE Calls/{Sid}.json": "supported", // CallService.Delete

	"GET Calls/{CallSid}/Recordings.json":          "supported", // GetCallRecordings
	"POST Calls/{CallSid}/Recordings.json":         "supported", // CreateRecording
	"GET Calls/{CallSid}/Recordings/{Sid}.json":    "supported", // GetRecording
	"POST Calls/{CallSid}/Recordings/{Sid}.json":   "supported", // UpdateRecording
	"DELETE Calls/{CallSid}/Recordings/{Sid}.json": "supported", // DeleteRecording

	"GET Calls/{CallSid}/Events.json":                                   "supported", // GetEvents
	"GET Calls/{CallSid}/Notifications.json":                            "supported", // GetNotifications
	"GET Calls/{CallSid}/Notifications/{Sid}.json":                      "supported", // GetNotification
	"POST Calls/{CallSid}/Payments.json":                                "supported", // CreatePayment
	"POST Calls/{CallSid}/Payments/{Sid}.json":                          "supported", // UpdatePayment
	"POST Calls/{CallSid}/Siprec.json":                                  "unimplemented: siprec start",
	"POST Calls/{CallSid}/Siprec/{Sid}.json":                            "unimplemented: siprec stop",
	"POST Calls/{CallSid}/Streams.json":                                 "unimplemented: media streams start",
	"POST Calls/{CallSid}/Streams/{Sid}.json":                           "unimplemented: media streams stop",
	"POST Calls/{CallSid}/Transcriptions.json":                          "unimplemented: realtime transcription start",
	"POST Calls/{CallSid}/Transcriptions/{Sid}.json":                    "unimplemented: realtime transcription stop",
	"POST Calls/{CallSid}/UserDefinedMessages.json":                     "unimplemented: user-defined messages",
	"POST Calls/{CallSid}/UserDefinedMessageSubscriptions.json":         "unimplemented: user-defined message subscriptions create",
	"DELETE Calls/{CallSid}/UserDefinedMessageSubscriptions/{Sid}.json": "unimplemented: user-defined message subscriptions delete",
}

func TestCallsOpenAPISchemaCoverage(t *testing.T) {
	assertSchemaCoverage(t, loadSpec(t), callSchemas)
}

func TestCallsOpenAPIExampleDecode(t *testing.T) {
	spec := loadSpec(t)
	if n := decodeExamples(t, spec, callPath, callSchemas); n == 0 {
		t.Fatal("no Calls instance examples were decoded; the spec layout may have changed")
	}
}

func TestCallsOpenAPIEndpoints(t *testing.T) {
	assertOperationsClassified(t, loadSpec(t), callPath, callOperations)
}
