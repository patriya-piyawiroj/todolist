package service

// TaskRequest
type CreateTaskRequest struct {
	Name        string `json:"name" form:"name" query:"name"`
	Description string `json:"desc" form:"desc" query:"desc"`
	Status      string `json:"status" form:"status" query:"status"`
}
