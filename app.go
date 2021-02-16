package main

import (
	"context"
	"todolist/repo"
	"todolist/service"

	"github.com/labstack/echo/v4"
)

type App struct {
	db   *repo.MongoDB
	echo *echo.Echo
}

func (a *App) Initialize(addr string, db string, collection string) {
	a.db = repo.NewRepo(context.Background(), addr, db, collection)
	a.echo = echo.New()

	// Init
	validator := service.NewValidator()
	srv := service.NewTaskService(a.db)
	handler := service.NewHandler(srv, validator)

	// Define all methods
	a.echo.POST("/v1/tasks", handler.CreateTaskHandler)
}

func (a *App) Run(addr string) {
	a.echo.Logger.Fatal(a.echo.Start(addr))
}
