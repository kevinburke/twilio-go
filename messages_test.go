package twilio

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Parallel()
	sid := "SM7c734f6e057ff829bb20c7211cfb3ce1"
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	msg, err := envClient.Messages.Get(ctx, sid)
	if err != nil {
		t.Fatal(err)
	}
	if msg.Sid != sid {
		t.Errorf("expected Sid to equal %s, got %s", sid, msg.Sid)
	}
}

func TestGetPage(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Parallel()
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	page, err := envClient.Messages.GetPage(ctx, url.Values{"PageSize": []string{"5"}})
	if err != nil {
		t.Fatal(err)
	}
	if len(page.Messages) != 5 {
		t.Fatalf("expected len(messages) to be 5, got %d", len(page.Messages))
	}
}

func TestSendMessage(t *testing.T) {
	t.Parallel()
	client, s := getServer(sendMessageResponse)
	defer s.Close()
	msg, err := client.Messages.SendMessage(from, to, "twilio-go testing!", nil)
	if err != nil {
		t.Fatal(err)
	}
	if msg.From != from {
		t.Errorf("expected From to be from, got error")
	}
	if msg.Body != "twilio-go testing!" {
		t.Errorf("expected Body to be twilio-go testing, got %s", msg.Body)
	}
	if msg.NumSegments != 1 {
		t.Errorf("expected NumSegments to be 1, got %d", msg.NumSegments)
	}
	if msg.Status != StatusQueued {
		t.Errorf("expected Status to be StatusQueued, got %s", msg.Status)
	}
}

func TestGetMessage(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Skip("broke because of message archiving rules")
	t.Parallel()
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	msg, err := envClient.Messages.Get(ctx, "SM7c734f6e057ff829bb20c7211cfb3ce1")
	if err != nil {
		t.Fatal(err)
	}
	if msg.ErrorCode != CodeUnknownDestination {
		t.Errorf("expected Code to be %d, got %d", CodeUnknownDestination, msg.ErrorCode)
	}
	if msg.ErrorMessage == "" {
		t.Errorf(`expected ErrorMessage to be non-empty, got ""`)
	}
}

func TestIterateAll(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Skip("broke because of message archiving rules")
	t.Parallel()
	iter := envClient.Messages.GetPageIterator(url.Values{"PageSize": []string{"500"}})
	count := 0
	start := uint(0)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	for {
		page, err := iter.Next(ctx)
		if err == NoMoreResults {
			break
		}
		if err != nil {
			t.Fatal(err)
			return
		}
		if count > 0 && (page.Start <= start || page.Start-start > 500) {
			t.Fatalf("expected page.Start to be greater than previous, got %d, previous %d", page.Start, start)
			return
		} else {
			start = page.Start
		}
		if err != nil {
			t.Fatal(err)
			break
		}
		count++
		if count > 20 {
			fmt.Println("count > 20")
			t.Fail()
			break
		}
	}
	if count < 10 {
		t.Errorf("Too small of a count - expected at least 10, got %d", count)
	}
}

func TestGetMediaURLs(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Parallel()
	t.Skip("twilio removed access to old messages")
	sid := os.Getenv("TWILIO_ACCOUNT_SID")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	urls, err := envClient.Messages.GetMediaURLs(ctx, "MM89a8c4a6891c53054e9cd604922bfb61", nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(urls) != 1 {
		t.Fatalf("Wrong number of URLs returned: %d", len(urls))
	}
	if !strings.HasPrefix(urls[0].String(), "https://s3-external-1.amazonaws.com/media.twiliocdn.com/"+sid) {
		t.Errorf("wrong url: %s", urls[0].String())
	}
}

func TestDecode(t *testing.T) {
	t.Parallel()
	msg := new(Message)
	if err := json.Unmarshal(getMessageResponse, &msg); err != nil {
		t.Fatal(err)
	}
	if msg.Sid != "SM26b3b00f8def53be77c5697183bfe95e" {
		t.Errorf("wrong sid")
	}
	got := msg.DateCreated.Time.Format(time.RFC3339)
	want := "2016-09-20T22:59:57Z"
	if got != want {
		t.Errorf("msg.DateCreated: got %s, want %s", got, want)
	}
	if msg.Direction != DirectionOutboundReply {
		t.Errorf("wrong direction")
	}
	if msg.Status != StatusDelivered {
		t.Errorf("wrong status")
	}
	if msg.Body != "Welcome to ZomboCom." {
		t.Errorf("wrong body")
	}
	if msg.From != PhoneNumber("+19253920364") {
		t.Errorf("wrong from")
	}
	if msg.FriendlyPrice() != "$0.0075" {
		t.Errorf("wrong friendly price %v, want %v", msg.FriendlyPrice(), "$0.00750")
	}
}

func TestStatusFriendly(t *testing.T) {
	t.Parallel()
	if StatusQueued.Friendly() != "Queued" {
		t.Errorf("expected StatusQueued.Friendly to equal Queued, got %s", StatusQueued.Friendly())
	}
	s := Status("in-progress")
	if f := s.Friendly(); f != "In Progress" {
		t.Errorf("expected In Progress.Friendly to equal In Progress, got %s", f)
	}
}

var whatsappMessage = []byte(`
{
    "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "api_version": "2010-04-01",
    "body": "Testing whatsapp integration! \ud83d\ude0e",
    "date_created": "Sat, 04 Aug 2018 03:35:27 +0000",
    "date_sent": null,
    "date_updated": "Sat, 04 Aug 2018 03:35:27 +0000",
    "direction": "outbound-api",
    "error_code": null,
    "error_message": null,
    "from": "whatsapp:+14155238886",
    "messaging_service_sid": null,
    "num_media": "0",
    "num_segments": "1",
    "price": null,
    "price_unit": null,
    "sid": "SM75347b88e19f41fc8a83db8aa32e37ea",
    "status": "queued",
    "subresource_uris": {
        "media": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Messages/SM75347b88e19f41fc8a83db8aa32e37ea/Media.json"
    },
    "to": "whatsapp:+19253245555",
    "uri": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Messages/SM75347b88e19f41fc8a83db8aa32e37ea.json"
}
`)

func TestWhatsappMessageParsing(t *testing.T) {
	t.Parallel()
	m := new(Message)
	if err := json.Unmarshal(whatsappMessage, m); err != nil {
		t.Fatal(err)
	}
	if m.To.Local() != "(925) 324-5555" {
		t.Errorf("bad Local: %v", m.To.Local())
	}
}
