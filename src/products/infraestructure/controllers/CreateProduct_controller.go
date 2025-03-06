package controllers

import (
	"mi-tienda-online/src/products/application"
	"mi-tienda-online/src/products/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateProductController struct {
	CreateProductUseCase application.CreateProductUseCase
}

func NewCreateProductController(CreateProductUseCase application.CreateProductUseCase) *CreateProductController {
	return &CreateProductController{CreateProductUseCase: CreateProductUseCase}
}

var validate = validator.New()

func (controller *CreateProductController) Execute(c *gin.Context) {
	var product entities.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := validate.Struct(product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.CreateProductUseCase.Execute(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully"})
}
