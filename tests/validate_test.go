package tests

import (
	"net/http/httptest"
	"strings"
	"testing"
	"todolist/service"

	"github.com/labstack/echo/v4"
)

func TestValidateGetTaskRequest(t *testing.T) {
	// e := echo.New()
	// v := service.NewValidator()
	// rec := httptest.NewRecorder()
	// getTaskRequest := new(service.GetTaskRequest)
	// getTaskRequest.IdString = "asdfasdf"

	// // Test without id specified

	// // Test with invalid OID
	// stringReader := strings.NewReader("")
	// req := httptest.NewRequest(echo.GET, "http://localhost:1234/v1/tasks/602b48c4d8d855b476be2b27", stringReader)
	// req.Header.Set("Content-Type", "application/json")
	// c := e.NewContext(req, rec)
	// err := v.ValidateGetTaskRequest(c, getTaskRequest)
	// if !strings.Contains(err.Error(), "Naot") {
	// 	t.Errorf("Should have resulted in invalid OID error")
	// }

	// // Test with proper OID
	// getTaskRequest.IdString = "602b48c4d8d855b476be2b27"
}

func TestValidateCreateTaskRequest(t *testing.T) {
	e := echo.New()
	v := service.NewValidator()
	emptyReq := new(service.CreateTaskRequest)

	malformedReq := `"name": "name", "desc": "desc", "status": "Todo"}`
	invalidTypeReq := `{"name": "name", "desc": "desc", "status": "INVALID"}`

	stringReader := strings.NewReader(malformedReq)
	req := httptest.NewRequest(echo.POST, "http://localhost:1234/v1/tasks", stringReader)
	req.Header.Set("X-Request-ID", "val")
	req.Header.Set("X-Request-DateTime", "Date: Wed, 16 Oct 2019 07:28:00 GMT ")
	req.Header.Set("Content-Type", "multipart/form-data")

	// Test with invalid Content-Type Header
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := v.ValidateCreateTaskRequest(c, emptyReq)
	if !strings.Contains(err.Error(), "Expected application/json") {
		t.Errorf("Should have resulted in invalid header error")
	}

	// Test with malformed request body
	req.Header.Set("Content-Type", "application/json")
	c = e.NewContext(req, rec)
	err = v.ValidateCreateTaskRequest(c, emptyReq)
	if !strings.Contains(err.Error(), "Could not bind to request") {
		t.Errorf("Should have resulted in malformed request error")
	}

	// Test with invalid status type field
	stringReader = strings.NewReader(invalidTypeReq)
	req = httptest.NewRequest(echo.POST, "http://localhost:1234/v1/tasks", stringReader)
	req.Header.Set("Content-Type", "application/json")
	c = e.NewContext(req, rec)
	err = v.ValidateCreateTaskRequest(c, emptyReq)
	if !strings.Contains(err.Error(), "Not a valid status type") {
		t.Errorf("Should have resulted in invalid status type error")
	}

}
