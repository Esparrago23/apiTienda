package controllers

import (
	"mi-tienda-online/src/products/application"
	"strconv"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteProductController struct {
	DeleteProductUseCase application.DeleteProductUseCase
}

func NewDeleteProductController(DeleteProductUseCase application.DeleteProductUseCase) *DeleteProductController {
	return &DeleteProductController{DeleteProductUseCase: DeleteProductUseCase}
}

func (controller *DeleteProductController) Execute(c *gin.Context) {
productid := c.Param("id")
id, err := strconv.Atoi(productid)
if err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	return
}
err = controller.DeleteProductUseCase.Execute(id)
if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
}
c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}