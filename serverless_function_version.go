package twilio

import (
	"context"
	"net/url"
	"path/filepath"
)

const functionVersionPathPart = "Versions"

type FunctionVersionService struct {
	client      *Client
	serviceSid  string
	functionSid string
}

type FunctionVersion struct {
	Sid         string     `json:"sid"`
	AccountSid  string     `json:"account_sid"`
	ServiceSid  string     `json:"service_sid"`
	FunctionSid string     `json:"function_sid"`
	Path        string     `json:"path"`
	Visibility  string     `json:"visibility"`
	DateCreated TwilioTime `json:"date_created"`
	URL         string     `json:"url"`
}

type FunctionVersionContent struct {
	Sid         string `json:"sid"`
	AccountSid  string `json:"account_sid"`
	ServiceSid  string `json:"service_sid"`
	FunctionSid string `json:"function_sid"`
	Content     string `json:"content"`
	URL         string `json:"url"`
}

type FunctionVersionPage struct {
	Page
	FunctionVersions []*FunctionVersion `json:"function_versions"`
}

// Get retrieves a Function Version by its sid.
//
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/function for
// more.
func (f *FunctionVersionService) Get(ctx context.Context, sid string) (*FunctionVersion, error) {
	fv := new(FunctionVersion)
	pathPart := filepath.Join(servicesPathPrefix, f.serviceSid, functionPathPart, f.functionSid, functionVersionPathPart)
	err := f.client.GetResource(ctx, pathPart, sid, fv)
	return fv, err
}

// Get retrieves a Function Version Content by its sid.
//
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/function-version/function-version-content
// for more.
func (f *FunctionVersionService) GetContent(ctx context.Context, sid string) (*FunctionVersionContent, error) {
	fvc := new(FunctionVersionContent)
	pathPart := filepath.Join(servicesPathPrefix, f.serviceSid, functionPathPart, f.functionSid, functionVersionPathPart)
	err := f.client.GetResource(ctx, pathPart, sid+"/Content", fvc)
	return fvc, err
}

// GetPage retrieves a FunctionVersionPage, filtered by the given data.
func (f *FunctionVersionService) GetPage(ctx context.Context, data url.Values) (*FunctionVersionPage, error) {
	iter := f.GetPageIterator(data)
	return iter.Next(ctx)
}

type FunctionVersionPageIterator struct {
	p *PageIterator
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (f *FunctionVersionService) GetPageIterator(data url.Values) *FunctionVersionPageIterator {
	pathPart := filepath.Join(servicesPathPrefix, f.serviceSid, functionPathPart, f.functionSid, functionVersionPathPart)
	iter := NewPageIterator(f.client, data, pathPart)
	return &FunctionVersionPageIterator{
		p: iter,
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (fpi *FunctionVersionPageIterator) Next(ctx context.Context) (*FunctionVersionPage, error) {
	cp := new(FunctionVersionPage)
	err := fpi.p.Next(ctx, cp)
	if err != nil {
		return nil, err
	}
	fpi.p.SetNextPageURI(cp.NextPageURI)
	return cp, nil
}
