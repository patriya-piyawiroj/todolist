package repo

import (
	"context"
	"todolist/models"
)

// Repository Interface for db
type Repository interface {
	Insert(ctx context.Context, task *models.Task) error
}
