package main

// Repository Interface for db
type Repository interface {
	Insert(task *Task) error
}
