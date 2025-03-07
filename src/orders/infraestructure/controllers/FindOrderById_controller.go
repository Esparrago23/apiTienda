package controllers

import (
	"mi-tienda-online/src/orders/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FindOrderByIdController struct {
	FindOrderByIdUseCase application.FindOrderByIdUseCase
}

func NewFindOrderByIdController(FindOrderByIdUseCase application.FindOrderByIdUseCase) *FindOrderByIdController {
	return &FindOrderByIdController{FindOrderByIdUseCase: FindOrderByIdUseCase}
}

func (controller *FindOrderByIdController) Execute(c *gin.Context) {
	orderIDParam := c.Param("id")
	orderID, err := strconv.Atoi(orderIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := controller.FindOrderByIdUseCase.Execute(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}
