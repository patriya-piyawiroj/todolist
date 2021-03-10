package tests

import (
	"net/http/httptest"
	"strings"
	"testing"
	"todolist/request"
	"todolist/service"

	"github.com/labstack/echo/v4"
)

func TestValidateUpdateTaskRequest(t *testing.T) {
	e := echo.New()
	v := service.NewValidator()
	reqBody := `{"name": "name", "desc": "desc", "status": "Todo"}`
	stringReader := strings.NewReader(reqBody)
	httpReq := httptest.NewRequest(echo.PUT, "http://localhost:1234/v1/tasks/602b48c4d8d855b476be2b27", stringReader)
	httpReq.Header.Set("Content-Type", "application/json")
	c := e.NewContext(httpReq, nil)
	updateTaskRequest := new(request.UpdateTaskRequest)
	updateTaskRequest.IDString = "602b48c4d8d855b476be2b27"
	err := v.ValidateUpdateTaskRequest(c, updateTaskRequest)
	if err != nil {
		t.Errorf("Expected no error but received %s", err.Error())
	}
}

func TestValidateDeleteTaskRequest(t *testing.T) {
	e := echo.New()
	v := service.NewValidator()
	httpReq := httptest.NewRequest(echo.DELETE, "http://localhost:1234/v1/tasks/602b48c4d8d855b476be2b27", nil)
	httpReq.Header.Set("Content-Type", "application/json")
	c := e.NewContext(httpReq, nil)
	deleteTaskRequest := new(request.DeleteTaskRequest)
	deleteTaskRequest.IDString = "602b48c4d8d855b476be2b27"
	err := v.ValidateDeleteTaskRequest(c, deleteTaskRequest)
	if err != nil {
		t.Errorf("Expected no error but received %s", err.Error())
	}
}

func TestValidateGetTaskRequest(t *testing.T) {
	// Init echo
	e := echo.New()
	v := service.NewValidator()

	// Test with invalid OID
	httpReq := httptest.NewRequest(echo.GET, "http://localhost:1234/v1/tasks/602b48c4d8d855b476be2b27", nil)
	httpReq.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)
	getTaskRequest := new(request.GetTaskRequest)
	getTaskRequest.IDString = "asdfasdf"
	err := v.ValidateGetTaskRequest(c, getTaskRequest)
	if !strings.Contains(err.Error(), "Invalid OID") {
		t.Errorf("Expected 'Invalid OID' but received %v", err.Error())
	}

	// Test with proper OID
	getTaskRequest.IDString = "602b48c4d8d855b476be2b27"
	err = v.ValidateGetTaskRequest(c, getTaskRequest)
	if err != nil {
		t.Errorf("Expected no error but received %s", err.Error())
	}
}

func TestValidateCreateTaskRequest(t *testing.T) {
	e := echo.New()
	v := service.NewValidator()
	emptyReq := new(request.CreateTaskRequest)

	malformedReq := `"name": "name", "desc": "desc", "status": "Todo"}`
	invalidTypeReq := `{"name": "name", "desc": "desc", "status": "INVALID"}`
	// properReq := `{"name": "name", "desc": "desc", "status": "Todo"}`

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

	// Test succesful
	// stringReader = strings.NewReader(properReq)
	// req = httptest.NewRequest(echo.POST, "http://localhost:1234/v1/tasks", stringReader)
	// req.Header.Set("Content-Type", "application/json")
	// c = e.NewContext(req, rec)
	// srv := service.NewTaskService(mocks.MockRepo{})
	// handler := service.NewHandler(srv, v)
	// handler.GetTaskHandler(c)
	// if !strings.Contains(rec.Body.String(), "200") {
	// 	t.Errorf("Expected 'reponse 200' but received %v", rec.Body.String())
	// }

}
