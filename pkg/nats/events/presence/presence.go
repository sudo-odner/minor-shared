package events

type PresenceStatusUpdatedEvent struct {
	UserID       string `json:"user_id"`
	Status       int32  `json:"status"`
	CustomStatus string `json:"custom_status"`
	LastActiveAt int64  `json:"last_active_at"`
}
