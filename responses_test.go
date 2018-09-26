package twilio

import (
	"flag"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"sync"
	"testing"
)

// the envClient is configured to use an Account Sid and Auth Token set in the
// environment. all non-short tests should use the envClient
var envClient *Client

func init() {
	flag.Parse()
	if !testing.Short() && os.Getenv("TWILIO_ACCOUNT_SID") == "" {
		os.Stderr.WriteString("warning: no TWILIO_ACCOUNT_SID configured, HTTP tests will probably fail...\n\n")
	}
	envClient = NewClient(os.Getenv("TWILIO_ACCOUNT_SID"), os.Getenv("TWILIO_AUTH_TOKEN"), nil)
}

type Server struct {
	s *httptest.Server
	// copied from httptest.Server
	URL string
	// URLs of incoming requests, in order
	URLs []*url.URL
	mu   sync.Mutex
}

func (s *Server) Close() {
	s.s.Close()
}

func (s *Server) CloseClientConnections() {
	s.s.CloseClientConnections()
}

func (s *Server) Start() {
	s.s.Start()
}

func newServer(response []byte, code int) *Server {
	serv := &Server{URLs: make([]*url.URL, 0)}
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serv.mu.Lock()
		serv.URLs = append(serv.URLs, r.URL)
		serv.mu.Unlock()
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(code)
		if _, err := w.Write(response); err != nil {
			panic(err)
		}
	}))
	serv.s = s
	serv.URL = s.URL
	return serv
}

// getServer returns a http server that returns the given bytes when requested,
// and a Client configured to make requests to that server.
func getServer(response []byte) (*Client, *Server) {
	s := newServer(response, 200)
	client := NewClient("AC123", "456", nil)
	client.Base = s.URL
	client.Monitor.Base = s.URL
	client.Pricing.Base = s.URL
	client.Fax.Base = s.URL
	client.Wireless.Base = s.URL
	client.Notify.Base = s.URL
	client.Lookup.Base = s.URL
	client.Video.Base = s.URL
	return client, s
}

func getServerCode(response []byte, code int) (*Client, *Server) {
	s := newServer(response, code)
	client := NewClient("AC123", "456", nil)
	client.Base = s.URL
	client.Fax.Base = s.URL
	client.Monitor.Base = s.URL
	client.Pricing.Base = s.URL
	client.Wireless.Base = s.URL
	client.Notify.Base = s.URL
	client.Lookup.Base = s.URL
	client.Video.Base = s.URL
	return client, s
}

// useful trick: highlight the JSON range and hit `python -m json.tool` to
// pretty format it.

var conferenceInstance = []byte(`
{
    "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "api_version": "2010-04-01",
    "date_created": "Fri, 23 Aug 2013 21:52:27 +0000",
    "date_updated": "Fri, 23 Aug 2013 21:52:34 +0000",
    "friendly_name": "testConference",
    "region": "us1",
    "sid": "CF169b5eebb07ec48e0f9f2ee904b385c5",
    "status": "completed",
    "subresource_uris": {
        "participants": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Conferences/CF169b5eebb07ec48e0f9f2ee904b385c5/Participants.json"
    },
    "uri": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Conferences/CF169b5eebb07ec48e0f9f2ee904b385c5.json"
}
`)
var conferenceInstanceSid = "CF169b5eebb07ec48e0f9f2ee904b385c5"

var conferencePage = []byte(`
{
    "conferences": [
        {
            "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
            "api_version": "2010-04-01",
            "date_created": "Fri, 13 Sep 2013 19:53:14 +0000",
            "date_updated": "Fri, 13 Sep 2013 19:53:34 +0000",
            "friendly_name": "turbovote",
            "region": "us1",
            "sid": "CF347aef00d0b0ba10eec6a77fabdd1c95",
            "status": "completed",
            "subresource_uris": {
                "participants": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Conferences/CF347aef00d0b0ba10eec6a77fabdd1c95/Participants.json"
            },
            "uri": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Conferences/CF347aef00d0b0ba10eec6a77fabdd1c95.json"
        },
        {
            "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
            "api_version": "2010-04-01",
            "date_created": "Fri, 23 Aug 2013 21:52:27 +0000",
            "date_updated": "Fri, 23 Aug 2013 21:52:34 +0000",
            "friendly_name": "testConference",
            "region": "us1",
            "sid": "CF169b5eebb07ec48e0f9f2ee904b385c5",
            "status": "completed",
            "subresource_uris": {
                "participants": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Conferences/CF169b5eebb07ec48e0f9f2ee904b385c5/Participants.json"
            },
            "uri": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Conferences/CF169b5eebb07ec48e0f9f2ee904b385c5.json"
        },
        {
            "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
            "api_version": "2010-04-01",
            "date_created": "Fri, 23 Aug 2013 21:48:44 +0000",
            "date_updated": "Fri, 23 Aug 2013 21:52:01 +0000",
            "friendly_name": "testConference",
            "region": "us1",
            "sid": "CFb2b77e00f9e97764746aff575a15bfbb",
            "status": "completed",
            "subresource_uris": {
                "participants": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Conferences/CFb2b77e00f9e97764746aff575a15bfbb/Participants.json"
            },
            "uri": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Conferences/CFb2b77e00f9e97764746aff575a15bfbb.json"
        }
    ],
    "end": 2,
    "first_page_uri": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Conferences.json?PageSize=3&Page=0",
    "next_page_uri": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Conferences.json?PageSize=3&Page=1&AfterSid=CFb2b77e00f9e97764746aff575a15bfbb",
    "page": 0,
    "page_size": 3,
    "previous_page_uri": null,
    "start": 0,
    "uri": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Conferences.json?PageSize=3"
}
`)

var credentialsPage = []byte(`
{
  "credentials": [
    {
      "sid": "CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
      "friendly_name": "MyCredential",
      "account_sid": "AC6bc21af903cc765a9d7f7e0467ec812a",
      "type": "fcm",
      "date_created": "2015-08-26T00:07:12Z",
      "date_updated": "2015-08-26T00:07:12Z",
      "url": "https://notify.twilio.com/v1/Credentials/CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
    },
    {
      "sid": "CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
      "account_sid": "AC6bc21af903cc765a9d7f7e0467ec812a",
      "friendly_name": "MyCredential2",
      "type": "apn",
      "date_created": "2015-08-26T00:07:43Z",
      "date_updated": "2015-08-26T00:07:43Z",
      "url": "https://notify.twilio.com/v1/Credentials/CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
    }
  ],
  "meta": {
    "page": 0,
    "page_size": 50,
    "first_page_url": "https://notify.twilio.com/v1/Credentials?PageSize=50&Page=0",
    "previous_page_url": null,
    "url": "https://notify.twilio.com/v1/Credentials?PageSize=50&Page=0",
    "next_page_url": null,
    "key": "credentials"
  }
}
`)

var notifyCredential = []byte(`
{
  "sid": "CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "account_sid": "AC6bc21af903cc765a9d7f7e0467ec812a",
  "friendly_name": "MyFCMCredential",
  "type": "fcm",
  "sandbox": null,
  "date_created": "2015-08-26T00:07:43Z",
  "date_updated": "2015-08-26T00:07:43Z",
  "url": "https://notify.twilio.com/v1/Credentials/CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}
`)

var makeCallResponse = []byte(`
{
    "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "annotation": null,
    "answered_by": null,
    "api_version": "2010-04-01",
    "caller_name": null,
    "date_created": null,
    "date_updated": null,
    "direction": "outbound-api",
    "duration": null,
    "end_time": null,
    "forwarded_from": null,
    "from": "+19253920364",
    "from_formatted": "(925) 392-0364",
    "group_sid": null,
    "parent_call_sid": null,
    "phone_number_sid": "PN5fb9ed903e184c8baa86c1fb7544ca0f",
    "price": null,
    "price_unit": "USD",
    "sid": "CA47b862ce3b99a6d79939320a9aa54a02",
    "start_time": null,
    "status": "queued",
    "subresource_uris": {
        "notifications": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Calls/CA47b862ce3b99a6d79939320a9aa54a02/Notifications.json",
        "recordings": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Calls/CA47b862ce3b99a6d79939320a9aa54a02/Recordings.json"
    },
    "to": "+19252717005",
    "to_formatted": "(925) 271-7005",
    "uri": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Calls/CA47b862ce3b99a6d79939320a9aa54a02.json"
}
`)

