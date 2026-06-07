package twilio

import (
	"context"
	"net/url"
)

// SIPDomainCredentialListMapping maps a CredentialList to a SIP Domain so that
// the credentials in that list can authenticate traffic to the domain.
type SIPDomainCredentialListMapping struct {
	Sid          string     `json:"sid"`
	AccountSid   string     `json:"account_sid"`
	DomainSid    string     `json:"domain_sid"`
	FriendlyName string     `json:"friendly_name"`
	DateCreated  TwilioTime `json:"date_created"`
	DateUpdated  TwilioTime `json:"date_updated"`
	URI          string     `json:"uri"`
}

type SIPDomainCredentialListMappingPage struct {
	Page
	CredentialListMappings []*SIPDomainCredentialListMapping `json:"credential_list_mappings"`
}

// SIPDomainCredentialListMappingService manages the CredentialLists mapped to a
// single SIP Domain.
//
// See https://www.twilio.com/docs/voice/sip/api/sip-credentiallistmapping-resource
// for more.
type SIPDomainCredentialListMappingService struct {
	client    *Client
	domainSid string
}

func (s *SIPDomainCredentialListMappingService) pathPart() string {
	return sipDomainsPathPart + "/" + s.domainSid + "/CredentialListMappings"
}

// Get retrieves the mapping with the given sid.
func (s *SIPDomainCredentialListMappingService) Get(ctx context.Context, sid string) (*SIPDomainCredentialListMapping, error) {
	m := new(SIPDomainCredentialListMapping)
	err := s.client.GetResource(ctx, s.pathPart(), sid, m)
	return m, err
}

// Create maps a CredentialList to the SIP Domain. The CredentialListSid
// parameter is required.
func (s *SIPDomainCredentialListMappingService) Create(ctx context.Context, data url.Values) (*SIPDomainCredentialListMapping, error) {
	m := new(SIPDomainCredentialListMapping)
	err := s.client.CreateResource(ctx, s.pathPart(), data, m)
	return m, err
}

// Delete removes the mapping with the given sid.
func (s *SIPDomainCredentialListMappingService) Delete(ctx context.Context, sid string) error {
	return s.client.DeleteResource(ctx, s.pathPart(), sid)
}

// GetPage retrieves a page of mappings, filtered by the given data.
func (s *SIPDomainCredentialListMappingService) GetPage(ctx context.Context, data url.Values) (*SIPDomainCredentialListMappingPage, error) {
	return s.GetPageIterator(data).Next(ctx)
}

