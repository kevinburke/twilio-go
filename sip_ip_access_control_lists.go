package twilio

import (
	"context"
	"net/url"
)

const sipIPAccessControlListsPathPart = "SIP/IpAccessControlLists"
const sipIPAddressesPathPart = "IpAddresses"

// SIPIPAccessControlListService manages SIP IpAccessControlList resources, lists
// of IP addresses from which Twilio will accept SIP traffic.
//
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollist-resource
// for more.
type SIPIPAccessControlListService struct {
	client *Client
}

// SIPIPAccessControlList is a list of IP addresses that are permitted to send
// SIP traffic to your account.
type SIPIPAccessControlList struct {
	Sid             string            `json:"sid"`
	AccountSid      string            `json:"account_sid"`
	FriendlyName    string            `json:"friendly_name"`
	DateCreated     TwilioTime        `json:"date_created"`
	DateUpdated     TwilioTime        `json:"date_updated"`
	SubresourceURIs map[string]string `json:"subresource_uris"`
	URI             string            `json:"uri"`
}

type SIPIPAccessControlListPage struct {
	Page
	IPAccessControlLists []*SIPIPAccessControlList `json:"ip_access_control_lists"`
}

// Get retrieves the SIPIPAccessControlList with the given sid.
func (s *SIPIPAccessControlListService) Get(ctx context.Context, sid string) (*SIPIPAccessControlList, error) {
	list := new(SIPIPAccessControlList)
	err := s.client.GetResource(ctx, sipIPAccessControlListsPathPart, sid, list)
	return list, err
}

// Create creates a new SIPIPAccessControlList. The FriendlyName parameter is
// required.
func (s *SIPIPAccessControlListService) Create(ctx context.Context, data url.Values) (*SIPIPAccessControlList, error) {
	list := new(SIPIPAccessControlList)
	err := s.client.CreateResource(ctx, sipIPAccessControlListsPathPart, data, list)
	return list, err
}

// Update updates the SIPIPAccessControlList with the given sid.
func (s *SIPIPAccessControlListService) Update(ctx context.Context, sid string, data url.Values) (*SIPIPAccessControlList, error) {
	list := new(SIPIPAccessControlList)
	err := s.client.UpdateResource(ctx, sipIPAccessControlListsPathPart, sid, data, list)
	return list, err
}

// Delete removes the SIPIPAccessControlList with the given sid.
func (s *SIPIPAccessControlListService) Delete(ctx context.Context, sid string) error {
	return s.client.DeleteResource(ctx, sipIPAccessControlListsPathPart, sid)
}

// GetPage retrieves a page of SIPIPAccessControlLists, filtered by the given
// data.
func (s *SIPIPAccessControlListService) GetPage(ctx context.Context, data url.Values) (*SIPIPAccessControlListPage, error) {
	return s.GetPageIterator(data).Next(ctx)
}

type SIPIPAccessControlListPageIterator struct {
	p *PageIterator
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (s *SIPIPAccessControlListService) GetPageIterator(data url.Values) *SIPIPAccessControlListPageIterator {
	return &SIPIPAccessControlListPageIterator{
		p: NewPageIterator(s.client, data, sipIPAccessControlListsPathPart),
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (i *SIPIPAccessControlListPageIterator) Next(ctx context.Context) (*SIPIPAccessControlListPage, error) {
	lp := new(SIPIPAccessControlListPage)
	err := i.p.Next(ctx, lp)
	if err != nil {
		return nil, err
	}
	i.p.SetNextPageURI(lp.NextPageURI)
	return lp, nil
}

// IPAddresses returns a service for managing the individual IP addresses that
// belong to the IpAccessControlList with the given sid.
func (s *SIPIPAccessControlListService) IPAddresses(ipAccessControlListSid string) *SIPIPAddressService {
	return &SIPIPAddressService{
		client:                 s.client,
		ipAccessControlListSid: ipAccessControlListSid,
	}
}

// SIPIPAddressService manages the individual IP addresses within a single SIP
// IpAccessControlList.
//
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaddress-resource for more.
type SIPIPAddressService struct {
	client                 *Client
	ipAccessControlListSid string
}

// SIPIPAddress is a single IP address (or CIDR range) within an
// IpAccessControlList.
type SIPIPAddress struct {
	Sid                    string     `json:"sid"`
	AccountSid             string     `json:"account_sid"`
	FriendlyName           string     `json:"friendly_name"`
	IPAddress              string     `json:"ip_address"`
	CIDRPrefixLength       int        `json:"cidr_prefix_length"`
	IPAccessControlListSid string     `json:"ip_access_control_list_sid"`
	DateCreated            TwilioTime `json:"date_created"`
	DateUpdated            TwilioTime `json:"date_updated"`
	URI                    string     `json:"uri"`
}

type SIPIPAddressPage struct {
	Page
	IPAddresses []*SIPIPAddress `json:"ip_addresses"`
}

func (s *SIPIPAddressService) pathPart() string {
	return sipIPAccessControlListsPathPart + "/" + s.ipAccessControlListSid + "/" + sipIPAddressesPathPart
}

// Get retrieves the SIPIPAddress with the given sid.
func (s *SIPIPAddressService) Get(ctx context.Context, sid string) (*SIPIPAddress, error) {
	addr := new(SIPIPAddress)
	err := s.client.GetResource(ctx, s.pathPart(), sid, addr)
	return addr, err
}

// Create adds a new IP address to the IpAccessControlList. The FriendlyName and
// IpAddress parameters are required.
func (s *SIPIPAddressService) Create(ctx context.Context, data url.Values) (*SIPIPAddress, error) {
	addr := new(SIPIPAddress)
	err := s.client.CreateResource(ctx, s.pathPart(), data, addr)
	return addr, err
}

// Update updates the SIPIPAddress with the given sid.
func (s *SIPIPAddressService) Update(ctx context.Context, sid string, data url.Values) (*SIPIPAddress, error) {
	addr := new(SIPIPAddress)
	err := s.client.UpdateResource(ctx, s.pathPart(), sid, data, addr)
	return addr, err
}

// Delete removes the SIPIPAddress with the given sid.
func (s *SIPIPAddressService) Delete(ctx context.Context, sid string) error {
	return s.client.DeleteResource(ctx, s.pathPart(), sid)
}

// GetPage retrieves a page of SIPIPAddresses, filtered by the given data.
func (s *SIPIPAddressService) GetPage(ctx context.Context, data url.Values) (*SIPIPAddressPage, error) {
	return s.GetPageIterator(data).Next(ctx)
}

type SIPIPAddressPageIterator struct {
	p *PageIterator
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (s *SIPIPAddressService) GetPageIterator(data url.Values) *SIPIPAddressPageIterator {
	return &SIPIPAddressPageIterator{
		p: NewPageIterator(s.client, data, s.pathPart()),
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (i *SIPIPAddressPageIterator) Next(ctx context.Context) (*SIPIPAddressPage, error) {
	ap := new(SIPIPAddressPage)
	err := i.p.Next(ctx, ap)
	if err != nil {
		return nil, err
	}
	i.p.SetNextPageURI(ap.NextPageURI)
	return ap, nil
}
