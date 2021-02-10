package service

import (
	"net/http"
	"todolist/models"

	"github.com/labstack/echo/v4"
)

type Validator struct {
}

func NewValidator() *Validator {
	return &Validator{}
}

// ValidateCreateTaskRequest
func (v Validator) ValidateCreateTaskRequest(c echo.Context, req *CreateTaskRequest) error {

	// Validate header
	// 4. X Header like http request id, request datetime
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
	if _, exist := models.StatusTypes[req.Status]; !exist {
		return NewError(ErrContentType,
			http.StatusBadRequest,
			"Not a valid status type",
			ErrValidationInstance)
	}

	return nil
}
