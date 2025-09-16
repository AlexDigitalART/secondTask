package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Task struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

type taskRequest struct {
	Task string `json:"task"`
}

type UpdateTaskRequest struct {
	Task string `json:"task"`
}

var tasks []Task

func getTask(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

func postTask(c echo.Context) error {
	var req taskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if req.Task == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "task cannot be empty"})
	}

	task := Task{
		ID:   uuid.NewString(),
		Task: req.Task,
	}

	tasks = append(tasks, task)

	return c.JSON(http.StatusCreated, task)
}

func updateTask(c echo.Context) error {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task ID format"})
	}

	for i, task := range tasks {
		if task.ID == id {
			var req UpdateTaskRequest
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
			}

			if req.Task == "" {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "task text cannot be empty"})
			}

			tasks[i].Task = req.Task

			return c.JSON(http.StatusOK, tasks[i])
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task ID format"})
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)

			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/tasks", getTask)
	e.POST("/tasks", postTask)
	e.PATCH("/tasks/:id", updateTask)
	e.DELETE("/tasks/:id", deleteTask)

	e.Start("127.0.0.1:8080")
}
