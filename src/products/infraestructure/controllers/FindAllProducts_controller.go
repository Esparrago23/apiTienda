package controllers

import (
	"mi-tienda-online/src/products/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FindAllProductsController struct {
	FindAllProductsUseCase application.FindAllProductsUseCase
}

func NewFindAllProductsController(FindAllProductsUseCase application.FindAllProductsUseCase) *FindAllProductsController {
	return &FindAllProductsController{FindAllProductsUseCase: FindAllProductsUseCase}
}


func (controller *FindAllProductsController) Execute(c *gin.Context) {
	products, err := controller.FindAllProductsUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}


