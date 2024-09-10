package models

import (
	"time"
)

type TaskStatus string

const (
	Pending   TaskStatus = "pending"
	Done      TaskStatus = "done"
	Cancelled TaskStatus = "cancelled"
	NotDone   TaskStatus = "not_done"
)

type Priority string

const (
	Normal Priority = "normal"
	Low    Priority = "low"
	High   Priority = "high"
)

type Task struct {
	ID          int
	Title       string
	Description *string
	Status      TaskStatus
	CreatedAt   time.Time
	Priority    Priority
	DoneAt      *time.Time
}
