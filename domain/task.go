package domain

import (
	"time"
)

type Task struct {
	ID          int `gorm:"primarykey"`
	UserId      string
	Name        string
	Description string
	Priority    int
	DueAt       time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TaskService interface {
	Task(id int) (*Task, error)
	Tasks() ([]*Task, error)
	CreateTask(t *Task) error
	DeleteTask(id int) error
	DeleteTasks() error
}
