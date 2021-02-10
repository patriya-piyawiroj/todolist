package tests

import (
	"testing"
	"todolist/mocks"
	"todolist/models"
	"todolist/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertTask(t *testing.T) {
	req := service.CreateTaskRequest{
		Name:        "test-name",
		Description: "test-desc",
		Status:      "status",
	}
	rsp := new(service.CreateTaskResponse)
	ID := primitive.NewObjectID()

	mocks.GetInsertFunc = func(task *models.Task) error {
		task.OID = ID
		return nil
	}
	s := service.NewService(mocks.MockRepo{})
	s.CreateTask(&req, rsp)
	if rsp.Location != ID.String() {
		t.Errorf("Did not properly return task")
	}
}
