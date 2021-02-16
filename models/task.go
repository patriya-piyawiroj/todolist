package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (

	// Task in TodoList
	Task struct {
		OID         primitive.ObjectID `json:"id" bson:"_id, omitempty"`
		Name        string             `json:"name" bson:"name"`
		Description string             `json:"description" bson:"description"`
		Status      string             `json:"status" bson:"status"`
		CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	}
)
