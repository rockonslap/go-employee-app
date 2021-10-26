package employee

import "time"

type CreateEmployeeResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type GetEmployeeResponse struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Age          int8      `json:"age"`
	Gender       string    `json:"gender"`
	Address      string    `json:"address"`
	Department   string    `json:"department"`
	MobileNumber string    `json:"mobile_number"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UpdateEmployeeResponse struct {
	ID int64 `json:"id"`
}

type DeleteEmployeeResponse = UpdateEmployeeResponse
