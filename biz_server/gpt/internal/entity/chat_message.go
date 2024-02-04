package entity

import "time"

type ChatMessageEntity struct {
	ID        int32
	ChatID    int32
	RequestID string
	Version   int8
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
