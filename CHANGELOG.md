# Changes

## 0.65

Support the Commands Wireless API.

Add cmd/report-data-usage and cmd/report-data-usage-server for reporting recent
data usage for Sims.

## 0.64

Partial support for wireless.twilio.com endpoints, including the Sim resource
and UsageRecords for Sims.

This version of the library requires github.com/kevinburke/go-types version
0.20 at least.

## 0.63

Support Update() for the IncomingPhoneNumber resource.

## 0.62

Fix error in release 0.61.

## 0.61

Support the Available Phone Numbers API. Thanks to Maksym Pavlenko for
contributing the patch.

## 0.59

Support the Twilio Fax API.

- Rename imports from github.com/saintpete/twilio-go to
  github.com/kevinburke/twilio-go.

## 0.58

Add `client.UseSecretKey(key string)` to handle secret keys
properly. For more information on the secret key API, see
https://www.twilio.com/docs/api/rest/keys.

Thanks to Andrew Watson for reporting.

## 0.57

Switch all imports to use the `"context"` library.

## 0.56

Use the new github.com/kevinburke/rest.DefaultTransport RoundTripper for
easy HTTP debugging. (The previous code set the RoundTripper to nil, so
kevinburke/rest wouldn't log anything).

## 0.55

Handle new HTTPS-friendly media URLs.

Support deleting/releasing phone numbers via IncomingNumbers.Release(ctx, sid).

Initial support for the Pricing API (https://pricing.twilio.com).

Add an AUTHORS.txt.

## 0.54

Add Recordings.GetTranscriptions() to get Transcriptions for a recording. The
Transcriptions resource doesn't support filtering by Recording Sid.

## 0.53

Add Alert.StatusCode() function for retrieving a HTTP status code (if one
exists) from an alert.

## 0.52

Copy url.Values for GetXInRange() functions before modifying them.

## 0.51

Implement GetNextConferencesInRange

## 0.50

Implement GetConferencesInRange. Fix paging error in
GetCallsInRange/GetMessagesInRange.

## 0.47

Implement GetNextXInRange - if you have a next page URI and want to get an
Iterator (instead of starting with a url.Values).

## 0.45

Fix go 1.6 (messages_example_test) relied on the stdlib Context package by
accident.

## 0.44

Support filtering Calls/Messages down to the nanosecond in a TZ-aware way, with
Calls.GetCallsInRange / Messages/GetMessagesInRange.

## 0.42

Add more Description fields based on errors I've received in the past. There
are probably more to be found, but this is a good start.

## 0.41

Use the same JWT library instead of using two different ones.

Add Description() for Alert bodies.

## 0.40

Fix next page URL's for Twilio Monitor

## 0.39

The data in Update() requests was silently being ignored. They are not ignored
any more.

Support the Accounts resource.

Add RequestOnBehalfOf function to make requests on behalf of a subaccount.

Fixes short tests that were broken in 0.38

## 0.37

Support Outgoing Caller ID's

## 0.36

Support Keys

## 0.35

Added Ended(), EndedUnsuccessfully() helpers to a Call, and FriendlyPrice() to
a Transcription.
