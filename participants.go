package twilio

import (
	"context"
	"net/url"
	"path/filepath"
)

// It's difficult to work on this API since Twilio doesn't return Participants
// after a conference ends.
//
// https://github.com/saintpete/logrole/issues/4

const participantPathPart = "Participants"

type ParticipantService struct {
	client *Client
	sid    string
}

type Participant struct {
	AccountSid             string     `json:"account_sid"`
	CallSid                string     `json:"call_sid"`
	ConferenceSid          string     `json:"conference_sid"`
	DateCreated            TwilioTime `json:"date_created"`
	DateUpdated            TwilioTime `json:"date_updated"`
	EndConferenceOnExit    bool       `json:"end_conference_on_exit"`
	Hold                   bool       `json:"hold"`
	Muted                  bool       `json:"muted"`
	StartConferenceOnEnter bool       `json:"start_conference_on_enter"`
	URI                    string     `json:"uri"`
}

type ParticipantPage struct {
	Meta        Meta    `json:"meta"`
	Participant []*Room `json:"participants"`
}

type ParticipantPageIterator struct {
	p *PageIterator
}

// Create a participant with the given values.
//
// https://www.twilio.com/docs/voice/api/conference-participant-resource#create-a-participant
func (s *ParticipantService) Create(ctx context.Context, data url.Values) (*Participant, error) {
	participant := &Participant{}
	err := s.client.CreateResource(ctx, filepath.Join(conferencePathPart, s.sid, participantPathPart), data, participant)
	return participant, err
}

// Get a participant with the given values.
//
// https://www.twilio.com/docs/voice/api/conference-participant-resource#fetch-a-participant-resource
func (s *ParticipantService) Get(ctx context.Context, pid string) (*Participant, error) {
	participant := &Participant{}
	err := s.client.GetResource(ctx, filepath.Join(conferencePathPart, s.sid, participantPathPart), pid, participant)
	return participant, err
}

// Update the participant with the given data. Valid parameters may be found here:
// https://www.twilio.com/docs/voice/api/conference-participant-resource#update-a-participant-resource
func (s *ParticipantService) Update(ctx context.Context, pid string, data url.Values) (*Participant, error) {
	participant := &Participant{}
	err := s.client.UpdateResource(ctx, filepath.Join(conferencePathPart, s.sid, participantPathPart), pid, data, participant)
	return participant, err
}

// Delete the participant. Valid parameters may be found here:
// https://www.twilio.com/docs/voice/api/conference-participant-resource#delete-a-participant-resource
func (s *ParticipantService) Delete(ctx context.Context, pid string, data url.Values) error {
	return s.client.DeleteResource(ctx, filepath.Join(conferencePathPart, s.sid, participantPathPart), pid)
}

// Returns a list of participants for a conference. For more information on valid values,
// see https://www.twilio.com/docs/voice/api/conference-participant-resource#read-multiple-participant-resources
func (s *ParticipantService) GetPage(ctx context.Context, data url.Values) (*ParticipantPage, error) {
	return s.GetPageIterator(data).Next(ctx)
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (s *ParticipantService) GetPageIterator(data url.Values) *ParticipantPageIterator {
	return &ParticipantPageIterator{
		p: NewPageIterator(s.client, data, filepath.Join(conferencePathPart, s.sid, participantPathPart)),
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (s *ParticipantPageIterator) Next(ctx context.Context) (*ParticipantPage, error) {
	rp := &ParticipantPage{}
	err := s.p.Next(ctx, rp)
	if err != nil {
		return nil, err
	}
	s.p.SetNextPageURI(rp.Meta.NextPageURL)
	return rp, nil
}

func (s *ParticipantService) CreateParticipant(ctx context.Context, from, to string) (*Participant, error) {
	data := make(url.Values)
	data.Set("From", from)
	data.Set("To", to)
	return s.Create(ctx, data)
}

func (s *ParticipantService) Mute(ctx context.Context, pid string) (*Participant, error) {
	data := make(url.Values)
	data.Set("Muted", "True")
	return s.Update(ctx, pid, data)
}

func (s *ParticipantService) Unmute(ctx context.Context, pid string) (*Participant, error) {
	data := make(url.Values)
	data.Set("Muted", "False")
	return s.Update(ctx, pid, data)
}

func (s *ParticipantService) HoldWithMusic(ctx context.Context, pid string, holdURL string) (*Participant, error) {
	data := make(url.Values)
	data.Set("Hold", "True")
	data.Set("HoldUrl", holdURL)
	return s.Update(ctx, pid, data)
}

func (s *ParticipantService) Unhold(ctx context.Context, pid string) (*Participant, error) {
	data := make(url.Values)
	data.Set("Hold", "False")
	return s.Update(ctx, pid, data)
}

func (s *ParticipantService) SendAnnouncement(ctx context.Context, pid string, announceURL string) (*Participant, error) {
	data := make(url.Values)
	data.Set("AnnounceUrl", announceURL)
	return s.Update(ctx, pid, data)
}
