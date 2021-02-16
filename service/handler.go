package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	srv       *TaskService
	validator *Validator
}

func NewHandler(srv *TaskService, validator *Validator) *Handler {
	return &Handler{
		srv:       srv,
		validator: validator,
	}
}

// e.POST("/v1/tasks", CreateTaskHandler)
func (h *Handler) CreateTaskHandler(c echo.Context) error {

	// Placeholders
	req := new(CreateTaskRequest)
	// rsp := new(CreateTaskResponse)

	// Validate request
	if err := h.validator.ValidateCreateTaskRequest(c, req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Perform request logic
	rsp, err := h.srv.CreateTask(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Return response
	return c.JSON(http.StatusOK, rsp)
}

// e.GET("v1/tasks/:id", GetTaskHandler)
func (h *Handler) GetTaskHandler(c echo.Context) error {
	// Validate request
	req := new(GetTaskRequest)
	req.IdString = c.Param("id")
	if err := h.validator.ValidateGetTaskRequest(c, req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Perform request logic
	rsp, err := h.srv.GetTaskByID(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Return response
	return c.JSON(http.StatusOK, rsp)
}

// // e.GET("/tasks", getTodoList)
// func getTodoList(c echo.Context) error {
// 	log.Println("Getting all tasks")
// 	return c.JSON(http.StatusOK, getAll(collection))
// }

// // e.DELETE("/tasks/:id", deleteTask)
// func deleteTask(c echo.Context) error {
// 	id := c.Param("id")
// 	log.Println("Deleting detail by id", id)
// 	delete(id, collection)
// 	return c.String(http.StatusOK, "Delete Task Successful")
// 	//echo.NewHTTPError(http.StatusInternalServerError)
// }

// // e.PUT("/tasks/:id", updateTask)
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
