package controllers

import (
	"mi-tienda-online/src/payments/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeletePaymentController struct {
	DeletePaymentUseCase application.DeletePaymentUseCase
}

func NewDeletePaymentController(DeletePaymentUseCase application.DeletePaymentUseCase) *DeletePaymentController {
	return &DeletePaymentController{DeletePaymentUseCase: DeletePaymentUseCase}
}

func (controller *DeletePaymentController) Execute(c *gin.Context) {
	paymentIDParam := c.Param("id")
	paymentID, err := strconv.Atoi(paymentIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	err = controller.DeletePaymentUseCase.Execute(paymentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment deleted successfully"})
}
