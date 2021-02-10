package main

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestValidateCreateTaskRequest(t *testing.T) {
	e := echo.New()
	v := NewValidator()
	emptyReq := new(TaskRequest)

	malformedReq := `"name": "name", "desc": "desc", "status": "Todo"}`
	invalidTypeReq := `{"name": "name", "desc": "desc", "status": "INVALID"}`

	stringReader := strings.NewReader(malformedReq)
	req := httptest.NewRequest(echo.POST, "http://localhost:1234/v1/tasks", stringReader)
	req.Header.Set("X-Request-ID", "val")
	req.Header.Set("X-Request-Header", "val")
	req.Header.Set("Content-Type", "multipart/form-data")

	// Test with invalid Content-Type Header
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := v.validateCreateTaskRequest(c, emptyReq)
	if !strings.Contains(err.Error(), "Expected application/json") {
		t.Errorf("Should have resulted in invalid header error")
	}

	// Test with malformed request body
	req.Header.Set("Content-Type", "application/json")
	c = e.NewContext(req, rec)
	err = v.validateCreateTaskRequest(c, emptyReq)
	if !strings.Contains(err.Error(), "Could not bind to request") {
		t.Errorf("Should have resulted in malformed request error")
	}

	// Test with invalid status type field
	stringReader = strings.NewReader(invalidTypeReq)
	req = httptest.NewRequest(echo.POST, "http://localhost:1234/v1/tasks", stringReader)
	req.Header.Set("Content-Type", "application/json")
	c = e.NewContext(req, rec)
	err = v.validateCreateTaskRequest(c, emptyReq)
	if !strings.Contains(err.Error(), "Not a valid status type") {
		t.Errorf("Should have resulted in invalid status type error")
	}

}