var sendMessageResponse = []byte(`
{
    "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "api_version": "2010-04-01",
    "body": "twilio-go testing!",
    "date_created": "Wed, 26 Oct 2016 18:12:20 +0000",
    "date_sent": null,
    "date_updated": "Wed, 26 Oct 2016 18:12:20 +0000",
    "direction": "outbound-api",
    "error_code": null,
    "error_message": null,
    "from": "+19253920364",
    "messaging_service_sid": null,
    "num_media": "0",
    "num_segments": "1",
    "price": null,
    "price_unit": "USD",
    "sid": "SM9b7db369463c439384fe9abb09751af8",
    "status": "queued",
    "subresource_uris": {
        "media": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Messages/SM9b7db369463c439384fe9abb09751af8/Media.json"
    },
    "to": "+19253920364",
    "uri": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Messages/SM9b7db369463c439384fe9abb09751af8.json"
}
`)

var getMessageResponse = []byte(`
{
    "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "api_version": "2010-04-01",
    "body": "Welcome to ZomboCom.",
    "date_created": "Tue, 20 Sep 2016 22:59:57 +0000",
    "date_sent": "Tue, 20 Sep 2016 22:59:57 +0000",
    "date_updated": "Tue, 20 Sep 2016 22:59:57 +0000",
    "direction": "outbound-reply",
    "error_code": null,
    "error_message": null,
    "from": "+19253920364",
    "messaging_service_sid": null,
    "num_media": "0",
    "num_segments": "1",
    "price": "-0.00750",
    "price_unit": "USD",
    "sid": "SM26b3b00f8def53be77c5697183bfe95e",
    "status": "delivered",
    "subresource_uris": {
        "media": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Messages/SM26b3b00f8def53be77c5697183bfe95e/Media.json"
    },
    "to": "+13365584092",
    "uri": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Messages/SM26b3b00f8def53be77c5697183bfe95e.json"
}
`)

var alertInstanceResponse = []byte(`
{
    "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "alert_text": "Msg&sourceComponent=12000&ErrorCode=11200&httpResponse=405&url=https%3A%2F%2Fkev.inburke.com%2Fzombo%2Fzombocom.mp3&LogLevel=ERROR",
    "api_version": "2010-04-01",
    "date_created": "2016-10-27T02:34:21Z",
    "date_generated": "2016-10-27T02:34:21Z",
    "date_updated": "2016-10-27T02:34:23Z",
    "error_code": "11200",
    "log_level": "error",
    "more_info": "https://www.twilio.com/docs/errors/11200",
    "request_headers": null,
    "request_method": "POST",
    "request_url": "https://kev.inburke.com/zombo/zombocom.mp3",
    "request_variables": "Called=%2B19252717005&ToState=CA&CallerCountry=US&Direction=outbound-api&CallerState=CA&ToZip=94596&CallSid=CA6d27370cbbfb605521fe8800bb73f2d2&To=%2B19252717005&CallerZip=94514&ToCountry=US&ApiVersion=2010-04-01&CalledZip=94596&CalledCity=PLEASANTON&CallStatus=in-progress&From=%2B19253920364&AccountSid=AC58f1e8f2b1c6b88ca90a012a4be0c279&CalledCountry=US&CallerCity=BRENTWOOD&Caller=%2B19253920364&FromCountry=US&ToCity=PLEASANTON&FromCity=BRENTWOOD&CalledState=CA&FromZip=94514&FromState=CA",
    "resource_sid": "CA6d27370cbbfb605521fe8800bb73f2d2",
    "response_body": "<html>\r\n<head><title>405 Not Allowed</title></head>\r\n<body bgcolor=\"white\">\r\n<center><h1>405 Not Allowed</h1></center>\r\n<hr><center>nginx</center>\r\n</body>\r\n</html>",
    "response_headers": "Transfer-Encoding=chunked&Server=cloudflare-nginx&CF-RAY=2f82bf9cb8102204-EWR&Set-Cookie=__cfduid%3Dd46f1cfd57d664c3038ae66f1c1de9e751477535661%3B+expires%3DFri%2C+27-Oct-17+02%3A34%3A21+GMT%3B+path%3D%2F%3B+domain%3D.inburke.com%3B+HttpOnly&Date=Thu%2C+27+Oct+2016+02%3A34%3A21+GMT&Content-Type=text%2Fhtml",
    "service_sid": null,
    "sid": "NO00ed1fb4aa449be2434d54ec8e492349",
    "url": "https://monitor.twilio.com/v1/Alerts/NO00ed1fb4aa449be2434d54ec8e492349"
}
`)

var alertListResponse = []byte(`
{
    "alerts": [
        {
            "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
            "alert_text": "Msg&sourceComponent=12000&ErrorCode=11200&httpResponse=405&url=https%3A%2F%2Fkev.inburke.com%2Fzombo%2Fzombocom.mp3&LogLevel=ERROR",
            "api_version": "2010-04-01",
            "date_created": "2016-10-27T02:34:21Z",
            "date_generated": "2016-10-27T02:34:21Z",
            "date_updated": "2016-10-27T02:34:23Z",
            "error_code": "11200",
            "log_level": "error",
            "more_info": "https://www.twilio.com/docs/errors/11200",
            "request_method": "POST",
            "request_url": "https://kev.inburke.com/zombo/zombocom.mp3",
            "resource_sid": "CA6d27370cbbfb605521fe8800bb73f2d2",
            "service_sid": null,
            "sid": "NO00ed1fb4aa449be2434d54ec8e492349",
            "url": "https://monitor.twilio.com/v1/Alerts/NO00ed1fb4aa449be2434d54ec8e492349"
        },
        {
            "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
            "alert_text": "sourceComponent=14100&ErrorCode=14101&LogLevel=ERROR&Msg=The+destination+number+for+a+TwiML+message+can+not+be+the+same+as+the+originating+number+of+an+incoming+message.&EmailNotification=false",
            "api_version": "2008-08-01",
            "date_created": "2016-10-26T18:12:20Z",
            "date_generated": "2016-10-26T18:12:20Z",
            "date_updated": "2016-10-26T18:12:24Z",
            "error_code": "14101",
            "log_level": "error",
            "more_info": "https://www.twilio.com/docs/errors/14101",
            "request_method": "POST",
            "request_url": "https://kev.inburke.com/zombo/zombo.php",
            "resource_sid": "SM07bb932105b6c969574be9ca5771d139",
            "service_sid": null,
            "sid": "NOffde6702f5035789b395f8618f0aa65a",
            "url": "https://monitor.twilio.com/v1/Alerts/NOffde6702f5035789b395f8618f0aa65a"
        }
    ],
    "meta": {
        "first_page_url": "https://monitor.twilio.com/v1/Alerts?PageSize=2&Page=0",
        "key": "alerts",
        "next_page_url": "https://monitor.twilio.com/v1/Alerts?PageSize=2&Page=1&PageToken=PANOffde6702f5035789b395f8618f0aa65a",
        "page": 0,
        "page_size": 2,
        "previous_page_url": null,
        "url": "https://monitor.twilio.com/v1/Alerts?PageSize=2&Page=0"
    }
}
`)

var transcriptionDeletedTwice = []byte(`
{
    "code": 20404,
    "message": "The requested resource /2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Transcriptions/TR4c7f9a71f19b7509cb1e8344af78fc82.json was not found",
    "more_info": "https://www.twilio.com/docs/errors/20404",
    "status": 404
}
`)

