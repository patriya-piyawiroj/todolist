package main

import "time"

type TaskResponse struct {
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"dateCreated"`
}
