package token

import (
	"testing"
	"time"
)

const (
	SERVICE_SID   = "123234567567"
	ENDPOINT_ID   = "asdfghjklpoiuytrewq"
	DEP_ROLE_SID  = "1qaz2wsx3edc4rfv5tgb6yhn7ujm8ik9ol"
	PUSH_CRED_SID = "cde3xsw2zaq1vfr4bgtnhy6mju78ijhgtf"
	PROFILE_SID   = "erfergrtugdifuovudsfhguidhgouidrhg"
	ROOM_SID      = "erfergrtugdifuovudsfhguidhgouidrhd"
)

func TestIPMessageGrant(t *testing.T) {
	t.Parallel()
	ipMsgGrant := NewIPMessageGrant(SERVICE_SID, ENDPOINT_ID, DEP_ROLE_SID, PUSH_CRED_SID)

	if ipMsgGrant.Key() != ipMessagingGrant {
		t.Errorf("key expected to be %s, got %s\n", ipMessagingGrant, ipMsgGrant.Key())
	}

	if ipMsgGrant.ToPayload()[keyServiceSid] != SERVICE_SID {
		t.Errorf("%s expected to be %s, got %s\n", keyServiceSid, SERVICE_SID, ipMsgGrant.ToPayload()[keyServiceSid])
	}

	if ipMsgGrant.ToPayload()[keyEndpointId] != ENDPOINT_ID {
		t.Errorf("%s expected to be %s, got %s\n", keyEndpointId, ENDPOINT_ID, ipMsgGrant.ToPayload()[keyEndpointId])
	}

	if ipMsgGrant.ToPayload()[keyDepRoleSide] != DEP_ROLE_SID {
		t.Errorf("%s expected to be %s, got %s\n", keyDepRoleSide, DEP_ROLE_SID, ipMsgGrant.ToPayload()[keyDepRoleSide])
	}

	if ipMsgGrant.ToPayload()[keyPushCredSid] != PUSH_CRED_SID {
		t.Errorf("%s expected to be %s, got %s\n", keyPushCredSid, PUSH_CRED_SID, ipMsgGrant.ToPayload()[keyPushCredSid])
	}
}

func TestConversationsGrant(t *testing.T) {
	t.Parallel()
	convGrant := NewConversationsGrant(PROFILE_SID)

	if convGrant.Key() != conversationsGrant {
		t.Errorf("key expected to be %s, got %s\n", conversationsGrant, convGrant.Key())
	}

	if convGrant.ToPayload()[keyConfProfSid] != PROFILE_SID {
		t.Errorf("%s expected to be %s, got %s\n", keyConfProfSid, PROFILE_SID, convGrant.ToPayload()[keyConfProfSid])
	}
}

func TestVoiceGrant(t *testing.T) {
	t.Parallel()
	params := map[string]interface{}{
		"extra":     "wefergdfgdf",
		"logged_id": true,
		"data": map[string]interface{}{
			"id":         101212434,
			"name":       "John",
			"created_at": time.Now(),
		},
	}
	vcGrnt := NewVoiceGrant(APP_SID, params, ENDPOINT_ID, PUSH_CRED_SID)

	if vcGrnt.Key() != voiceGrant {
		t.Errorf("key expected to be %s, got %s\n", voiceGrant, vcGrnt.Key())
	}

	if vcGrnt.ToPayload()[keyVoiceOutgoing] == nil {
		t.Errorf("expected payload %s to exist", keyVoiceOutgoing)
	}

	payload := vcGrnt.ToPayload()
	outgoingMap := payload[keyVoiceOutgoing].(map[string]interface{})
	if outgoingMap[keyAppSid] != APP_SID {
		t.Errorf("Expected payload [%s][%s] to be %s, got %s\n", keyVoiceOutgoing, keyAppSid, APP_SID, outgoingMap[keyAppSid])
	}

	if outgoingMap[keyVoiceParams] == nil {
		t.Errorf("Expected payload [%s][%s] to exist\n", keyVoiceOutgoing, keyVoiceParams)
	}

	endpointId := vcGrnt.ToPayload()[keyEndpointId]
	if endpointId != ENDPOINT_ID {
		t.Errorf("Expected payload %s to be %s, got %s\n", keyEndpointId, ENDPOINT_ID, endpointId)
	}

	pushCredSid := vcGrnt.ToPayload()[keyPushCredSid]
	if pushCredSid != PUSH_CRED_SID {
		t.Errorf("Expected payload %s to be %s, got %s\n", keyPushCredSid, PUSH_CRED_SID, pushCredSid)
	}
}

func TestVideoGrant(t *testing.T) {
	t.Parallel()
	vdGrnt := NewVideoGrant(keyRoomSid)

	if vdGrnt.Key() != videoGrant {
		t.Errorf("key expected to be %s, got %s\n", videoGrant, vdGrnt.Key())
	}

	if vdGrnt.ToPayload()[keyRoomSid] != keyRoomSid {
		t.Errorf("%s expected to be %s, got %s\n", keyRoomSid, ROOM_SID, vdGrnt.ToPayload()[keyRoomSid])
	}
}

func TestChatGrant(t *testing.T) {
	t.Parallel()
	chGrnt := NewChatGrant(SERVICE_SID)

	if chGrnt.Key() != chatGrant {
		t.Errorf("key expected to be %s, got %s\n", chatGrant, chGrnt.Key())
	}

	if chGrnt.ToPayload()[keyServiceSid] != SERVICE_SID {
		t.Errorf("%s expected to be %s, got %s\n", keyServiceSid, SERVICE_SID, chGrnt.ToPayload()[keyServiceSid])
	}
}
