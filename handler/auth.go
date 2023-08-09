package handler

import (
	authdto "TestBE/dto/auth"
	dto "TestBE/dto/result"
	"TestBE/models"
	"TestBE/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handlerAuth struct {
	AuthRepository repository.AuthRepository
}

func HandlerAuth(AuthRepository repository.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(c *gin.Context) {
	request := new(authdto.RegisterRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	Name := c.Request.FormValue("name")

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	pictureEmployee, _ := c.Get("UploadedFile")

	employee := models.Employee{
		Name:     Name,
		Picture:  pictureEmployee.(string),
		CreateAt: time.Now(),
	}

	data, err := h.AuthRepository.Register(employee)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}
