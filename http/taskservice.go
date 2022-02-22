package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ragoncsa/todo/domain"
)

type JsonResponse struct {
	Type    string        `json:"type"`
	Data    []domain.Task `json:"data"`
	Message string        `json:"message"`
}

type TaskService struct {
	Service domain.TaskService
	Server  *mux.Router
}

func (t *TaskService) RegisterRoutes() {
	t.Server.HandleFunc("/tasks/", t.GetTasks).Methods("GET")
	t.Server.HandleFunc("/tasks/", t.CreateTask).Methods("POST")
	t.Server.HandleFunc("/tasks/{taskid}", t.DeleteTask).Methods("DELETE")
	t.Server.HandleFunc("/tasks/", t.DeleteTasks).Methods("DELETE")
}

func (t *TaskService) GetTasks(w http.ResponseWriter, r *http.Request) {
}

func (t *TaskService) CreateTask(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("taskname")

	var response = JsonResponse{}
	defer json.NewEncoder(w).Encode(response)

	if name == "" {
		response = JsonResponse{Type: "error", Message: "You are missing taskname parameter."}
		return
	}
	t.Service.CreateTask(&domain.Task{Name: name})

	response = JsonResponse{Type: "success", Message: "task created successfully"}
}

func (t *TaskService) DeleteTask(w http.ResponseWriter, r *http.Request) {
}

func (t *TaskService) DeleteTasks(w http.ResponseWriter, r *http.Request) {
}
