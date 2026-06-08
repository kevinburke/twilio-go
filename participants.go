package twilio

import (
	"context"
	"net/url"
)

// ParticipantService manages the Participants of a single conference. Obtain
// one with Client.Conferences.Participants(conferenceSid).
//
// Note: Twilio does not return Participants after a conference ends, so these
// methods are only useful while the conference is in progress.
// https://github.com/saintpete/logrole/issues/4
type ParticipantService struct {
	client        *Client
	conferenceSid string
}

type Participant struct {
	AccountSid             string     `json:"account_sid"`
	CallSid                string     `json:"call_sid"`
	CallSidToCoach         string     `json:"call_sid_to_coach"`
	Coaching               bool       `json:"coaching"`
	ConferenceSid          string     `json:"conference_sid"`
	DateCreated            TwilioTime `json:"date_created"`
	DateUpdated            TwilioTime `json:"date_updated"`
	EndConferenceOnExit    bool       `json:"end_conference_on_exit"`
	Hold                   bool       `json:"hold"`
	Label                  string     `json:"label"`
	Muted                  bool       `json:"muted"`
	QueueTime              string     `json:"queue_time"`
	StartConferenceOnEnter bool       `json:"start_conference_on_enter"`
	Status                 string     `json:"status"`
	URI                    string     `json:"uri"`
}

type ParticipantPage struct {
	Page
	Participants []*Participant `json:"participants"`
}

func (s *ParticipantService) pathPart() string {
	return conferencePathPart + "/" + s.conferenceSid + "/Participants"
}

// Get returns the Participant with the given callSid. Participants are
// addressed by the Call SID of their leg, not by a participant SID.
func (s *ParticipantService) Get(ctx context.Context, callSid string) (*Participant, error) {
	participant := new(Participant)
	err := s.client.GetResource(ctx, s.pathPart(), callSid, participant)
	return participant, err
}

// Create adds a Participant to the conference by dialing out. The From and To
// parameters are required; see
// https://www.twilio.com/docs/voice/api/conference-participant-resource#create-a-participant
// for the full list.
func (s *ParticipantService) Create(ctx context.Context, data url.Values) (*Participant, error) {
	participant := new(Participant)
	err := s.client.CreateResource(ctx, s.pathPart(), data, participant)
	return participant, err
}

// Update changes the Participant with the given callSid, for example to hold,
// mute, or coach them. See
// https://www.twilio.com/docs/voice/api/conference-participant-resource#update-a-participant-resource
// for the available parameters.
func (s *ParticipantService) Update(ctx context.Context, callSid string, data url.Values) (*Participant, error) {
	participant := new(Participant)
	err := s.client.UpdateResource(ctx, s.pathPart(), callSid, data, participant)
	return participant, err
}

// Delete removes the Participant with the given callSid from the conference
// (kicking them out). If they are not in the conference, Delete returns nil.
func (s *ParticipantService) Delete(ctx context.Context, callSid string) error {
	return s.client.DeleteResource(ctx, s.pathPart(), callSid)
}

// GetPage returns a single page of participants for the conference. To retrieve
// multiple pages, use GetPageIterator.
func (s *ParticipantService) GetPage(ctx context.Context, data url.Values) (*ParticipantPage, error) {
	return s.GetPageIterator(data).Next(ctx)
}

type ParticipantPageIterator struct {
	p *PageIterator
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (s *ParticipantService) GetPageIterator(data url.Values) *ParticipantPageIterator {
	return &ParticipantPageIterator{
		p: NewPageIterator(s.client, data, s.pathPart()),
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (i *ParticipantPageIterator) Next(ctx context.Context) (*ParticipantPage, error) {
	pp := new(ParticipantPage)
	err := i.p.Next(ctx, pp)
	if err != nil {
		return nil, err
	}
	i.p.SetNextPageURI(pp.NextPageURI)
	return pp, nil
}
