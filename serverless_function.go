package twilio

import (
	"context"
	"net/url"
	"path/filepath"
)

const functionPathPart = "Functions"

type FunctionService struct {
	client     *Client
	serviceSid string

	Versions *FunctionVersionService
}

type Function struct {
	Sid          string     `json:"sid"`
	AccountSid   string     `json:"account_sid"`
	ServiceSid   string     `json:"service_sid"`
	FriendlyName string     `json:"friendly_name"`
	DateCreated  TwilioTime `json:"date_created"`
	DateUpdated  TwilioTime `json:"date_updated"`
	URL          string     `json:"url"`
}

type FunctionPage struct {
	Page
	Functions []*Function `json:"functions"`
}

// Get retrieves a Function by its sid.
//
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/function for
// more.
func (f *FunctionService) Get(ctx context.Context, sid string) (*Function, error) {
	function := new(Function)
	pathPart := filepath.Join(servicesPathPrefix, f.serviceSid, functionPathPart)
	err := f.client.GetResource(ctx, pathPart, sid, function)
	return function, err
}

// Create creates a new Function.
//
// For a list of valid parameters see
// https://www.twilio.com/docs/runtime/functions-assets-api/api/function.
func (f *FunctionService) Create(ctx context.Context, data url.Values) (*Function, error) {
	function := new(Function)
	pathPart := filepath.Join(servicesPathPrefix, f.serviceSid, functionPathPart)
	err := f.client.CreateResource(ctx, pathPart, data, function)
	return function, err
}

// Delete deletes a Function.
//
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/function for
// more.
func (f *FunctionService) Delete(ctx context.Context, sid string) error {
	pathPart := filepath.Join(servicesPathPrefix, f.serviceSid, functionPathPart)
	return f.client.DeleteResource(ctx, pathPart, sid)
}

// Update updates a Function, using the given data.
//
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/function for
// more.
func (f *FunctionService) Update(ctx context.Context, sid string, data url.Values) (*Function, error) {
	function := new(Function)
	pathPart := filepath.Join(servicesPathPrefix, f.serviceSid, functionPathPart)
	err := f.client.UpdateResource(ctx, pathPart, sid, data, function)
	return function, err
}

// GetPage retrieves a FunctionPage, filtered by the given data.
func (f *FunctionService) GetPage(ctx context.Context, data url.Values) (*FunctionPage, error) {
	iter := f.GetPageIterator(data)
	return iter.Next(ctx)
}

type FunctionPageIterator struct {
	p *PageIterator
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (f *FunctionService) GetPageIterator(data url.Values) *FunctionPageIterator {
	pathPart := filepath.Join(servicesPathPrefix, f.serviceSid, functionPathPart)
	iter := NewPageIterator(f.client, data, pathPart)
	return &FunctionPageIterator{
		p: iter,
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (fpi *FunctionPageIterator) Next(ctx context.Context) (*FunctionPage, error) {
	cp := new(FunctionPage)
	err := fpi.p.Next(ctx, cp)
	if err != nil {
		return nil, err
	}
	fpi.p.SetNextPageURI(cp.NextPageURI)
	return cp, nil
}
