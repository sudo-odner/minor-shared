package events

type UserCreatedEvent struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserUpdatedEvent struct {
	UserID    string  `json:"user_id"`
	Username  string  `json:"username"`
	AvatarURL *string `json:"avatar_url"`
	Bio       string  `json:"bio"`
}

type UserDeletedEvent struct {
	UserID string `json:"user_id"`
}

type RelationshipUpdatedEvent struct {
	UserID   string `json:"user_id"`
	TargetID string `json:"target_id"`
	Status   int16  `json:"status"`
}

type RelationshipDeletedEvent struct {
	UserID   string `json:"user_id"`
	TargetID string `json:"target_id"`
}
