package controllers

import (
	"mi-tienda-online/src/users/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteUserController struct {
	DeleteUserUseCase application.DeleteUserUseCase
}

func NewDeleteUserController(DeleteUserUseCase application.DeleteUserUseCase) *DeleteUserController {
	return &DeleteUserController{DeleteUserUseCase: DeleteUserUseCase}
}

func (controller *DeleteUserController) Execute(c *gin.Context) {
	userIDParam := c.Param("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = controller.DeleteUserUseCase.Execute(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
