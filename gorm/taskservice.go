package gorm

import (
	"github.com/ragoncsa/todo/domain"
	"gorm.io/gorm"
)

type TaskService struct {
	DB *gorm.DB
}

func (s *TaskService) Task(id int) (*domain.Task, error) {
	var t domain.Task
	s.DB.First(&t, id)
	return &t, nil
}

func (s *TaskService) CreateTask(t *domain.Task) error {
	s.DB.Create(t)
	// error ?
	return nil
}
