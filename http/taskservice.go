package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ragoncsa/todo/domain"
)

type TaskService struct {
	Service domain.TaskService
}

// GetTasks godoc
// @Summary  Get all tasks
// @Schemes
// @Description  Reads and returns all the tasks.
// @Tags         read
// @Accept       json
// @Produce      json
// @Success      200      {array}   domain.Task
// @Failure      default  {string}  string  "unexpected error"
// @Router       /tasks/ [get]
func (t *TaskService) GetTasks(c *gin.Context) {
	tasks, err := t.Service.Tasks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

// GetTask godoc
// @Summary  Get task
// @Schemes
// @Description  Reads a single task and returns it.
// @Tags         read
// @Accept       json
// @Produce      json
// @Param        taskid   path      int  true  "Task ID"
// @Success      200      {object}  domain.Task
// @Failure      401      {string}  string  "not found"
// @Failure      default  {string}  string  "unexpected error"
// @Router       /tasks/{taskid} [get]
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

// CreateTask godoc
// @Summary  Creates task
// @Schemes
// @Description  Creates a new task.
// @Tags         write
// @Accept       json
// @Produce      json
// @Param        task  body  CreateTaskRequest  true  "New task"
// @Success      200
// @Router       /tasks/ [post]
func (t *TaskService) CreateTask(c *gin.Context) {
	var request CreateTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := t.Service.CreateTask(request.Task.httpToModel()); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	c.IndentedJSON(http.StatusCreated, request.Task)
}

// DeleteTasks godoc
// @Summary  Deletes task
// @Schemes
// @Description  Deletes a single task.
// @Tags         write
// @Accept       json
// @Produce      json
// @Param        taskid  path  int  true  "Task ID"
// @Success      200
// @Router       /tasks/{taskid} [delete]
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

// DeleteTask godoc
// @Summary  Delete all tasks
// @Schemes
// @Description  Deletes all the tasks.
// @Tags         write
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /tasks/ [delete]
func (t *TaskService) DeleteTasks(c *gin.Context) {
	err := t.Service.DeleteTasks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	c.IndentedJSON(http.StatusOK, struct{}{})
}
