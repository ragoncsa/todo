package http

import (
	"strings"
	"time"

	"github.com/ragoncsa/todo/domain"
)

// This section is implemented based on
// https://github.com/swaggo/swag#use-swaggertype-tag-to-supported-custom-type
type TimestampTime struct {
	time.Time
}

///implement encoding.JSON.Marshaler interface
func (t *TimestampTime) MarshalJSON() ([]byte, error) {
	bin := make([]byte, 0, len("2019-10-12T07:20:50.52Z"))
	s := "\"" + t.Format(time.RFC3339) + "\""
	bin = append(bin, s...)
	return bin, nil
}

func (t *TimestampTime) UnmarshalJSON(bin []byte) error {
	s := strings.Trim(string(bin), string([]byte{0}))
	s = strings.Trim(s, "\"")
	parsedTime, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}

// Task is a struct with a subset of the fields of domain.Task. It is used when
// task needs to be provided as an input for task creation. So it excludes
// auto-generated fields.
type Task struct {
	UserId      string        `json:"userId" example:"johndoe"`
	Name        string        `json:"name" example:"my-task-1"`
	Description string        `json:"description" example:"description of my-task-1"`
	Priority    int           `json:"priority" example:"1" format:"int64"`
	DueAt       TimestampTime `json:"dueAt" swaggertype:"primitive,string" example:"2019-10-12T07:20:50.52Z"`
}

type CreateTaskRequest struct {
	Task *Task `json:"task" binding:"required"`
}

func (t *Task) httpToModel() *domain.Task {
	return &domain.Task{
		UserId:      t.UserId,
		Name:        t.Name,
		Description: t.Description,
		Priority:    t.Priority,
		DueAt:       t.DueAt.Time,
	}
}
