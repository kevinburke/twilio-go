package twilio

import (
	"context"
	"net/url"
)

const servicesPathPart = "Services"
const verificationsPathPart = "Verifications"
const verificationCheckPart = "VerificationCheck"

type VerifyPhoneNumberService struct {
	client *Client
}

type VerifyPhoneNumber struct {
	Sid         string      `json:"sid"`
	ServiceSid  string      `json:"service_sid"`
	AccountSid  string      `json:"account_sid"`
	To          PhoneNumber `json:"to"`
	Channel     string      `json:"channel"`
	Status      string      `json:"status"`
	Valid       bool        `json:"valid"`
	Lookup      PhoneLookup `json:"lookup"`
	Amount      string      `json:"amount"`
	Payee       string      `json:"payee"`
	DateCreated TwilioTime  `json:"date_created"`
	DateUpdated TwilioTime  `json:"date_updated"`
	URL         string      `json:"url"`
}

type CheckPhoneNumber struct {
	Sid         string     `json:"sid"`
	ServiceSid  string     `json:"service_sid"`
	AccountSid  string     `json:"account_sid"`
	To          string     `json:"to"`
	Channel     string     `json:"channel"`
	Status      string     `json:"status"`
	Valid       bool       `json:"valid"`
	Amount      string     `json:"amount"`
	Payee       string     `json:"payee"`
	DateCreated TwilioTime `json:"date_created"`
	DateUpdated TwilioTime `json:"date_updated"`
}

func (v *VerifyPhoneNumberService) Create(ctx context.Context, verifyServiceID string, data url.Values) (*VerifyPhoneNumber, error) {
	verify := new(VerifyPhoneNumber)
	err := v.client.CreateResource(ctx, servicesPathPart+"/"+verifyServiceID+"/"+verificationsPathPart, data, verify)
	return verify, err
}

func (v *VerifyPhoneNumberService) Get(ctx context.Context, verifyServiceID string, sid string) (*VerifyPhoneNumber, error) {
	verify := new(VerifyPhoneNumber)
	err := v.client.GetResource(ctx, servicesPathPart+"/"+verifyServiceID+"/"+verificationsPathPart, sid, verify)
	return verify, err
}

func (v *VerifyPhoneNumberService) Check(ctx context.Context, verifyServiceID string, data url.Values) (*CheckPhoneNumber, error) {
	check := new(CheckPhoneNumber)
	err := v.client.CreateResource(ctx, servicesPathPart+"/"+verifyServiceID+"/"+verificationCheckPart, data, check)
	return check, err
}
