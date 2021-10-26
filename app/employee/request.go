package employee

type CreateEmployeeRequest struct {
	Name         string `form:"name" json:"name" binding:"required"`
	Age          int8   `form:"age" json:"age" binding:"gte=1"`
	Gender       string `form:"gender" json:"gender" binding:"oneof=M F"`
	Address      string `form:"address" json:"address"`
	Department   string `form:"department" json:"department"`
	MobileNumber string `form:"mobile_number" json:"mobile_number"`
}

type UpdateEmployeeRequest = CreateEmployeeRequest
