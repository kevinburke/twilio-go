package twilio

import (
	"net/url"

	"golang.org/x/net/context"
)

const availableNumbersPath = "AvailablePhoneNumbers"

type AvailableNumberBase struct {
	client   *Client
	pathPart string
}

type AvailableNumberService struct {
	Local    *AvailableNumberBase
	Mobile   *AvailableNumberBase
	TollFree *AvailableNumberBase
}

// The subresources of the AvailableNumbers resource let you search for local, toll-free and
// mobile phone numbers that are available for you to purchase.
// See https://www.twilio.com/docs/api/rest/available-phone-numbers for details
type AvailableNumber struct {
	FriendlyName        string            `json:"friendly_name"`
	PhoneNumber         PhoneNumber       `json:"phone_number"`
	Lata                string            `json:"lata"`
	RateCenter          string            `json:"rate_center"`
	Latitude            string            `json:"latitude"`
	Longitude           string            `json:"longitude"`
	Region              string            `json:"region"`
	PostalCode          string            `json:"postal_code"`
	ISOCountry          string            `json:"iso_country"`
	Capabilities        *NumberCapability `json:"capabilities"`
	AddressRequirements string            `json:"address_requirements"`
	Beta                bool              `json:"beta"`
}

type AvailableNumberPage struct {
	Uri     string             `json:"uri"`
	Numbers []*AvailableNumber `json:"available_phone_numbers"`
}

// GetPage returns a page of available phone numbers.
//
// For more information, see the Twilio documentation:
// https://www.twilio.com/docs/api/rest/available-phone-numbers#local
// https://www.twilio.com/docs/api/rest/available-phone-numbers#toll-free
// https://www.twilio.com/docs/api/rest/available-phone-numbers#mobile
func (s *AvailableNumberBase) GetPage(ctx context.Context, isoCountry string, filters url.Values) (*AvailableNumberPage, error) {
	sr := new(AvailableNumberPage)
	path := availableNumbersPath + "/" + isoCountry + "/" + s.pathPart
	err := s.client.ListResource(ctx, path, filters, sr)
	if err != nil {
		return nil, err
	}

	return sr, nil
}
