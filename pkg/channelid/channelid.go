package channelid

import (
	"github.com/google/uuid"
)

type Type byte

const (
	TypeDM        Type = 0x00 // 0000 0000
	TypeCommunity Type = 0x20 // 0010 0000

	typeMask Type = 0x20
)

func New(t Type) (uuid.UUID, error) {
	/*
		 0                   1                   2                   3
		 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		|                           unix_ts_ms                          |
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		|          unix_ts_ms           |  ver  |  rand_a (12 bit seq)  |
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		|var|T|                      rand_b                             |
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		|                            rand_b                             |
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

		T - custom channel type
	*/
	id, err := uuid.NewV7()
	if err != nil {
		return uuid.Nil, err
	}

	// Byte 8: [var: 2 bit] [Т: 1 bit] [rand_b: 5 bit]
	// Save bit var (0xC0) and rand_b (0x1F), cahnge bit type
	id[8] = (id[8] & ^byte(typeMask)) | byte(t)

	return id, nil
}

func Resolve(id uuid.UUID) Type {
	return Type(id[8] & byte(typeMask))
}
