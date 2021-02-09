package main

import "log"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s Service) CreateTask(request *TaskRequest, response *TaskResponse) error {
	log.Println("Startins ervice")
	t := Task{
		Name:        "namem",
		Description: "asfd",
		Status:      "TODO",
	}
	s.repo.Insert(t)
	return nil
}
