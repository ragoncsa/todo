package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ragoncsa/todo/authz"
	"github.com/ragoncsa/todo/domain"
)

type TaskService struct {
	Service     domain.TaskService
	AuthzClient authz.Client
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
// @Param        CallerId  header  string  false  "the id of the caller"  "johndoe"
// @Security     JWT
func (t *TaskService) GetTasks(c *gin.Context) {
	tasks, err := t.Service.Tasks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	dreqVal, ok := c.Get("decisionRequest")
	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	dreq := dreqVal.(*authz.DecisionRequest)
	temp := tasks[:0]
	// inefficient way of sending authorization requests sequentially - ok for demoing
	for _, v := range tasks {
		dreq.Owner = v.UserId
		dreq.TaskID = strconv.Itoa(v.ID)
		allowed, err := t.AuthzClient.IsAllowed(dreq)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
			return
		}
		if allowed {
			temp = append(temp, v)
		}
	}
	tasks = temp
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
// @Param        CallerId  header  string  false  "the id of the caller"  "johndoe"
// @Security     JWT
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
	dreqVal, ok := c.Get("decisionRequest")
	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	dreq := dreqVal.(*authz.DecisionRequest)
	dreq.Owner = task.UserId
	dreq.TaskID = strconv.Itoa(task.ID)
	allowed, err := t.AuthzClient.IsAllowed(dreq)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
	} else if !allowed {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "forbidden"})
	} else {
		c.IndentedJSON(http.StatusOK, task)
	}
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
// @Param        CallerId  header  string  false  "the id of the caller"  "johndoe"
// @Security     JWT
func (t *TaskService) CreateTask(c *gin.Context) {
	var request CreateTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dreqVal, ok := c.Get("decisionRequest")
	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	dreq := dreqVal.(*authz.DecisionRequest)
	dreq.Owner = request.Task.UserId
	allowed, err := t.AuthzClient.IsAllowed(dreq)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
	} else if !allowed {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "forbidden"})
	} else if err := t.Service.CreateTask(request.Task.httpToModel()); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
	} else {
		c.IndentedJSON(http.StatusCreated, request.Task)
	}
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
// @Param        CallerId  header  string  false  "the id of the caller"  "johndoe"
// @Security     JWT
func (t *TaskService) DeleteTask(c *gin.Context) {
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
	dreqVal, ok := c.Get("decisionRequest")
	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	dreq := dreqVal.(*authz.DecisionRequest)
	dreq.Owner = task.UserId
	dreq.TaskID = strconv.Itoa(task.ID)
	allowed, err := t.AuthzClient.IsAllowed(dreq)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
	} else if !allowed {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "forbidden"})
	} else if err = t.Service.DeleteTask(idInt); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("task with id %d not found", idInt)})
	} else {
		c.IndentedJSON(http.StatusOK, struct{}{})
	}
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
// @Param        CallerId  header  string  false  "the id of the caller"  "johndoe"
// @Security     JWT
func (t *TaskService) DeleteTasks(c *gin.Context) {
	dreqVal, ok := c.Get("decisionRequest")
	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	dreq := dreqVal.(*authz.DecisionRequest)
	allowed, err := t.AuthzClient.IsAllowed(dreq)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
	} else if !allowed {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "forbidden"})
	} else if err := t.Service.DeleteTasks(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
	} else {
		c.IndentedJSON(http.StatusOK, struct{}{})
	}
}
