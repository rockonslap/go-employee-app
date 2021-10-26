package employee

import (
	"context"
)

func Create(request CreateEmployeeRequest) (EmployeeModel, error) {
	employeeRepo := NewEmployeeRepository()
	ctx := context.Background()

	employee := EmployeeModel{
		Name:         request.Name,
		Age:          request.Age,
		Gender:       request.Gender,
		Address:      request.Address,
		Department:   request.Department,
		MobileNumber: request.MobileNumber,
	}

	err := employeeRepo.Create(ctx, &employee)
	if err != nil {
		return employee, err
	}

	return employee, nil
}

func GetById(id int64) (EmployeeModel, error) {
	employeeRepo := NewEmployeeRepository()
	ctx := context.Background()

	employee, getErr := employeeRepo.GetById(ctx, id)
	if getErr != nil {
		return employee, getErr
	}

	return employee, nil
}

func Update(id int64, request UpdateEmployeeRequest) (EmployeeModel, error) {
	employeeRepo := NewEmployeeRepository()
	ctx := context.Background()

	employee, getErr := employeeRepo.GetById(ctx, id)
	if getErr != nil {
		return employee, getErr
	}

	employee.Name = request.Name
	employee.Age = request.Age
	employee.Gender = request.Gender
	employee.Address = request.Address
	employee.Department = request.Department
	employee.MobileNumber = request.MobileNumber

	uodateErr := employeeRepo.Update(ctx, &employee)
	if uodateErr != nil {
		return employee, uodateErr
	}

	return employee, nil
}

func Delete(id int64) error {
	employeeRepo := NewEmployeeRepository()
	ctx := context.Background()

	employee, getErr := employeeRepo.GetById(ctx, id)
	if getErr != nil {
		return getErr
	}

	deleteErr := employeeRepo.Delete(ctx, &employee)
	if deleteErr != nil {
		return deleteErr
	}

	return nil
}
