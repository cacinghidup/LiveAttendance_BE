package routes

import (
	"TestBE/handler"
	"TestBE/pkg/middleware"
	"TestBE/pkg/mysql"
	"TestBE/repository"

	"github.com/gin-gonic/gin"
)

func AttendanceRoutes(r *gin.RouterGroup) {
	attendanceRepository := repository.RepositoryAttendance(mysql.DB)
	h := handler.HandlerAttendance(attendanceRepository)

	r.GET("/attendance", h.FindAll)
	r.POST("/attendance/checkin/:id", middleware.PhotoChecker(h.CreateAttendanceCheckIn))
	r.POST("/attendance/checkout/:id", h.CreateAttendanceCheckOut)
}
