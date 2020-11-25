package twilio

import (
	"context"
	"net/url"
	"testing"
)

func TestGetVoicePriceUS(t *testing.T) {
	t.Parallel()
	client, server := getServer(voicePriceUS)
	defer server.Close()

	isoCountry := "US"
	expectedCountryName := "United States"
	expectedPriceUnit := "USD"

	voicePrice, err := client.Pricing.Voice.Countries.Get(context.Background(), isoCountry, nil)
	if err != nil {
		t.Fatal(err)
	}
	if voicePrice == nil {
		t.Fatal("expected voice price to be returned")
	}
	if voicePrice.Country != expectedCountryName {
		t.Errorf("Expected voice price country to be %s, but got %s\n", expectedCountryName, voicePrice.Country)
	}
	if voicePrice.IsoCountry != isoCountry {
		t.Errorf("Expected voice price iso country to be %s, but got %s\n", isoCountry, voicePrice.IsoCountry)
	}
	if voicePrice.PriceUnit != expectedPriceUnit {
		t.Errorf("Expected voice price unit to be %s, but got %s\n", expectedPriceUnit, voicePrice.PriceUnit)
	}
	if voicePrice.InboundCallPrices == nil {
		t.Error("Expected voice price to contain InboundCallPrices")
	}
	if voicePrice.OutboundPrefixPrices == nil {
		t.Error("Expected voice price to contain OutboundPrefixPrices")
	}

	inboundPriceMap := make(map[string]bool)
	for _, inPrice := range voicePrice.InboundCallPrices {
		inboundPriceMap[inPrice.NumberType] = true
	}
	// inboundPriceMap => map[local:true toll free:true]
	_, ok := inboundPriceMap["local"]
	if ok == false {
		t.Error("Expected inbound price to contain a price for local calls")
	}

	outboundDestinationPrefixPriceMap := make(map[string]bool)
	outboundOriginationPrefixPriceMap := make(map[string]bool)
	for _, outPrice := range voicePrice.OutboundPrefixPrices {
		for _, prefix := range outPrice.DestinationPrefixes {
			outboundDestinationPrefixPriceMap[prefix] = true
		}
		for _, prefix := range outPrice.OriginationPrefixes {
			outboundOriginationPrefixPriceMap[prefix] = true
		}
	}
	// outboundDestinationPrefixPriceMap => map[1907:true 1844:true 1866:true 1877:true 1:true 1808:true 1800:true 1855:true 1888:true]
	// outboundOriginationPrefixPriceMap => map[ALL:true]
	_, ok = outboundDestinationPrefixPriceMap["1"]
	if ok == false {
		t.Error("Expected outbound price to contain a price for the destination prefix 1")
	}

	_, ok = outboundOriginationPrefixPriceMap["ALL"]
	if ok == false {
		t.Error("Expected outbound price to contain a price for the origination prefix ALL")
	}
}

func TestGetVoicePriceGB(t *testing.T) {
	t.Parallel()
	client, server := getServer(voicePricesGB)
	defer server.Close()

	isoCountry := "GB"
	expectedCountryName := "United Kingdom"
	expectedPriceUnit := "USD"

	voicePrice, err := client.Pricing.Voice.Countries.Get(context.Background(), isoCountry, nil)
	if err != nil {
		t.Fatal(err)
	}
	if voicePrice == nil {
		t.Fatal("expected voice price to be returned")
	}
	if voicePrice.Country != expectedCountryName {
		t.Errorf("Expected voice price country to be %s, but got %s\n", expectedCountryName, voicePrice.Country)
	}
	if voicePrice.IsoCountry != isoCountry {
		t.Errorf("Expected voice price iso country to be %s, but got %s\n", isoCountry, voicePrice.IsoCountry)
	}
	if voicePrice.PriceUnit != expectedPriceUnit {
		t.Errorf("Expected voice price unit to be %s, but got %s\n", expectedPriceUnit, voicePrice.PriceUnit)
	}
	if voicePrice.InboundCallPrices == nil {
		t.Error("Expected voice price to contain InboundCallPrices")
	}
	if voicePrice.OutboundPrefixPrices == nil {
		t.Error("Expected voice price to contain OutboundPrefixPrices")
	}

	inboundPriceMap := make(map[string]bool)
	for _, inPrice := range voicePrice.InboundCallPrices {
		inboundPriceMap[inPrice.NumberType] = true
	}
	_, ok := inboundPriceMap["local"]
	if ok == false {
		t.Error("Expected inbound price to contain a price for local calls")
	}

	outboundDestinationPrefixPriceMap := make(map[string]bool)
	outboundOriginationPrefixPriceMap := make(map[string]bool)
	for _, outPrice := range voicePrice.OutboundPrefixPrices {
		for _, prefix := range outPrice.DestinationPrefixes {
			outboundDestinationPrefixPriceMap[prefix] = true
		}
		for _, prefix := range outPrice.OriginationPrefixes {
			outboundOriginationPrefixPriceMap[prefix] = true
		}
	}
	_, ok = outboundDestinationPrefixPriceMap["44"]
	if ok == false {
		t.Error("Expected outbound price to contain a price for the prefix 44")
	}

	_, ok = outboundOriginationPrefixPriceMap["ALL"]
	if ok == false {
		t.Error("Expected outbound price to contain a price for the origination prefix ALL")
	}
}

