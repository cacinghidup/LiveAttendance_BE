package repository

import (
	"TestBE/models"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type AttendanceRepository interface {
	FindAll() ([]models.Attendance, error)
	CreateAttendanceCheckIn(attendance models.Attendance, ID int, Status string) (models.Attendance, error)
	CreateAttendanceCheckOut(attendance models.Attendance, ID int, Status string) (models.Attendance, error)
	GetByID(id int) (models.Employee, error)
}

type attendanceRepository struct {
	db *gorm.DB
}

func RepositoryAttendance(db *gorm.DB) *attendanceRepository {
	return &attendanceRepository{db}
}

func (r *attendanceRepository) FindAll() ([]models.Attendance, error) {
	var attendance []models.Attendance
	err := r.db.Preload("Employee").Find(&attendance).Error

	return attendance, err
}

func (r *attendanceRepository) CreateAttendanceCheckIn(attendance models.Attendance, ID int, Status string) (models.Attendance, error) {

	dateCheck := time.Now().Format("2006-01-02")
	EmployeeId := strconv.Itoa(ID)
	var attendanceCheck models.Attendance

	err := r.db.Where("date = ? AND employee_id = ? AND status = ?", dateCheck, EmployeeId, Status).First(&attendanceCheck).Error

	// fmt.Println("database uhuyyyy", err)

	if err == nil {
		return attendanceCheck, fmt.Errorf("attendance already exist")
	}

	err = r.db.Create(&attendance).Error

	return attendance, err
}

func (r *attendanceRepository) CreateAttendanceCheckOut(attendance models.Attendance, ID int, Status string) (models.Attendance, error) {

	dateCheck := time.Now().Format("2006-01-02")
	EmployeeId := strconv.Itoa(ID)
	var attendanceCheck models.Attendance

	err := r.db.Where("date = ? AND employee_id = ? AND status = ?", dateCheck, EmployeeId, Status).First(&attendanceCheck).Error

	// fmt.Println("database uhuyyyy", err)

	if err == nil {
		return attendanceCheck, fmt.Errorf("attendance already exist")
	}

	err = r.db.Create(&attendance).Error

	return attendance, err
}

func (r *attendanceRepository) GetByID(id int) (models.Employee, error) {
	var employee models.Employee
	err := r.db.Where("id = ?", id).Find(&employee).Error

	return employee, err
}
