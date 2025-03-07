package controllers

import (
	"mi-tienda-online/src/orders/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FindAllOrdersController struct {
	FindAllOrdersUseCase application.FindAllOrdersUseCase
}

func NewFindAllOrdersController(FindAllOrdersUseCase application.FindAllOrdersUseCase) *FindAllOrdersController {
	return &FindAllOrdersController{FindAllOrdersUseCase: FindAllOrdersUseCase}
}

func (controller *FindAllOrdersController) Execute(c *gin.Context) {
	orders, err := controller.FindAllOrdersUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}
