package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ragoncsa/todo/domain"
)

type CreateTaskRequest struct {
	Name string `json:"name" binding:"required"`
}

type TaskService struct {
	Service domain.TaskService
}

func (t *TaskService) GetTasks(c *gin.Context) {
	tasks, err := t.Service.Tasks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func (t *TaskService) GetTask(c *gin.Context) {
	id := c.Param("taskid")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "taskid must be an integer"})
		return
	}
	task, err := t.Service.Task(idInt)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("task with id %d not found", idInt)})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func (t *TaskService) CreateTask(c *gin.Context) {
	var request CreateTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task := &domain.Task{Name: request.Name}
	t.Service.CreateTask(task)
	c.IndentedJSON(http.StatusCreated, task)
}

func (t *TaskService) DeleteTask(c *gin.Context) {
	id := c.Param("taskid")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "taskid must be an integer"})
		return
	}
	err = t.Service.DeleteTask(idInt)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("task with id %d not found", idInt)})
		return
	}
	c.IndentedJSON(http.StatusOK, struct{}{})
}

func (t *TaskService) DeleteTasks(c *gin.Context) {
	err := t.Service.DeleteTasks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	c.IndentedJSON(http.StatusOK, struct{}{})
}
