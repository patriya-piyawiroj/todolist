package main

import (
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
		return NewError(ErrContentType,
			http.StatusBadRequest,
			"Expected application/json",
			ErrValidationInstance)
	}

	// Bind request
	if err := c.Bind(req); err != nil {
		return NewError(ErrContentType,
			http.StatusBadRequest,
			"Could not bind to request",
			ErrValidationInstance)
	}

	// Validate fields
	if _, exist := StatusTypes[req.Status]; !exist {
		return NewError(ErrContentType,
			http.StatusBadRequest,
			"Not a valid status type",
			ErrValidationInstance)
	}

	return nil
}

// func (v Validator) validateGetTaskRequest(req *CreateTaskRequest) error {
// 	// TODO: Check ID length
// 	return nil
// }
