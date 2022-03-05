package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ragoncsa/todo/domain"
	"github.com/ragoncsa/todo/mock"
)

func TestGetTask(t *testing.T) {
	var ts mock.TaskService
	tsHTTP := &TaskService{Service: &ts}

	// Mock Task() call.
	ts.TaskFn = func(id int) (*domain.Task, error) {
		if id != 100 {
			t.Fatalf("unexpected id: %d", id)
		}
		return &domain.Task{ID: 100, Name: "my-task-1"}, nil
	}

	// Invoke the handler.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/tasks/100", nil)

	server := InitServer()
	server.RegisterRoutes(tsHTTP)
	server.router.ServeHTTP(w, r)

	// Validate mock.
	if !ts.TaskInvoked {
		t.Fatal("expected Task() to be invoked")
	}
}

func TestGetTasks(t *testing.T) {
	var ts mock.TaskService
	tsHTTP := &TaskService{Service: &ts}

	// Mock Tasks() call.
	ts.TasksFn = func() ([]*domain.Task, error) {
		return []*domain.Task{{ID: 100, Name: "my-task-1"}}, nil
	}

	// Invoke the handler.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/tasks/", nil)

	server := InitServer()
	server.RegisterRoutes(tsHTTP)
	server.router.ServeHTTP(w, r)

	// Validate mock.
	if !ts.TasksInvoked {
		t.Fatal("expected Tasks() to be invoked")
	}
}

func TestCreateTask(t *testing.T) {
	var ts mock.TaskService
	tsHTTP := &TaskService{Service: &ts}

	// Mock our CreateTask() call.
	ts.CreateTaskFn = func(task *domain.Task) error {
		if task.Name != "my-task-1" {
			t.Fatalf("unexpected name: %s", task.Name)
		}
		return nil
	}

	// Invoke the handler.
	w := httptest.NewRecorder()
	request, err := json.Marshal(&CreateTaskRequest{&Task{Name: "my-task-1"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
		return
	}
	reader := strings.NewReader(string(request))
	r, _ := http.NewRequest("POST", "/tasks/", reader)

	server := InitServer()
	server.RegisterRoutes(tsHTTP)
	server.router.ServeHTTP(w, r)

	// Validate mock.
	if !ts.CreateTaskInvoked {
		t.Fatal("expected CreateTask() to be invoked")
	}
}

func TestDeleteTask(t *testing.T) {
	var ts mock.TaskService
	tsHTTP := &TaskService{Service: &ts}

	// Mock DeleteTask() call.
	ts.DeleteTaskFn = func(id int) error {
		if id != 100 {
			t.Fatalf("unexpected id: %d", id)
		}
		return nil
	}

	// Invoke the handler.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/tasks/100", nil)

	server := InitServer()
	server.RegisterRoutes(tsHTTP)
	server.router.ServeHTTP(w, r)

	// Validate mock.
	if !ts.DeleteTaskInvoked {
		t.Fatal("expected DeleteTask() to be invoked")
	}
}

func TestDeleteTasks(t *testing.T) {
	var ts mock.TaskService
	tsHTTP := &TaskService{Service: &ts}

	// Mock Tasks() call.
	ts.DeleteTasksFn = func() error {
		return nil
	}

	// Invoke the handler.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/tasks/", nil)

	server := InitServer()
	server.RegisterRoutes(tsHTTP)
	server.router.ServeHTTP(w, r)

	// Validate mock.
	if !ts.DeleteTasksInvoked {
		t.Fatal("expected DeleteTasks() to be invoked")
	}
}
