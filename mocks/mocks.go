package mocks

import (
	"context"
	"todolist/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockRepo struct {
	InsertFunc func(task *models.Task) error
}

func (m MockRepo) Insert(ctx context.Context, task *models.Task) error {
	return GetInsertFunc(ctx, task)
}

func (m MockRepo) GetByID(ctx context.Context, id primitive.ObjectID) (models.Task, error) {
	return GetByIDFunc(ctx, id)
}

var (
	GetInsertFunc func(ctx context.Context, task *models.Task) error
	GetByIDFunc   func(ctx context.Context, id primitive.ObjectID) (models.Task, error)
)
