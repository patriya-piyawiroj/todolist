package service

import (
	"context"
	"net/http"
	"time"
	"todolist/models"
	"todolist/repo"
)

// add interface
type TaskService struct {
	repo repo.Repository
}

// NewService
func NewTaskService(repo repo.Repository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

// CreateTask
func (s *TaskService) CreateTask(ctx context.Context, req *CreateTaskRequest) (*CreateTaskResponse, error) {
	t := models.Task{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
		CreatedAt:   time.Now(),
	}
	if err := s.repo.Insert(ctx, &t); err != nil {
		return nil, NewError(ErrRepositoryInsert,
			http.StatusInternalServerError,
			err.Error(),
			ErrServiceInstance)
	}
	rsp := new(CreateTaskResponse)
	rsp.CreatedAt = t.CreatedAt
	rsp.Location = t.OID.String()
	return rsp, nil
}
