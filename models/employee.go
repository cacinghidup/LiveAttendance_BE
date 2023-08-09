package models

import "time"

type Employee struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Picture  string    `json:"picture"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type EmployeeResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (EmployeeResponse) TableName() string {
	return "employees"
}
