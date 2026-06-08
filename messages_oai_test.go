package twilio

// Messages and Media conformance checks against twilio-oai. Shared machinery
// lives in oai_test.go.

import (
	"reflect"
	"strings"
	"testing"
)

var messageSchemas = map[string]reflect.Type{
	"api.v2010.account.message":                  reflect.TypeFor[Message](),
	"api.v2010.account.message.media":            reflect.TypeFor[Media](),
	"api.v2010.account.message.message_feedback": reflect.TypeFor[MessageFeedback](),
}

// messagePath matches the Messages resource and its Media and Feedback
// sub-resources.
func messagePath(path string) bool {
	return strings.Contains(path, "/Messages.json") || strings.Contains(path, "/Messages/")
}

// messageOperations classifies every Messages operation in the spec. See
// assertOperationsClassified for the contract. Keep this in sync with the spec;
// an unclassified operation fails the test.
var messageOperations = map[string]string{
	"POST Messages.json":                            "supported", // MessageService.Create
	"GET Messages.json":                             "supported", // MessageService.GetPage
	"GET Messages/{Sid}.json":                       "supported", // MessageService.Get
	"DELETE Messages/{Sid}.json":                    "supported", // MessageService.Delete
	"POST Messages/{Sid}.json":                      "supported", // MessageService.Update / Redact
	"GET Messages/{MessageSid}/Media.json":          "supported", // MediaService.GetPage
	"GET Messages/{MessageSid}/Media/{Sid}.json":    "supported", // MediaService.Get / GetURL
	"DELETE Messages/{MessageSid}/Media/{Sid}.json": "supported", // MediaService.Delete
	"POST Messages/{MessageSid}/Feedback.json":      "supported", // MessageService.CreateFeedback
}

func TestMessagesOpenAPISchemaCoverage(t *testing.T) {
	assertSchemaCoverage(t, loadSpec(t), messageSchemas)
}

func TestMessagesOpenAPIExampleDecode(t *testing.T) {
	spec := loadSpec(t)
	if n := decodeExamples(t, spec, messagePath, messageSchemas); n == 0 {
		t.Fatal("no Messages instance examples were decoded; the spec layout may have changed")
	}
}

func TestMessagesOpenAPIEndpoints(t *testing.T) {
	assertOperationsClassified(t, loadSpec(t), messagePath, messageOperations)
}
