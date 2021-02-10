package mocks

import (
	"todolist/models"
)

type MockRepo struct {
	InsertFunc func(task *models.Task) error
}

func (m MockRepo) Insert(task *models.Task) error {
	return GetInsertFunc(task)
}

var (
	GetInsertFunc func(task *models.Task) error
)
