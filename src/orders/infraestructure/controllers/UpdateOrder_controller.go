package controllers

import (
	"mi-tienda-online/src/orders/application"
	"mi-tienda-online/src/orders/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateOrderController struct {
	UpdateOrderUseCase application.UpdateOrderUseCase
}

func NewUpdateOrderController(UpdateOrderUseCase application.UpdateOrderUseCase) *UpdateOrderController {
	return &UpdateOrderController{UpdateOrderUseCase: UpdateOrderUseCase}
}

func (controller *UpdateOrderController) Execute(c *gin.Context) {
	var order entities.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := validate.Struct(order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.UpdateOrderUseCase.Execute(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}
