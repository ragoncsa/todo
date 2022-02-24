package mock

import "github.com/ragoncsa/todo/domain"

type TaskService struct {
	TaskFn      func(id int) (*domain.Task, error)
	TaskInvoked bool

	TasksFn      func() ([]*domain.Task, error)
	TasksInvoked bool

	CreateTaskFn      func(t *domain.Task) error
	CreateTaskInvoked bool

	DeleteTaskFn      func(id int) error
	DeleteTaskInvoked bool
}

func (s *TaskService) Task(id int) (*domain.Task, error) {
	s.TaskInvoked = true
	return s.TaskFn(id)
}

func (s *TaskService) Tasks() ([]*domain.Task, error) {
	s.TasksInvoked = true
	return s.TasksFn()
}

func (s *TaskService) CreateTask(t *domain.Task) error {
	s.CreateTaskInvoked = true
	return s.CreateTaskFn(t)
}

func (s *TaskService) DeleteTask(id int) error {
	s.DeleteTaskInvoked = true
	return s.DeleteTaskFn(id)
}