var applicationInstance = []byte(`
{
    "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "api_version": "2010-04-01",
    "date_created": "Sat, 01 Oct 2011 07:19:51 +0000",
    "date_updated": "Thu, 11 Jul 2013 05:06:50 +0000",
    "friendly_name": "Hackpack for Heroku and Flask",
    "message_status_callback": "",
    "sid": "AP7d6fd7b9a8894e36877dc2355da381c8",
    "sms_fallback_method": "POST",
    "sms_fallback_url": "",
    "sms_method": "POST",
    "sms_status_callback": "",
    "sms_url": "http://twilio-amaze-client.herokuapp.com/sms",
    "status_callback": "",
    "status_callback_method": "POST",
    "uri": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Applications/AP7d6fd7b9a8894e36877dc2355da381c8.json",
    "voice_caller_id_lookup": false,
    "voice_fallback_method": "POST",
    "voice_fallback_url": "",
    "voice_method": "POST",
    "voice_url": "http://twilio-amaze-client.herokuapp.com/client/incoming"
}
`)

var callerIDInstance = []byte(`
{
    "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "date_created": "Sat, 01 Feb 2014 05:30:57 +0000",
    "date_updated": "Sat, 01 Feb 2014 05:30:57 +0000",
    "friendly_name": "foo",
    "phone_number": "+19252717005",
    "sid": "PNca86cf94c7d4f89e0bd45bfa7d9b9e7d",
    "uri": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/OutgoingCallerIds/PNca86cf94c7d4f89e0bd45bfa7d9b9e7d.json"
}
`)

var callerIDVerify = []byte(`
{
    "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "call_sid": "CA6662a69ccb58ef2e162098861f1892b5",
    "friendly_name": "test friendly name",
    "phone_number": "+14105551234",
    "validation_code": "531628"
}
`)

var accountInstance = []byte(`
{
    "auth_token": "[redacted]",
    "date_created": "Fri, 18 Feb 2011 00:51:02 +0000",
    "date_updated": "Mon, 12 Sep 2016 22:17:12 +0000",
    "friendly_name": "kevin account woo",
    "owner_account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "status": "active",
    "subresource_uris": {
        "applications": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Applications.json",
        "authorized_connect_apps": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/AuthorizedConnectApps.json",
        "available_phone_numbers": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/AvailablePhoneNumbers.json",
        "calls": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Calls.json",
        "conferences": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Conferences.json",
        "connect_apps": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/ConnectApps.json",
        "incoming_phone_numbers": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/IncomingPhoneNumbers.json",
        "media": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Media.json",
        "messages": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Messages.json",
        "notifications": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Notifications.json",
        "outgoing_caller_ids": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/OutgoingCallerIds.json",
        "queues": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Queues.json",
        "recordings": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Recordings.json",
        "sandbox": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Sandbox.json",
        "sip": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/SIP.json",
        "sms_messages": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/SMS/Messages.json",
        "transcriptions": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Transcriptions.json",
        "usage": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/Usage.json"
    },
    "type": "Full",
    "uri": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279.json"
}
`)

var accountList = []byte(`
{
    "accounts": [
        {
            "auth_token": "[redacted]",
            "date_created": "Fri, 23 Aug 2013 21:46:14 +0000",
            "date_updated": "Mon, 12 Sep 2016 22:18:33 +0000",
            "friendly_name": "TestAccountUno",
            "owner_account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
            "sid": "AC0cd9be8fd5e6e4fa0a04f50ac1caca4e",
            "status": "active",
            "subresource_uris": {
                "applications": "/2010-04-01/Accounts/AC0cd9be8fd5e6e4fa0a04f50ac1caca4e/Applications.json",
                "authorized_connect_apps": "/2010-04-01/Accounts/AC0cd9be8fd5e6e4fa0a04f50ac1caca4e/AuthorizedConnectApps.json",
                "available_phone_numbers": "/2010-04-01/Accounts/AC0cd9be8fd5e6e4fa0a04f50ac1caca4e/AvailablePhoneNumbers.json",
                "calls": "/2010-04-01/Accounts/AC0cd9be8fd5e6e4fa0a04f50ac1caca4e/Calls.json",
                "conferences": "/2010-04-01/Accounts/AC0cd9be8fd5e6e4fa0a04f50ac1caca4e/Conferences.json",
                "connect_apps": "/2010-04-01/Accounts/AC0cd9be8fd5e6e4fa0a04f50ac1caca4e/ConnectApps.json",
                "incoming_phone_numbers": "/2010-04-01/Accounts/AC0cd9be8fd5e6e4fa0a04f50ac1caca4e/IncomingPhoneNumbers.json",
                "media": "/2010-04-01/Accounts/AC0cd9be8fd5e6e4fa0a04f50ac1caca4e/Media.json",
                "messages": "/2010-04-01/Accounts/AC0cd9be8fd5e6e4fa0a04f50ac1caca4e/Messages.json",
                "notifications": "/2010-04-01/Accounts/AC0cd9be8fd5e6e4fa0a04f50ac1caca4e/Notifications.json",
                "outgoing_caller_ids": "/2010-04-01/Accounts/AC0cd9be8fd5e6e4fa0a04f50ac1caca4e/OutgoingCallerIds.json",
                "queues": "/2010-04-01/Accounts/AC0cd9be8fd5e6e4fa0a04f50ac1caca4e/Queues.json",
                "recordings": "/2010-04-01/Accounts/AC0cd9be8fd5e6e4fa0a04f50ac1caca4e/Recordings.json",
                "sandbox": "/2010-04-01/Accounts/AC0cd9be8fd5e6e4fa0a04f50ac1caca4e/Sandbox.json",
                "sip": "/2010-04-01/Accounts/AC0cd9be8fd5e6e4fa0a04f50ac1caca4e/SIP.json",
                "sms_messages": "/2010-04-01/Accounts/AC0cd9be8fd5e6e4fa0a04f50ac1caca4e/SMS/Messages.json",
                "transcriptions": "/2010-04-01/Accounts/AC0cd9be8fd5e6e4fa0a04f50ac1caca4e/Transcriptions.json",
                "usage": "/2010-04-01/Accounts/AC0cd9be8fd5e6e4fa0a04f50ac1caca4e/Usage.json"
            },
            "type": "Full",
            "uri": "/2010-04-01/Accounts/AC0cd9be8fd5e6e4fa0a04f50ac1caca4e.json"
        },
        {
            "auth_token": "[redacted]",
            "date_created": "Fri, 23 Aug 2013 21:47:12 +0000",
            "date_updated": "Mon, 12 Sep 2016 22:18:33 +0000",
            "friendly_name": "TestAccountUno",
            "owner_account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
            "sid": "ACdd54a711c3d4031ac500c5236ab121d7",
            "status": "active",
            "subresource_uris": {
                "applications": "/2010-04-01/Accounts/ACdd54a711c3d4031ac500c5236ab121d7/Applications.json",
                "authorized_connect_apps": "/2010-04-01/Accounts/ACdd54a711c3d4031ac500c5236ab121d7/AuthorizedConnectApps.json",
                "available_phone_numbers": "/2010-04-01/Accounts/ACdd54a711c3d4031ac500c5236ab121d7/AvailablePhoneNumbers.json",
                "calls": "/2010-04-01/Accounts/ACdd54a711c3d4031ac500c5236ab121d7/Calls.json",
                "conferences": "/2010-04-01/Accounts/ACdd54a711c3d4031ac500c5236ab121d7/Conferences.json",
                "connect_apps": "/2010-04-01/Accounts/ACdd54a711c3d4031ac500c5236ab121d7/ConnectApps.json",
                "incoming_phone_numbers": "/2010-04-01/Accounts/ACdd54a711c3d4031ac500c5236ab121d7/IncomingPhoneNumbers.json",
                "media": "/2010-04-01/Accounts/ACdd54a711c3d4031ac500c5236ab121d7/Media.json",
                "messages": "/2010-04-01/Accounts/ACdd54a711c3d4031ac500c5236ab121d7/Messages.json",
                "notifications": "/2010-04-01/Accounts/ACdd54a711c3d4031ac500c5236ab121d7/Notifications.json",
                "outgoing_caller_ids": "/2010-04-01/Accounts/ACdd54a711c3d4031ac500c5236ab121d7/OutgoingCallerIds.json",
                "queues": "/2010-04-01/Accounts/ACdd54a711c3d4031ac500c5236ab121d7/Queues.json",
                "recordings": "/2010-04-01/Accounts/ACdd54a711c3d4031ac500c5236ab121d7/Recordings.json",
                "sandbox": "/2010-04-01/Accounts/ACdd54a711c3d4031ac500c5236ab121d7/Sandbox.json",
                "sip": "/2010-04-01/Accounts/ACdd54a711c3d4031ac500c5236ab121d7/SIP.json",
                "sms_messages": "/2010-04-01/Accounts/ACdd54a711c3d4031ac500c5236ab121d7/SMS/Messages.json",
                "transcriptions": "/2010-04-01/Accounts/ACdd54a711c3d4031ac500c5236ab121d7/Transcriptions.json",
                "usage": "/2010-04-01/Accounts/ACdd54a711c3d4031ac500c5236ab121d7/Usage.json"
            },
            "type": "Full",
            "uri": "/2010-04-01/Accounts/ACdd54a711c3d4031ac500c5236ab121d7.json"
        }
    ],
    "end": 1,
    "first_page_uri": "/2010-04-01/Accounts.json?PageSize=2&Page=0",
    "next_page_uri": "/2010-04-01/Accounts.json?PageSize=2&Page=1&AfterSid=ACdd54a711c3d4031ac500c5236ab121d7",
    "page": 0,
    "page_size": 2,
    "previous_page_uri": null,
    "start": 0,
    "uri": "/2010-04-01/Accounts.json?PageSize=2"
}
`)

