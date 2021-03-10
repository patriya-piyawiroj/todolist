package service

import "fmt"

var (
	ErrContentType = "Incorrect content type"
	ErrRepository  = "Database error"

	ErrValidationInstance = "/todolist/validator"
	ErrServiceInstance    = "/todolist/service"
)

type Error struct {
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

func NewError(title string, status int, detail string, instance string) *Error {
	return &Error{
		Title:    title,
		Status:   status,
		Detail:   detail,
		Instance: instance,
	}
}

func (r *Error) Error() string {
	return fmt.Sprintf("Error %d %v: %v at %v", r.Status, r.Title, r.Detail, r.Instance)
}
