package twiml

import (
	"encoding/xml"
	"fmt"
	"net/url"
)

// TwiML xml datagram
type TwiML struct {
	XMLName struct{} `xml:"Response"`
	Dial    *Dial    `xml:",omitempty"`
	Say     *Say     `xml:",omitempty"`
	Play    *Play    `xml:",omitempty"`
	Enqueue *Enqueue `xml:",omitempty"`
}

func (ml *TwiML) String() string {
	b, err := xml.Marshal(ml)
	if err != nil {
		return fmt.Sprintf("<!-- error: %v -->", err)
	}

	return string(b)
}

func (o *TwiML) ApplyValues(v url.Values) {
	v.Add("twiml", o.String())
}

// Dial TwiML xml datagram
// docs: https://www.twilio.com/docs/voice/twiml/dial
type Dial struct {
	Dial string `xml:",chardata"`

	Action                        string `xml:"action,attr,omitempty"`
	AnswerOnBridge                *bool  `xml:"answerOnBridge,attr,omitempty"`
	CallerId                      string `xml:"callerId,attr,omitempty"`
	CallReason                    string `xml:"callReason,attr,omitempty"`
	HangupOnStar                  *bool  `xml:"hangupOnStar,attr,omitempty"`
	Method                        string `xml:"method,attr,omitempty"`
	Record                        string `xml:"record,attr,omitempty"`
	RecordingStatusCallback       string `xml:"recordingStatusCallback,attr,omitempty"`
	RecordingStatusCallbackMethod string `xml:"recordingStatusCallbackMethod,attr,omitempty"`
	RecordingStatusCallbackEvent  string `xml:"recordingStatusCallbackEvent,attr,omitempty"`
	RecordingTrack                string `xml:"recordingTrack,attr,omitempty"`
	ReferURL                      string `xml:"referURL,attr,omitempty"`
	ReferMethod                   string `xml:"referMethod,attr,omitempty"`
	RingTone                      string `xml:"ringTone,attr,omitempty"`
	TimeLimit                     *uint  `xml:"timeLimit,attr,omitempty"`
	Timeout                       *uint  `xml:"timeout,attr,omitempty"`
	Trim                          string `xml:"trim,attr,omitempty"`

	Client     string      `xml:",omitempty"`
	Conference *Conference `xml:",omitempty"`
	Number     string      `xml:",omitempty"`
	Queue      string      `xml:",omitempty"`
	SIM        string      `xml:",omitempty"`
	SIP        string      `xml:",omitempty"`
}

// Say TwiML xml datagram
// docs: https://www.twilio.com/docs/voice/twiml/say
type Say struct {
	Say string `xml:",chardata"`

	Voice    string `xml:"voice,attr,omitempty"`
	Loop     int    `xml:"loop,attr,omitempty"`
	Language string `xml:"language,attr,omitempty"`
}

// Conference TwiML xml datagram
// docs: https://www.twilio.com/docs/voice/twiml/conference
type Conference struct {
	Conference string `xml:",chardata"`

	EventCallbackURL       string `xml:"eventCallbackURL,attr,omitempty"`
	StartConferenceOnEnter *bool  `xml:"startConferenceOnEnter,attr,omitempty"`
	EndConferenceOnExit    *bool  `xml:"endConferenceOnExit,attr,omitempty"`
	StatusCallback         string `xml:"statusCallback,attr,omitempty"`
	StatusCallbackEvent    string `xml:"statusCallbackEvent,attr,omitempty"`
	StatusCallbackMethod   string `xml:"statusCallbackMethod,attr,omitempty"`

	Muted            *bool  `xml:"muted,attr,omitempty"`
	Beep             *bool  `xml:"beep,attr,omitempty"`
	ParticipantLabel string `xml:"participantLabel,attr,omitempty"`
	JitterBufferSize string `xml:"jitterBufferSize,attr,omitempty"`
	MaxParticipants  *int   `xml:"maxParticipants,attr,omitempty"`
	Region           string `xml:"region,attr,omitempty"`
	Trim             string `xml:"trim,attr,omitempty"`
	Coach            string `xml:"coach,attr,omitempty"`

	WaitUrl    string `xml:"waitUrl,attr,omitempty"`
	WaitMethod string `xml:"waitMethod,attr,omitempty"`

	Record                        string `xml:"record,attr,omitempty"`
	RecordingStatusCallback       string `xml:"recordingStatusCallback,attr,omitempty"`
	RecordingStatusCallbackMethod string `xml:"recordingStatusCallbackMethod,attr,omitempty"`
	RecordingStatusCallbackEvent  string `xml:"recordingStatusCallbackEvent,attr,omitempty"`
}

// Play TwiML xml datagram
// docs: https://www.twilio.com/docs/voice/twiml/play
type Play struct {
	Play string `xml:",chardata"`
}

// Enqueue TwiML xml datagram
// docs: https://www.twilio.com/docs/voice/twiml/enqueue
type Enqueue struct {
	Enqueue string `xml:",chardata"`

	Action      string `xml:"action,attr,omitempty"`
	Method      string `xml:"method,attr,omitempty"`
	WaitURL     string `xml:"waitUrl,attr,omitempty"`
	WorkflowSID string `xml:"workflowSid,attr,omitempty"`

	Task *Task `xml:",omitempty"`
}

// Task TwiML xml datagram
// docs: Task = The attributes to be set for the newly created task, formatted as JSON
type Task struct {
	Task     string `xml:",innerxml"`
	Priority string `xml:",innerxml"`
}
