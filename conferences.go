package twilio

import (
	"context"
	"fmt"
	"maps"
	"net/url"
	"slices"
	"time"

	types "github.com/kevinburke/go-types"
)

const conferencePathPart = "Conferences"

type ConferenceService struct {
	client *Client
}

type Conference struct {
	Sid string `json:"sid"`
	// Call status, StatusInProgress or StatusCompleted
	Status       Status `json:"status"`
	FriendlyName string `json:"friendly_name"`
	// The conference region, probably "us1"
	Region                  string            `json:"region"`
	DateCreated             TwilioTime        `json:"date_created"`
	AccountSid              string            `json:"account_sid"`
	APIVersion              string            `json:"api_version"`
	DateUpdated             TwilioTime        `json:"date_updated"`
	URI                     string            `json:"uri"`
	SubresourceURIs         map[string]string `json:"subresource_uris"`
	CallSidEndingConference types.NullString  `json:"call_sid_ending_conference"`
	ReasonConferenceEnded   string            `json:"reason_conference_ended"`
}

type ConferencePage struct {
	Page
	Conferences []*Conference
}

func (c *ConferenceService) Get(ctx context.Context, sid string) (*Conference, error) {
	conference := new(Conference)
	err := c.client.GetResource(ctx, conferencePathPart, sid, conference)
	return conference, err
}

// Update the conference with the given sid, for example to end it
// (Status=completed) or to play an announcement. See
// https://www.twilio.com/docs/voice/api/conference-resource#update-a-conference-resource
// for the available parameters.
func (c *ConferenceService) Update(ctx context.Context, sid string, data url.Values) (*Conference, error) {
	conference := new(Conference)
	err := c.client.UpdateResource(ctx, conferencePathPart, sid, data, conference)
	return conference, err
}

// Participants returns a service for managing the Participants of the
// conference with the given sid.
func (c *ConferenceService) Participants(conferenceSid string) *ParticipantService {
	return &ParticipantService{client: c.client, conferenceSid: conferenceSid}
}

func conferenceRecordingsPathPart(conferenceSid string) string {
	return conferencePathPart + "/" + conferenceSid + "/" + recordingsPathPart
}

// GetRecordings returns a single page of recordings for the conference with the
// given conferenceSid. Conference recordings share the Recording representation.
func (c *ConferenceService) GetRecordings(ctx context.Context, conferenceSid string, data url.Values) (*RecordingPage, error) {
	return c.GetRecordingsIterator(conferenceSid, data).Next(ctx)
}

// GetRecordingsIterator returns an iterator over the recordings for the given
// conference.
func (c *ConferenceService) GetRecordingsIterator(conferenceSid string, data url.Values) *RecordingPageIterator {
	return &RecordingPageIterator{
		p: NewPageIterator(c.client, data, conferenceRecordingsPathPart(conferenceSid)),
	}
}

// GetRecording returns the conference recording with the given sid.
func (c *ConferenceService) GetRecording(ctx context.Context, conferenceSid string, sid string) (*Recording, error) {
	recording := new(Recording)
	err := c.client.GetResource(ctx, conferenceRecordingsPathPart(conferenceSid), sid, recording)
	return recording, err
}

// UpdateRecording changes the state of an in-progress conference recording, for
// example to pause, resume, or stop it. Set the Status parameter in data.
func (c *ConferenceService) UpdateRecording(ctx context.Context, conferenceSid string, sid string, data url.Values) (*Recording, error) {
	recording := new(Recording)
	err := c.client.UpdateResource(ctx, conferenceRecordingsPathPart(conferenceSid), sid, data, recording)
	return recording, err
}

// DeleteRecording deletes the conference recording with the given sid. If it
// has already been deleted, or does not exist, DeleteRecording returns nil.
func (c *ConferenceService) DeleteRecording(ctx context.Context, conferenceSid string, sid string) error {
	return c.client.DeleteResource(ctx, conferenceRecordingsPathPart(conferenceSid), sid)
}

func (c *ConferenceService) GetPage(ctx context.Context, data url.Values) (*ConferencePage, error) {
	return c.GetPageIterator(data).Next(ctx)
}

