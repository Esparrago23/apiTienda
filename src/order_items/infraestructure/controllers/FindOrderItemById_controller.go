package controllers

import (
	"mi-tienda-online/src/order_items/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FindOrderItemByIdController struct {
	FindOrderItemByIdUseCase application.FindOrderItemByIdUseCase
}

func NewFindOrderItemByIdController(FindOrderItemByIdUseCase application.FindOrderItemByIdUseCase) *FindOrderItemByIdController {
	return &FindOrderItemByIdController{FindOrderItemByIdUseCase: FindOrderItemByIdUseCase}
}

func (controller *FindOrderItemByIdController) Execute(c *gin.Context) {
	orderItemIDParam := c.Param("id")
	orderItemID, err := strconv.Atoi(orderItemIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order item ID"})
		return
	}

	orderItem, err := controller.FindOrderItemByIdUseCase.Execute(orderItemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orderItem)
}
