package token

const (
	// See https://www.twilio.com/docs/api/chat/rest/messages#properties
	// https://www.twilio.com/docs/api/rest/access-tokens
	ipMessagingGrant   = "ip_messaging"
	conversationsGrant = "rtc"
	voiceGrant         = "voice"
	videoGrant         = "video"
	keyServiceSid      = "service_sid"
	keyEndpointId      = "endpoint_id"
	keyDepRoleSide     = "deployment_role_sid"
	keyPushCredSid     = "push_credential_sid"
	keyVoiceOutgoing   = "outgoing"
	keyConfProfSid     = "configuration_profile_sid"
	keyAppSid          = "application_sid"
	keyVoiceParams     = "params"
	// https://www.twilio.com/docs/video/tutorials/user-identity-access-tokens
	keyRoomSid = "room"
	chatGrant  = "chat"
)

// Grant is a Twilio SID resource that can be added to an AccessToken for extra
// services. Implement this interface to create a custom Grant.
type Grant interface {
	ToPayload() map[string]interface{}
	Key() string
}

// IPMessageGrant is a grant for accessing Twilio IP Messaging
type IPMessageGrant struct {
	serviceSid        string
	endpointID        string
	deploymentRoleSid string
	pushCredentialSid string
}

func NewIPMessageGrant(serviceSid, endpointID, deploymentRoleSid, pushCredentialSid string) *IPMessageGrant {
	return &IPMessageGrant{
		serviceSid:        serviceSid,
		endpointID:        endpointID,
		deploymentRoleSid: deploymentRoleSid,
		pushCredentialSid: pushCredentialSid,
	}
}

func (gr *IPMessageGrant) ToPayload() map[string]interface{} {
	grant := make(map[string]interface{})
	if len(gr.serviceSid) > 0 {
		grant[keyServiceSid] = gr.serviceSid
	}
	if len(gr.endpointID) > 0 {
		grant[keyEndpointId] = gr.endpointID
	}
	if len(gr.deploymentRoleSid) > 0 {
		grant[keyDepRoleSide] = gr.deploymentRoleSid
	}
	if len(gr.pushCredentialSid) > 0 {
		grant[keyPushCredSid] = gr.pushCredentialSid
	}
	return grant
}

func (gr *IPMessageGrant) Key() string {
	return ipMessagingGrant
}

// ConversationsGrant is for Twilio Conversations
type ConversationsGrant struct {
	configurationProfileSid string
}

func NewConversationsGrant(sid string) *ConversationsGrant {
	return &ConversationsGrant{configurationProfileSid: sid}
}

func (gr *ConversationsGrant) ToPayload() map[string]interface{} {
	if len(gr.configurationProfileSid) > 0 {
		return map[string]interface{}{
			keyConfProfSid: gr.configurationProfileSid,
		}
	}
	return make(map[string]interface{})
}

func (gr *ConversationsGrant) Key() string {
	return conversationsGrant
}

// VoiceGrant is a grant for accessing Twilio IP Messaging
type VoiceGrant struct {
	outgoingApplicationSid    string                 // application sid to call when placing outgoing call
	outgoingApplicationParams map[string]interface{} // request params to pass to the application
	endpointID                string                 // Specify an endpoint identifier for this device, which will allow the developer to direct calls to a specific endpoint when multiple devices are associated with a single identity
	pushCredentialSid         string                 // Push Credential Sid to use when registering to receive incoming call notifications
}

func NewVoiceGrant(outAppSid string, outAppParams map[string]interface{}, endpointID string, pushCredentialSid string) *VoiceGrant {
	return &VoiceGrant{
		outgoingApplicationSid:    outAppSid,
		outgoingApplicationParams: outAppParams,
		endpointID:                endpointID,
		pushCredentialSid:         pushCredentialSid,
	}
}

func (gr *VoiceGrant) ToPayload() map[string]interface{} {
	outVoice := make(map[string]interface{})
	if len(gr.outgoingApplicationSid) > 0 {
		outVoice[keyAppSid] = gr.outgoingApplicationSid
	}
	if len(gr.outgoingApplicationParams) > 0 {
		outVoice[keyVoiceParams] = gr.outgoingApplicationParams
	}

	grant := make(map[string]interface{})
	grant[keyVoiceOutgoing] = outVoice
	if len(gr.endpointID) > 0 {
		grant[keyEndpointId] = gr.endpointID
	}
	if len(gr.pushCredentialSid) > 0 {
		grant[keyPushCredSid] = gr.pushCredentialSid
	}
	return grant
}

func (gr *VoiceGrant) Key() string {
	return voiceGrant
}

// VideoGrant is for Twilio Programmable Video access
type VideoGrant struct {
	roomSID string
}

func NewVideoGrant(sid string) *VideoGrant {
	return &VideoGrant{roomSID: sid}
}

func (gr *VideoGrant) ToPayload() map[string]interface{} {
	if len(gr.roomSID) > 0 {
		return map[string]interface{}{
			keyRoomSid: gr.roomSID,
		}
	}
	return make(map[string]interface{})
}

func (gr *VideoGrant) Key() string {
	return videoGrant
}

// ChatGrant is for Twilio Programmable Chat
type ChatGrant struct {
	serviceSid string
}

func NewChatGrant(sid string) *ChatGrant {
	return &ChatGrant{serviceSid: sid}
}

func (cg *ChatGrant) ToPayload() map[string]interface{} {
	if len(cg.serviceSid) > 0 {
		return map[string]interface{}{
			keyServiceSid: cg.serviceSid,
		}
	}
	return make(map[string]interface{})
}

func (cg *ChatGrant) Key() string {
	return chatGrant
}
