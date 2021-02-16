package service

// Repository Interface for db
type Repository interface {
	CreateTask(req *CreateTaskRequest) (*CreateTaskResponse, error)
}
