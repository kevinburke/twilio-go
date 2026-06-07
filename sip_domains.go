package twilio

import (
	"context"
	"net/url"
)

const sipDomainsPathPart = "SIP/Domains"

// SIPDomainService manages SIP Domain resources. A SIP Domain is a custom SIP
// hostname (ending in sip.twilio.com) that accepts SIP traffic for your
// account; registered SIP endpoints and inbound calls are routed to its
// configured Voice URL.
//
// See https://www.twilio.com/docs/voice/sip/api/sip-domain-resource for more.
type SIPDomainService struct {
	client *Client
}

// SIPDomain is a custom SIP hostname reserved on Twilio.
type SIPDomain struct {
	Sid                       string            `json:"sid"`
	AccountSid                string            `json:"account_sid"`
	APIVersion                string            `json:"api_version"`
	AuthType                  string            `json:"auth_type"`
	DomainName                string            `json:"domain_name"`
	FriendlyName              string            `json:"friendly_name"`
	SIPRegistration           bool              `json:"sip_registration"`
	EmergencyCallingEnabled   bool              `json:"emergency_calling_enabled"`
	Secure                    bool              `json:"secure"`
	BYOCTrunkSid              string            `json:"byoc_trunk_sid"`
	EmergencyCallerSid        string            `json:"emergency_caller_sid"`
	VoiceURL                  string            `json:"voice_url"`
	VoiceMethod               string            `json:"voice_method"`
	VoiceFallbackURL          string            `json:"voice_fallback_url"`
	VoiceFallbackMethod       string            `json:"voice_fallback_method"`
	VoiceStatusCallbackURL    string            `json:"voice_status_callback_url"`
	VoiceStatusCallbackMethod string            `json:"voice_status_callback_method"`
	DateCreated               TwilioTime        `json:"date_created"`
	DateUpdated               TwilioTime        `json:"date_updated"`
	SubresourceURIs           map[string]string `json:"subresource_uris"`
	URI                       string            `json:"uri"`
}

type SIPDomainPage struct {
	Page
	Domains []*SIPDomain `json:"domains"`
}

// Get retrieves the SIPDomain with the given sid.
func (s *SIPDomainService) Get(ctx context.Context, sid string) (*SIPDomain, error) {
	domain := new(SIPDomain)
	err := s.client.GetResource(ctx, sipDomainsPathPart, sid, domain)
	return domain, err
}

// Create creates a new SIPDomain. The DomainName parameter is required and must
// end with "sip.twilio.com".
func (s *SIPDomainService) Create(ctx context.Context, data url.Values) (*SIPDomain, error) {
	domain := new(SIPDomain)
	err := s.client.CreateResource(ctx, sipDomainsPathPart, data, domain)
	return domain, err
}

// Update updates the SIPDomain with the given sid.
func (s *SIPDomainService) Update(ctx context.Context, sid string, data url.Values) (*SIPDomain, error) {
	domain := new(SIPDomain)
	err := s.client.UpdateResource(ctx, sipDomainsPathPart, sid, data, domain)
	return domain, err
}

// Delete removes the SIPDomain with the given sid.
func (s *SIPDomainService) Delete(ctx context.Context, sid string) error {
	return s.client.DeleteResource(ctx, sipDomainsPathPart, sid)
}

// GetPage retrieves a page of SIPDomains, filtered by the given data.
func (s *SIPDomainService) GetPage(ctx context.Context, data url.Values) (*SIPDomainPage, error) {
	return s.GetPageIterator(data).Next(ctx)
}

type SIPDomainPageIterator struct {
	p *PageIterator
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (s *SIPDomainService) GetPageIterator(data url.Values) *SIPDomainPageIterator {
	return &SIPDomainPageIterator{
		p: NewPageIterator(s.client, data, sipDomainsPathPart),
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (i *SIPDomainPageIterator) Next(ctx context.Context) (*SIPDomainPage, error) {
	dp := new(SIPDomainPage)
	err := i.p.Next(ctx, dp)
	if err != nil {
		return nil, err
	}
	i.p.SetNextPageURI(dp.NextPageURI)
	return dp, nil
}

// CredentialListMappings returns a service for managing the mappings between
// the SIP Domain with the given sid and the CredentialLists used to
// authenticate calls to that domain.
func (s *SIPDomainService) CredentialListMappings(domainSid string) *SIPDomainCredentialListMappingService {
	return &SIPDomainCredentialListMappingService{client: s.client, domainSid: domainSid}
}

// IPAccessControlListMappings returns a service for managing the mappings
// between the SIP Domain with the given sid and the IpAccessControlLists used
// to authenticate calls to that domain.
func (s *SIPDomainService) IPAccessControlListMappings(domainSid string) *SIPDomainIPAccessControlListMappingService {
	return &SIPDomainIPAccessControlListMappingService{client: s.client, domainSid: domainSid}
}

// AuthCallsCredentialListMappings returns a service for managing the
// CredentialList mappings that authenticate inbound calls (the "Calls" auth
// type) on the SIP Domain with the given sid.
func (s *SIPDomainService) AuthCallsCredentialListMappings(domainSid string) *SIPAuthCallsCredentialListMappingService {
	return &SIPAuthCallsCredentialListMappingService{client: s.client, domainSid: domainSid}
}

// AuthCallsIPAccessControlListMappings returns a service for managing the
// IpAccessControlList mappings that authenticate inbound calls (the "Calls"
// auth type) on the SIP Domain with the given sid.
func (s *SIPDomainService) AuthCallsIPAccessControlListMappings(domainSid string) *SIPAuthCallsIPAccessControlListMappingService {
	return &SIPAuthCallsIPAccessControlListMappingService{client: s.client, domainSid: domainSid}
}

// AuthRegistrationsCredentialListMappings returns a service for managing the
// CredentialList mappings that authenticate SIP registrations on the SIP Domain
// with the given sid.
func (s *SIPDomainService) AuthRegistrationsCredentialListMappings(domainSid string) *SIPAuthRegistrationsCredentialListMappingService {
	return &SIPAuthRegistrationsCredentialListMappingService{client: s.client, domainSid: domainSid}
}