func TestGetVoicePriceNumber(t *testing.T) {
	t.Parallel()
	client, server := getServer(voicePriceNumberUS)
	defer server.Close()

	expectedIsoCountry := "US"
	expectedCountryName := "United States"
	expectedPriceUnit := "USD"

	data := url.Values{}
	data.Set("OriginationNumber", from)
	voicePriceNum, err := client.Pricing.Voice.Numbers.Get(context.Background(), to, data)
	if err != nil {
		t.Fatal(err)
	}
	if voicePriceNum == nil {
		t.Fatal("expected voice price to be returned")
	}
	if voicePriceNum.OriginationNumber != from {
		t.Errorf("Expected voice price number to be %s, but got %s\n", from, voicePriceNum.OriginationNumber)
	}
	if voicePriceNum.DestinationNumber != to {
		t.Errorf("Expected voice price number to be %s, but got %s\n", to, voicePriceNum.DestinationNumber)
	}
	if voicePriceNum.Country != expectedCountryName {
		t.Errorf("Expected voice price country to be %s, but got %s\n", expectedCountryName, voicePriceNum.Country)
	}
	if voicePriceNum.IsoCountry != expectedIsoCountry {
		t.Errorf("Expected voice price iso country to be %s, but got %s\n", expectedIsoCountry, voicePriceNum.IsoCountry)
	}
	if voicePriceNum.PriceUnit != expectedPriceUnit {
		t.Errorf("Expected voice price unit to be %s, but got %s\n", expectedPriceUnit, voicePriceNum.PriceUnit)
	}
	if voicePriceNum.InboundCallPrice.BasePrice != "" {
		t.Error("Expected voice price to not contain an InboundCallPrice")
	}
	if len(voicePriceNum.OutboundCallPrices) != 1 {
		t.Error("Expected voice price to contain an OutboundPrefixPrices")
	}
	if voicePriceNum.OutboundCallPrices[0].BasePrice == "" {
		t.Error("Expected outbound call prices to contain an BasePrice")
	}
}

func TestGetVoicePricePage(t *testing.T) {
	t.Parallel()
	client, server := getServer(voiceCountriesPage)
	defer server.Close()

	data := url.Values{"PageSize": []string{"10"}}
	page, err := client.Pricing.Voice.Countries.GetPage(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}
	if len(page.Countries) == 0 {
		t.Error("expected to get a list of countries, got back 0")
	}
	if len(page.Countries) != 10 {
		t.Errorf("expected 10 countries, got %d", len(page.Countries))
	}
	if page.Meta.Key != "countries" {
		t.Errorf("expected Key to be 'countries', got %s", page.Meta.Key)
	}
	if page.Meta.PageSize != 10 {
		t.Errorf("expected PageSize to be 10, got %d", page.Meta.PageSize)
	}
	if page.Meta.Page != 0 {
		t.Errorf("expected Page to be 0, got %d", page.Meta.Page)
	}
	if page.Meta.PreviousPageURL.Valid != false {
		t.Errorf("expected previousPage.Valid to be false, got true")
	}
}
