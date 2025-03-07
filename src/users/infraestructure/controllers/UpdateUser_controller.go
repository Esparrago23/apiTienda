package controllers

import (
	"mi-tienda-online/src/users/application"
	"mi-tienda-online/src/users/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateUserController struct {
	UpdateUserUseCase application.UpdateUserUseCase
}

func NewUpdateUserController(UpdateUserUseCase application.UpdateUserUseCase) *UpdateUserController {
	return &UpdateUserController{UpdateUserUseCase: UpdateUserUseCase}
}

func (controller *UpdateUserController) Execute(c *gin.Context) {
	var user entities.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.UpdateUserUseCase.Execute(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
