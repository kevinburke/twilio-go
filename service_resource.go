package twilio

import (
	"context"
	"fmt"
	"net/url"
)

const (
	messagingServicePathPart       = "Services"
	phoneNumberServicePathTemplate = messagingServicePathPart + "/%s/PhoneNumbers"
	alphaSenderServicePathTemplate = messagingServicePathPart + "/%s/AlphaSenders"
)

type ServiceResourceService struct {
	*MessagingService
	PhoneNumber      *PhoneNumberService
	AlphaSender      *AlphaSenderService
}

type MessagingService struct {
	client *Client
}

type PhoneNumberService struct {
	client *Client
}

type AlphaSenderService struct {
	client *Client
}

type MessagingResourceLinks struct {
	PhoneNumbers string `json:"phone_numbers"`
	ShortCodes   string `json:"short_codes"`
	AlphaSenders string `json:"alpha_senders"`
	Messages     string `json:"messages"`
	Broadcasts   string `json:"broadcasts"`
}

type MessagingResource struct {
	AccountSid            string                  `json:"account_sid"`
	Sid                   string                  `json:"sid"`
	DateCreated           TwilioTime              `json:"date_created"`
	DateUpdated           TwilioTime              `json:"date_updated"`
	FriendlyName          string                  `json:"friendly_name"`
	InboundRequestURL     string                  `json:"inbound_request_url"`
	InboundMethod         string                  `json:"inbound_method"`
	FallbackURL           string                  `json:"fallback_url"`
	FallbackMethod        string                  `json:"fallback_method"`
	StatusCallback        string                  `json:"status_callback"`
	StickySender          bool                    `json:"sticky_sender"`
	SmartEncoding         bool                    `json:"smart_encoding"`
	MMSConverter          bool                    `json:"mms_converter"`
	FallbackToLongCode    bool                    `json:"fallback_to_long_code"`
	ScanMessageContent    string                  `json:"scan_message_content"`
	AreaCodeGeomatch      bool                    `json:"area_code_geomatch"`
	ValidityPeriod        int                     `json:"validity_period"`
	SynchronousValidation bool                    `json:"synchronous_validation"`
	Links                 *MessagingResourceLinks `json:"links"`
	URL                   string                  `json:"url"`
}

type MessagingResourcePage struct {
	Meta              Meta                 `json:"meta"`
	MessagingResource []*MessagingResource `json:"services"`
}

// Create a service resource
//
// https://www.twilio.com/docs/sms/services/api#create-a-service-resource
func (s *MessagingService) Create(ctx context.Context, data url.Values) (*MessagingResource, error) {
	messagingService := new(MessagingResource)
	err := s.client.CreateResource(ctx, messagingServicePathPart, data, messagingService)
	return messagingService, err
}

// Fetch a service resource
//
// https://www.twilio.com/docs/sms/services/api#fetch-a-service-resource
func (s *MessagingService) Fetch(ctx context.Context, sid string) (*MessagingResource, error) {
	messagingService := new(MessagingResource)
	err := s.client.GetResource(ctx, messagingServicePathPart, sid, messagingService)
	return messagingService, err
}

// Update a service resource
//
// https://www.twilio.com/docs/sms/services/api#update-a-service-resource
func (s *MessagingService) Update(ctx context.Context, sid string, data url.Values) (*MessagingResource, error) {
	messagingService := new(MessagingResource)
	err := s.client.UpdateResource(ctx, messagingServicePathPart, sid, data, messagingService)
	return messagingService, err
}

// Delete a service resource
//
// https://www.twilio.com/docs/sms/services/api#delete-a-service-resource
func (s *MessagingService) Delete(ctx context.Context, sid string, data url.Values) error {
	return s.client.DeleteResource(ctx, messagingServicePathPart, sid)
}

// GetPage retrieves an MessagingResourcePage, filtered by the given data.
func (s *MessagingService) GetPage(ctx context.Context, data url.Values) (*MessagingResourcePage, error) {
	iter := s.GetPageIterator(data)
	return iter.Next(ctx)
}

