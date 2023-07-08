package models

import "time"

type Task struct {
	Name        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Tasks []Task
