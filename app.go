package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

type App struct {
	// Router *mux.Router
	// DB     *sql.DB
	service *Service
	db      *MongoDB
	echo    *echo.Echo
}

func (a *App) Initialize(user, password, dbname string) {
	a.db = NewMongoDB()
	a.echo = echo.New()

	// Init
	a.db = NewMongoDB()
	a.service = NewService(a.db)
	handler := NewHandler(a.service)

	// Define all methods
	a.echo.POST("/v1/tasks", handler.createTaskHandler)
	log.Println(a.db, a.service, handler)
}

func (a *App) Run(addr string) {
	a.echo.Logger.Fatal(a.echo.Start(addr))
}
