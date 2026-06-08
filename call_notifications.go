package twilio

import (
	"context"
	"net/url"
)

// CallNotification is a log entry (an alert or error) Twilio recorded while
// handling a Call. The list representation omits RequestVariables, ResponseBody,
// and ResponseHeaders; the instance representation includes them.
//
// See https://www.twilio.com/docs/voice/api/call-notification-resource for more.
type CallNotification struct {
	Sid           string     `json:"sid"`
	AccountSid    string     `json:"account_sid"`
	APIVersion    string     `json:"api_version"`
	CallSid       string     `json:"call_sid"`
	ErrorCode     string     `json:"error_code"`
	Log           string     `json:"log"`
	MessageText   string     `json:"message_text"`
	MessageDate   TwilioTime `json:"message_date"`
	MoreInfo      string     `json:"more_info"`
	RequestMethod string     `json:"request_method"`
	RequestURL    string     `json:"request_url"`
	DateCreated   TwilioTime `json:"date_created"`
	DateUpdated   TwilioTime `json:"date_updated"`
	URI           string     `json:"uri"`

	// Only present on the instance (single-notification) representation.
	RequestVariables string `json:"request_variables"`
	ResponseBody     string `json:"response_body"`
	ResponseHeaders  string `json:"response_headers"`
}

type CallNotificationPage struct {
	Page
	Notifications []*CallNotification `json:"notifications"`
}

func callNotificationsPathPart(callSid string) string {
	return callsPathPart + "/" + callSid + "/Notifications"
}

// GetNotification returns the CallNotification with the given sid for the given
// call. The instance representation includes the request and response detail
// fields.
func (c *CallService) GetNotification(ctx context.Context, callSid string, sid string) (*CallNotification, error) {
	n := new(CallNotification)
	err := c.client.GetResource(ctx, callNotificationsPathPart(callSid), sid, n)
	return n, err
}

// GetNotifications returns a single page of notifications for the call with the
// given callSid. To retrieve multiple pages, use GetNotificationsIterator.
func (c *CallService) GetNotifications(ctx context.Context, callSid string, data url.Values) (*CallNotificationPage, error) {
	return c.GetNotificationsIterator(callSid, data).Next(ctx)
}

type CallNotificationPageIterator struct {
	p *PageIterator
}

// GetNotificationsIterator returns an iterator over the notifications for the
// given call.
func (c *CallService) GetNotificationsIterator(callSid string, data url.Values) *CallNotificationPageIterator {
	return &CallNotificationPageIterator{
		p: NewPageIterator(c.client, data, callNotificationsPathPart(callSid)),
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (i *CallNotificationPageIterator) Next(ctx context.Context) (*CallNotificationPage, error) {
	np := new(CallNotificationPage)
	err := i.p.Next(ctx, np)
	if err != nil {
		return nil, err
	}
	i.p.SetNextPageURI(np.NextPageURI)
	return np, nil
}
