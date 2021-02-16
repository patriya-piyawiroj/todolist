package tests

import (
	"context"
	"testing"
	"todolist/mocks"
	"todolist/models"
	"todolist/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TestCreateTask with: go test ./... -v -coverpkg=../service/
func TestCreateTask(t *testing.T) {
	req := service.CreateTaskRequest{
		Name:        "test-name",
		Description: "test-desc",
		Status:      "status",
	}
	rsp := new(service.CreateTaskResponse)
	ID := primitive.NewObjectID()

	mocks.GetInsertFunc = func(ctx context.Context, task *models.Task) error {
		task.OID = ID
		ctx = context.Background()
		return nil
	}
	s := service.NewTaskService(mocks.MockRepo{})
	s.CreateTask(context.Background(), &req)
	if rsp.Location != ID.String() {
		t.Errorf("Expected %s but received %s", ID.String(), rsp.Location)
	}
}

func TestGetTask(t *testing.T) {
	oid, _ := primitive.ObjectIDFromHex("602b48c4d8d855b476be2b27")
	req := service.GetTaskRequest{
		OID:      oid,
		IdString: "602b48c4d8d855b476be2b27",
	}
	var task models.Task

	mocks.GetByIDFunc = func(ctx context.Context, id primitive.ObjectID) (models.Task, error) {
		task.OID = req.OID
		ctx = context.Background()
		return task, nil
	}
	s := service.NewTaskService(mocks.MockRepo{})
	rsp, _ := s.GetTaskByID(context.Background(), &req)
	if rsp.Task.OID != oid {
		t.Errorf("Expected %s but received %s", oid, rsp.Task.OID)
	}
}
