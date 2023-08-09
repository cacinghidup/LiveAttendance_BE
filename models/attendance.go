package models

type Attendance struct {
	Id         int              `json:"id" gorm:"primary_key"`
	Date       string           `json:"date"`
	CheckIn    string           `json:"check_in"`
	CheckOut   string           `json:"check_out"`
	Status     string           `json:"status"`
	StatusNote string           `json:"status_note"`
	EmployeeId int              `json:"employee_id"`
	Employee   EmployeeResponse `gnorm:"foreignKey:EmployeeId"`
}
