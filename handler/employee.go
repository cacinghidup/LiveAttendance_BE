package handler

import (
	dto "TestBE/dto/result"
	"TestBE/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	EmployeeRepository repository.EmployeeRepository
}

func NewHandler(EmployeeRepository repository.EmployeeRepository) *handler {
	return &handler{EmployeeRepository}
}

func (h *handler) FindAll(c *gin.Context) {
	employee, err := h.EmployeeRepository.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: employee,
	})
}

func (h *handler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	employee, err := h.EmployeeRepository.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: employee,
	})
}
