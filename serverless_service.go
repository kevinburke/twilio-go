package twilio

import (
	"context"
	"net/url"
)

const servicesPathPrefix = "Services"

type ServiceService struct {
	client *Client
}

type Service struct {
	Sid                string     `json:"sid"`
	AccountSid         string     `json:"account_sid"`
	FriendlyName       string     `json:"friendly_name"`
	UniqueName         string     `json:"unique_name"`
	IncludeCredentials bool       `json:"include_credentials"`
	UIEditable         bool       `json:"ui_editable"`
	DateCreated        TwilioTime `json:"date_created"`
	DateUpdated        TwilioTime `json:"date_updated"`
	URL                string     `json:"url"`
}

type ServicePage struct {
	Page
	Services []*Service `json:"services"`
}

// Get retrieves a Service by its sid.
//
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/Service for
// more.
func (s *ServiceService) Get(ctx context.Context, sid string) (*Service, error) {
	service := new(Service)
	err := s.client.GetResource(ctx, servicesPathPrefix, sid, service)
	return service, err
}

// Create creates a new Service.
//
// For a list of valid parameters see
// https://www.twilio.com/docs/runtime/functions-assets-api/api/Service.
func (s *ServiceService) Create(ctx context.Context, data url.Values) (*Service, error) {
	service := new(Service)
	err := s.client.CreateResource(ctx, servicesPathPrefix, data, service)
	return service, err
}

// Delete deletes an Service.
//
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/Service for
// more.
func (s *ServiceService) Delete(ctx context.Context, sid string) error {
	return s.client.DeleteResource(ctx, servicesPathPrefix, sid)
}

// Update updates an Service, using the given data.
//
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/Service for
// more.
func (s *ServiceService) Update(ctx context.Context, sid string, data url.Values) (*Service, error) {
	service := new(Service)
	err := s.client.UpdateResource(ctx, servicesPathPrefix, sid, data, service)
	return service, err
}

// GetPage retrieves an ServicePage, filtered by the given data.
func (s *ServiceService) GetPage(ctx context.Context, data url.Values) (*ServicePage, error) {
	iter := s.GetPageIterator(data)
	return iter.Next(ctx)
}

type ServicePageIterator struct {
	p *PageIterator
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (s *ServiceService) GetPageIterator(data url.Values) *ServicePageIterator {
	iter := NewPageIterator(s.client, data, servicesPathPrefix)
	return &ServicePageIterator{
		p: iter,
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (spi *ServicePageIterator) Next(ctx context.Context) (*ServicePage, error) {
	sp := new(ServicePage)
	err := spi.p.Next(ctx, sp)
	if err != nil {
		return nil, err
	}
	spi.p.SetNextPageURI(sp.NextPageURI)
	return sp, nil
}
