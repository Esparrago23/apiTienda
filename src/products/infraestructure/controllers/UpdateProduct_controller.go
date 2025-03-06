package controllers

import (
	"mi-tienda-online/src/products/application"
	"mi-tienda-online/src/products/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
type UpdateProductController struct {
	UpdateProductUseCase application.UpdateProductUseCase
}

func NewUpdateProductController(UpdateProductUseCase application.UpdateProductUseCase) *UpdateProductController {
	return &UpdateProductController{UpdateProductUseCase : UpdateProductUseCase}
}



func (controller *UpdateProductController) Execute(c *gin.Context) {
	var product entities.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := validate.Struct(product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	productid := c.Param("id")
	id, err := strconv.Atoi(productid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	product.ProductID = id

	err = controller.UpdateProductUseCase.Execute(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}