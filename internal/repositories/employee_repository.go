package repositories

import (
	"context"
	"codeid.hr-api/internal/domain/model"
	"codeid.hr-api/internal/domain/query"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	FindAll(ctx context.Context) ([]*model.Employee, error)
	FindByID(ctx context.Context, id int32) (*model.Employee, error)
	Create(ctx context.Context, employee *model.Employee) error
	Update(ctx context.Context, employee *model.Employee) error
	Delete(ctx context.Context, id int32) error
	SearchByName(ctx context.Context, name string) ([]*model.Employee, error)
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{
		Q: query.Use(db)}
}

type employeeRepository struct {
	Q *query.Query
}

// SearchByName implements EmployeeRepository.
func (d *employeeRepository) SearchByName(ctx context.Context, name string) ([]*model.Employee, error) {
	departments, err := d.Q.Employee.WithContext(ctx).
		Where(d.Q.Employee.FirstName.Like("%" + name + "%")).Find()
	if err != nil {
		return nil, err
	}
	return departments, nil
}

// Create implements EmployeeRepository.
func (d *employeeRepository) Create(ctx context.Context, employee *model.Employee) error {
	return d.Q.Employee.WithContext(ctx).Create(employee)
}

// Delete implements EmployeeRepository.
// func (d *employeeRepository) Delete(ctx context.Context, id uint) error {
// 	employeeID, err := (id, err)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = d.Q.Employee.WithContext(ctx).Where(d.Q.Employee.EmployeeID.Eq(int32(employeeID))).Delete(&model.Employee{})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// Delete implements EmployeeRepository.
func (d *employeeRepository) Delete(ctx context.Context, id int32) error {
	_, err := d.Q.Employee.WithContext(ctx).Where(d.Q.Employee.EmployeeID.Eq(int32(id))).Delete(&model.Employee{})
	if err != nil {
		return err
	}
	return nil
}

// FindAll implements EmployeeRepository.
func (d *employeeRepository) FindAll(ctx context.Context) ([]*model.Employee, error) {
	var employees []*model.Employee
	employees, err := d.Q.Employee.WithContext(ctx).Find()
	if err != nil {
		return nil, err
	}
	return employees, err
}

// FindByID implements EmployeeRepository.
func (d *employeeRepository) FindByID(ctx context.Context, id int32) (*model.Employee, error) {
	employee, err := d.Q.Employee.WithContext(ctx).Where(d.Q.Employee.EmployeeID.Eq(int32(id))).First()
	if err != nil {
		return nil, err
	}
	return employee, nil
}

// Update implements EmployeeRepository.
func (d *employeeRepository) Update(ctx context.Context, employee *model.Employee) error {
	return d.Q.Employee.WithContext(ctx).Save(employee)
}
