package main

import (
	"firstTask/internal/db"
	"firstTask/internal/handlers"
	"firstTask/internal/taskService"
	"firstTask/internal/userService"
	"firstTask/internal/web/tasks"
	"firstTask/internal/web/users"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	dbConn, err := db.InitDB()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	db.DB = dbConn

	tasksRepo := taskService.NewTaskRepository(db.DB)
	tasksService := taskService.NewTaskService(tasksRepo)

	usersRepo := userService.NewUserRepository(db.DB)
	usersService := userService.NewUserService(usersRepo, tasksService)

	tasksHandler := handlers.NewTaskHandler(tasksService)
	usersHandler := handlers.NewUserHandler(usersService)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	tasksStrictHandler := tasks.NewStrictHandler(tasksHandler, nil)
	tasks.RegisterHandlers(e, tasksStrictHandler)

	usersStrictHandler := users.NewStrictHandler(usersHandler, nil)
	users.RegisterHandlers(e, usersStrictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
