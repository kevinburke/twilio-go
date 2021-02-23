package twilio

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"testing"
	"time"
)

var pnTestCases = []struct {
	in       PhoneNumber
	expected string
}{
	{PhoneNumber("+41446681800"), "+41 44 668 18 00"},
	{PhoneNumber("+14105554092"), "+1 410-555-4092"},
	{PhoneNumber("blah"), "blah"},
}

func TestPhoneNumberFriendly(t *testing.T) {
	t.Parallel()
	for _, tt := range pnTestCases {
		if f := tt.in.Friendly(); f != tt.expected {
			t.Errorf("Friendly(%v): got %s, want %s", tt.in, f, tt.expected)
		}
	}
}

var pnParseTestCases = []struct {
	in  string
	out PhoneNumber
	err error
}{
	{"+14105551234", PhoneNumber("+14105551234"), nil},
	{"410 555 1234", PhoneNumber("+14105551234"), nil},
	{"(410) 555-1234", PhoneNumber("+14105551234"), nil},
	{"+41 44 6681800", PhoneNumber("+41446681800"), nil},
	{"foobarbang", PhoneNumber(""), errors.New("twilio: Invalid phone number: foobarbang")},
	{"22", PhoneNumber("+122"), nil},
	{"", PhoneNumber(""), ErrEmptyNumber},
}

func TestNewPhoneNumber(t *testing.T) {
	t.Parallel()
	for _, tt := range pnParseTestCases {
		pn, err := NewPhoneNumber(tt.in)
		name := fmt.Sprintf("ParsePhoneNumber(%v)", tt.in)
		if tt.err != nil {
			if err == nil {
				t.Errorf("%s: expected %v, got nil", name, tt.err)
				continue
			}
			if err.Error() != tt.err.Error() {
				t.Errorf("%s: expected error %v, got %v", name, tt.err, err)
			}
		} else if pn != tt.out {
			t.Errorf("%s: expected %v, got %v", name, tt.out, pn)
		}
	}
}

var timeTests = []struct {
	in    string
	valid bool
	time  time.Time
}{
	{`"Tue, 20 Sep 2016 22:59:57 +0000"`, true, time.Date(2016, 9, 20, 22, 59, 57, 0, time.UTC)},
	{`"2016-10-27T02:34:21Z"`, true, time.Date(2016, 10, 27, 2, 34, 21, 0, time.UTC)},
	{`null`, false, time.Time{}},
}

func TestUnmarshalTime(t *testing.T) {
	t.Parallel()
	for _, test := range timeTests {
		in := []byte(test.in)
		var tt TwilioTime
		if err := json.Unmarshal(in, &tt); err != nil {
			t.Errorf("json.Unmarshal(%v): got error %v", test.in, err)
		}
		if tt.Valid != test.valid {
			t.Errorf("json.Unmarshal(%v): expected Valid=%t, got %t", test.in, test.valid, tt.Valid)
		}
		if !tt.Time.Equal(test.time) {
			t.Errorf("json.Unmarshal(%v): expected time=%v, got %v", test.in, test.time, tt.Time)
		}
	}
}

func TestNewTwilioTime(t *testing.T) {
	t.Parallel()
	v := NewTwilioTime("foo")
	if v.Valid == true {
		t.Errorf("expected time to be invalid, got true")
	}
	in := "Tue, 20 Sep 2016 22:59:57 +0000"
	v = NewTwilioTime(in)
	if v.Valid == false {
		t.Errorf("expected %s to be valid time, got false", in)
	}
	if v.Time.Minute() != 59 {
		t.Errorf("wrong minute")
	}
	expected := "2016-09-20T22:59:57Z"
	if f := v.Time.Format(time.RFC3339); f != expected {
		t.Errorf("wrong format: got %v, want %v", f, expected)
	}
}

var priceTests = []struct {
	unit     string
	amount   string
	expected string
}{
	{"USD", "-0.0075", "$0.0075"},
	{"usd", "-0.0075", "$0.0075"},
	{"EUR", "0.0075", "€-0.0075"},
	{"UNK", "2.45", "UNK -2.45"},
	{"", "2.45", "-2.45"},
	{"USD", "-0.75000000", "$0.75"},
	{"USD", "-0.750", "$0.75"},
	{"USD", "-5000.00", "$5000"},
	{"USD", "-5000.", "$5000"},
	{"USD", "-5000", "$5000"},
}

