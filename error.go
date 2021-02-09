package main

import "errors"

var (
	ErrIDNotFound = errors.New("ID no found")
)

type Error struct {
	title    string `json:"title"`
	status   int    `json:"status"`
	detail   string `json:"detail"`
	instance string `json:"instance"`
}

func NewError(title string, status int, detail string, instance string) *Error {
	return &Error{
		title:    title,
		status:   status,
		detail:   detail,
		instance: instance,
	}
}

// func (r *Error) Error() string {
// 	return fmt.Sprintf("status %d: err %v", r.status, r.title)
// }
