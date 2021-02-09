package main

// TaskRequest
type TaskRequest struct {
	Name        string `json:"name" form:"name" query:"name"`
	Description string `json:"desc" form:"desc" query:"desc"`
	Status      string `json:"status" form:"status" query:"status"`
}

// func (cv *CustomValidator) Validate(i interface{}) error {
// 	return cv.validator.Struct(i)
// }
