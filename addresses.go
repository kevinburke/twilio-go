package twilio

import (
	"context"
	"net/url"
)

const addressesPathPart = "Addresses"

type AddressService struct {
	client *Client
}

// Address is a customer's or business's physical mailing address, used to
// satisfy local regulatory requirements for some phone numbers and for
// emergency calling.
//
// See https://www.twilio.com/docs/usage/api/address for more.
type Address struct {
	Sid              string     `json:"sid"`
	AccountSid       string     `json:"account_sid"`
	FriendlyName     string     `json:"friendly_name"`
	CustomerName     string     `json:"customer_name"`
	Street           string     `json:"street"`
	StreetSecondary  string     `json:"street_secondary"`
	City             string     `json:"city"`
	Region           string     `json:"region"`
	PostalCode       string     `json:"postal_code"`
	IsoCountry       string     `json:"iso_country"`
	EmergencyEnabled bool       `json:"emergency_enabled"`
	Validated        bool       `json:"validated"`
	Verified         bool       `json:"verified"`
	DateCreated      TwilioTime `json:"date_created"`
	DateUpdated      TwilioTime `json:"date_updated"`
	URI              string     `json:"uri"`
}

type AddressPage struct {
	Page
	Addresses []*Address `json:"addresses"`
}

// Create a new Address. The CustomerName, Street, City, Region, PostalCode, and
// IsoCountry parameters are required. Valid parameters may be found here:
// https://www.twilio.com/docs/usage/api/address#create-an-address-resource
func (a *AddressService) Create(ctx context.Context, data url.Values) (*Address, error) {
	addr := new(Address)
	err := a.client.CreateResource(ctx, addressesPathPart, data, addr)
	return addr, err
}

// Get retrieves a single Address.
func (a *AddressService) Get(ctx context.Context, sid string) (*Address, error) {
	addr := new(Address)
	err := a.client.GetResource(ctx, addressesPathPart, sid, addr)
	return addr, err
}

// GetPage retrieves an AddressPage, filtered by the given data.
func (a *AddressService) GetPage(ctx context.Context, data url.Values) (*AddressPage, error) {
	ap := new(AddressPage)
	err := a.client.ListResource(ctx, addressesPathPart, data, ap)
	return ap, err
}

// Update the Address with the given sid. Valid parameters may be found here:
// https://www.twilio.com/docs/usage/api/address#update-an-address-resource
func (a *AddressService) Update(ctx context.Context, sid string, data url.Values) (*Address, error) {
	addr := new(Address)
	err := a.client.UpdateResource(ctx, addressesPathPart, sid, data, addr)
	return addr, err
}

// Delete the Address with the given sid. If the Address has already been
// deleted, or does not exist, Delete returns nil. If another error or a timeout
// occurs, the error is returned.
func (a *AddressService) Delete(ctx context.Context, sid string) error {
	return a.client.DeleteResource(ctx, addressesPathPart, sid)
}

// AddressPageIterator lets you retrieve consecutive pages of resources.
type AddressPageIterator struct {
	p *PageIterator
}

// GetPageIterator returns an AddressPageIterator with the given page filters.
// Call iterator.Next() to get the first page of resources (and again to
// retrieve subsequent pages).
func (a *AddressService) GetPageIterator(data url.Values) *AddressPageIterator {
	iter := NewPageIterator(a.client, data, addressesPathPart)
	return &AddressPageIterator{
		p: iter,
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (a *AddressPageIterator) Next(ctx context.Context) (*AddressPage, error) {
	ap := new(AddressPage)
	err := a.p.Next(ctx, ap)
	if err != nil {
		return nil, err
	}
	a.p.SetNextPageURI(ap.NextPageURI)
	return ap, nil
}

// DependentPhoneNumber is a phone number whose regulatory requirements are
// satisfied by an Address. Deleting an Address that still has dependent phone
// numbers is not permitted.
//
// See https://www.twilio.com/docs/usage/api/address#dependentphonenumber for
// more.
type DependentPhoneNumber struct {
	Sid                  string            `json:"sid"`
	AccountSid           string            `json:"account_sid"`
	FriendlyName         string            `json:"friendly_name"`
	PhoneNumber          PhoneNumber       `json:"phone_number"`
	APIVersion           string            `json:"api_version"`
	AddressRequirements  string            `json:"address_requirements"`
	Capabilities         *NumberCapability `json:"capabilities"`
	EmergencyAddressSid  string            `json:"emergency_address_sid"`
	EmergencyStatus      string            `json:"emergency_status"`
	SMSApplicationSid    string            `json:"sms_application_sid"`
	SMSFallbackMethod    string            `json:"sms_fallback_method"`
	SMSFallbackURL       string            `json:"sms_fallback_url"`
	SMSMethod            string            `json:"sms_method"`
	SMSURL               string            `json:"sms_url"`
	StatusCallback       string            `json:"status_callback"`
	StatusCallbackMethod string            `json:"status_callback_method"`
	TrunkSid             string            `json:"trunk_sid"`
	VoiceApplicationSid  string            `json:"voice_application_sid"`
	VoiceCallerIDLookup  bool              `json:"voice_caller_id_lookup"`
	VoiceFallbackMethod  string            `json:"voice_fallback_method"`
	VoiceFallbackURL     string            `json:"voice_fallback_url"`
	VoiceMethod          string            `json:"voice_method"`
	VoiceURL             string            `json:"voice_url"`
	DateCreated          TwilioTime        `json:"date_created"`
	DateUpdated          TwilioTime        `json:"date_updated"`
	URI                  string            `json:"uri"`
}

type DependentPhoneNumberPage struct {
	Page
	DependentPhoneNumbers []*DependentPhoneNumber `json:"dependent_phone_numbers"`
}

func dependentPhoneNumbersPathPart(addressSid string) string {
	return addressesPathPart + "/" + addressSid + "/DependentPhoneNumbers"
}

// GetDependentPhoneNumbers returns a single page of phone numbers that depend
// on the given Address. To retrieve multiple pages, use
// GetDependentPhoneNumbersIterator.
func (a *AddressService) GetDependentPhoneNumbers(ctx context.Context, addressSid string, data url.Values) (*DependentPhoneNumberPage, error) {
	return a.GetDependentPhoneNumbersIterator(addressSid, data).Next(ctx)
}

// DependentPhoneNumberPageIterator lets you retrieve consecutive pages of
// resources.
type DependentPhoneNumberPageIterator struct {
	p *PageIterator
}

// GetDependentPhoneNumbersIterator returns an iterator over the phone numbers
// that depend on the given Address.
func (a *AddressService) GetDependentPhoneNumbersIterator(addressSid string, data url.Values) *DependentPhoneNumberPageIterator {
	return &DependentPhoneNumberPageIterator{
		p: NewPageIterator(a.client, data, dependentPhoneNumbersPathPart(addressSid)),
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (d *DependentPhoneNumberPageIterator) Next(ctx context.Context) (*DependentPhoneNumberPage, error) {
	dp := new(DependentPhoneNumberPage)
	err := d.p.Next(ctx, dp)
	if err != nil {
		return nil, err
	}
	d.p.SetNextPageURI(dp.NextPageURI)
	return dp, nil
}
