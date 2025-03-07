package controllers

import (
	"mi-tienda-online/src/payments/application"
	"mi-tienda-online/src/payments/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdatePaymentController struct {
	UpdatePaymentUseCase application.UpdatePaymentUseCase
}

func NewUpdatePaymentController(UpdatePaymentUseCase application.UpdatePaymentUseCase) *UpdatePaymentController {
	return &UpdatePaymentController{UpdatePaymentUseCase: UpdatePaymentUseCase}
}


func (controller *UpdatePaymentController) Execute(c *gin.Context) {
	var payment entities.Payment
	if err := c.BindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := validate.Struct(payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.UpdatePaymentUseCase.Execute(&payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment updated successfully"})
}
