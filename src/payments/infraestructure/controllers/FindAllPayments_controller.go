package controllers

import (
	"mi-tienda-online/src/payments/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FindAllPaymentsController struct {
	FindAllPaymentsUseCase application.FindAllPaymentsUseCase
}

func NewFindAllPaymentsController(FindAllPaymentsUseCase application.FindAllPaymentsUseCase) *FindAllPaymentsController {
	return &FindAllPaymentsController{FindAllPaymentsUseCase: FindAllPaymentsUseCase}
}

func (controller *FindAllPaymentsController) Execute(c *gin.Context) {
	payments, err := controller.FindAllPaymentsUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payments)
}