var accountCreateResponse = []byte(`
{
    "auth_token": "[redacted]",
    "date_created": "Wed, 02 Nov 2016 16:44:41 +0000",
    "date_updated": "Wed, 02 Nov 2016 16:44:42 +0000",
    "friendly_name": "new account name 1478105087",
    "owner_account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "sid": "ACde8301520edc3b9171b8a68420d6e149",
    "status": "active",
    "subresource_uris": {
        "applications": "/2010-04-01/Accounts/ACde8301520edc3b9171b8a68420d6e149/Applications.json",
        "authorized_connect_apps": "/2010-04-01/Accounts/ACde8301520edc3b9171b8a68420d6e149/AuthorizedConnectApps.json",
        "available_phone_numbers": "/2010-04-01/Accounts/ACde8301520edc3b9171b8a68420d6e149/AvailablePhoneNumbers.json",
        "calls": "/2010-04-01/Accounts/ACde8301520edc3b9171b8a68420d6e149/Calls.json",
        "conferences": "/2010-04-01/Accounts/ACde8301520edc3b9171b8a68420d6e149/Conferences.json",
        "connect_apps": "/2010-04-01/Accounts/ACde8301520edc3b9171b8a68420d6e149/ConnectApps.json",
        "incoming_phone_numbers": "/2010-04-01/Accounts/ACde8301520edc3b9171b8a68420d6e149/IncomingPhoneNumbers.json",
        "media": "/2010-04-01/Accounts/ACde8301520edc3b9171b8a68420d6e149/Media.json",
        "messages": "/2010-04-01/Accounts/ACde8301520edc3b9171b8a68420d6e149/Messages.json",
        "notifications": "/2010-04-01/Accounts/ACde8301520edc3b9171b8a68420d6e149/Notifications.json",
        "outgoing_caller_ids": "/2010-04-01/Accounts/ACde8301520edc3b9171b8a68420d6e149/OutgoingCallerIds.json",
        "queues": "/2010-04-01/Accounts/ACde8301520edc3b9171b8a68420d6e149/Queues.json",
        "recordings": "/2010-04-01/Accounts/ACde8301520edc3b9171b8a68420d6e149/Recordings.json",
        "sandbox": "/2010-04-01/Accounts/ACde8301520edc3b9171b8a68420d6e149/Sandbox.json",
        "sip": "/2010-04-01/Accounts/ACde8301520edc3b9171b8a68420d6e149/SIP.json",
        "sms_messages": "/2010-04-01/Accounts/ACde8301520edc3b9171b8a68420d6e149/SMS/Messages.json",
        "transcriptions": "/2010-04-01/Accounts/ACde8301520edc3b9171b8a68420d6e149/Transcriptions.json",
        "usage": "/2010-04-01/Accounts/ACde8301520edc3b9171b8a68420d6e149/Usage.json"
    },
    "type": "Full",
    "uri": "/2010-04-01/Accounts/ACde8301520edc3b9171b8a68420d6e149.json"
}
`)

var voicePriceUS = []byte(`{
    "country": "United States",
    "iso_country": "US",
    "outbound_prefix_prices": [
        {
            "prefixes": [
                "1907"
            ],
            "friendly_name": "Programmable Outbound Minute - United States - Alaska",
            "base_price": "0.090",
            "current_price": "0.090"
        },
        {
            "prefixes": [
                "1808"
            ],
            "friendly_name": "Programmable Outbound Minute - United States - Hawaii",
            "base_price": "0.015",
            "current_price": "0.015"
        },
        {
            "prefixes": [
                "1800",
                "1844",
                "1855",
                "1866",
                "1877",
                "1888"
            ],
            "friendly_name": "Programmable Outbound Minute - United States - Toll Free",
            "base_price": "0.015",
            "current_price": "0.015"
        },
        {
            "prefixes": [
                "1"
            ],
            "friendly_name": "Programmable Outbound Minute - United States & Canada",
            "base_price": "0.015",
            "current_price": "0.015"
        }
    ],
    "inbound_call_prices": [
        {
            "number_type": "local",
            "base_price": "0.0075",
            "current_price": "0.0075"
        },
        {
            "number_type": "toll free",
            "base_price": "0.0275",
            "current_price": "0.0275"
        }
    ],
    "price_unit": "USD",
    "url": "https://pricing.twilio.com/v1/Voice/Countries/US"
}`)

