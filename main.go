package main

import (
	"firstTask/internal/db"
	"firstTask/internal/handlers"
	"firstTask/internal/taskService"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	e := echo.New()

	taskRepo := taskService.NewTaskRepository(database)
	taskServices := taskService.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(taskServices)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/tasks", taskHandlers.GetTask)
	e.POST("/tasks", taskHandlers.PostTask)
	e.PATCH("/tasks/:id", taskHandlers.UpdateTask)
	e.DELETE("/tasks/:id", taskHandlers.DeleteTask)

	e.Start("127.0.0.1:8080")
}
