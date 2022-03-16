package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ragoncsa/todo/config"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// TaskServiceHTTPHandlers defines all the handlers the TaskService needs. It's
// possible to register routes for a different implementation (like a mock).
type TaskServiceHTTPHandlers interface {
	GetTask(c *gin.Context)
	GetTasks(c *gin.Context)
	CreateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	DeleteTasks(c *gin.Context)
}

type Server struct {
	router *gin.Engine
	port   int
}

func InitServer(conf *config.Config) *Server {
	return &Server{
		router: gin.Default(),
		port:   conf.Server.Port,
	}
}

func (s Server) RegisterRoutes(t TaskServiceHTTPHandlers) {
	s.router.GET("/tasks/", t.GetTasks)
	s.router.GET("/tasks/:taskid", t.GetTask)
	s.router.POST("/tasks/", t.CreateTask)
	s.router.DELETE("/tasks/:taskid", t.DeleteTask)
	s.router.DELETE("/tasks/", t.DeleteTasks)

	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))

}

func (s Server) Start() {
	fmt.Printf("Server at %d\n", s.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.router))
}
