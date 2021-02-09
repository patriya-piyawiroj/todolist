package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Validator struct {
}

func NewValidator() *Validator {
	return &Validator{}
}

func (v Validator) validateCreateTaskRequest(c echo.Context, req *TaskRequest) error {

	// Validate header
	contentType := c.Request().Header.Get("Content-Type")
	if contentType != "application/json" {
		errString := "incorrect content type: Expected application/json but received " + contentType
		return c.JSON(http.StatusBadRequest, errString)
	}

	// Bind request
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Validate fields
	if _, exist := StatusTypes[req.Status]; !exist {
		log.Println("Invalid status")
		return errors.New("Not a valid status")
	}

	return nil
}

// func (v Validator) validateGetTaskRequest(req *CreateTaskRequest) error {
// 	// TODO: Check ID length
// 	return nil
// }
