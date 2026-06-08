package twilio

import (
	"context"
	"encoding/json"
	"net/url"
)

// CallRecording is a recording made of a single Call leg. It differs from the
// account-level Recording in that it carries the Track that was recorded and is
// addressed under the Call that produced it.
//
// See https://www.twilio.com/docs/voice/api/recording for more.
type CallRecording struct {
	Sid               string          `json:"sid"`
	AccountSid        string          `json:"account_sid"`
	APIVersion        string          `json:"api_version"`
	CallSid           string          `json:"call_sid"`
	ConferenceSid     string          `json:"conference_sid"`
	Status            string          `json:"status"`
	Source            string          `json:"source"`
	Track             string          `json:"track"`
	Channels          uint            `json:"channels"`
	StartTime         TwilioTime      `json:"start_time"`
	DateCreated       TwilioTime      `json:"date_created"`
	DateUpdated       TwilioTime      `json:"date_updated"`
	Duration          string          `json:"duration"`
	Price             string          `json:"price"`
	PriceUnit         string          `json:"price_unit"`
	ErrorCode         int             `json:"error_code"`
	EncryptionDetails json.RawMessage `json:"encryption_details"`
	URI               string          `json:"uri"`
}

type CallRecordingPage struct {
	Page
	Recordings []*CallRecording `json:"recordings"`
}

func callRecordingsPathPart(callSid string) string {
	return callsPathPart + "/" + callSid + "/" + recordingsPathPart
}

// GetCallRecordings returns a single page of recordings for the given call
// using Twilio's call-scoped Recordings endpoint. To retrieve multiple pages,
// use GetCallRecordingsIterator.
func (c *CallService) GetCallRecordings(ctx context.Context, callSid string, data url.Values) (*CallRecordingPage, error) {
	return c.GetCallRecordingsIterator(callSid, data).Next(ctx)
}

type CallRecordingPageIterator struct {
	p *PageIterator
}

// GetCallRecordingsIterator returns an iterator over recordings for the given
// call using Twilio's call-scoped Recordings endpoint.
func (c *CallService) GetCallRecordingsIterator(callSid string, data url.Values) *CallRecordingPageIterator {
	return &CallRecordingPageIterator{
		p: NewPageIterator(c.client, data, callRecordingsPathPart(callSid)),
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (i *CallRecordingPageIterator) Next(ctx context.Context) (*CallRecordingPage, error) {
	rp := new(CallRecordingPage)
	err := i.p.Next(ctx, rp)
	if err != nil {
		return nil, err
	}
	i.p.SetNextPageURI(rp.NextPageURI)
	return rp, nil
}

// CreateRecording starts a recording on the in-progress call with the given
// callSid. See
// https://www.twilio.com/docs/voice/api/recording#create-a-recording-resource
// for the available parameters.
func (c *CallService) CreateRecording(ctx context.Context, callSid string, data url.Values) (*CallRecording, error) {
	rec := new(CallRecording)
	err := c.client.CreateResource(ctx, callRecordingsPathPart(callSid), data, rec)
	return rec, err
}

// GetRecording returns the CallRecording with the given sid for the given call.
func (c *CallService) GetRecording(ctx context.Context, callSid string, sid string) (*CallRecording, error) {
	rec := new(CallRecording)
	err := c.client.GetResource(ctx, callRecordingsPathPart(callSid), sid, rec)
	return rec, err
}

// UpdateRecording changes the state of an in-progress CallRecording, for
// example to pause, resume, or stop it. Set the Status parameter (and
// optionally PauseBehavior) in data.
func (c *CallService) UpdateRecording(ctx context.Context, callSid string, sid string, data url.Values) (*CallRecording, error) {
	rec := new(CallRecording)
	err := c.client.UpdateResource(ctx, callRecordingsPathPart(callSid), sid, data, rec)
	return rec, err
}

// DeleteRecording deletes the CallRecording with the given sid. If it has
// already been deleted, or does not exist, DeleteRecording returns nil.
func (c *CallService) DeleteRecording(ctx context.Context, callSid string, sid string) error {
	return c.client.DeleteResource(ctx, callRecordingsPathPart(callSid), sid)
}
