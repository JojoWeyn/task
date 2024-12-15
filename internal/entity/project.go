package entity

import "time"

type Project struct {
	ID          uint
	Name        string
	Description string
	CreatedBy   uint
	CreatedAt   time.Time
}
