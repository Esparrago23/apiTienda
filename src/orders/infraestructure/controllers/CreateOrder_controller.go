package controllers

import (
	"mi-tienda-online/src/orders/application"
	"mi-tienda-online/src/orders/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateOrderController struct {
	CreateOrderUseCase application.CreateOrderUseCase
}

func NewCreateOrderController(CreateOrderUseCase application.CreateOrderUseCase) *CreateOrderController {
	return &CreateOrderController{CreateOrderUseCase: CreateOrderUseCase}
}

var validate = validator.New()

func (controller *CreateOrderController) Execute(c *gin.Context) {
	var order entities.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := validate.Struct(order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.CreateOrderUseCase.Execute(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}
