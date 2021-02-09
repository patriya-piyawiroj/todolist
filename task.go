package main

//https://kscm.kasikornbank.com:8443/781/go-loyalty-new/-/blob/PRD/internal/pkg/ais/model.go

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// StatusType for a task
	StatusType string

	// Task in TodoList
	Task struct {
		OID         primitive.ObjectID `json:"id" bson:"omitempty"`
		Name        string             `json:"name" bson:"name"`
		Description string             `json:"description" bson:"description"`
		Status      StatusType         `json:"status" bson:"status"`
		CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	}
)
