package phonenumber

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/ttacon/libphonenumber"
)

type PhoneNumber string

var ErrEmptyNumber = errors.New("twilio: The provided phone number was empty")

// NewPhoneNumber parses the given value as a phone number or returns an error
// if it cannot be parsed as one. If a phone number does not begin with a plus
// sign, we assume it's a US national number. Numbers are stored in E.164
// format.
func NewPhoneNumber(pn string) (PhoneNumber, error) {
	if len(pn) == 0 {
		return "", ErrEmptyNumber
	}
	num, err := libphonenumber.Parse(pn, "US")
	// Add some better error messages - the ones in libphonenumber are generic
	switch {
	case err == libphonenumber.ErrNotANumber:
		return "", fmt.Errorf("twilio: Invalid phone number: %s", pn)
	case err == libphonenumber.ErrInvalidCountryCode:
		return "", fmt.Errorf("twilio: Invalid country code for number: %s", pn)
	case err != nil:
		return "", err
	}
	return PhoneNumber(libphonenumber.Format(num, libphonenumber.E164)), nil
}

// Friendly returns a friendly international representation of the phone
// number, for example, "+14105554092" is returned as "+1 410-555-4092". If the
// phone number is not in E.164 format, we try to parse it as a US number. If
// we cannot parse it as a US number, it is returned as is.
func (pn PhoneNumber) Friendly() string {
	num, err := libphonenumber.Parse(string(pn), "US")
	if err != nil {
		return string(pn)
	}
	return libphonenumber.Format(num, libphonenumber.INTERNATIONAL)
}

// Local returns a friendly national representation of the phone number, for
// example, "+14105554092" is returned as "(410) 555-4092". If the phone number
// is not in E.164 format, we try to parse it as a US number. If we cannot
// parse it as a US number, it is returned as is.
func (pn PhoneNumber) Local() string {
	num, err := libphonenumber.Parse(string(pn), "US")
	if err != nil {
		return string(pn)
	}
	return libphonenumber.Format(num, libphonenumber.NATIONAL)
}