var voicePricesGB = []byte(`{
    "country": "United Kingdom",
    "iso_country": "GB",
    "outbound_prefix_prices": [
        {
            "prefixes": [
                "44",
                "44203",
                "44207",
                "44208"
            ],
            "friendly_name": "Programmable Outbound Minute - United Kingdom",
            "base_price": "0.0175",
            "current_price": "0.0175"
        },
        {
            "prefixes": [
                "447",
                "4470",
                "4474408",
                "44870",
                "44871",
                "44872",
                "44873"
            ],
            "friendly_name": "Programmable Outbound Minute - United Kingdom - Other",
            "base_price": "0.32",
            "current_price": "0.32"
        },
        {
            "prefixes": [
                "443",
                "445",
                "44551107"
            ],
            "friendly_name": "Programmable Outbound Minute - United Kingdom - Other - Service ",
            "base_price": "0.045",
            "current_price": "0.045"
        },
        {
            "prefixes": [
                "44843",
                "44844",
                "44845"
            ],
            "friendly_name": "Programmable Outbound Minute - United Kingdom - Other - Local",
            "base_price": "0.25",
            "current_price": "0.25"
        },
        {
            "prefixes": [
                "447106",
                "447107",
                "447340",
                "447341",
                "447342",
                "4473970",
                "4473971",
                "4473972",
                "4473973",
                "4473975",
                "4473976",
                "4473977",
                "4473978",
                "4473979",
                "447398",
                "447399",
                "447400",
                "447401",
                "447402",
                "447403",
                "447407",
                "447409",
                "447410",
                "447411",
                "447412",
                "447413",
                "447414",
                "447415",
                "447416",
                "4474170",
                "4474180",
                "447419",
                "447420",
                "447421",
                "447422",
                "447423",
                "447425",
                "447426",
                "447427",
                "447428",
                "447429",
                "447430",
                "447431",
                "447432",
                "447433",
                "447434",
                "447435",
                "447436",
                "447437",
                "447442",
                "447443",
                "447444",
                "447445",
                "447446",
                "447447",
                "447449",
                "447450",
                "4474527",
                "4474528",
                "4474529",
                "447453",
                "447454",
                "447455",
                "447456",
                "4474586",
                "4474589",
                "447460",
                "447461",
                "447462",
                "447463",
                "447464",
                "4474652",
                "4474654",
                "4474656",
                "4474657",
                "4474658",
                "4474659",
                "447467",
                "447468",
                "447469",
                "447470",
                "447471",
                "447472",
                "447473",
                "447474",
                "447475",
                "447476",
                "447477",
                "447478",
                "447479",
                "447480",
                "447481",
                "447482",
                "447483",
                "447484",
                "447485",
                "447486",
                "447487",
                "4474884",
                "4474885",
                "4474887",
                "4474889",
                "447489",
                "447490",
                "447491",
                "447492",
                "447493",
                "447494",
                "447495",
                "447496",
                "447497",
                "447498",
                "447499",
                "447500",
                "447501",
                "447502",
                "447503",
                "447504",
                "447505",
                "447506",
                "447507",
                "447508",
                "44751",
                "447521",
                "447522",
                "447523",
                "447525",
                "447526",
                "447527",
                "447528",
                "447529",
                "447530",
                "447531",
                "4475320",
                "4475321",
                "4475322",
                "4475323",
                "4475324",
                "4475326",
                "4475327",
                "4475328",
                "447533",
                "447534",
                "447535",
                "447536",
                "4475374",
                "4475378",
                "4475379",
                "447538",
                "447539",
                "44754",
                "447550",
                "447551",
                "447552",
                "447553",
                "447554",
                "447555",
                "447556",
                "447557",
                "44756",
                "447570",
                "447572",
                "447573",
                "447574",
                "447575",
                "447576",
                "447577",
                "447578",
                "447579",
                "447580",
                "447581",
                "447582",
                "447583",
                "447584",
                "447585",
                "447586",
                "447587",
                "447588",
                "44759",
                "447701",
                "447702",
                "447703",
                "447704",
                "447705",
                "447706",
                "447707",
                "447708",
                "447709",
                "447710",
                "447711",
                "447712",
                "447713",
                "447714",
                "447715",
                "447716",
                "447717",
                "447718",
                "447719",
                "447720",
                "447721",
                "447722",
                "447723",
                "447724",
                "447725",
                "447726",
                "447727",
                "447728",
                "447729",
                "447730",
                "447731",
                "447732",
                "447733",
                "447734",
                "447735",
                "447736",
                "447737",
                "447738",
                "447739",
                "447740",
                "447741",
                "447742",
                "447743",
                "447745",
                "447746",
                "447747",
                "447748",
                "447749",
                "447750",
                "447751",
                "447752",
                "4477531",
                "4477532",
                "4477533",
                "4477534",
                "4477535",
                "4477536",
                "4477537",
                "4477538",
                "4477539",
                "447754",
                "447756",
                "447757",
                "447758",
                "447759",
                "447760",
                "447761",
                "447762",
                "447763",
                "447764",
                "447765",
                "447766",
                "447767",
                "447768",
                "447769",
                "447770",
                "447771",
                "447772",
                "447773",
                "447774",
                "447775",
                "447776",
                "447778",
                "447779",
                "447780",
                "447782",
                "447783",
                "447784",
                "447785",
                "447786",
                "447787",
                "447788",
                "447789",
                "447790",
                "447791",
                "447792",
                "447793",
                "447794",
                "447795",
                "447796",
                "447798",
                "447799",
                "447800",
                "447801",
                "447802",
                "447803",
                "447804",
                "447805",
                "447806",
                "447807",
                "447808",
                "447809",
                "447810",
                "447811",
                "447812",
                "447813",
                "447814",
                "447815",
                "447816",
                "447817",
                "447818",
                "447819",
                "447820",
                "447821",
                "447823",
                "447824",
                "447825",
                "447826",
                "447827",
                "447828",
                "4478290",
                "4478291",
                "4478292",
                "4478293",
                "4478294",
                "4478295",
                "4478296",
                "447830",
                "447831",
                "447832",
                "447833",
                "447834",
                "447835",
                "447836",
                "447837",
                "447838",
                "447840",
                "447841",
                "447842",
                "447843",
                "447844",
                "447845",
                "447846",
                "447847",
                "447848",
                "447849",
                "447850",
                "447851",
                "447852",
                "447853",
                "447854",
                "447855",
                "447856",
                "447857",
                "447858",
                "447859",
                "447860",
                "447861",
                "447862",
                "447863",
                "4478640",
                "4478641",
                "4478642",
                "4478643",
                "4478645",
                "4478646",
                "4478647",
                "4478648",
                "4478649",
                "447865",
                "447866",
                "447867",
                "447868",
                "447869",
                "447870",
                "447871",
                "4478720",
                "4478721",
                "4478723",
                "4478724",
                "4478725",
                "4478726",
                "4478728",
                "4478729",
                "4478731",
                "4478732",
                "4478733",
                "4478734",
                "4478735",
                "4478736",
                "4478737",
                "4478738",
                "4478739",
                "4478740",
                "4478741",
                "4478742",
                "4478743",
                "4478746",
                "4478747",
                "4478748",
                "4478749",
                "447875",
                "447876",
                "447877",
                "447878",
                "447879",
                "447880",
                "447881",
                "447882",
                "447883",
                "447884",
                "447885",
                "447886",
                "447887",
                "447888",
                "447889",
                "447890",
                "447891",
                "4478923",
                "4478924",
                "4478926",
                "4478927",
                "4478928",
                "4478929",
                "4478932",
                "4478934",
                "4478935",
                "4478936",
                "4478937",
                "447894",
                "447895",
                "447896",
                "447897",
                "447898",
                "447899",
                "447900",
                "447901",
                "447902",
                "447903",
                "447904",
                "447905",
                "447906",
                "447907",
                "447908",
                "447909",
                "447910",
                "447912",
                "447913",
                "447914",
                "447915",
                "447916",
                "447917",
                "447918",
                "447919",
                "447920",
                "447921",
                "447922",
                "447923",
                "447925",
                "447926",
                "447927",
                "447928",
                "447929",
                "447930",
                "447931",
                "447932",
                "447933",
                "447934",
                "447935",
                "447936",
                "447938",
                "447939",
                "44794",
                "447950",
                "447951",
                "447952",
                "447953",
                "447954",
                "447955",
                "447956",
                "447957",
                "447958",
                "447959",
                "447960",
                "447961",
                "447962",
                "447963",
                "447964",
                "447965",
                "447966",
                "447967",
                "447968",
                "447969",
                "447970",
                "447971",
                "447972",
                "447973",
                "447974",
                "447975",
                "447976",
                "447977",
                "447979",
                "447980",
                "447981",
                "447982",
                "447983",
                "447984",
                "447985",
                "447986",
                "447987",
                "447988",
                "447989",
                "447990",
                "447999"
            ],
            "friendly_name": "Programmable Outbound Minute - United Kingdom - Mobile",
            "base_price": "0.040",
            "current_price": "0.040"
        }
    ],
    "inbound_call_prices": [
        {
            "number_type": "local",
            "base_price": "0.0075",
            "current_price": "0.0075"
        },
        {
            "number_type": "national",
            "base_price": "0.0075",
            "current_price": "0.0075"
        },
        {
            "number_type": "toll free",
            "base_price": "0.0575",
            "current_price": "0.0575"
        }
    ],
    "price_unit": "USD",
    "url": "https://pricing.twilio.com/v1/Voice/Countries/GB"
}`)

