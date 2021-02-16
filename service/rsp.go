package service

import (
	"time"
	"todolist/models"
)

// CreateTaskResponse
type CreateTaskResponse struct {
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"dateCreated"`
}

// GetTaskResponse
type GetTaskResponse struct {
	Task models.Task `json:"task" form:"task" query:"task"`
}
