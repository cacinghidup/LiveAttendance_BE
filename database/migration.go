package database

import (
	"TestBE/models"
	"TestBE/pkg/mysql"
	"fmt"
)

func MigrationDB() {
	err := mysql.DB.AutoMigrate(
		&models.Employee{},
		&models.Attendance{},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Migration Success")
}
