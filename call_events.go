package twilio

import (
	"context"
	"encoding/json"
	"net/url"
)

// CallEventLog is a single event in the lifecycle of a Call: the request Twilio
// made (or received) and the response. Both are free-form JSON objects whose
// shape depends on the event, so they are exposed as raw JSON.
//
// It models the core API's Call Event resource. The shorter name CallEvent is
// already taken by the Voice Insights events type, which this single-package
// library cannot share; a rename to CallEvent is queued for v3 (see V3.md).
// https://www.twilio.com/docs/voice/api/call-event-resource
type CallEventLog struct {
	Request  json.RawMessage `json:"request"`
	Response json.RawMessage `json:"response"`
}

type CallEventLogPage struct {
	Page
	Events []*CallEventLog `json:"events"`
}

func callEventsPathPart(callSid string) string {
	return callsPathPart + "/" + callSid + "/Events"
}

// GetEvents returns a single page of events for the call with the given
// callSid. To retrieve multiple pages, use GetEventsIterator.
func (c *CallService) GetEvents(ctx context.Context, callSid string, data url.Values) (*CallEventLogPage, error) {
	return c.GetEventsIterator(callSid, data).Next(ctx)
}

type CallEventLogPageIterator struct {
	p *PageIterator
}

// GetEventsIterator returns an iterator over the events for the given call.
func (c *CallService) GetEventsIterator(callSid string, data url.Values) *CallEventLogPageIterator {
	return &CallEventLogPageIterator{
		p: NewPageIterator(c.client, data, callEventsPathPart(callSid)),
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (i *CallEventLogPageIterator) Next(ctx context.Context) (*CallEventLogPage, error) {
	ep := new(CallEventLogPage)
	err := i.p.Next(ctx, ep)
	if err != nil {
		return nil, err
	}
	i.p.SetNextPageURI(ep.NextPageURI)
	return ep, nil
}
