package main

import (
	"todolist/repo"
	"todolist/service"

	"github.com/labstack/echo/v4"
)

type App struct {
	db   *repo.MongoDB
	echo *echo.Echo
}

func (a *App) Initialize() {
	a.db = repo.NewMongoDB()
	a.echo = echo.New()

	// Init
	validator := service.NewValidator()
	srv := service.NewService(a.db)
	handler := service.NewHandler(srv, validator)

	// Define all methods
	a.echo.POST("/v1/tasks", handler.CreateTaskHandler)
}

func (a *App) Run(addr string) {
	a.echo.Logger.Fatal(a.echo.Start(addr))
}
