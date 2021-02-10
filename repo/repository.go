package repo

import "todolist/models"

// Repository Interface for db
type Repository interface {
	Insert(task *models.Task) error
}
