package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// Pointer to Collection
	collection *mongo.Collection
)

type (
	// StatusType for a task
	StatusType string

	// Task in TodoList
	Task struct {
		ID          primitive.ObjectID `json:"id" bson:"omitempty"`
		Name        string             `json:"name" bson:"name"`
		Description string             `json:"description" bson:"description"`
		Status      StatusType         `json:"status" bson:"status"`
	}
)

// Const
const (
	TestID = "601a45fde33637ef17c32275"

	Todo       StatusType = "TODO"
	InProgress StatusType = "IN_PROGRESS"
	Done       StatusType = "DONE"
)

// Main function
func main() {
	e := echo.New()

	// Define all methods
	e.GET("/getDetailByID/:id", getDetailByID)
	e.GET("/getTodoList", getTodoList)
	e.POST("/createTask", createTask)
	e.PUT("/updateTask/:id", updateTask)
	e.DELETE("/deleteTask/:id", deleteTask)

	// Open server connection
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error connecting to db:27017", err)
		os.Exit(1)
	}
	collection = client.Database("todolist").Collection("tasks")
	log.Println("Connected to mongo", client)

	// Start service
	e.Logger.Fatal(e.Start(":1234"))
}

// e.GET("/getDetailByID/:id", getDetailById)
func getDetailByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		id = TestID
	}
	log.Println("Getting detail by id", id)
	return c.JSON(http.StatusOK, getByID(id, collection))
}

// e.GET("/getTodoList", getTodoList)
func getTodoList(c echo.Context) error {
	log.Println("Getting all tasks")
	return c.JSON(http.StatusOK, getAll(collection))
}

// e.POST("/createTask", createTask)
func createTask(c echo.Context) error {
	log.Println("Creating new task")
	t := Task{
		Name:        c.QueryParam("name"),
		Description: c.QueryParam("description"),
		Status:      StatusType(c.QueryParam("status")),
	}
	insert(t, collection)
	return c.String(http.StatusOK, "Successful")
}

// e.DELETE("/deleteTask/:id", deleteTask)
func deleteTask(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		id = TestID
	}
	log.Println("Deleting detail by id", id)
	delete(id, collection)
	return c.String(http.StatusOK, "Successful")
}

// e.PUT("/updateTask/:id", updateTask)
func updateTask(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		id = TestID
	}
	log.Println("Updating task", id)
	t := Task{
		Name:        c.QueryParam("name"),
		Description: c.QueryParam("description"),
		Status:      StatusType(c.QueryParam("status")),
	}
	update(id, t, collection)
	return c.String(http.StatusOK, "Successful")
}
