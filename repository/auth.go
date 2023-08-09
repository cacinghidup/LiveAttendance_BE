package repository

import (
	"TestBE/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(employee models.Employee) (models.Employee, error)
}

type authRepository struct {
	db *gorm.DB
}

func RepositoryAuth(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

func (a *authRepository) Register(employee models.Employee) (models.Employee, error) {
	err := a.db.Create(&employee).Error

	return employee, err
}
