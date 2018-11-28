package twilio

import (
	"context"
	"net/url"
	"strings"
)

const voicePathPart = "Voice"

type VoicePriceService struct {
	Countries *CountryVoicePriceService
	Numbers   *NumberVoicePriceService
}

type CountryVoicePriceService struct {
	client *Client
}

type NumberVoicePriceService struct {
	client *Client
}

type PrefixPrice struct {
	BasePrice           string   `json:"base_price"`
	CurrentPrice        string   `json:"current_price"`
	FriendlyName        string   `json:"friendly_name"`
	DestinationPrefixes []string `json:"destination_prefixes"`
	OriginationPrefixes []string `json:"origination_prefixes"`
}

type InboundPrice struct {
	BasePrice    string `json:"base_price"`
	CurrentPrice string `json:"current_price"`
	NumberType   string `json:"number_type"`
}

type OutboundCallPrice struct {
	BasePrice           string   `json:"base_price"`
	CurrentPrice        string   `json:"current_price"`
	OriginationPrefixes []string `json:"origination_prefixes"`
}

type VoicePrice struct {
	Country              string         `json:"country"`
	IsoCountry           string         `json:"iso_country"`
	OutboundPrefixPrices []PrefixPrice  `json:"outbound_prefix_prices"`
	InboundCallPrices    []InboundPrice `json:"inbound_call_prices"`
	PriceUnit            string         `json:"price_unit"`
	URL                  string         `json:"url"`
}

type VoiceNumberPrice struct {
	Country            string              `json:"country"`
	IsoCountry         string              `json:"iso_country"`
	DestinationNumber  string              `json:"destination_number"`
	OriginationNumber  string              `json:"origination_number"`
	InboundCallPrice   InboundPrice        `json:"inbound_call_price"`
	OutboundCallPrices []OutboundCallPrice `json:"outbound_call_prices"`
	PriceUnit          string              `json:"price_unit"`

	URL string `json:"url"`
}

// https://www.twilio.com/docs/voice/pricing#pricing-voice-country-instance-resource
// returns the call price by country
func (cvps *CountryVoicePriceService) Get(ctx context.Context, isoCountry string) (*VoicePrice, error) {
	voicePrice := new(VoicePrice)
	err := cvps.client.GetResource(ctx, voicePathPart+"/Countries", isoCountry, voicePrice)
	return voicePrice, err
}

// https://www.twilio.com/docs/voice/pricing#list-uri-by-origination-number
// returns the call price by number or numbers pair
func (nvps *NumberVoicePriceService) Get(ctx context.Context, destinationNumber string, data url.Values) (*VoiceNumberPrice, error) {
	voiceNumPrice := new(VoiceNumberPrice)
	pathPart := strings.Join([]string{voicePathPart, "Numbers", destinationNumber}, "/")
	err := nvps.client.ListResource(ctx, pathPart, data, voiceNumPrice)
	return voiceNumPrice, err
}

// returns a list of countries where Twilio voice services are available and the corresponding URL
// for retrieving the country specific voice prices.
func (cvps *CountryVoicePriceService) GetPage(ctx context.Context, data url.Values) (*CountriesPricePage, error) {
	return cvps.GetPageIterator(data).Next(ctx)
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (cvps *CountryVoicePriceService) GetPageIterator(data url.Values) *CountryPricePageIterator {
	iter := NewPageIterator(cvps.client, data, voicePathPart+"/Countries")
	return &CountryPricePageIterator{
		p: iter,
	}
}