var voicePriceNumberUS = []byte(`{
    "number": "+19253920364",
    "country": "United States",
    "iso_country": "US",
    "outbound_call_price": {
        "base_price": "0.015",
        "current_price": "0.015"
    },
    "inbound_call_price": {
        "number_type": null,
        "base_price": null,
        "current_price": null
    },
    "price_unit": "USD",
    "url": "https://pricing.twilio.com/v1/Voice/Numbers/+19253920364"
}`)

var messagePriceGB = []byte(`{
    "url": "https://pricing.twilio.com/v1/Messaging/Countries/GB",
    "country": "United Kingdom",
    "iso_country": "GB",
    "price_unit": "USD",
    "outbound_sms_prices": [
        {
            "mcc": "234",
            "mnc": "55",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "3",
            "carrier": "Vodafone",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "50",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "58",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "18",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "9",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "33",
            "carrier": "Orange",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "30",
            "carrier": "T-Mobile",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "2",
            "carrier": "O2",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "10",
            "carrier": "O2",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "0",
            "carrier": "Vodafone",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "20",
            "carrier": "3",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "7",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "994",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "26",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "1",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "19",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "5",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "6",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "8",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "11",
            "carrier": "O2",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "14",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "15",
            "carrier": "Vodafone",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "16",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "17",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "22",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "24",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "25",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "31",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "32",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "34",
            "carrier": "Orange",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "35",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "36",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "37",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "51",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "76",
            "carrier": "Vodafone",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "78",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "235",
            "mnc": "0",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "235",
            "mnc": "1",
            "carrier": "Everything Everywhere",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "235",
            "mnc": "2",
            "carrier": "Everything Everywhere",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "235",
            "mnc": "77",
            "carrier": "Vodafone",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "235",
            "mnc": "91",
            "carrier": "Vodafone",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "235",
            "mnc": "92",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "235",
            "mnc": "94",
            "carrier": "3",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "999",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "shortcode",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "235",
            "mnc": "999",
            "carrier": "Other",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "27",
            "carrier": "Teleena",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "28",
            "carrier": "Marathon Telecom",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "99",
            "carrier": "Lleida.net",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "39",
            "carrier": "SSE Energy Supply",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "38",
            "carrier": "Virgin",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "23",
            "carrier": "Vectofone Mobile",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        },
        {
            "mcc": "234",
            "mnc": "53",
            "carrier": "Limitless Mobile",
            "prices": [
                {
                    "number_type": "mobile",
                    "base_price": "0.040",
                    "current_price": "0.040"
                },
                {
                    "number_type": "local",
                    "base_price": "0.040",
                    "current_price": "0.040"
                }
            ]
        }
    ],
    "inbound_sms_prices": [
        {
            "number_type": "local",
            "base_price": "0.0075",
            "current_price": "0.0075"
        },
        {
            "number_type": "mobile",
            "base_price": "0.0075",
            "current_price": "0.0075"
        },
        {
            "number_type": "shortcode",
            "base_price": "0.0075",
            "current_price": "0.0075"
        }
    ]
}`)

var phoneNumberPriceGB = []byte(`{
    "country": "United Kingdom",
    "iso_country": "GB",
    "phone_number_prices": [
        {
            "number_type": "local",
            "base_price": "1.00",
            "current_price": "1.00"
        },
        {
            "number_type": "mobile",
            "base_price": "1.00",
            "current_price": "1.00"
        },
        {
            "number_type": "national",
            "base_price": "1.00",
            "current_price": "1.00"
        },
        {
            "number_type": "toll free",
            "base_price": "2.00",
            "current_price": "2.00"
        }
    ],
    "price_unit": "USD",
    "url": "https://pricing.twilio.com/v1/PhoneNumbers/Countries/GB"
}`)

var phoneNumberCountriesPage = []byte(`{
    "meta": {
        "page": 0,
        "page_size": 10,
        "first_page_url": "https://pricing.twilio.com/v1/PhoneNumbers/Countries?PageSize=10&Page=0",
        "previous_page_url": null,
        "url": "https://pricing.twilio.com/v1/PhoneNumbers/Countries?PageSize=10&Page=0",
        "next_page_url": "https://pricing.twilio.com/v1/PhoneNumbers/Countries?PageSize=10&Page=1&PageToken=DNCZ",
        "key": "countries"
    },
    "countries": [
        {
            "country": "Austria",
            "iso_country": "AT",
            "url": "https://pricing.twilio.com/v1/PhoneNumbers/Countries/AT"
        },
        {
            "country": "Australia",
            "iso_country": "AU",
            "url": "https://pricing.twilio.com/v1/PhoneNumbers/Countries/AU"
        },
        {
            "country": "Belgium",
            "iso_country": "BE",
            "url": "https://pricing.twilio.com/v1/PhoneNumbers/Countries/BE"
        },
        {
            "country": "Bulgaria",
            "iso_country": "BG",
            "url": "https://pricing.twilio.com/v1/PhoneNumbers/Countries/BG"
        },
        {
            "country": "Brazil",
            "iso_country": "BR",
            "url": "https://pricing.twilio.com/v1/PhoneNumbers/Countries/BR"
        },
        {
            "country": "Canada",
            "iso_country": "CA",
            "url": "https://pricing.twilio.com/v1/PhoneNumbers/Countries/CA"
        },
        {
            "country": "Switzerland",
            "iso_country": "CH",
            "url": "https://pricing.twilio.com/v1/PhoneNumbers/Countries/CH"
        },
        {
            "country": "Chile",
            "iso_country": "CL",
            "url": "https://pricing.twilio.com/v1/PhoneNumbers/Countries/CL"
        },
        {
            "country": "Cyprus",
            "iso_country": "CY",
            "url": "https://pricing.twilio.com/v1/PhoneNumbers/Countries/CY"
        },
        {
            "country": "Czech Republic",
            "iso_country": "CZ",
            "url": "https://pricing.twilio.com/v1/PhoneNumbers/Countries/CZ"
        }
    ]
}`)

var messagingCountriesPage = []byte(`{
    "meta": {
        "page": 0,
        "page_size": 10,
        "first_page_url": "https://pricing.twilio.com/v1/Messaging/Countries?PageSize=10&Page=0",
        "previous_page_url": null,
        "url": "https://pricing.twilio.com/v1/Messaging/Countries?PageSize=10&Page=0",
        "next_page_url": "https://pricing.twilio.com/v1/Messaging/Countries?PageSize=10&Page=1&PageToken=DNAQ",
        "key": "countries"
    },
    "countries": [
        {
            "country": "Andorra",
            "iso_country": "AD",
            "url": "https://pricing.twilio.com/v1/Messaging/Countries/AD"
        },
        {
            "country": "United Arab Emirates",
            "iso_country": "AE",
            "url": "https://pricing.twilio.com/v1/Messaging/Countries/AE"
        },
        {
            "country": "Afghanistan",
            "iso_country": "AF",
            "url": "https://pricing.twilio.com/v1/Messaging/Countries/AF"
        },
        {
            "country": "Antigua and Barbuda",
            "iso_country": "AG",
            "url": "https://pricing.twilio.com/v1/Messaging/Countries/AG"
        },
        {
            "country": "Anguilla",
            "iso_country": "AI",
            "url": "https://pricing.twilio.com/v1/Messaging/Countries/AI"
        },
        {
            "country": "Albania",
            "iso_country": "AL",
            "url": "https://pricing.twilio.com/v1/Messaging/Countries/AL"
        },
        {
            "country": "Armenia",
            "iso_country": "AM",
            "url": "https://pricing.twilio.com/v1/Messaging/Countries/AM"
        },
        {
            "country": "Netherlands Antilles",
            "iso_country": "AN",
            "url": "https://pricing.twilio.com/v1/Messaging/Countries/AN"
        },
        {
            "country": "Angola",
            "iso_country": "AO",
            "url": "https://pricing.twilio.com/v1/Messaging/Countries/AO"
        },
        {
            "country": "Antarctica",
            "iso_country": "AQ",
            "url": "https://pricing.twilio.com/v1/Messaging/Countries/AQ"
        }
    ]
}`)

