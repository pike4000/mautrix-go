package mautrix

import (
	"encoding/json"
)

// RespError is the standard JSON error response from Homeservers. It also implements the Golang "error" interface.
// See http://matrix.org/docs/spec/client_server/r0.2.0.html#api-standards
type RespError struct {
	ErrCode string `json:"errcode"`
	Err     string `json:"error"`
}

// Error returns the errcode and error message.
func (e RespError) Error() string {
	return e.ErrCode + ": " + e.Err
}

// RespCreateFilter is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-user-userid-filter
type RespCreateFilter struct {
	FilterID string `json:"filter_id"`
}

// RespVersions is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#get-matrix-client-versions
type RespVersions struct {
	Versions []string `json:"versions"`
}

// RespJoinRoom is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-join
type RespJoinRoom struct {
	RoomID string `json:"room_id"`
}

// RespLeaveRoom is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-leave
type RespLeaveRoom struct{}

// RespForgetRoom is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-forget
type RespForgetRoom struct{}

// RespInviteUser is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-invite
type RespInviteUser struct{}

// RespKickUser is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-kick
type RespKickUser struct{}

// RespBanUser is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-ban
type RespBanUser struct{}

// RespUnbanUser is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-unban
type RespUnbanUser struct{}

// RespTyping is the JSON response for https://matrix.org/docs/spec/client_server/r0.2.0.html#put-matrix-client-r0-rooms-roomid-typing-userid
type RespTyping struct{}

// RespJoinedRooms is the JSON response for https://matrix.org/docs/spec/client_server/r0.4.0.html#get-matrix-client-r0-joined-rooms
type RespJoinedRooms struct {
	JoinedRooms []string `json:"joined_rooms"`
}

// RespJoinedMembers is the JSON response for https://matrix.org/docs/spec/client_server/r0.4.0.html#get-matrix-client-r0-joined-rooms
type RespJoinedMembers struct {
	Joined map[string]struct {
		DisplayName *string `json:"display_name"`
		AvatarURL   *string `json:"avatar_url"`
	} `json:"joined"`
}

// RespMessages is the JSON response for https://matrix.org/docs/spec/client_server/r0.2.0.html#get-matrix-client-r0-rooms-roomid-messages
type RespMessages struct {
	Start string   `json:"start"`
	Chunk []*Event `json:"chunk"`
	State []*Event `json:"state"`
	End   string   `json:"end"`
}

// RespSendEvent is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#put-matrix-client-r0-rooms-roomid-send-eventtype-txnid
type RespSendEvent struct {
	EventID string `json:"event_id"`
}

// RespMediaUpload is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-media-r0-upload
type RespMediaUpload struct {
	ContentURI string `json:"content_uri"`
}

// RespUserInteractive is the JSON response for https://matrix.org/docs/spec/client_server/r0.2.0.html#user-interactive-authentication-api
type RespUserInteractive struct {
	Flows []struct {
		Stages []string `json:"stages"`
	} `json:"flows"`
	Params    map[string]interface{} `json:"params"`
	Session   string                 `json:"string"`
	Completed []string               `json:"completed"`
	ErrCode   string                 `json:"errcode"`
	Error     string                 `json:"error"`
}

// HasSingleStageFlow returns true if there exists at least 1 Flow with a single stage of stageName.
func (r RespUserInteractive) HasSingleStageFlow(stageName string) bool {
	for _, f := range r.Flows {
		if len(f.Stages) == 1 && f.Stages[0] == stageName {
			return true
		}
	}
	return false
}

// RespUserDisplayName is the JSON response for https://matrix.org/docs/spec/client_server/r0.2.0.html#get-matrix-client-r0-profile-userid-displayname
type RespUserDisplayName struct {
	DisplayName string `json:"displayname"`
}

// RespRegister is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-register
type RespRegister struct {
	AccessToken  string `json:"access_token"`
	DeviceID     string `json:"device_id"`
	HomeServer   string `json:"home_server"`
	RefreshToken string `json:"refresh_token"`
	UserID       string `json:"user_id"`
}

type RespLoginFlows struct {
	Flows []struct {
		Type string `json:"type"`
	} `json:"flows"`
}

// RespLogin is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-login
type RespLogin struct {
	AccessToken string `json:"access_token"`
	DeviceID    string `json:"device_id"`
	HomeServer  string `json:"home_server"`
	UserID      string `json:"user_id"`
}

// RespLogout is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-logout
type RespLogout struct{}

// RespCreateRoom is the JSON response for https://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-createroom
type RespCreateRoom struct {
	RoomID string `json:"room_id"`
}

type RespMembers struct {
	Chunk []*Event `json:"chunk"`
}

type LazyLoadSummary struct {
	Heroes             []string `json:"m.heroes,omitempty"`
	JoinedMemberCount  *int     `json:"m.joined_member_count,omitempty"`
	InvitedMemberCount *int     `json:"m.invited_member_count,omitempty"`
}

// RespSync is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#get-matrix-client-r0-sync
type RespSync struct {
	NextBatch   string `json:"next_batch"`
	AccountData struct {
		Events []json.RawMessage `json:"events"`
	} `json:"account_data"`
	Presence struct {
		Events []json.RawMessage `json:"events"`
	} `json:"presence"`
	Rooms struct {
		Leave map[string]struct {
			Summary LazyLoadSummary `json:"summary"`
			State   struct {
				Events []json.RawMessage `json:"events"`
			} `json:"state"`
			Timeline struct {
				Events    []json.RawMessage `json:"events"`
				Limited   bool              `json:"limited"`
				PrevBatch string            `json:"prev_batch"`
			} `json:"timeline"`
		} `json:"leave"`
		Join map[string]struct {
			Summary LazyLoadSummary `json:"summary"`
			State   struct {
				Events []json.RawMessage `json:"events"`
			} `json:"state"`
			Timeline struct {
				Events    []json.RawMessage `json:"events"`
				Limited   bool              `json:"limited"`
				PrevBatch string            `json:"prev_batch"`
			} `json:"timeline"`
			Ephemeral struct {
				Events []json.RawMessage `json:"events"`
			} `json:"ephemeral"`
			AccountData struct {
				Events []json.RawMessage `json:"events"`
			} `json:"account_data"`
		} `json:"join"`
		Invite map[string]struct {
			Summary LazyLoadSummary `json:"summary"`
			State   struct {
				Events []json.RawMessage `json:"events"`
			} `json:"invite_state"`
		} `json:"invite"`
	} `json:"rooms"`
}

type RespTurnServer struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	TTL      int      `json:"ttl"`
	URIs     []string `json:"uris"`
}

type RespAliasCreate struct{}
type RespAliasDelete struct{}
type RespAliasResolve struct {
	RoomID  string   `json:"room_id"`
	Servers []string `json:"servers"`
}
