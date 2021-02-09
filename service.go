package main

import (
	"log"
	"net/http"
	"time"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s Service) CreateTask(req *TaskRequest, rsp *TaskResponse) error {
	log.Println("Starting service")
	t := Task{
		Name:        req.Name,
		Description: req.Description,
		Status:      StatusType(req.Status),
		CreatedAt:   time.Now(),
	}
	if err := s.repo.Insert(&t); err != nil {
		return NewError(ErrRepositoryInsert,
			http.StatusInternalServerError,
			err.Error(),
			ErrServiceInstance)
	}
	rsp.CreatedAt = t.CreatedAt
	rsp.Location = t.OID.String()
	return nil
}
