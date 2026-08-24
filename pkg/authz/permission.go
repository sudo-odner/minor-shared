package authz

type Permission uint64

const (
	// Base rules
	PermViewChannel Permission = 1 << 0 // View channel / read history

	// Text channel
	PermSendMessages  Permission = 1 << 5 // Send/edit/delete your message
	PermAttachFiles   Permission = 1 << 6 // Attach files/image
	PermMenageMessage Permission = 1 << 7 // [Modereate] Can delete all message in channel

	// Голосовые права

	// PermVoiceConnect Позволяет входить в голосовой канал
	PermVoiceConnect Permission = 1 << 10

	// PermVoiceSpeak Позволяет говорить в голосовом канал
	PermVoiceSpeak Permission = 1 << 11

	// PermVoiceStream Позволяет транслировать видео/экран
	PermVoiceStream Permission = 1 << 12

	// PermVoiceMute позволяет выключать голос, видео
	PermVoiceMute Permission = 1 << 13

	// Общая модерация

	// PermKickMembers позволяет выгонять или банить участников чата/сервера
	PermKickMembers Permission = 1 << 20

	// PermManageChannels позволяет создавать, удалять или переименовыввать каналы
	PermManageChannels Permission = 1 << 21

	// PermManageNicknames позволяет изменить ник вем ползователям
	PermManageNicknames Permission = 1 << 22

	// PermManageNicknames позвоялет изменять ник только себе
	PermChangeNicknames Permission = 1 << 22

	// PermManageRole позволяет управлять ролями
	PermManageRole Permission = 1 << 23
)

// Has verification based on all of the criteria
func Has(mask, perm Permission) bool {
	return (mask & perm) == perm
}

// HasAll verification based on all of the criteria
func HasAll(mask Permission, perms ...Permission) bool {
	var required Permission
	for _, p := range perms {
		required |= p
	}
	return (mask & required) == required
}

// HasAny verification based on one of the criteria
func HasAny(mask Permission, perms ...Permission) bool {
	for _, p := range perms {
		if (mask & p) == p {
			return true
		}
	}
	return false
}
