package mocks

import (
	"context"
	"todolist/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockRepo struct {
	InsertFunc func(ctx context.Context, task *models.Task) error
}

func (m MockRepo) Insert(ctx context.Context, task *models.Task) error {
	return GetInsertFunc(ctx, task)
}

func (m MockRepo) GetByID(ctx context.Context, id primitive.ObjectID) (models.Task, error) {
	return GetByIDFunc(ctx, id)
}

func (m MockRepo) GetAll(ctx context.Context) ([]models.Task, error) {
	return GetAllFunc(ctx)
}

func (m MockRepo) Delete(ctx context.Context, id primitive.ObjectID) (string, error) {
	return GetDeleteFunc(ctx, id)
}

func (m MockRepo) Update(ctx context.Context, id primitive.ObjectID, task *models.Task) (models.Task, error) {
	return GetUpdateFunc(ctx, id, task)
}

var (
	GetInsertFunc func(ctx context.Context, task *models.Task) error
	GetByIDFunc   func(ctx context.Context, id primitive.ObjectID) (models.Task, error)
	GetAllFunc    func(ctx context.Context) ([]models.Task, error)
	GetDeleteFunc func(ctx context.Context, id primitive.ObjectID) (string, error)
	GetUpdateFunc func(ctx context.Context, id primitive.ObjectID, task *models.Task) (models.Task, error)
)
