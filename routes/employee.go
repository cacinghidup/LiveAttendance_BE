package routes

import (
	"TestBE/handler"
	"TestBE/pkg/mysql"
	"TestBE/repository"

	"github.com/gin-gonic/gin"
)

func EmployeeRoutes(e *gin.RouterGroup) {
	employeeRepository := repository.RepositoryEmployee(mysql.DB)
	h := handler.NewHandler(employeeRepository)

	e.GET("/employee", h.FindAll)
	e.GET("/employee/:id", h.GetByID)
}
