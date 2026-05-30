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
