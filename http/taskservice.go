package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ragoncsa/todo/domain"
)

type JsonResponse struct {
	Type    string         `json:"type"`
	Data    []*domain.Task `json:"data"`
	Message string         `json:"message"`
}

type CreateTaskRequest struct {
	Name string `json:"name"`
}

type TaskService struct {
	Service domain.TaskService
}

func (t *TaskService) GetTasks(w http.ResponseWriter, r *http.Request) {
}

func (t *TaskService) GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["taskid"]

	var response = JsonResponse{}
	defer func() { json.NewEncoder(w).Encode(response) }()

	if id == "" {
		response = JsonResponse{Type: "error", Message: "taskid missing"}
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response = JsonResponse{Type: "error", Message: "taskid must be an integer"}
		return
	}
	task, err := t.Service.Task(idInt)
	if err != nil {
		response = JsonResponse{Type: "error", Message: err.Error()}
		return
	}

	response = JsonResponse{
		Type:    "success",
		Message: "task retrieved successfully",
		Data:    []*domain.Task{task},
	}
}

func (t *TaskService) CreateTask(w http.ResponseWriter, r *http.Request) {

	var response *JsonResponse
	defer func() { json.NewEncoder(w).Encode(&response) }()

	// name := r.FormValue("taskname")
	decoder := json.NewDecoder(r.Body)
	var request CreateTaskRequest
	err := decoder.Decode(&request)
	if err != nil {
		response = &JsonResponse{Type: "error", Message: "invalid JSON in request"}
		return
	}
	if request.Name == "" {
		response = &JsonResponse{Type: "error", Message: "taskname missing"}
		return
	}
	t.Service.CreateTask(&domain.Task{Name: request.Name})

	response = &JsonResponse{Type: "success", Message: "task created successfully"}
}

func (t *TaskService) DeleteTask(w http.ResponseWriter, r *http.Request) {
}

func (t *TaskService) DeleteTasks(w http.ResponseWriter, r *http.Request) {
}
