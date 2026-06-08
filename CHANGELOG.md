# Changes

## Unreleased

Added support for the SIP configuration resources under the core API's `/SIP`
path, reachable via `Client.SIP`: Credential Lists and their Credentials, SIP
Domains, IP Access Control Lists and their IP Addresses, the Domain
CredentialList and IpAccessControlList mappings, and the Domain Auth mappings
for calls and registrations. Together these cover the provisioning needed to
register a SIP endpoint to a Programmable Voice SIP Domain.

A new conformance test (`sip_oai_test.go`) checks these structs against
Twilio's official [OpenAPI definitions][twilio-oai]. It is scoped to the SIP
resources and skips unless a twilio-oai checkout is present; set
`TWILIO_OAI_DIR` to point at one. It verifies every SIP schema property is
modeled by a struct field and that the spec's response examples decode with
unknown fields disallowed.

Extended the [OpenAPI][twilio-oai] conformance tests beyond SIP to Messages,
Calls, phone numbers (incoming, available, outgoing caller IDs), Recordings,
Transcriptions, and Conferences (participants and conference recordings), with
shared machinery in `oai_test.go`. Each resource gets three checks: schema
coverage (every spec property is modeled by a struct field, with simple schema
types checked against Go field types), example decoding (every response example
decodes with unknown fields disallowed, catching missing fields and example type
mismatches), and endpoint classification (every spec operation is triaged
supported or unimplemented, failing if the spec gains an operation the test has
not triaged, and documenting the gaps). The decoder
handles both instance responses and the items of list responses, so resources
that expose only list endpoints (such as available numbers) are covered too.

The checks surfaced one bug and a number of missing fields, all fixed here:

- `NullAnsweredBy` now implements `UnmarshalJSON`. Previously a non-null
  `answered_by` value (such as `"human"` or `"machine"`) failed to decode the
  entire `Call`, because the JSON string could not be unmarshaled into the
  struct shape. Existing fixtures only exercised the null case, so the bug went
  unnoticed.
- `Call` gained `from_formatted`, `to_formatted`, `trunk_sid`, and
  `subresource_uris`.
- `NumberCapability` gained `fax` (shared by `IncomingPhoneNumber` and
  `AvailableNumber`).
- `IncomingPhoneNumber` gained `address_sid`, `bundle_sid`,
  `emergency_address_status`, `identity_sid`, `origin`, `status`, `type`, and
  `voice_receive_mode`.
- `AvailableNumber` gained `locality`.
- `Recording` gained `conference_sid`, `source`, `start_time`, `error_code`,
  `media_url`, `encryption_details`, and `subresource_uris`.
- `Conference` gained `reason_conference_ended`.
- `Participant` gained `call_sid_to_coach`, `coaching`, `label`, `queue_time`,
  and `status`.

Implemented a set of endpoints the conformance tests had flagged as missing:

- Messages: `MessageService.Update` and `Redact`, `MessageService.CreateFeedback`
  (with a new `MessageFeedback` type), and `MediaService.Delete`.
- Calls: `CallService.Delete`; call-scoped recordings (`GetCallRecordings`,
  `CreateRecording`, `GetRecording`, `UpdateRecording`, `DeleteRecording`, with a
  new `CallRecording` type); call events (`GetEvents`, with `CallEventLog` — named to
  avoid colliding with the Voice Insights `CallEvent`); call notifications
  (`GetNotifications`, `GetNotification`, with `CallNotification`); and `<Pay>`
  payments (`CreatePayment`, `UpdatePayment`, with `Payment`).
- Conferences: `ConferenceService.Update`; conference recordings (`GetRecordings`,
  `GetRecording`, `UpdateRecording`, `DeleteRecording`, reusing `Recording`); and
  the participant lifecycle via `Conferences.Participants(conferenceSid)`, which
  turns the previously-empty `ParticipantService` into a real service with
  `Create`, `Get`, `Update`, `Delete`, and `GetPage`.

The phone-number sub-resources (typed lists, Mobile purchasing, AssignedAddOns)
and the remaining call sub-resources (siprec, streams, realtime transcriptions,
user-defined messages) are still unimplemented.

[twilio-oai]: https://github.com/twilio/twilio-oai

## 2.11.0

The module path now includes the required `/v2` major-version suffix:
`github.com/kevinburke/twilio-go/v2`. Previous v2.x releases shipped under
the unsuffixed path, which meant `go get github.com/kevinburke/twilio-go@latest`
silently resolved to a v1.x tag (or a pre-v2 pseudo-version) instead of the
2.x line. From 2.11.0 on, v2 tags are resolvable normally.

To upgrade, rewrite imports to add the `/v2` suffix, for example:

	import twilio "github.com/kevinburke/twilio-go"

