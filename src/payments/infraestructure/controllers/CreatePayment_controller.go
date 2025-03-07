package controllers

import (
	"mi-tienda-online/src/payments/application"
	"mi-tienda-online/src/payments/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreatePaymentController struct {
	CreatePaymentUseCase application.CreatePaymentUseCase
}

func NewCreatePaymentController(CreatePaymentUseCase application.CreatePaymentUseCase) *CreatePaymentController {
	return &CreatePaymentController{CreatePaymentUseCase: CreatePaymentUseCase}
}

var validate = validator.New()

func (controller *CreatePaymentController) Execute(c *gin.Context) {
	var payment entities.Payment
	if err := c.BindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := validate.Struct(payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.CreatePaymentUseCase.Execute(&payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment created successfully"})
}
