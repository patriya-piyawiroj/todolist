package service

import "context"

// Repository Interface for db
type Repository interface {
	CreateTask(ctx context.Context, req *CreateTaskRequest) (*CreateTaskResponse, error)
	GetTaskByID(ctx context.Context, req *GetTaskRequest) (*GetTaskResponse, error)
}
