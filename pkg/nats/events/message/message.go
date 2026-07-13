package message

import (
	"time"

	"github.com/google/uuid"
)

type MessageCreatedEvent struct {
	MessageID uuid.UUID  `json:"message_id"`
	ChannelID uuid.UUID  `json:"channel_id"`
	AuthorID  uuid.UUID  `json:"author_id"`
	Username   string     `json:"username"`
	Content   string     `json:"content"`
	ReplyTo   *uuid.UUID `json:"reply_to,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
}

type MessageDeletedEvent struct {
	MessageID uuid.UUID `json:"message_id"`
	ChannelID uuid.UUID `json:"channel_id"`
}
