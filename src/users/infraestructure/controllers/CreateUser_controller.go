package controllers

import (
	"mi-tienda-online/src/users/application"
	"mi-tienda-online/src/users/domain/entities"
	"net/http"

	//_"mi-tienda-online/src/users/infraestructure/rabbitmq"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateUserController struct {
	CreateUserUseCase application.CreateUserUseCase
}

func NewCreateUserController(CreateUserUseCase application.CreateUserUseCase) *CreateUserController {
	return &CreateUserController{CreateUserUseCase: CreateUserUseCase}
}

var validate = validator.New()

func (controller *CreateUserController) Execute(c *gin.Context) {
	var user entities.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.CreateUserUseCase.Execute(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//rabbitmq.PublishUser("Nuevo usuario creado")
	//c.JSON(http.StatusOK, gin.H{"message": "Usuario creado y enviado a RabbitMQ"})
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
