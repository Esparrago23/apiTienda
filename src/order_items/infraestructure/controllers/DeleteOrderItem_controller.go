package controllers

import (
	"mi-tienda-online/src/order_items/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteOrderItemController struct {
	DeleteOrderItemUseCase application.DeleteOrderItemUseCase
}

func NewDeleteOrderItemController(DeleteOrderItemUseCase application.DeleteOrderItemUseCase) *DeleteOrderItemController {
	return &DeleteOrderItemController{DeleteOrderItemUseCase: DeleteOrderItemUseCase}
}

func (controller *DeleteOrderItemController) Execute(c *gin.Context) {
	orderItemIDParam := c.Param("id")
	orderItemID, err := strconv.Atoi(orderItemIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order item ID"})
		return
	}

	err = controller.DeleteOrderItemUseCase.Execute(orderItemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order item deleted successfully"})
}
