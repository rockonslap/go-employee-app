package employee

import (
	"context"
	database "employee-app/config"
	"time"

	"github.com/uptrace/bun"
)

type EmployeeModel struct {
	bun.BaseModel `bun:"employee"`
	ID            int64
	Name          string
	Age           int8
	Gender        string
	Address       string
	Department    string
	MobileNumber  string
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

type EmployeeRepository interface {
	Create(context.Context, *EmployeeModel) error
	GetById(context.Context, int64) (EmployeeModel, error)
	Update(context.Context, *EmployeeModel) error
	Delete(context.Context, *EmployeeModel) error
}

func NewEmployeeRepository() EmployeeRepository {
	return new(repo)
}

type repo struct {
}

func (repo *repo) Create(ctx context.Context, employee *EmployeeModel) error {
	_, err := database.PGConnection.NewInsert().Model(employee).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (repo *repo) GetById(ctx context.Context, id int64) (EmployeeModel, error) {
	employee := new(EmployeeModel)

	err := database.PGConnection.NewSelect().Model(employee).Where("id = ?", id).Scan(ctx, employee)
	if err != nil {
		return *employee, err
	}

	return *employee, nil
}

func (repo *repo) Update(ctx context.Context, employee *EmployeeModel) error {
	_, err := database.PGConnection.NewUpdate().Model(employee).Where("id = ?", employee.ID).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (repo *repo) Delete(ctx context.Context, employee *EmployeeModel) error {
	_, err := database.PGConnection.NewDelete().Model(employee).Where("id = ?", employee.ID).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
