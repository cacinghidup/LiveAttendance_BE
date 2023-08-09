package routes

import (
	"TestBE/handler"
	"TestBE/pkg/middleware"
	"TestBE/pkg/mysql"
	"TestBE/repository"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	authRepository := repository.RepositoryAuth(mysql.DB)
	h := handler.HandlerAuth(authRepository)

	r.POST("/register", middleware.UploadFile(h.Register))
}
