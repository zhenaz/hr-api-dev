package services

import (
	"context"
	"errors"

	"codeid.hr-api/internal/domain/model"
	"codeid.hr-api/internal/repositories"
)

// 4. define implementation methods, use tombol Ctrl + . saat kursor ada di
// NewRegionRepository, supy bisa generate all method interface.
type EmployeeService interface {
	GetAllEmployees(ctx context.Context) ([]*model.Employee, error)
	GetEmployeeByID(ctx context.Context, id int32) (*model.Employee, error)
	CreateEmployee(ctx context.Context, employee *model.Employee) error
	UpdateEmployee(ctx context.Context, employee *model.Employee) error
	DeleteEmployee(ctx context.Context, id int32) error
}
type employeeService struct {
	employeeRepo repositories.EmployeeRepository
}

func NewEmployeeService(employeeRepo repositories.EmployeeRepository) EmployeeService {
	return &employeeService{
		employeeRepo: employeeRepo}
}
func (s *employeeService) GetAllEmployees(ctx context.Context) ([]*model.Employee, error) {
	return s.employeeRepo.FindAll(ctx)

	// employeesPtr, err := s.employeeRepo.FindAll(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	// employees := make([]model.Employee, len(employeesPtr))
	// for i, empPtr := range employeesPtr {
	// 	if empPtr != nil {
	// 		employees[i] = *empPtr
	// 	}
	// }
	// return employees, nil
}
func (s *employeeService) GetEmployeeByID(ctx context.Context, id int32) (*model.Employee, error) {
	if id == 0 {
		return nil, errors.New("employee ID cannot be empty")
	}
	return s.employeeRepo.FindByID(ctx,(id))
}

func (s *employeeService) CreateEmployee(ctx context.Context, employee *model.Employee) error {
	if *employee.FirstName == "" {
		return errors.New("employee first name cannot be empty")
	}
	if len(*employee.FirstName) > 25 {
		return errors.New("employee first name cannot exceed 25 characters")
	}
	return s.employeeRepo.Create(ctx, employee)
}
func (s *employeeService) UpdateEmployee(ctx context.Context, employee *model.Employee) error {
	if employee.EmployeeID == 0 {
		return errors.New("employee ID cannot be empty")
	}
	if *employee.FirstName == "" {
		return errors.New("employee first name cannot be empty")
	}
	if _, err := s.employeeRepo.FindByID(ctx, employee.EmployeeID); err != nil {
		return errors.New("employee doesn't exists")
	}
	return s.employeeRepo.Update(ctx, employee)
}
func (s *employeeService) DeleteEmployee(ctx context.Context, id int32) error {
	if id == 0 {
		return errors.New("employee ID cannot be empty")
	}
	if _, err := s.employeeRepo.FindByID(ctx, id); err != nil {
		return errors.New("employee doesn't exists")
	}
	return s.employeeRepo.Delete(ctx, id)
}
