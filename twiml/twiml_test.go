package twiml

import (
	"encoding/xml"
	"net/http"
	"testing"

	"github.com/matryer/is"
)

type twimlTestCase struct {
	in  *TwiML
	out string
}

func TestDialTwiML(t *testing.T) {
	tests := []twimlTestCase{
		{ // Simple dial
			in:  &TwiML{Dial: &Dial{Dial: "123-456-7890"}, Say: &Say{Say: "Goodbye"}},
			out: "<Response><Dial>123-456-7890</Dial><Say>Goodbye</Say></Response>",
		},
		{ //Dial a phone number from a Twilio Client
			in:  &TwiML{Dial: &Dial{Number: "123-456-7890"}, Say: &Say{Say: "Goodbye"}},
			out: "<Response><Dial><Number>123-456-7890</Number></Dial><Say>Goodbye</Say></Response>",
		},
		{ // Specify an action URL and method
			in:  &TwiML{Dial: &Dial{Number: "+15558675310", CallerId: "+15551112222"}},
			out: `<Response><Dial callerId="+15551112222"><Number>+15558675310</Number></Dial></Response>`,
		},
		{ // Set Dual-Channel Recording on <Dial>
			in:  &TwiML{Dial: &Dial{Dial: "415-123-4567", Action: "/handleDialCallStatus", Method: http.MethodGet}, Say: &Say{Say: "I am unreachable"}},
			out: `<Response><Dial action="/handleDialCallStatus" method="GET">415-123-4567</Dial><Say>I am unreachable</Say></Response>`,
		},
		{ // SIP REFER Inbound to Twilio using referUrl
			in:  &TwiML{Dial: &Dial{Number: "+15558675310", Record: "record-from-ringing-dual", RecordingStatusCallback: "www.myexample.com"}},
			out: `<Response><Dial record="record-from-ringing-dual" recordingStatusCallback="www.myexample.com"><Number>+15558675310</Number></Dial></Response>`,
		},
		{ // Use a nested <Number> with Dial
			in: &TwiML{Dial: &Dial{
				SIP:      "sip:AgentA@xyz.sip.us1.twilio.com?User-to-User=123456789%3Bencoding%3Dhex&amp;X-Name=Agent%2C+A",
				ReferURL: "https://example.com/handler"},
			},
			out: `<Response><Dial referURL="https://example.com/handler"><SIP>sip:AgentA@xyz.sip.us1.twilio.com?User-to-User=123456789%3Bencoding%3Dhex&amp;amp;X-Name=Agent%2C+A</SIP></Dial></Response>`,
		},
		{ // Set Dual-Channel Recording on <Dial> with nested <Conference>;
			in: &TwiML{Dial: &Dial{
				Record:                  "record-from-ringing-dual",
				RecordingStatusCallback: "www.myexample.com",
				Conference:              &Conference{Conference: "myteamroom"},
			}},
			out: `<Response><Dial record="record-from-ringing-dual" recordingStatusCallback="www.myexample.com"><Conference>myteamroom</Conference></Dial></Response>`,
		},
	}
	testTwiML(t, tests)
}
func TestSayTwiML(t *testing.T) {
	tests := []twimlTestCase{

		{ // Using attributes in a Say verb
			in:  &TwiML{Say: &Say{Say: "Chapeau!", Voice: "woman", Language: "fr-FR"}},
			out: `<Response><Say voice="woman" language="fr-FR">Chapeau!</Say></Response>`,
		},
	}
	testTwiML(t, tests)
}
func TestConferenceTwiML(t *testing.T) {

	TRUE := true
	FALSE := false

	tests := []twimlTestCase{
		{ // A Simple Conference
			in:  &TwiML{Dial: &Dial{Conference: &Conference{Conference: "Room 1234"}}},
			out: `<Response><Dial><Conference>Room 1234</Conference></Dial></Response>`,
		},
		{ // A Moderated Conference
			in:  &TwiML{Dial: &Dial{Conference: &Conference{Conference: "moderated-conference-room", StartConferenceOnEnter: &FALSE}}},
			out: `<Response><Dial><Conference startConferenceOnEnter="false">moderated-conference-room</Conference></Dial></Response>`,
		},
		{ // A Moderated Conference (begin on enter)
			in: &TwiML{Dial: &Dial{Conference: &Conference{
				Conference:             "moderated-conference-room",
				StartConferenceOnEnter: &TRUE,
				EndConferenceOnExit:    &TRUE,
			}}},
			out: `<Response><Dial><Conference startConferenceOnEnter="true" endConferenceOnExit="true">moderated-conference-room</Conference></Dial></Response>`,
		},
		{ // Join an Evented Conference
			in: &TwiML{Dial: &Dial{Conference: &Conference{
				Conference:          "EventedConf",
				StatusCallback:      "https://myapp.com/events",
				StatusCallbackEvent: "start end join leave mute hold",
			}}},
			out: `<Response><Dial><Conference statusCallback="https://myapp.com/events" statusCallbackEvent="start end join leave mute hold">EventedConf</Conference></Dial></Response>`,
		},
		{ // Join a Conference Muted
			in: &TwiML{Dial: &Dial{Conference: &Conference{
				Conference: "EventedConf",
				Muted:      &TRUE,
			}}},
			out: `<Response><Dial><Conference muted="true">EventedConf</Conference></Dial></Response>`,
		},
		{ // Bridging Calls
			in: &TwiML{Dial: &Dial{Conference: &Conference{
				Conference:             "NoMusicNoBeepRoom",
				StartConferenceOnEnter: &TRUE,
				EndConferenceOnExit:    &TRUE,
				Beep:                   &FALSE,
				WaitUrl:                "http://your-webhook-host.com",
			}}},
			out: `<Response><Dial><Conference startConferenceOnEnter="true" endConferenceOnExit="true" beep="false" waitUrl="http://your-webhook-host.com">NoMusicNoBeepRoom</Conference></Dial></Response>`,
		},
		{ // Call on Hold
			in: &TwiML{Dial: &Dial{Conference: &Conference{
				Conference: "Customer Waiting Room",
				Beep:       &FALSE,
			}}},
			out: `<Response><Dial><Conference beep="false">Customer Waiting Room</Conference></Dial></Response>`,
		},
		{ // Call on Hold (end on exit)
			in: &TwiML{Dial: &Dial{Conference: &Conference{
				Conference:          "Customer Waiting Room",
				Beep:                &FALSE,
				EndConferenceOnExit: &TRUE,
			}}},
			out: `<Response><Dial><Conference endConferenceOnExit="true" beep="false">Customer Waiting Room</Conference></Dial></Response>`,
		},
		{ // Combining with Dial attributes
			in: &TwiML{Dial: &Dial{
				Action:       "handleLeaveConference.php",
				Method:       "POST",
				HangupOnStar: &TRUE,
				Conference:   &Conference{Conference: "LoveTwilio"},
			}},
			out: `<Response><Dial action="handleLeaveConference.php" hangupOnStar="true" method="POST"><Conference>LoveTwilio</Conference></Dial></Response>`,
		},
		{ // Recording a Conference
			in: &TwiML{Dial: &Dial{Conference: &Conference{
				Conference:              "LoveTwilio",
				Record:                  "record-from-start",
				RecordingStatusCallback: "www.myexample.com",
			}}},
			out: `<Response><Dial><Conference record="record-from-start" recordingStatusCallback="www.myexample.com">LoveTwilio</Conference></Dial></Response>`,
		},
	}

	testTwiML(t, tests)
}

func testTwiML(t *testing.T, tests []twimlTestCase) {
	t.Helper()
	is := is.New(t)

	for _, tt := range tests {
		b, err := xml.Marshal(tt.in)
		is.NoErr(err)
		t.Log(string(b))
		is.Equal(string(b), tt.out)
	}
}
