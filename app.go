package main

import (
	"github.com/labstack/echo/v4"
)

type App struct {
	db   *MongoDB
	echo *echo.Echo
}

func (a *App) Initialize() {
	a.db = NewMongoDB()
	a.echo = echo.New()

	// Init
	validator := NewValidator()
	service := NewService(a.db)
	handler := NewHandler(service, validator)

	// Define all methods
	a.echo.POST("/v1/tasks", handler.createTaskHandler)
}

func (a *App) Run(addr string) {
	a.echo.Logger.Fatal(a.echo.Start(addr))
}
