package controllers

import (
	"api/dtos"
	"api/models"
	"api/repositories"
	"api/services"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskController struct {
	taskRepository *repositories.TaskRepository
	taskService    *services.TaskService
}

func NewTaskController(
	tr *repositories.TaskRepository,
	ts *services.TaskService,
) *TaskController {
	return &TaskController{tr, ts}
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var (
		dto  dtos.CreateTaskDto
		err  error
		task *models.Task
	)

	if err := c.BindJSON(&dto); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if task, err = tc.taskService.CreateTask(c, dto.Description); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"task": task,
	})
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	var (
		err error
		id  uuid.UUID
		dto dtos.UpdateTaskDto
	)

	if err := c.BindJSON(&dto); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if taskId, ok := c.Params.Get("taskId"); !ok {
		c.Status(http.StatusBadRequest)
		return
	} else {
		id, err = uuid.Parse(taskId)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
	}

	tc.taskRepository.UpdateTask(id, dto)
	c.Status(http.StatusNoContent)
}

func (tc *TaskController) FindTasks(c *gin.Context) {
	var (
		err    error
		offset int64 = 0
		size   int64 = 10
		query  string
	)

	offsetQ := c.Query("offset")
	sizeQ := c.Query("size")
	query = c.Query("query")

	if offsetQ != "" {
		offset, err = strconv.ParseInt(offsetQ, 10, 64)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
	}
	if sizeQ != "" {
		size, err = strconv.ParseInt(sizeQ, 10, 64)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
	}

	if size > 50 {
		c.Status(http.StatusBadRequest)
		return
	}
	tasks, err := tc.taskService.FindTasks(c, offset, size, query)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

func (tc *TaskController) FindTaskById(c *gin.Context) {
	var (
		taskUuid uuid.UUID
		err      error
	)
	if taskId, ok := c.Params.Get("taskId"); !ok {
		c.Status(http.StatusBadRequest)
		return
	} else {
		taskUuid, err = uuid.Parse(taskId)
	}

	task, err := tc.taskService.FindTaskById(c, taskUuid)
	if err != nil {
		if errors.Is(err, repositories.ErrResourceNotFound) {
			c.Status(http.StatusNotFound)
		} else {
			c.Status(http.StatusBadRequest)
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	var (
		err      error
		taskUuid uuid.UUID
	)

	if taskId, ok := c.Params.Get("taskId"); !ok {
		c.Status(http.StatusBadRequest)
		return
	} else {
		taskUuid, err = uuid.Parse(taskId)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
	}

	_, err = tc.taskService.FindTaskById(c, taskUuid)
	if err != nil {
		if errors.Is(err, repositories.ErrResourceNotFound) {
			c.Status(http.StatusNotFound)
		} else {
			c.Status(http.StatusBadRequest)
		}
		return
	}

	err = tc.taskService.DeleteTask(c, taskUuid)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.Status(http.StatusNoContent)
}
