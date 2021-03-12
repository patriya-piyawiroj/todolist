package main

import (
	"context"
	"fmt"
	"os"
	"todolist/app"
	"todolist/configs"
	"todolist/handler"
	"todolist/repo"
	"todolist/service"

	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v3"
)

// Main function
func main() {

	// ===================== init configs =============================
	f, err := os.Open("configs/config.yml")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	var conf configs.Configs
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Println(err)
	}

	// ===================== init redis =============================
	app.InitRedis(&conf)

	// ===================== init db =============================
	db := repo.NewRepo(context.Background(), conf.Database.Address, conf.Database.DB, conf.Database.Collection)
	validator := service.NewValidator()
	srv := service.NewTaskService(db, conf)
	handler := handler.NewHandler(srv, validator)

	//  ===================== define routes =============================
	e := echo.New()
	e.POST("/v1/tasks", handler.CreateTaskHandler)
	e.GET("/v1/tasks/:id", handler.GetTaskHandler)
	e.GET("/v1/tasks", handler.GetAllHandler)
	e.DELETE("/v1/tasks/:id", handler.DeleteTaskHandler)
	e.PUT("/v1/tasks/:id", handler.UpdateTaskHandler)

	e.Logger.Fatal(e.Start(conf.Server.Port))
}
