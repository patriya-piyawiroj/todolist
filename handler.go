package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	srv       *Service
	validator *Validator
}

func NewHandler(srv *Service, validator *Validator) *Handler {
	return &Handler{
		srv:       srv,
		validator: validator,
	}
}

// e.POST("/v1/tasks", createTaskHandler)
func (h Handler) createTaskHandler(c echo.Context) error {

	// Placeholders
	req := new(TaskRequest)
	rsp := new(TaskResponse)

	// Validate request
	if err := h.validator.validateCreateTaskRequest(c, req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Perform request logic
	if err := h.srv.CreateTask(req, rsp); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Return response
	return c.JSON(http.StatusOK, rsp)
}

type TaskResponse struct {
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"dateCreated"`
}

// // e.GET("/getDetailByID/:id", getDetailById)
// func getDetailByID(c echo.Context) error {
// 	id := c.Param("id")
// 	log.Println("Getting detail by id", id)
// 	return c.JSON(http.StatusOK, getByID(id, collection))
// }

// // e.GET("/getTodoList", getTodoList)
// func getTodoList(c echo.Context) error {
// 	log.Println("Getting all tasks")
// 	return c.JSON(http.StatusOK, getAll(collection))
// }

// // e.DELETE("/deleteTask/:id", deleteTask)
// func deleteTask(c echo.Context) error {
// 	id := c.Param("id")
// 	log.Println("Deleting detail by id", id)
// 	delete(id, collection)
// 	return c.String(http.StatusOK, "Delete Task Successful")
// 	//echo.NewHTTPError(http.StatusInternalServerError)
// }

// // e.PUT("/updateTask/:id", updateTask)
// func updateTask(c echo.Context) error {
// 	id := c.Param("id")
// 	log.Println("Updating task", id)
// 	t := Task{
// 		Name:        c.QueryParam("name"),
// 		Description: c.QueryParam("description"),
// 		Status:      StatusType(c.QueryParam("status")),
// 	}
// 	update(id, t, collection)
// 	return c.String(http.StatusOK, "Successful")
// }

// Handler
// func(c echo.Context) error {
// 	name := c.FormValue("name")
// 	return c.String(http.StatusOK, name)
// }

// type Timestamp time.Time

// func (t *Timestamp) UnmarshalParam(src string) error {
// 	ts, err := time.Parse(time.RFC3339, src)
// 	*t = Timestamp(ts)
// 	return err
// }

// type TaskRequest struct {}

// func (cb *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {
// 	// You may use default binder
// 	db := new(echo.DefaultBinder)
// 	if err = db.Bind(i, c); err != echo.ErrUnsupportedMediaType {
// 		return
// 	}

// 	// Define your custom implementation

// 	return
// }
