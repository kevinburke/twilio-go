package twilio

import (
	"context"
	"net/url"
)

const sipCredentialListsPathPart = "SIP/CredentialLists"
const sipCredentialsPathPart = "Credentials"

// SIPCredentialListService manages SIP CredentialList resources, lists of
// usernames and passwords used to authenticate SIP traffic.
//
// See https://www.twilio.com/docs/voice/sip/api/sip-credentiallist-resource for
// more.
type SIPCredentialListService struct {
	client *Client
}

// SIPCredentialList is a list of credentials used to authenticate SIP traffic.
type SIPCredentialList struct {
	Sid             string            `json:"sid"`
	AccountSid      string            `json:"account_sid"`
	FriendlyName    string            `json:"friendly_name"`
	DateCreated     TwilioTime        `json:"date_created"`
	DateUpdated     TwilioTime        `json:"date_updated"`
	SubresourceURIs map[string]string `json:"subresource_uris"`
	URI             string            `json:"uri"`
}

type SIPCredentialListPage struct {
	Page
	CredentialLists []*SIPCredentialList `json:"credential_lists"`
}

// Get retrieves the SIPCredentialList with the given sid.
func (s *SIPCredentialListService) Get(ctx context.Context, sid string) (*SIPCredentialList, error) {
	list := new(SIPCredentialList)
	err := s.client.GetResource(ctx, sipCredentialListsPathPart, sid, list)
	return list, err
}

// Create creates a new SIPCredentialList. The FriendlyName parameter is
// required.
func (s *SIPCredentialListService) Create(ctx context.Context, data url.Values) (*SIPCredentialList, error) {
	list := new(SIPCredentialList)
	err := s.client.CreateResource(ctx, sipCredentialListsPathPart, data, list)
	return list, err
}

// Update updates the SIPCredentialList with the given sid.
func (s *SIPCredentialListService) Update(ctx context.Context, sid string, data url.Values) (*SIPCredentialList, error) {
	list := new(SIPCredentialList)
	err := s.client.UpdateResource(ctx, sipCredentialListsPathPart, sid, data, list)
	return list, err
}

// Delete removes the SIPCredentialList with the given sid.
func (s *SIPCredentialListService) Delete(ctx context.Context, sid string) error {
	return s.client.DeleteResource(ctx, sipCredentialListsPathPart, sid)
}

// GetPage retrieves a page of SIPCredentialLists, filtered by the given data.
func (s *SIPCredentialListService) GetPage(ctx context.Context, data url.Values) (*SIPCredentialListPage, error) {
	return s.GetPageIterator(data).Next(ctx)
}

type SIPCredentialListPageIterator struct {
	p *PageIterator
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (s *SIPCredentialListService) GetPageIterator(data url.Values) *SIPCredentialListPageIterator {
	return &SIPCredentialListPageIterator{
		p: NewPageIterator(s.client, data, sipCredentialListsPathPart),
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (i *SIPCredentialListPageIterator) Next(ctx context.Context) (*SIPCredentialListPage, error) {
	cp := new(SIPCredentialListPage)
	err := i.p.Next(ctx, cp)
	if err != nil {
		return nil, err
	}
	i.p.SetNextPageURI(cp.NextPageURI)
	return cp, nil
}

// Credentials returns a service for managing the individual SIP Credentials
// (username/password pairs) that belong to the CredentialList with the given
// sid.
func (s *SIPCredentialListService) Credentials(credentialListSid string) *SIPCredentialService {
	return &SIPCredentialService{
		client:            s.client,
		credentialListSid: credentialListSid,
	}
}

// SIPCredentialService manages the individual username/password Credentials
// within a single SIP CredentialList.
//
// See https://www.twilio.com/docs/voice/sip/api/sip-credential-resource for
// more.
type SIPCredentialService struct {
	client            *Client
	credentialListSid string
}

// SIPCredential is a single username/password pair within a CredentialList. The
// password is never returned by the API.
type SIPCredential struct {
	Sid               string     `json:"sid"`
	AccountSid        string     `json:"account_sid"`
	CredentialListSid string     `json:"credential_list_sid"`
	Username          string     `json:"username"`
	DateCreated       TwilioTime `json:"date_created"`
	DateUpdated       TwilioTime `json:"date_updated"`
	URI               string     `json:"uri"`
}

type SIPCredentialPage struct {
	Page
	Credentials []*SIPCredential `json:"credentials"`
}

func (s *SIPCredentialService) pathPart() string {
	return sipCredentialListsPathPart + "/" + s.credentialListSid + "/" + sipCredentialsPathPart
}

// Get retrieves the SIPCredential with the given sid.
func (s *SIPCredentialService) Get(ctx context.Context, sid string) (*SIPCredential, error) {
	cred := new(SIPCredential)
	err := s.client.GetResource(ctx, s.pathPart(), sid, cred)
	return cred, err
}

// Create adds a new Credential to the CredentialList. The Username and Password
// parameters are required.
func (s *SIPCredentialService) Create(ctx context.Context, data url.Values) (*SIPCredential, error) {
	cred := new(SIPCredential)
	err := s.client.CreateResource(ctx, s.pathPart(), data, cred)
	return cred, err
}

// Update changes the password of the SIPCredential with the given sid. Set the
// Password parameter in data.
func (s *SIPCredentialService) Update(ctx context.Context, sid string, data url.Values) (*SIPCredential, error) {
	cred := new(SIPCredential)
	err := s.client.UpdateResource(ctx, s.pathPart(), sid, data, cred)
	return cred, err
}

// Delete removes the SIPCredential with the given sid.
func (s *SIPCredentialService) Delete(ctx context.Context, sid string) error {
	return s.client.DeleteResource(ctx, s.pathPart(), sid)
}

// GetPage retrieves a page of SIPCredentials, filtered by the given data.
func (s *SIPCredentialService) GetPage(ctx context.Context, data url.Values) (*SIPCredentialPage, error) {
	return s.GetPageIterator(data).Next(ctx)
}

type SIPCredentialPageIterator struct {
	p *PageIterator
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (s *SIPCredentialService) GetPageIterator(data url.Values) *SIPCredentialPageIterator {
	return &SIPCredentialPageIterator{
		p: NewPageIterator(s.client, data, s.pathPart()),
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (i *SIPCredentialPageIterator) Next(ctx context.Context) (*SIPCredentialPage, error) {
	cp := new(SIPCredentialPage)
	err := i.p.Next(ctx, cp)
	if err != nil {
		return nil, err
	}
	i.p.SetNextPageURI(cp.NextPageURI)
	return cp, nil
}
