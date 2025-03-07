package controllers

import (
	"mi-tienda-online/src/order_items/application"
	"mi-tienda-online/src/order_items/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateOrderItemController struct {
	UpdateOrderItemUseCase application.UpdateOrderItemUseCase
}

func NewUpdateOrderItemController(UpdateOrderItemUseCase application.UpdateOrderItemUseCase) *UpdateOrderItemController {
	return &UpdateOrderItemController{UpdateOrderItemUseCase: UpdateOrderItemUseCase}
}

func (controller *UpdateOrderItemController) Execute(c *gin.Context) {
	var orderItem entities.OrderItem
	if err := c.BindJSON(&orderItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := validate.Struct(orderItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.UpdateOrderItemUseCase.Execute(&orderItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order item updated successfully"})
}
