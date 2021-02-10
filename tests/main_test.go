package tests

import (
	"testing"
	"todolist/mocks"
	"todolist/models"
	"todolist/service"
)

func TestInsertTask(t *testing.T) {
	s := service.NewService(mocks.MockRepo{})

	mocks.GetInsertFunc = func(task *models.Task) error {
		return nil
	}
}
