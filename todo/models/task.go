package models

import "time"

type Task struct {
	Id          int
	Title       string
	Description string
	IsCompleted bool
	DueDate     time.Time
	UserId      int
}

func (task *Task) markAsCompleted() {

	task.IsCompleted = true
}
