package controllers

import (
	"mi-tienda-online/src/orders/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteOrderController struct {
	DeleteOrderUseCase application.DeleteOrderUseCase
}

func NewDeleteOrderController(DeleteOrderUseCase application.DeleteOrderUseCase) *DeleteOrderController {
	return &DeleteOrderController{DeleteOrderUseCase: DeleteOrderUseCase}
}

func (controller *DeleteOrderController) Execute(c *gin.Context) {
	orderIDParam := c.Param("id")
	orderID, err := strconv.Atoi(orderIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	err = controller.DeleteOrderUseCase.Execute(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
