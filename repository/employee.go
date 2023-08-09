package repository

import (
	"TestBE/models"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	FindAll() ([]models.Employee, error)
	GetByID(id int) (models.Employee, error)
}

type employeeRepository struct {
	db *gorm.DB
}

func RepositoryEmployee(db *gorm.DB) *employeeRepository {
	return &employeeRepository{db}
}

func (r *employeeRepository) FindAll() ([]models.Employee, error) {
	var employees []models.Employee
	err := r.db.Find(&employees).Error

	return employees, err
}

func (r *employeeRepository) GetByID(id int) (models.Employee, error) {
	var employee models.Employee
	err := r.db.Where("id = ?", id).Find(&employee).Error

	return employee, err
}
