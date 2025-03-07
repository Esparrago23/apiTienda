package controllers

import (
	"mi-tienda-online/src/order_items/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FindAllOrderItemsController struct {
	FindAllOrderItemsUseCase application.FindAllOrderItemsUseCase
}

func NewFindAllOrderItemsController(FindAllOrderItemsUseCase application.FindAllOrderItemsUseCase) *FindAllOrderItemsController {
	return &FindAllOrderItemsController{FindAllOrderItemsUseCase: FindAllOrderItemsUseCase}
}

func (controller *FindAllOrderItemsController) Execute(c *gin.Context) {
	orderItems, err := controller.FindAllOrderItemsUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orderItems)
}
