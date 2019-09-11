package twilio

import (
	"context"
	"net/url"
)

const messagingServicePathPart = "Services"

type ServiceResource struct {
	client   *Client
	pathPart string
}

type MessagingServiceLinks struct {
	PhoneNumbers string `json:"phone_numbers"`
	ShortCodes   string `json:"short_codes"`
	AlphaSenders string `json:"alpha_senders"`
	Messages     string `json:"messages"`
	Broadcasts   string `json:"broadcasts"`
}

type MessagingService struct {
	AccountSid            string                 `json:"account_sid"`
	Sid                   string                 `json:"sid"`
	DateCreated           TwilioTime             `json:"date_created"`
	DateUpdated           TwilioTime             `json:"date_updated"`
	FriendlyName          string                 `json:"friendly_name"`
	InboundRequestURL     string                 `json:"inbound_request_url"`
	InboundMethod         string                 `json:"inbound_method"`
	FallbackURL           string                 `json:"fallback_url"`
	FallbackMethod        string                 `json:"fallback_method"`
	StatusCallback        string                 `json:"status_callback"`
	StickySender          bool                   `json:"sticky_sender"`
	SmartEncoding         bool                   `json:"smart_encoding"`
	MMSConverter          bool                   `json:"mms_converter"`
	FallbackToLongCode    bool                   `json:"fallback_to_long_code"`
	ScanMessageContent    string                 `json:"scan_message_content"`
	AreaCodeGeomatch      string                 `json:"area_code_geomatch"`
	ValidityPeriod        int                    `json:"validity_period"`
	SynchronousValidation bool                   `json:"synchronous_validation"`
	Links                 *MessagingServiceLinks `json:"links"`
	URL                   string                 `json:"url"`
}

type MessagingServicePage struct {
	Meta              Meta                `json:"meta"`
	MessagingServices []*MessagingService `json:"services"`
}

// Create a service resource
//
// https://www.twilio.com/docs/sms/services/api#create-a-service-resource
func (s *ServiceResource) Create(ctx context.Context, data url.Values) (*MessagingService, error) {
	messagingService := new(MessagingService)
	err := s.client.CreateResource(ctx, messagingServicePathPart, data, messagingService)
	return messagingService, err
}

// Fetch a service resource
//
// https://www.twilio.com/docs/sms/services/api#fetch-a-service-resource
func (s *ServiceResource) Fetch(ctx context.Context, sid string) (*MessagingService, error) {
	messagingService := new(MessagingService)
	err := s.client.GetResource(ctx, messagingServicePathPart, sid, messagingService)
	return messagingService, err
}

// Update a service resource
//
// https://www.twilio.com/docs/sms/services/api#update-a-service-resource
func (s *ServiceResource) Update(ctx context.Context, sid string, data url.Values) (*MessagingService, error) {
	messagingService := new(MessagingService)
	err := s.client.UpdateResource(ctx, messagingServicePathPart, sid, data, messagingService)
	return messagingService, err
}

// Delete a service resource
//
// https://www.twilio.com/docs/sms/services/api#delete-a-service-resource
func (s *ServiceResource) Delete(ctx context.Context, sid string, data url.Values) error {
	return s.client.DeleteResource(ctx, messagingServicePathPart, sid)
}

// GetPage retrieves an IncomingPhoneNumberPage, filtered by the given data.
func (ins *ServiceResource) GetPage(ctx context.Context, data url.Values) (*MessagingServicePage, error) {
	iter := ins.GetPageIterator(data)
	return iter.Next(ctx)
}

type MessagingServiceIterator struct {
	p *PageIterator
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (c *ServiceResource) GetPageIterator(data url.Values) *MessagingServiceIterator {
	iter := NewPageIterator(c.client, data, messagingServicePathPart)
	return &MessagingServiceIterator{
		p: iter,
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (c *MessagingServiceIterator) Next(ctx context.Context) (*MessagingServicePage, error) {
	cp := new(MessagingServicePage)
	err := c.p.Next(ctx, cp)
	if err != nil {
		return nil, err
	}
	c.p.SetNextPageURI(cp.Meta.NextPageURL)
	return cp, nil
}
