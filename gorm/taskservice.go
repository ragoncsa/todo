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
	tx := s.DB.First(&t, id)
	return &t, tx.Error
}

func (s *TaskService) CreateTask(t *domain.Task) error {
	tx := s.DB.Create(t)
	return tx.Error
}

func (s *TaskService) Tasks() ([]*domain.Task, error) {
	return nil, nil
}

func (s *TaskService) DeleteTask(id int) error {
	return nil
}