func TestPrice(t *testing.T) {
	t.Parallel()
	for _, tt := range priceTests {
		out := price(tt.unit, tt.amount)
		if out != tt.expected {
			t.Errorf("price(%v, %v): got %v, want %v", tt.unit, tt.amount, out, tt.expected)
		}
	}
}

func TestTwilioDuration(t *testing.T) {
	t.Parallel()
	in := []byte(`"88"`)
	var td TwilioDuration
	if err := json.Unmarshal(in, &td); err != nil {
		t.Fatal(err)
	}
	if td != TwilioDuration(88*time.Second) {
		t.Errorf("got wrong duration: %v, wanted 88 seconds", td)
	}
}

var marshalUnmarshalTwilioDurationTests = []TwilioDuration{
	0,
	1 * TwilioDuration(time.Second),
	88 * TwilioDuration(time.Second),
}

func TestMarshalUnmarshalTwilioDuration(t *testing.T) {
	t.Parallel()
	for _, tt := range marshalUnmarshalTwilioDurationTests {
		b, err := json.Marshal(tt)
		if err != nil {
			t.Fatal(err)
		}

		var td TwilioDuration
		if err := json.Unmarshal(b, &td); err != nil {
			t.Fatal(err)
		}

		if tt != td {
			t.Fatalf("expected %d, got %d", tt, td)
		}
	}
}

var hdr = `"Transfer-Encoding=chunked&Server=cloudflare-nginx&CF-RAY=2f82bf9cb8102204-EWR&Set-Cookie=__cfduid%3Dd46f1cfd57d664c3038ae66f1c1de9e751477535661%3B+expires%3DFri%2C+27-Oct-17+02%3A34%3A21+GMT%3B+path%3D%2F%3B+domain%3D.inburke.com%3B+HttpOnly&Date=Thu%2C+27+Oct+2016+02%3A34%3A21+GMT&Content-Type=text%2Fhtml&CF-RAY=two"`

func TestUnmarshalHeader(t *testing.T) {
	t.Parallel()
	h := new(Values)
	if err := json.Unmarshal([]byte(hdr), h); err != nil {
		t.Fatal(err)
	}
	if h == nil {
		t.Fatal("nil h")
	}
	if val := h.Get("Transfer-Encoding"); val != "chunked" {
		t.Errorf("expected Transfer-Encoding: chunked header, got %s", val)
	}
	if vals := h.Values["CF-RAY"]; len(vals) < 2 {
		t.Errorf("expected to parse two CF-RAY headers, got less than 2")
	}
	if vals := h.Values["CF-RAY"]; vals[1] != "two" {
		t.Errorf("expected second header to be 'two', got %v", vals[1])
	}
}

var marshalUnmarshalSegmentsTests = []Segments{
	0,
	100,
	math.MaxUint32,
	Segments(^uint(0)),
}

func TestMarshalUnmarshalSegments(t *testing.T) {
	t.Parallel()
	for _, tt := range marshalUnmarshalSegmentsTests {
		b, err := json.Marshal(tt)
		if err != nil {
			t.Fatal(err)
		}

		var seg Segments
		if err := json.Unmarshal(b, &seg); err != nil {
			t.Fatal(err)
		}

		if tt != seg {
			t.Fatalf("expected %d, got %d", tt, seg)
		}
	}
}

var marshalUnmarshalNumMediaTests = []NumMedia{
	0,
	100,
	math.MaxUint32,
	NumMedia(^uint(0)),
}

func TestMarshalUnmarshalNumMedia(t *testing.T) {
	t.Parallel()
	for _, tt := range marshalUnmarshalNumMediaTests {
		b, err := json.Marshal(tt)
		if err != nil {
			t.Fatal(err)
		}

		var numMedia NumMedia
		if err := json.Unmarshal(b, &numMedia); err != nil {
			t.Fatal(err)
		}

		if tt != numMedia {
			t.Fatalf("expected %d, got %d", tt, numMedia)
		}
	}
}
