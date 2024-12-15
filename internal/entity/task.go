package entity

import "time"

type Task struct {
	ID          uint
	ProjectID   uint
	Name        string
	Description string
	Status      string
	Deadline    time.Time
	AssignedTo  uint
	CreatedAt   time.Time
}
