package twilio

import (
	"context"
	"testing"
	"time"
)

var participantInstance = []byte(`
{
	"account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"call_sid": "CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"label": "customer",
	"conference_sid": "CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"date_created": "Fri, 18 Feb 2011 21:07:19 +0000",
	"date_updated": "Fri, 18 Feb 2011 21:07:19 +0000",
	"end_conference_on_exit": false,
	"muted": false,
	"hold": false,
	"status": "complete",
	"start_conference_on_enter": true,
	"coaching": false,
	"call_sid_to_coach": "CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"uri": "/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json"
}
`)
var participantInstanceSid = "CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"

var participantHoldInstance = []byte(`
{
	"account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"call_sid": "CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"label": null,
	"conference_sid": "CF169b5eebb07ec48e0f9f2ee904b385c5",
	"date_created": "Fri, 18 Feb 2011 21:07:19 +0000",
	"date_updated": "Fri, 18 Feb 2011 21:07:19 +0000",
	"end_conference_on_exit": false,
	"muted": true,
	"hold": true,
	"status": "complete",
	"start_conference_on_enter": true,
	"coaching": false,
	"call_sid_to_coach": "CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"uri": "/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json"
}
`)

var participantInstanceHoldURL = "http://www.myapp.com/hold"

func TestCreateParticipant(t *testing.T) {
	t.Parallel()
	client, s := getServer(participantInstance)
	defer s.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	participant, err := client.Conferences.ParticipantService(conferenceInstanceSid).CreateParticipant(ctx, "+15017122661", "+15558675310")
	if err != nil {
		t.Fatal(err)
	}
	if participant.CallSid != participantInstanceSid {
		t.Errorf("expected Sid to be %s, got %s", participantInstanceSid, participant.CallSid)
	}
}

func TestHoldParticipant(t *testing.T) {
	t.Parallel()
	client, s := getServer(participantHoldInstance)
	defer s.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	participant, err := client.Conferences.ParticipantService(conferenceInstanceSid).HoldWithMusic(ctx, participantInstanceSid, participantInstanceHoldURL)
	if err != nil {
		t.Fatal(err)
	}
	if participant.ConferenceSid != conferenceInstanceSid {
		t.Errorf("expected Sid to be %s, got %s", conferenceInstanceSid, participant.ConferenceSid)
	}
	if participant.CallSid != participantInstanceSid {
		t.Errorf("expected Sid to be %s, got %s", participantInstanceSid, participant.CallSid)
	}
	if participant.Hold != true {
		t.Errorf("expected participant to be on hold, got %t", participant.Hold)
	}
}
