package routes

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	EmployeeRoutes(r)
	AuthRoutes(r)
	AttendanceRoutes(r)
}