// GetConferencesInRange gets an Iterator containing conferences in the range
// [start, end), optionally further filtered by data. GetConferencesInRange
// panics if start is not before end. Any date filters provided in data will
// be ignored. If you have an end, but don't want to specify a start, use
// twilio.Epoch for start. If you have a start, but don't want to specify an
// end, use twilio.HeatDeath for end.
//
// Assumes that Twilio returns resources in chronological order, latest
// first. If this assumption is incorrect, your results will not be correct.
//
// Returned ConferencePages will have at most PageSize results, but may have fewer,
// based on filtering.
func (c *ConferenceService) GetConferencesInRange(start time.Time, end time.Time, data url.Values) ConferencePageIterator {
	if start.After(end) {
		panic("start date is after end date")
	}
	d := url.Values{}
	maps.Copy(d, data)
	d.Del("DateCreated")
	d.Del("Page") // just in case
	if start != Epoch {
		startFormat := start.UTC().Format(APISearchLayout)
		d.Set("DateCreated>", startFormat)
	}
	if end != HeatDeath {
		// If you specify "StartTime<=YYYY-MM-DD", the *latest* result returned
		// will be midnight (the earliest possible second) on DD. We want all
		// of the results for DD so we need to specify DD+1 in the API.
		//
		// TODO validate midnight-instant math more closely, since I don't think
		// Twilio returns the correct results for that instant.
		endFormat := end.UTC().Add(24 * time.Hour).Format(APISearchLayout)
		d.Set("DateCreated<", endFormat)
	}
	iter := NewPageIterator(c.client, d, conferencePathPart)
	return &conferenceDateIterator{
		start: start,
		end:   end,
		p:     iter,
	}
}

// GetNextConferencesInRange retrieves the page at the nextPageURI and continues
// retrieving pages until any results are found in the range given by start or
// end, or we determine there are no more records to be found in that range.
//
// If ConferencePage is non-nil, it will have at least one result.
func (c *ConferenceService) GetNextConferencesInRange(start time.Time, end time.Time, nextPageURI string) ConferencePageIterator {
	if nextPageURI == "" {
		panic("nextpageuri is empty")
	}
	iter := NewNextPageIterator(c.client, conferencePathPart)
	iter.SetNextPageURI(types.NullString{Valid: true, String: nextPageURI})
	return &conferenceDateIterator{
		start: start,
		end:   end,
		p:     iter,
	}
}

type conferenceDateIterator struct {
	p     *PageIterator
	start time.Time
	end   time.Time
}

// Next returns the next page of resources. We may need to fetch multiple
// pages from the Twilio API before we find one in the right date range, so
// latency may be higher than usual. If page is non-nil, it contains at least
// one result.
func (c *conferenceDateIterator) Next(ctx context.Context) (*ConferencePage, error) {
	var page *ConferencePage
	for {
		// just wipe it clean every time to avoid remnants hanging around
		page = new(ConferencePage)
		if err := c.p.Next(ctx, page); err != nil {
			return nil, err
		}
		if len(page.Conferences) == 0 {
			return nil, NoMoreResults
		}
		times := make([]time.Time, len(page.Conferences))
		for i, conference := range page.Conferences {
			if !conference.DateCreated.Valid {
				// we really should not ever hit this case but if we can't parse
				// a date, better to give you back an error than to give you back
				// a list of conferences that may or may not be in the time range
				return nil, fmt.Errorf("twilio: couldn't verify the date of conference: %#v", conference)
			}
			times[i] = conference.DateCreated.Time
		}
		if containsResultsInRange(c.start, c.end, times) {
			indexesToDelete := indexesOutsideRange(c.start, c.end, times)
			// reverse order so we don't delete the wrong index
			for _, index := range slices.Backward(indexesToDelete) {
				page.Conferences = append(page.Conferences[:index], page.Conferences[index+1:]...)
			}
			c.p.SetNextPageURI(page.NextPageURI)
			return page, nil
		}
		if shouldContinuePaging(c.start, times) {
			c.p.SetNextPageURI(page.NextPageURI)
			continue
		} else {
			// should not continue paging and no results in range, stop
			return nil, NoMoreResults
		}
	}
}

type ConferencePageIterator interface {
	// Next returns the next page of resources. If there are no more resources,
	// NoMoreResults is returned.
	Next(context.Context) (*ConferencePage, error)
}

// ConferencePageIterator lets you retrieve consecutive ConferencePages.
type conferencePageIterator struct {
	p *PageIterator
}

// GetPageIterator returns a ConferencePageIterator with the given page
// filters. Call iterator.Next() to get the first page of resources (and again
// to retrieve subsequent pages).
func (c *ConferenceService) GetPageIterator(data url.Values) ConferencePageIterator {
	return &conferencePageIterator{
		p: NewPageIterator(c.client, data, conferencePathPart),
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (c *conferencePageIterator) Next(ctx context.Context) (*ConferencePage, error) {
	cp := new(ConferencePage)
	err := c.p.Next(ctx, cp)
	if err != nil {
		return nil, err
	}
	c.p.SetNextPageURI(cp.NextPageURI)
	return cp, nil
}
