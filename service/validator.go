package service

import (
	"net/http"
	"todolist/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Validator struct {
}

func NewValidator() *Validator {
	return &Validator{}
}

func validateHeader(c echo.Context) error {
	// Validate header
	// 4. X Header like http request id, request datetime
	contentType := c.Request().Header.Get("Content-Type")
	if contentType != "application/json" {
		return NewError(ErrContentType,
			http.StatusBadRequest,
			"Expected application/json",
			ErrValidationInstance)
	}
	return nil
}

// ValidateCreateTaskRequest
func (v Validator) ValidateCreateTaskRequest(c echo.Context, req *CreateTaskRequest) error {

	// Validate Header
	if err := validateHeader(c); err != nil {
		return err
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

// ValidateGetTaskRequest
func (v Validator) ValidateGetTaskRequest(c echo.Context, req *GetTaskRequest) error {
	// Validate Header
	if err := validateHeader(c); err != nil {
		return err
	}

	// Bind request
	if err := c.Bind(req); err != nil {
		return NewError(ErrContentType,
			http.StatusBadRequest,
			"Could not bind to request",
			ErrValidationInstance)
	}

	// Check for valid OID number
	if oid, err := primitive.ObjectIDFromHex(req.idString); err != nil {
		return NewError(ErrContentType,
			http.StatusBadRequest,
			"Invalid OID",
			ErrValidationInstance)
	} else {
		req.OID = oid
	}
	return nil
}
