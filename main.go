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
	tasksHandler := handlers.NewTaskHandler(tasksService)

	usersRepo := userService.NewUserRepository(db.DB)
	usersService := userService.NewUserService(usersRepo)
	usersHandler := handlers.NewUserHandler(usersService)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictHandler := tasks.NewStrictHandler(tasksHandler, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictHandler)

	usersStrictHandler := users.NewStrictHandler(usersHandler, nil)
	users.RegisterHandlers(e, usersStrictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
