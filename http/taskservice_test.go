package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ragoncsa/todo/domain"
	"github.com/ragoncsa/todo/mock"
)

func TestHandler(t *testing.T) {
	// Inject our mock into our handler.
	var ts mock.TaskService
	tsHTTP := &TaskService{Service: &ts}

	// Mock our User() call.
	ts.TaskFn = func(id int) (*domain.Task, error) {
		if id != 100 {
			t.Fatalf("unexpected id: %d", id)
		}
		return &domain.Task{ID: 100, Name: "my-task-1"}, nil
	}

	// Invoke the handler.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/tasks/100", nil)

	router := InitServer()
	router.HandleFunc("/tasks/{taskid}", tsHTTP.GetTask)
	router.ServeHTTP(w, r)

	// Validate mock.
	if !ts.TaskInvoked {
		t.Fatal("expected Task() to be invoked")
	}
}
