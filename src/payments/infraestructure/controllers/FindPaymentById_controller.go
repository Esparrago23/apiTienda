package controllers

import (
	"mi-tienda-online/src/payments/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FindPaymentByIdController struct {
	FindPaymentByIdUseCase application.FindPaymentByIdUseCase
}

func NewFindPaymentByIdController(FindPaymentByIdUseCase application.FindPaymentByIdUseCase) *FindPaymentByIdController {
	return &FindPaymentByIdController{FindPaymentByIdUseCase: FindPaymentByIdUseCase}
}

func (controller *FindPaymentByIdController) Execute(c *gin.Context) {
	paymentIDParam := c.Param("id")
	paymentID, err := strconv.Atoi(paymentIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	payment, err := controller.FindPaymentByIdUseCase.Execute(paymentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payment)
}
