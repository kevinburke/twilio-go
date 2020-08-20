package twilio

import (
	"context"
	"github.com/kevinburke/go-types"
	"net/url"
	"strings"
)

type ServiceService struct {
	client *Client
}

type Service struct {
	Sid                   string           `json:"sid"`
	AccountSid            string           `json:"account_sid"`
	FriendlyName          string           `json:"friendly_name"`
	DateCreated           TwilioTime       `json:"date_created"`
	DateUpdated           TwilioTime       `json:"date_updated"`
	InboundRequestURL     types.NullString `json:"inbound_request_url"`
	InboundMethod         string           `json:"inbound_method"`
	FallbackURL           string           `json:"fallback_url"`
	FallbackMethod        string           `json:"fallback_method"`
	StatusCallback        string           `json:"status_callback"`
	StickySender          bool             `json:"sticky_sender"`
	MMSConverter          bool             `json:"mms_converter"`
	SmartEncoding         bool             `json:"smart_encoding"`
	FallbackToLongCode    bool             `json:"fallbackToLongCode"`
	AreaCodeGeomatch      bool             `json:"area_code_geomatch"`
	ValidityPeriod        uint             `json:"validity_period"`
	SynchronousValidation bool             `json:"synchronous_validation"`
}

type ServicePhoneNumber struct {
	Sid         string      `json:"sid"`
	PhoneNumber PhoneNumber `json:"phone_number"`
	ServiceSid  string      `json:"service_sid"`
	DateCreated TwilioTime  `json:"date_created"`
	AccountSid  string      `json:"account_sid"`
	DateUpdated TwilioTime  `json:"date_updated"`
	CountryCode string      `json:"country_code"`
	URL         string      `json:"url"`
}

// A ServicePage contains a Page of services.
type ServicePage struct {
	Page
	Services []*Service `json:"services"`
}

// A ServicePhoneNumberPage contains a Page of service phone numbers
type ServicePhoneNumberPage struct {
	Page
	PhoneNumbers []*ServicePhoneNumber `json:"phone_numbers"`
}

// Create a service with the given url.Values. For more information on valid
// values, see https://www.twilio.com/docs/sms/services/api#create-a-service-resource
func (s *ServiceService) Create(ctx context.Context, data url.Values) (*Service, error) {
	msg := new(Service)
	err := s.client.CreateResource(ctx, servicesPathPart, data, msg)
	return msg, err
}

// Tries to update the service's properties, and returns the updated resource representation if successful.
// https://www.twilio.com/docs/sms/services/api#update-a-service-resource
func (s *ServiceService) Update(ctx context.Context, sid string, data url.Values) (*Service, error) {
	service := new(Service)
	err := s.client.UpdateResource(ctx, servicesPathPart, sid, data, service)
	return service, err
}

// Adds the given IncomingPhoneNumber to the messaging Service
func (s *ServiceService) CreatePhoneNumber(ctx context.Context, serviceSid string, phoneNumberSid string) (*ServicePhoneNumber, error) {
	pn := new(ServicePhoneNumber)
	v := url.Values{}
	v.Set("PhoneNumberSid", phoneNumberSid)
	path := strings.Join([]string{servicesPathPart, serviceSid, phoneNumbersPathPart}, "/")
	err := s.client.CreateResource(ctx, path, v, pn)
	return pn, err
}

// Gets the ServicePhoneNumber by phoneNumberSid associated with the given Service by serviceSid
func (s *ServiceService) GetPhoneNumber(ctx context.Context, serviceSid string, phoneNumberSid string) (*ServicePhoneNumber, error) {
	msg := new(ServicePhoneNumber)
	path := strings.Join([]string{servicesPathPart, serviceSid, phoneNumbersPathPart}, "/")
	err := s.client.GetResource(ctx, path, phoneNumberSid, msg)
	return msg, err
}

// Removes the ServicePhoneNumber with the given phoneNumberSid from the Service with the given serviceSid.
// If the PhoneNumber has already been deleted, or does not exist, DeletePhoneNumber returns nil. If another error or a
// timeout occurs, the error is returned.
func (s *ServiceService) DeletePhoneNumber(ctx context.Context, serviceSid string, phoneNumberSid string) error {
	path := strings.Join([]string{servicesPathPart, serviceSid, phoneNumbersPathPart}, "/")
	return s.client.DeleteResource(ctx, path, phoneNumberSid)
}

// ServicePageIterator lets you retrieve consecutive pages of resources.
type ServicePageIterator interface {
	// Next returns the next page of resources. If there are no more resources,
	// NoMoreResults is returned.
	Next(context.Context) (*ServicePage, error)
}

type servicePageIterator struct {
	p *PageIterator
}

// ServicePageIterator lets you retrieve consecutive pages of resources.
type ServicePhoneNumberPageIterator interface {
	// Next returns the next page of resources. If there are no more resources,
	// NoMoreResults is returned.
	Next(context.Context) (*ServicePhoneNumberPage, error)
}

type servicePhoneNumberPageIterator struct {
	p *PageIterator
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (s *servicePageIterator) Next(ctx context.Context) (*ServicePage, error) {
	mp := new(ServicePage)
	err := s.p.Next(ctx, mp)
	if err != nil {
		return nil, err
	}
	s.p.SetNextPageURI(mp.NextPageURI)
	return mp, nil
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (s *servicePhoneNumberPageIterator) Next(ctx context.Context) (*ServicePhoneNumberPage, error) {
	mp := new(ServicePhoneNumberPage)
	err := s.p.Next(ctx, mp)
	if err != nil {
		return nil, err
	}
	s.p.SetNextPageURI(mp.NextPageURI)
	return mp, nil
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (s *ServiceService) GetPageIterator(data url.Values) ServicePageIterator {
	iter := NewPageIterator(s.client, data, servicesPathPart)
	return &servicePageIterator{
		p: iter,
	}
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (s *ServiceService) GetPhoneNumberPageIterator(serviceSid string, data url.Values) ServicePhoneNumberPageIterator {
	path := strings.Join([]string{servicesPathPart, serviceSid, phoneNumbersPathPart}, "/")
	iter := NewPageIterator(s.client, data, path)
	return &servicePhoneNumberPageIterator{
		p: iter,
	}
}

func (s *ServiceService) Get(ctx context.Context, sid string) (*Service, error) {
	msg := new(Service)
	err := s.client.GetResource(ctx, servicesPathPart, sid, msg)
	return msg, err
}

// GetPage returns a single page of resources. To retrieve multiple pages, use
// GetPageIterator.
func (s *ServiceService) GetPage(ctx context.Context, data url.Values) (*ServicePage, error) {
	iter := s.GetPageIterator(data)
	return iter.Next(ctx)
}

// GetPhoneNumberPage returns a single page of resources. To retrieve multiple pages, use
// GetPageIterator.
func (s *ServiceService) GetPhoneNumberPage(ctx context.Context, serviceSid string, data url.Values) (*ServicePhoneNumberPage, error) {
	iter := s.GetPhoneNumberPageIterator(serviceSid, data)
	return iter.Next(ctx)
}

// Delete the Service with the given sid. If the Service has already been
// deleted, or does not exist, Delete returns nil. If another error or a
// timeout occurs, the error is returned.
func (s *ServiceService) Delete(ctx context.Context, sid string) error {
	return s.client.DeleteResource(ctx, servicesPathPart, sid)
}
