package community

import (
	"time"

	"github.com/google/uuid"
)

type ChannelDTO struct {
	ID        uuid.UUID  `json:"id"`
	ServerID  uuid.UUID  `json:"server_id"`
	Name      string     `json:"name"`
	Type      int        `json:"type"`
	ParentID  *uuid.UUID `json:"parent_id,omitempty"`
	Position  int        `json:"position"`
	CreatedAt time.Time  `json:"created_at"`
}

type ShortChannelDTO struct {
	ID       uuid.UUID  `json:"id"`
	ParentID *uuid.UUID `json:"parent_id,omitempty"`
	Position int        `json:"position"`
}

type ChannelEvent struct {
	ServerID uuid.UUID  `json:"server_id"`
	Channel  ChannelDTO `json:"channel"`
}

type ChannelDeletedEvent struct {
	ServerID  uuid.UUID `json:"server_id"`
	ChannelID uuid.UUID `json:"channel_id"`
}

type ChannelPositionUpdateEvent struct {
	ServerID uuid.UUID         `json:"server_id"`
	Channels []ShortChannelDTO `json:"channels"`
}

type UserDeleteEvent struct {
	UserID uuid.UUID `json:"user_id"`
}

type MemberDTO struct {
	ServerID uuid.UUID `json:"server_id"`
	UserID   uuid.UUID `json:"user_id"`
	Nickname *string   `json:"nickname,omitempty"`
	JoinedAt time.Time `json:"joined_at"`
}

type MemberAddedEvent struct {
	ServerID uuid.UUID `json:"server_id"`
	Member   MemberDTO `json:"member"`
}

type MemberRemovedEvent struct {
	ServerID uuid.UUID `json:"server_id"`
	UserID   uuid.UUID `json:"user_id"`
}
