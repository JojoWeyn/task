package entity

import "time"

type User struct {
	ID           uint
	Email        string
	PasswordHash string

	IsActive       bool
	ActivationCode string
	ResetCode      string
	CreatedAt      time.Time

	FirstName   string
	LastName    string
	MiddleName  string
	Position    string
	TelegramID  string
	PhoneNumber string
}
