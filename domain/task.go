package domain

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	UserId      string
	Name        string
	Description string
	Priority    int
	DueAt       time.Time
}
