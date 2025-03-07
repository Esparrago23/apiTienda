package controllers

import (
	"mi-tienda-online/src/users/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FindUserByIdController struct {
	FindUserByIdUseCase application.FindUserByIdUseCase
}

func NewFindUserByIdController(FindUserByIdUseCase application.FindUserByIdUseCase) *FindUserByIdController {
	return &FindUserByIdController{FindUserByIdUseCase: FindUserByIdUseCase}
}

func (controller *FindUserByIdController) Execute(c *gin.Context) {
	userIDParam := c.Param("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := controller.FindUserByIdUseCase.Execute(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
