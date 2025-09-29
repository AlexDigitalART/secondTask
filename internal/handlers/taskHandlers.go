package handlers

import (
	"firstTask/internal/taskService"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service taskService.TasksService
}

func NewTaskHandler(service taskService.TasksService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) GetTask(c echo.Context) error {
	task, err := h.service.GetAllTask()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not get calculations"})
	}

	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) PostTask(c echo.Context) error {
	var req taskService.TaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if req.Task == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "task cannot be empty"})
	}

	task, err := h.service.CreateTask(req.Task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not create task"})
	}

	return c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) UpdateTask(c echo.Context) error {
	id := c.Param("id")

	var req taskService.UpdateTaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	updatedTask, err := h.service.UpdateTask(id, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not update task"})
	}

	return c.JSON(http.StatusOK, updatedTask)
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task ID format"})
	}

	err := h.service.DeleteTask(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not delete task"})
	}

	return c.NoContent(http.StatusNoContent)
}
