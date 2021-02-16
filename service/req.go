package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateTaskRequest
type CreateTaskRequest struct {
	Name        string `json:"name" form:"name" query:"name"`
	Description string `json:"desc" form:"desc" query:"desc"`
	Status      string `json:"status" form:"status" query:"status"`
}

// GetTaskRequest
type GetTaskRequest struct {
	OID      primitive.ObjectID `json:"oid" form:"oid" query:"oid"`
	idString string             `json:"id" form:"id" query:"id"`
}
