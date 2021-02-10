package service

import "time"

type CreateTaskResponse struct {
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"dateCreated"`
}
