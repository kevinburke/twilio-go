package twilio

import (
	"context"
	"net/url"
)

// Payment represents a pay-as-you-go <Pay> session capturing payment
// information during a Call. Twilio returns only the identifying fields over
// REST; the captured data is delivered to your application via TwiML callbacks.
//
// See https://www.twilio.com/docs/voice/api/payment-resource for more.
type Payment struct {
	Sid         string     `json:"sid"`
	AccountSid  string     `json:"account_sid"`
	CallSid     string     `json:"call_sid"`
	DateCreated TwilioTime `json:"date_created"`
	DateUpdated TwilioTime `json:"date_updated"`
	URI         string     `json:"uri"`
}

func callPaymentsPathPart(callSid string) string {
	return callsPathPart + "/" + callSid + "/Payments"
}

// CreatePayment starts a <Pay> session on the in-progress call with the given
// callSid. The IdempotencyKey and StatusCallback parameters are required; see
// https://www.twilio.com/docs/voice/api/payment-resource#create-a-payment-resource
// for the full list.
func (c *CallService) CreatePayment(ctx context.Context, callSid string, data url.Values) (*Payment, error) {
	payment := new(Payment)
	err := c.client.CreateResource(ctx, callPaymentsPathPart(callSid), data, payment)
	return payment, err
}

// UpdatePayment advances or completes the <Pay> session with the given sid, for
// example to capture the next data field or to complete the session. See
// https://www.twilio.com/docs/voice/api/payment-resource#update-a-payment-resource
// for the available parameters.
func (c *CallService) UpdatePayment(ctx context.Context, callSid string, sid string, data url.Values) (*Payment, error) {
	payment := new(Payment)
	err := c.client.UpdateResource(ctx, callPaymentsPathPart(callSid), sid, data, payment)
	return payment, err
}
