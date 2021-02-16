package repo

import (
	"context"
	"todolist/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Repository Interface for db
type Repository interface {
	Insert(ctx context.Context, task *models.Task) error
	GetByID(ctx context.Context, id primitive.ObjectID) (models.Task, error)
}
