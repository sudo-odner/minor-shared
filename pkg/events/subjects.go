package events

// Community Service NATS Subjects
const (
	SubjectChannelCreated          = "community.channel.created"
	SubjectChannelUpdated          = "community.channel.updated"
	SubjectChannelDeleted          = "community.channel.deleted"
	SubjectChannelPositionsUpdated = "community.channel.positions_updated"

	SubjectMemberAdded   = "community.member.added"
	SubjectMemberRemoved = "community.member.removed"
)

// Message Service NATS Subjects
const (
	SubjectMessageCreated = "chat.message.created"
	SubjectMessageDeleted = "chat.message.deleted"
)

// User Service NATS Subjects
const (
	SubjectUserCreated = "user.created"
	SubjectUserUpdated = "user.updated"
	SubjectUserDeleted = "user.deleted"

	SubjectRelationshipUpdated = "relationship.updated"
	SubjectRelationshipDeleted = "relationship.deleted"
)

// Presence Service NATS Subjects
const (
	SubjectPresenceStatusUpdated = "presence.status.updated"
)