becomes

	import twilio "github.com/kevinburke/twilio-go/v2"

and likewise for the subpackages (`/v2/token`, `/v2/twilioclient`,
`/v2/datausage`). No API changes; this is purely a module-path rename. A
one-shot sed will handle most codebases:

	find . -name '*.go' -exec sed -i '' \
	    -e 's|"github.com/kevinburke/twilio-go"|"github.com/kevinburke/twilio-go/v2"|g' \
	    -e 's|github.com/kevinburke/twilio-go/\(token\|twilioclient\|datausage\)|github.com/kevinburke/twilio-go/v2/\1|g' \
	    {} +

then `go mod tidy`. Pinned v1.x users (`v1.8` and earlier) are unaffected;
those tags continue to resolve against the old import path.

## 2.10.0 (2026-05-22)

Fix a data race in `token.AccessToken.JWT`. Although the 2.8 release removed
the `jwt-go` dependency, the `token` package still builds Access Token JWTs by
hand; concurrent calls could share and mutate the package-level encoded header
buffer while constructing the signing input.

## 2.8

JWT functionality has been removed. We did not want to continue to support
functionality that may break or be compromised easily, and is difficult to use
correctly.

To achieve the same functionality, you can integrate your own JWT code, or look
at the first draft of pull request #99.

## 2.7

Dependency update

## 2.6

Add "incoming" option for VoiceGrant.

Tags in VoiceCallSummary are a []string, not a map[string]string (the docs have
them as "null", so this wasn't clear).

Segments, NumMedia and TwilioTime can be marshaled back to JSON as strings.

ChatGrant now accepts a PushCredentialSid as a second argument.

The Verify client now has support for AccessTokens and Challenges.

Calls now support the "queue_time" attribute, with help from a new
"TwilioDurationMS" type.

IntervalMetrics response object now marshals types correctly from the Twilio API
(the response format changed).

Use Go modules to track dependencies.

Update source code to be formatted with Go 1.19 doc styles.

## 2.5

Use a new version of github.com/kevinburke/rest that reduces the number of
imports necessary to run the library.

## 2.4

Add Voice Insights API (thank you @yeoji)

Add more fields to Conference struct (thank you @zmei95 for the issue report).

## 2.3

The Twilio Calls API recently started returning calls by StartTime instead of by
Date Created, which broke our GetCallsInRange in-memory filtering. Use StartTime
as the sort order key if present, falling back to DateCreated. Update the tests
to match.

Add Programmable Grant chat type to the list of tokens. Thanks Kelmer Perez for
the patch.

## 2.2

Implement Twilio Verify API. Thanks Elijah Oyekunle for the patch.

## 2.1

- Support deleting messages via `client.Messages.Delete(sid)`.

- Support a different default region for local phone number parsing than "US" by
  overriding the value of `twilio.DefaultRegion` in an init function.

Thanks Kevin Golding for the suggestions and issue reports.

## 2.0

Update Pricing API to v2. This is a breaking change.

To upgrade from the v1 Pricing API:

* `VoicePrice` is now `VoicePrices`.
* `VoiceNumberPrice` is now `VoiceNumberPrices`.
* `NumberPrice` is now `NumberPrices`.
* `MessagePrice` is now `MessagePrices`
* The `Prefixes` field in the `PrefixPrice` struct is now `DestinationPrefixes`.
* `OriginationPrefixes` field has been added to the `PrefixPrice` struct.
* The `OutboundCallPrice` field in the `VoiceNumberPrices` struct is now
`OutboundCallPrices`. Add the old `OutboundCallPrice` as the first entry in the
array.
* `OriginationPrefixes` field has been added to the `OutboundCallPrice` struct.

* `NumberVoicePriceService.Get` now has 2 parameters: `destinationNumber` and
`data` url.Values. `destinationNumber` is same as the previous `number` field.
`data` can contain additional parameters like `OriginationNumber`.
* Several other `*PriceService.Get` functions now have a third `url.Values`
  parameter. Pass `nil` for this third parameter to maintain compatibility with
  the v1 API.

## 1.8

Add Task Router Workers API.

## 1.7

Add Task Router Workflow API.

## 1.6

Add Task Router TaskQueue API.

## 1.5

Add Task Router Activity API.

## 1.4

Add Supported Countries service. (via Andrei Schneider @megaherz)

## 1.3

Remove dependency on jwt-go (functionality is the same, we just generate JWT's
by hand now.)

## 1.2

Fix key for ConferenceSid.

## 1.1

Add video and video recordings API

## 1.0

Update kevinburke/rest to version 2.0.

## 0.66

Support the Notify Credentials API. Thanks to Maksym Pavlenko for implementing
this change.

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