type MessagingServiceIterator struct {
	p *PageIterator
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (s *MessagingService) GetPageIterator(data url.Values) *MessagingServiceIterator {
	iter := NewPageIterator(s.client, data, messagingServicePathPart)
	return &MessagingServiceIterator{
		p: iter,
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (c *MessagingServiceIterator) Next(ctx context.Context) (*MessagingResourcePage, error) {
	cp := new(MessagingResourcePage)
	err := c.p.Next(ctx, cp)
	if err != nil {
		return nil, err
	}
	c.p.SetNextPageURI(cp.Meta.NextPageURL)
	return cp, nil
}

type PhoneNumberResource struct {
	Sid          string      `json:"sid"`
	AccountSid   string      `json:"account_sid"`
	ServiceSid   string      `json:"service_sid"`
	DateCreated  TwilioTime  `json:"date_created"`
	DateUpdated  TwilioTime  `json:"date_updated"`
	PhoneNumber  PhoneNumber `json:"phone_number"`
	CountryCode  string      `json:"country_code"`
	Capabilities []string    `json:"capabilities"`
	URL          string      `json:"url"`
}

type PhoneNumberResourcePage struct {
	Meta         Meta                   `json:"meta"`
	PhoneNumbers []*PhoneNumberResource `json:"phone_numbers"`
}

// Create a phone number resource for messaging service
//
// https://www.twilio.com/docs/sms/services/api/phonenumber-resource#create-a-phonenumber-resource
func (s *PhoneNumberService) Create(ctx context.Context, serviceSID string, data url.Values) (*PhoneNumberResource, error) {
	messagingService := new(PhoneNumberResource)
	pathPart := fmt.Sprintf(phoneNumberServicePathTemplate, serviceSID)
	err := s.client.CreateResource(ctx, pathPart, data, messagingService)
	return messagingService, err
}

// Fetch a phone number resource for messaging service
//
// https://www.twilio.com/docs/sms/services/api/phonenumber-resource#fetch-a-phonenumber-resource
func (s *PhoneNumberService) Fetch(ctx context.Context, serviceSID, sid string) (*PhoneNumberResource, error) {
	messagingService := new(PhoneNumberResource)
	pathPart := fmt.Sprintf(phoneNumberServicePathTemplate, serviceSID)
	err := s.client.GetResource(ctx, pathPart, sid, messagingService)
	return messagingService, err
}

// Delete a phone number resource from messaging service
//
// https://www.twilio.com/docs/sms/services/api/phonenumber-resource#delete-a-phonenumber-resource
func (s *PhoneNumberService) Delete(ctx context.Context, serviceSID, sid string, data url.Values) error {
	pathPart := fmt.Sprintf(phoneNumberServicePathTemplate, serviceSID)
	return s.client.DeleteResource(ctx, pathPart, sid)
}

// GetPage retrieves an PhoneNumberResourcePage for given messaging service, filtered by the given data.
func (s *PhoneNumberService) GetPage(ctx context.Context, serviceSID string, data url.Values) (*PhoneNumberResourcePage, error) {
	iter := s.GetPageIterator(serviceSID, data)
	return iter.Next(ctx)
}

type PhoneNumberIterator struct {
	p *PageIterator
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (s *PhoneNumberService) GetPageIterator(serviceSID string, data url.Values) *PhoneNumberIterator {
	pathPart := fmt.Sprintf(phoneNumberServicePathTemplate, serviceSID)
	iter := NewPageIterator(s.client, data, pathPart)
	return &PhoneNumberIterator{
		p: iter,
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (c *PhoneNumberIterator) Next(ctx context.Context) (*PhoneNumberResourcePage, error) {
	cp := new(PhoneNumberResourcePage)
	err := c.p.Next(ctx, cp)
	if err != nil {
		return nil, err
	}
	c.p.SetNextPageURI(cp.Meta.NextPageURL)
	return cp, nil
}

type AlphaSenderResource struct {
	Sid          string     `json:"sid"`
	AccountSid   string     `json:"account_sid"`
	ServiceSid   string     `json:"service_sid"`
	DateCreated  TwilioTime `json:"date_created"`
	DateUpdated  TwilioTime `json:"date_updated"`
	AlphaSender  string     `json:"alpha_sender"`
	Capabilities []string   `json:"capabilities"`
	URL          string     `json:"url"`
}

type AlphaSenderResourcePage struct {
	Meta         Meta                   `json:"meta"`
	AlphaSenders []*AlphaSenderResource `json:"alpha_senders"`
}

// Create a alpha sender resource for messaging service
//
// https://www.twilio.com/docs/sms/services/api/alphasender-resource#create-an-alphasender-resource
func (s *AlphaSenderService) Create(ctx context.Context, serviceSID string, data url.Values) (*AlphaSenderResource, error) {
	messagingService := new(AlphaSenderResource)
	pathPart := fmt.Sprintf(alphaSenderServicePathTemplate, serviceSID)
	err := s.client.CreateResource(ctx, pathPart, data, messagingService)
	return messagingService, err
}

// Fetch a alpha sender resource for messaging service
//
// https://www.twilio.com/docs/sms/services/api/alphasender-resource#fetch-an-alphasender-resource
func (s *AlphaSenderService) Fetch(ctx context.Context, serviceSID, sid string) (*AlphaSenderResource, error) {
	messagingService := new(AlphaSenderResource)
	pathPart := fmt.Sprintf(alphaSenderServicePathTemplate, serviceSID)
	err := s.client.GetResource(ctx, pathPart, sid, messagingService)
	return messagingService, err
}

// Delete a alpha sender resource from messaging service
//
// https://www.twilio.com/docs/sms/services/api/alphasender-resource#delete-an-alphasender-resource
func (s *AlphaSenderService) Delete(ctx context.Context, serviceSID, sid string, data url.Values) error {
	pathPart := fmt.Sprintf(alphaSenderServicePathTemplate, serviceSID)
	return s.client.DeleteResource(ctx, pathPart, sid)
}

// GetPage retrieves an AlphaSenderResourcePage for given messaging service, filtered by the given data.
func (s *AlphaSenderService) GetPage(ctx context.Context, serviceSID string, data url.Values) (*AlphaSenderResourcePage, error) {
	iter := s.GetPageIterator(serviceSID, data)
	return iter.Next(ctx)
}

type AlphaSenderIterator struct {
	p *PageIterator
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (s *AlphaSenderService) GetPageIterator(serviceSID string, data url.Values) *AlphaSenderIterator {
	pathPart := fmt.Sprintf(alphaSenderServicePathTemplate, serviceSID)
	iter := NewPageIterator(s.client, data, pathPart)
	return &AlphaSenderIterator{
		p: iter,
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (c *AlphaSenderIterator) Next(ctx context.Context) (*AlphaSenderResourcePage, error) {
	cp := new(AlphaSenderResourcePage)
	err := c.p.Next(ctx, cp)
	if err != nil {
		return nil, err
	}
	c.p.SetNextPageURI(cp.Meta.NextPageURL)
	return cp, nil
}