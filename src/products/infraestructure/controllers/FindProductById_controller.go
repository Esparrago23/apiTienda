package controllers

import (
	"mi-tienda-online/src/products/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FindProductByIdController struct {
	FindProductByIdUseCase application.FindProductByIdUseCase
}

func NewFindProductByIdController(FindProductByIdUseCase application.FindProductByIdUseCase) *FindProductByIdController {
	return &FindProductByIdController{FindProductByIdUseCase: FindProductByIdUseCase}
}

func (controller *FindProductByIdController) Execute(c *gin.Context) {
	productid := c.Param("id")
	
	id, err := strconv.Atoi(productid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	product, err := controller.FindProductByIdUseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}