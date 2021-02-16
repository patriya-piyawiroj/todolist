package service

import "context"

// Service Interface
type Service interface {
	CreateTask(ctx context.Context, req *CreateTaskRequest) (*CreateTaskResponse, error)
	GetTaskByID(ctx context.Context, req *GetTaskRequest) (*GetTaskResponse, error)
}
