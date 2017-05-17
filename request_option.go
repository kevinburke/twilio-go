package twilio

import (
	"net/url"
	"strconv"
)

type RequestOption func(values url.Values)

func WithFriendlyName(name string) RequestOption {
	return WithOption("FriendlyName", name)
}

func WithAreaCode(areaCode string) RequestOption {
	return WithOption("AreaCode", areaCode)
}

func WithPageSize(size int) RequestOption {
	return WithOption("PageSize", strconv.Itoa(size))
}

func WithPhoneNumber(phoneNumber string) RequestOption {
	return WithOption("PhoneNumber", phoneNumber)
}

func WithSmsEnabled(enabled bool) RequestOption {
	return WithOption("SmsEnabled", strconv.FormatBool(enabled))
}

func WithMmsEnabled(enabled bool) RequestOption {
	return WithOption("MmsEnabled", strconv.FormatBool(enabled))
}

func WithVoiceEnabled(enabled bool) RequestOption {
	return WithOption("VoiceEnabled", strconv.FormatBool(enabled))
}

func WithFaxEnabled(enabled bool) RequestOption {
	return WithOption("FaxEnabled", strconv.FormatBool(enabled))
}

func WithStatus(status string) RequestOption {
	return WithOption("Status", status)
}

func WithSid(sid string) RequestOption {
	return WithOption("Sid", sid)
}

func WithAccountSid(accountSid string) RequestOption {
	return WithOption("AccountSid", accountSid)
}

func WithTrunkSid(trunkSid string) RequestOption {
	return WithOption("TrunkSid", trunkSid)
}

func WithStreet(street string) RequestOption {
	return WithOption("Street", street)
}

func WithCity(city string) RequestOption {
	return WithOption("City", city)
}

func WithRegion(region string) RequestOption {
	return WithOption("Region", region)
}

func WithPostalCode(postalCode string) RequestOption {
	return WithOption("PostalCode", postalCode)
}

func WithOption(key string, val string) RequestOption {
	return func(values url.Values) {
		values.Add(key, val)
	}
}

func getValues(opts ...RequestOption) url.Values {
	v := url.Values{}
	for _, f := range opts {
		f(v)
	}

	return v
}
