package controllers

import (
	"mi-tienda-online/src/users/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FindAllUsersController struct {
	FindAllUsersUseCase application.FindAllUsersUseCase
}

func NewFindAllUsersController(FindAllUsersUseCase application.FindAllUsersUseCase) *FindAllUsersController {
	return &FindAllUsersController{FindAllUsersUseCase: FindAllUsersUseCase}
}

func (controller *FindAllUsersController) Execute(c *gin.Context) {
	users, err := controller.FindAllUsersUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
