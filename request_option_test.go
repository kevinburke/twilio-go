package twilio

import "testing"

func TestRequestOptions(t *testing.T) {
	t.Parallel()

	vals := getValues(
		WithFriendlyName("test"),
		WithAreaCode("562"),
		WithSmsEnabled(true),
		WithMmsEnabled(false),
		WithVoiceEnabled(true),
		WithFaxEnabled(false),
		WithStatus("active"),
		WithSid("123"),
		WithAccountSid("321"),
		WithTrunkSid("213"),
		WithStreet("S"),
		WithCity("C"),
		WithRegion("R"),
		WithPostalCode("02123"),
		WithOption("Custom", "111"),
	)

	expected := map[string][]string{
		"FriendlyName": {"test"},
		"AreaCode":     {"562"},
		"SmsEnabled":   {"true"},
		"MmsEnabled":   {"false"},
		"VoiceEnabled": {"true"},
		"FaxEnabled":   {"false"},
		"Status":       {"active"},
		"Sid":          {"123"},
		"AccountSid":   {"321"},
		"TrunkSid":     {"213"},
		"Street":       {"S"},
		"City":         {"C"},
		"Region":       {"R"},
		"PostalCode":   {"02123"},
		"Custom":       {"111"},
	}

	if len(expected) != len(vals) {
		t.Fatal("invalid map size")
	}

	for k, expected := range expected {
		actual, ok := vals[k]
		if !ok {
			t.Fatalf("expected key not found %s", k)
		}

		if len(expected) != len(actual) {
			t.Fatalf("unexpected len for key: %s", k)
		}

		for i, x := range expected {
			if x != actual[i] {
				t.Errorf("not equal values at %d for key '%s' ('%v' != '%v')", i, k, x, actual[i])
			}
		}
	}
}