var voiceCountriesPage = []byte(`{
    "meta": {
        "page": 0,
        "page_size": 10,
        "first_page_url": "https://pricing.twilio.com/v1/Voice/Countries?PageSize=10&Page=0",
        "previous_page_url": null,
        "url": "https://pricing.twilio.com/v1/Voice/Countries?PageSize=10&Page=0",
        "next_page_url": "https://pricing.twilio.com/v1/Voice/Countries?PageSize=10&Page=1&PageToken=DNAQ",
        "key": "countries"
    },
    "countries": [
        {
            "country": "Andorra",
            "iso_country": "AD",
            "url": "https://pricing.twilio.com/v1/Voice/Countries/AD"
        },
        {
            "country": "United Arab Emirates",
            "iso_country": "AE",
            "url": "https://pricing.twilio.com/v1/Voice/Countries/AE"
        },
        {
            "country": "Afghanistan",
            "iso_country": "AF",
            "url": "https://pricing.twilio.com/v1/Voice/Countries/AF"
        },
        {
            "country": "Antigua and Barbuda",
            "iso_country": "AG",
            "url": "https://pricing.twilio.com/v1/Voice/Countries/AG"
        },
        {
            "country": "Anguilla",
            "iso_country": "AI",
            "url": "https://pricing.twilio.com/v1/Voice/Countries/AI"
        },
        {
            "country": "Albania",
            "iso_country": "AL",
            "url": "https://pricing.twilio.com/v1/Voice/Countries/AL"
        },
        {
            "country": "Armenia",
            "iso_country": "AM",
            "url": "https://pricing.twilio.com/v1/Voice/Countries/AM"
        },
        {
            "country": "Netherlands Antilles",
            "iso_country": "AN",
            "url": "https://pricing.twilio.com/v1/Voice/Countries/AN"
        },
        {
            "country": "Angola",
            "iso_country": "AO",
            "url": "https://pricing.twilio.com/v1/Voice/Countries/AO"
        },
        {
            "country": "Antarctica",
            "iso_country": "AQ",
            "url": "https://pricing.twilio.com/v1/Voice/Countries/AQ"
        }
    ]
}`)

var faxCreateResponse = []byte(`
{
    "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "api_version": "v1",
    "date_created": "2017-03-31T09:44:20Z",
    "date_updated": "2017-03-31T09:44:20Z",
    "direction": "outbound",
    "duration": null,
    "from": "+19252717005",
    "media_url": null,
    "num_pages": null,
    "price": null,
    "price_unit": null,
    "quality": "fine",
    "sid": "FXeb76f282888a074547beba3516552174",
    "status": "queued",
    "to": "+18326327228",
    "url": "https://fax.twilio.com/v1/Faxes/FXeb76f282888a074547beba3516552174"
}
`)

var faxGetResponse = []byte(`
{
    "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "api_version": "v1",
    "date_created": "2017-03-31T09:44:20Z",
    "date_updated": "2017-03-31T09:45:55Z",
    "direction": "outbound",
    "duration": 64,
    "from": "+19252717005",
    "media_url": "https://media.twiliocdn.com/fax/AC58f1e8f2b1c6b88ca90a012a4be0c279/c2d4d301762880a519202e142fe486ff56edbcbfe0e9c5742109b0cd1ac2fa91d58be998c7a4c5905703e590339c7f54196481920eb252bd3fd9837a74094aab?x-amz-security-token=FQoDYXdzEHkaDGaGfJdmYVkoWfLy8yK3AyFTXB6SaYtdA53nCYTajFsMzjM%2Fc%2B7QLS7TqTU2nNMnwJPB4oFQelC6tfdvNGZVvdEhZW3He%2BY9FIVrOT5nk9WMOrcOZFEZ2NdCCpRsGPo5ZcjTbVQKbxQ1IhZvrZCz2MDja3G5RGpBweDtB0UfC6U%2Bj%2BUv8LQChDjS3wu%2FZCsWhWcz1cSSLoPOVDzOtENSlS7SjCITACiJDXerUjRvOxNqRvW2GAUSe7tn8S0%2FGjplc0qXGDRjDTwsIitIE2mA%2BrzMzEx8zXJ%2BblMSKrUaQh3qRfXWtvrZQt8JZj03fIZ2CzL7BJUjoUOATbIRRUlRDhhLCtP07rgMP2yyFNpDV%2BO5B3ZnNdonXmTJuVg4VycaRC%2B5roj%2FJ%2FSnStURa9RApFTPWRXM97CZqiK5R8WWKwBx6D2bK67cHlyrCrnN0fq0aylhR%2BC%2F%2BEM9elZTzlLK7WdjzhEkV5LMgEiWeETWxxGtUiT1R569idTJZVcnAQKkWmRVtAgrlsPSWK1qTkiQew2Lugd0X1rsWgy6OM9S1moeTsVGnf4X%2BvsXppNgYBFocyXD7AzOuIHH2s%2BCGiw6%2BlzC%2B3VHxbwo9Pn5xgU%3D&AWSAccessKeyId=ASIAIHZ6SRB3PKVIGYGA&Expires=1490986753&Signature=ZqFu5tgPrBL5c5PqL07RIGhANao%3D",
    "num_pages": 1,
    "price": "0.014",
    "price_unit": "USD",
    "quality": "fine",
    "sid": "FXeb76f282888a074547beba3516552174",
    "status": "delivered",
    "to": "+18326327228",
    "url": "https://fax.twilio.com/v1/Faxes/FXeb76f282888a074547beba3516552174"
}
`)

var faxGetPageResponse = []byte(`
{
    "faxes": [
        {
            "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
            "api_version": "v1",
            "date_created": "2017-03-31T09:44:20Z",
            "date_updated": "2017-03-31T09:45:55Z",
            "direction": "outbound",
            "duration": 64,
            "from": "+19252717005",
            "media_url": "https://media.twiliocdn.com/fax/AC58f1e8f2b1c6b88ca90a012a4be0c279/c2d4d301762880a519202e142fe486ff56edbcbfe0e9c5742109b0cd1ac2fa91d58be998c7a4c5905703e590339c7f54196481920eb252bd3fd9837a74094aab?x-amz-security-token=FQoDYXdzEHkaDGaGfJdmYVkoWfLy8yK3AyFTXB6SaYtdA53nCYTajFsMzjM%2Fc%2B7QLS7TqTU2nNMnwJPB4oFQelC6tfdvNGZVvdEhZW3He%2BY9FIVrOT5nk9WMOrcOZFEZ2NdCCpRsGPo5ZcjTbVQKbxQ1IhZvrZCz2MDja3G5RGpBweDtB0UfC6U%2Bj%2BUv8LQChDjS3wu%2FZCsWhWcz1cSSLoPOVDzOtENSlS7SjCITACiJDXerUjRvOxNqRvW2GAUSe7tn8S0%2FGjplc0qXGDRjDTwsIitIE2mA%2BrzMzEx8zXJ%2BblMSKrUaQh3qRfXWtvrZQt8JZj03fIZ2CzL7BJUjoUOATbIRRUlRDhhLCtP07rgMP2yyFNpDV%2BO5B3ZnNdonXmTJuVg4VycaRC%2B5roj%2FJ%2FSnStURa9RApFTPWRXM97CZqiK5R8WWKwBx6D2bK67cHlyrCrnN0fq0aylhR%2BC%2F%2BEM9elZTzlLK7WdjzhEkV5LMgEiWeETWxxGtUiT1R569idTJZVcnAQKkWmRVtAgrlsPSWK1qTkiQew2Lugd0X1rsWgy6OM9S1moeTsVGnf4X%2BvsXppNgYBFocyXD7AzOuIHH2s%2BCGiw6%2BlzC%2B3VHxbwo9Pn5xgU%3D&AWSAccessKeyId=ASIAIHZ6SRB3PKVIGYGA&Expires=1490987185&Signature=aG8SXCkaYlIDfH7zGTSS7oYeIXk%3D",
            "num_pages": 1,
            "price": "0.014",
            "price_unit": "USD",
            "quality": "fine",
            "sid": "FXeb76f282888a074547beba3516552174",
            "status": "delivered",
            "to": "+18326327228",
            "url": "https://fax.twilio.com/v1/Faxes/FXeb76f282888a074547beba3516552174"
        }
    ],
    "meta": {
        "first_page_url": "https://fax.twilio.com/v1/Faxes?PageSize=50&Page=0",
        "key": "faxes",
        "next_page_url": null,
        "page": 0,
        "page_size": 50,
        "previous_page_url": null,
        "url": "https://fax.twilio.com/v1/Faxes?PageSize=50&Page=0"
    }
}`)

