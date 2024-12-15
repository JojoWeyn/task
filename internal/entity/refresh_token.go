package entity

import "time"

type RefreshToken struct {
	ID        uint
	Token     string
	UserID    uint
	CreatedAt time.Time
	ExpiresAt time.Time
}