type SIPDomainCredentialListMappingPageIterator struct {
	p *PageIterator
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (s *SIPDomainCredentialListMappingService) GetPageIterator(data url.Values) *SIPDomainCredentialListMappingPageIterator {
	return &SIPDomainCredentialListMappingPageIterator{
		p: NewPageIterator(s.client, data, s.pathPart()),
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (i *SIPDomainCredentialListMappingPageIterator) Next(ctx context.Context) (*SIPDomainCredentialListMappingPage, error) {
	mp := new(SIPDomainCredentialListMappingPage)
	err := i.p.Next(ctx, mp)
	if err != nil {
		return nil, err
	}
	i.p.SetNextPageURI(mp.NextPageURI)
	return mp, nil
}

// SIPDomainIPAccessControlListMapping maps an IpAccessControlList to a SIP
// Domain so that the IP addresses in that list can authenticate traffic to the
// domain.
type SIPDomainIPAccessControlListMapping struct {
	Sid          string     `json:"sid"`
	AccountSid   string     `json:"account_sid"`
	DomainSid    string     `json:"domain_sid"`
	FriendlyName string     `json:"friendly_name"`
	DateCreated  TwilioTime `json:"date_created"`
	DateUpdated  TwilioTime `json:"date_updated"`
	URI          string     `json:"uri"`
}

type SIPDomainIPAccessControlListMappingPage struct {
	Page
	IPAccessControlListMappings []*SIPDomainIPAccessControlListMapping `json:"ip_access_control_list_mappings"`
}

// SIPDomainIPAccessControlListMappingService manages the IpAccessControlLists
// mapped to a single SIP Domain.
//
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollistmapping-resource
// for more.
type SIPDomainIPAccessControlListMappingService struct {
	client    *Client
	domainSid string
}

func (s *SIPDomainIPAccessControlListMappingService) pathPart() string {
	return sipDomainsPathPart + "/" + s.domainSid + "/IpAccessControlListMappings"
}

// Get retrieves the mapping with the given sid.
func (s *SIPDomainIPAccessControlListMappingService) Get(ctx context.Context, sid string) (*SIPDomainIPAccessControlListMapping, error) {
	m := new(SIPDomainIPAccessControlListMapping)
	err := s.client.GetResource(ctx, s.pathPart(), sid, m)
	return m, err
}

// Create maps an IpAccessControlList to the SIP Domain. The
// IpAccessControlListSid parameter is required.
func (s *SIPDomainIPAccessControlListMappingService) Create(ctx context.Context, data url.Values) (*SIPDomainIPAccessControlListMapping, error) {
	m := new(SIPDomainIPAccessControlListMapping)
	err := s.client.CreateResource(ctx, s.pathPart(), data, m)
	return m, err
}

// Delete removes the mapping with the given sid.
func (s *SIPDomainIPAccessControlListMappingService) Delete(ctx context.Context, sid string) error {
	return s.client.DeleteResource(ctx, s.pathPart(), sid)
}

// GetPage retrieves a page of mappings, filtered by the given data.
func (s *SIPDomainIPAccessControlListMappingService) GetPage(ctx context.Context, data url.Values) (*SIPDomainIPAccessControlListMappingPage, error) {
	return s.GetPageIterator(data).Next(ctx)
}

type SIPDomainIPAccessControlListMappingPageIterator struct {
	p *PageIterator
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (s *SIPDomainIPAccessControlListMappingService) GetPageIterator(data url.Values) *SIPDomainIPAccessControlListMappingPageIterator {
	return &SIPDomainIPAccessControlListMappingPageIterator{
		p: NewPageIterator(s.client, data, s.pathPart()),
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (i *SIPDomainIPAccessControlListMappingPageIterator) Next(ctx context.Context) (*SIPDomainIPAccessControlListMappingPage, error) {
	mp := new(SIPDomainIPAccessControlListMappingPage)
	err := i.p.Next(ctx, mp)
	if err != nil {
		return nil, err
	}
	i.p.SetNextPageURI(mp.NextPageURI)
	return mp, nil
}

// SIPAuthMapping is a credential-list or IP-access-control-list mapping under a
// SIP Domain's Auth sub-resource. The three Auth mapping resources (calls
// credential list, calls IP access control list, and registrations credential
// list) share this representation.
type SIPAuthMapping struct {
	Sid          string     `json:"sid"`
	AccountSid   string     `json:"account_sid"`
	FriendlyName string     `json:"friendly_name"`
	DateCreated  TwilioTime `json:"date_created"`
	DateUpdated  TwilioTime `json:"date_updated"`
}

// SIPAuthMappingPage is a page of Auth mappings. The API returns these under the
// "contents" key for all three Auth mapping resources.
type SIPAuthMappingPage struct {
	Page
	Contents []*SIPAuthMapping `json:"contents"`
}

type SIPAuthMappingPageIterator struct {
	p *PageIterator
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (i *SIPAuthMappingPageIterator) Next(ctx context.Context) (*SIPAuthMappingPage, error) {
	mp := new(SIPAuthMappingPage)
	err := i.p.Next(ctx, mp)
	if err != nil {
		return nil, err
	}
	i.p.SetNextPageURI(mp.NextPageURI)
	return mp, nil
}

// SIPAuthCallsCredentialListMappingService manages the CredentialList mappings
// that authenticate inbound calls on a SIP Domain.
//
// See https://www.twilio.com/docs/voice/sip/api/sip-domain-auth-calls-credentiallistmapping
// for more.
type SIPAuthCallsCredentialListMappingService struct {
	client    *Client
	domainSid string
}

func (s *SIPAuthCallsCredentialListMappingService) pathPart() string {
	return sipDomainsPathPart + "/" + s.domainSid + "/Auth/Calls/CredentialListMappings"
}

// Get retrieves the mapping with the given sid.
func (s *SIPAuthCallsCredentialListMappingService) Get(ctx context.Context, sid string) (*SIPAuthMapping, error) {
	m := new(SIPAuthMapping)
	err := s.client.GetResource(ctx, s.pathPart(), sid, m)
	return m, err
}

// Create maps a CredentialList to the SIP Domain for call authentication. The
// CredentialListSid parameter is required.
func (s *SIPAuthCallsCredentialListMappingService) Create(ctx context.Context, data url.Values) (*SIPAuthMapping, error) {
	m := new(SIPAuthMapping)
	err := s.client.CreateResource(ctx, s.pathPart(), data, m)
	return m, err
}

// Delete removes the mapping with the given sid.
func (s *SIPAuthCallsCredentialListMappingService) Delete(ctx context.Context, sid string) error {
	return s.client.DeleteResource(ctx, s.pathPart(), sid)
}

// GetPage retrieves a page of mappings, filtered by the given data.
func (s *SIPAuthCallsCredentialListMappingService) GetPage(ctx context.Context, data url.Values) (*SIPAuthMappingPage, error) {
	return s.GetPageIterator(data).Next(ctx)
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (s *SIPAuthCallsCredentialListMappingService) GetPageIterator(data url.Values) *SIPAuthMappingPageIterator {
	return &SIPAuthMappingPageIterator{
		p: NewPageIterator(s.client, data, s.pathPart()),
	}
}

// SIPAuthCallsIPAccessControlListMappingService manages the IpAccessControlList
// mappings that authenticate inbound calls on a SIP Domain.
//
// See https://www.twilio.com/docs/voice/sip/api/sip-domain-auth-calls-ipaccesscontrollistmapping
// for more.
type SIPAuthCallsIPAccessControlListMappingService struct {
	client    *Client
	domainSid string
}

func (s *SIPAuthCallsIPAccessControlListMappingService) pathPart() string {
	return sipDomainsPathPart + "/" + s.domainSid + "/Auth/Calls/IpAccessControlListMappings"
}

// Get retrieves the mapping with the given sid.
func (s *SIPAuthCallsIPAccessControlListMappingService) Get(ctx context.Context, sid string) (*SIPAuthMapping, error) {
	m := new(SIPAuthMapping)
	err := s.client.GetResource(ctx, s.pathPart(), sid, m)
	return m, err
}

// Create maps an IpAccessControlList to the SIP Domain for call authentication.
// The IpAccessControlListSid parameter is required.
func (s *SIPAuthCallsIPAccessControlListMappingService) Create(ctx context.Context, data url.Values) (*SIPAuthMapping, error) {
	m := new(SIPAuthMapping)
	err := s.client.CreateResource(ctx, s.pathPart(), data, m)
	return m, err
}

// Delete removes the mapping with the given sid.
func (s *SIPAuthCallsIPAccessControlListMappingService) Delete(ctx context.Context, sid string) error {
	return s.client.DeleteResource(ctx, s.pathPart(), sid)
}

// GetPage retrieves a page of mappings, filtered by the given data.
func (s *SIPAuthCallsIPAccessControlListMappingService) GetPage(ctx context.Context, data url.Values) (*SIPAuthMappingPage, error) {
	return s.GetPageIterator(data).Next(ctx)
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (s *SIPAuthCallsIPAccessControlListMappingService) GetPageIterator(data url.Values) *SIPAuthMappingPageIterator {
	return &SIPAuthMappingPageIterator{
		p: NewPageIterator(s.client, data, s.pathPart()),
	}
}

// SIPAuthRegistrationsCredentialListMappingService manages the CredentialList
// mappings that authenticate SIP registrations on a SIP Domain. This is the
// mapping you create to let a SIP endpoint register with the domain.
//
// See https://www.twilio.com/docs/voice/sip/api/sip-domain-auth-registrations-credentiallistmapping
// for more.
type SIPAuthRegistrationsCredentialListMappingService struct {
	client    *Client
	domainSid string
}

func (s *SIPAuthRegistrationsCredentialListMappingService) pathPart() string {
	return sipDomainsPathPart + "/" + s.domainSid + "/Auth/Registrations/CredentialListMappings"
}

// Get retrieves the mapping with the given sid.
func (s *SIPAuthRegistrationsCredentialListMappingService) Get(ctx context.Context, sid string) (*SIPAuthMapping, error) {
	m := new(SIPAuthMapping)
	err := s.client.GetResource(ctx, s.pathPart(), sid, m)
	return m, err
}

// Create maps a CredentialList to the SIP Domain for registration
// authentication. The CredentialListSid parameter is required.
func (s *SIPAuthRegistrationsCredentialListMappingService) Create(ctx context.Context, data url.Values) (*SIPAuthMapping, error) {
	m := new(SIPAuthMapping)
	err := s.client.CreateResource(ctx, s.pathPart(), data, m)
	return m, err
}

// Delete removes the mapping with the given sid.
func (s *SIPAuthRegistrationsCredentialListMappingService) Delete(ctx context.Context, sid string) error {
	return s.client.DeleteResource(ctx, s.pathPart(), sid)
}

// GetPage retrieves a page of mappings, filtered by the given data.
func (s *SIPAuthRegistrationsCredentialListMappingService) GetPage(ctx context.Context, data url.Values) (*SIPAuthMappingPage, error) {
	return s.GetPageIterator(data).Next(ctx)
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (s *SIPAuthRegistrationsCredentialListMappingService) GetPageIterator(data url.Values) *SIPAuthMappingPageIterator {
	return &SIPAuthMappingPageIterator{
		p: NewPageIterator(s.client, data, s.pathPart()),
	}
}