var availablePhoneNumbers = []byte(`
{
    "available_phone_numbers": [
        {
            "address_requirements": "none",
            "beta": false,
            "capabilities": {
                "MMS": true,
                "SMS": true,
                "fax": true,
                "voice": true
            },
            "friendly_name": "(571) 200-0596",
            "iso_country": "US",
            "lata": "236",
            "latitude": "38.804800",
            "longitude": "-77.046900",
            "phone_number": "+15712000596",
            "postal_code": "22301",
            "rate_center": "WSNGTNZN08",
            "region": "VA"
        }
    ],
    "uri": "/2010-04-01/Accounts/AC58f1e8f2b1c6b88ca90a012a4be0c279/AvailablePhoneNumbers/US/Local.json?Contains=571%2A%2A%2A%2A%2A%2A%2A&InRegion=VA&PageSize=1&SmsEnabled=true&VoiceEnabled=true"
}
`)

var supportedCountries = []byte(`
{
  "uri": "/2010-04-01/Accounts/ACd05ccdab3d3a79ea461fd212ee6975af/AvailablePhoneNumbers.json",
  "countries": [
    {
      "country_code": "ZA",
      "country": "South Africa",
      "uri": "/2010-04-01/Accounts/ACd05ccdab3d3a79ea461fd212ee6975af/AvailablePhoneNumbers/ZA.json",
      "beta": false,
      "subresource_uris": {
        "local": "/2010-04-01/Accounts/ACd05ccdab3d3a79ea461fd212ee6975af/AvailablePhoneNumbers/ZA/Local.json",
        "toll_free": "/2010-04-01/Accounts/ACd05ccdab3d3a79ea461fd212ee6975af/AvailablePhoneNumbers/ZA/TollFree.json"
      }
    },
    {
      "country_code": "PE",
      "country": "Peru",
      "uri": "/2010-04-01/Accounts/ACd05ccdab3d3a79ea461fd212ee6975af/AvailablePhoneNumbers/PE.json",
      "beta": false,
      "subresource_uris": {
        "local": "/2010-04-01/Accounts/ACd05ccdab3d3a79ea461fd212ee6975af/AvailablePhoneNumbers/PE/Local.json"
      }
    }]
}`)

var simGetResponse = []byte(`
{
    "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "commands_callback_method": "POST",
    "commands_callback_url": null,
    "date_created": "2017-06-22T21:44:48Z",
    "date_updated": "2017-06-22T21:45:45Z",
    "e_id": null,
    "friendly_name": null,
    "iccid": "8901260852290182427",
    "links": {
        "rate_plan": "https://wireless.twilio.com/v1/RatePlans/WPfaf4062b7697612c807c1d5e9316ffbc",
        "usage_records": "https://wireless.twilio.com/v1/Sims/DEe10f758e920e43318ad80677505fcf90/UsageRecords"
    },
    "rate_plan_sid": "WPfaf4062b7697612c807c1d5e9316ffbc",
    "sid": "DEe10f758e920e43318ad80677505fcf90",
    "sms_fallback_method": null,
    "sms_fallback_url": null,
    "sms_method": null,
    "sms_url": null,
    "status": "active",
    "unique_name": "iPhone v13",
    "url": "https://wireless.twilio.com/v1/Sims/DEe10f758e920e43318ad80677505fcf90",
    "voice_fallback_method": null,
    "voice_fallback_url": null,
    "voice_method": null,
    "voice_url": null
}
`)

var cmdCreateResponse = []byte(`
{
    "account_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "command": "twilio-go testing!",
    "command_mode": "text",
    "date_created": "2017-07-13T15:49:02Z",
    "date_updated": "2017-07-13T15:49:02Z",
    "direction": "to_sim",
    "sid": "DC5c94b39819e367e06417f695f138781f",
    "sim_sid": "DEe10f758e920e43318ad80677505fcf90",
    "status": "queued",
    "url": "https://wireless.twilio.com/v1/Commands/DC5c94b39819e367e06417f695f138781f"
}
`)

var phoneLookupResponse = []byte(`
{
    "country_code": "US",
    "phone_number": "+14157012312",
    "national_format": "(415) 701-2312",
    "url": "https://lookups.twilio.com/v1/PhoneNumber/+14157012312",
    "caller_name": {
        "caller_name": "CCSF",
        "caller_type": "BUSINESS",
        "error_code": null
    },    
    "carrier": {
        "type": "landline",
        "error_code": null,
        "mobile_network_code": null,
        "mobile_country_code": null,
        "name": "Pacific Bell"
    }
}	
`)

var roomResponse = []byte(`
{
    "api_key_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "date_created": "2015-07-30T20:00:00Z",
    "date_updated": "2015-07-30T20:00:00Z",
    "status": "in-progress",
    "type": "peer-to-peer",
    "sid": "RMca86cf94c7d4f89e0bd45bfa7d9b9e7d",
    "enable_turn": false,
    "unique_name": "DailyStandup",
    "max_participants": 10,
    "duration": 0,
    "status_callback_method": "POST",
    "status_callback": "",
    "record_participants_on_connect": false,
    "end_time": "2015-07-30T20:00:00Z",
    "url": "https://video.twilio.com/v1/Rooms/RMca86cf94c7d4f89e0bd45bfa7d9b9e7d",
    "links": {
        "recordings": "https://video.twilio.com/v1/Rooms/RMca86cf94c7d4f89e0bd45bfa7d9b9e7d/Recordings"
    }
}
`)

var videoRecordingResponse = []byte(`
{
    "api_key_sid": "AC58f1e8f2b1c6b88ca90a012a4be0c279",
    "status": "processing",
    "date_created": "2015-07-30T20:00:00Z",
    "sid": "RT63868a235fc1cf3987e6a2b67346273f",
    "source_sid": "MT58f1e8f2b1c6b88ca90a012a4be0c279",
    "size": 0,
    "url": "https://video.twilio.com/v1/Recordings/RT58f1e8f2b1c6b88ca90a012a4be0c279",
    "type": "audio",
    "duration": 20,
    "container_format": "mka",
    "codec": "OPUS",
    "grouping_sids": {
        "room_sid" : "RM58f1e8f2b1c6b88ca90a012a4be0c279"
    },
    "links": {
        "media": "https://video.twilio.com/v1/Recordings/RT58f1e8f2b1c6b88ca90a012a4be0c279/Media"
    }
}
`)

const from = "+19253920364"
const to = "+19253920364"
