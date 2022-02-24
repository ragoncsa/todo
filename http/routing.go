package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// TaskServiceHTTPHandlers defines all the handlers the TaskService needs. It's
// possible to register routes for a different implementation (like a mock).
type TaskServiceHTTPHandlers interface {
	GetTask(w http.ResponseWriter, r *http.Request)
	GetTasks(w http.ResponseWriter, r *http.Request)
	CreateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
	DeleteTasks(w http.ResponseWriter, r *http.Request)
}

type Server struct {
	router *mux.Router
}

func InitServer() *Server {
	return &Server{router: mux.NewRouter()}
}

func (s Server) RegisterRoutes(t TaskServiceHTTPHandlers) {
	s.router.HandleFunc("/tasks/", t.GetTasks).Methods("GET")
	s.router.HandleFunc("/tasks/{taskid}", t.GetTask).Methods("GET")
	s.router.HandleFunc("/tasks/", t.CreateTask).Methods("POST")
	s.router.HandleFunc("/tasks/{taskid}", t.DeleteTask).Methods("DELETE")
	s.router.HandleFunc("/tasks/", t.DeleteTasks).Methods("DELETE")
}

func (s Server) Start() {
	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8080", s.router))
}
