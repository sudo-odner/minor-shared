package dm

const (
	SubjectChannelDeleted     = "channel.dm.deleted"
	SubjectGroupMemberAdded   = "channel.dm.group.member.added"
	SubjectGroupMemberRemoved = "channel.dm.group.member.removed"
)

type ChannelDeleted struct{}

func (ChannelDeleted) Subject() string {
	return SubjectChannelDeleted
}

type GroupMemberAdded struct{}

func (GroupMemberAdded) Subject() string {
	return SubjectGroupMemberAdded
}

type GroupMemeberRemoved struct{}

func (GroupMemeberRemoved) Subject() string {
	return SubjectGroupMemberRemoved
}
