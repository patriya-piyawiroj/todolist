package main

import (
	"log"
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
	err := s.repo.Insert(t)
	// if err != nil {
	// 	return NewError("Could not create new task", 500, err, RepoLayerError)
	// }
	return err
}
