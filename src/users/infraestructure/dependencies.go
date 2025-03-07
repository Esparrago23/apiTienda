package infraestructure

import (
	"mi-tienda-online/src/users/application"
	"mi-tienda-online/src/users/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	us := NewMySQL()

	createUserService := application.NewCreateUserUseCase(us)
	deleteUserService := application.NewDeleteUserUseCase(us)
	findUserByIdService := application.NewFindUserByIdUseCase(us)
	findAllUsersService := application.NewFindAllUsersUseCase(us)
	updateUserService := application.NewUpdateUserUseCase(us)

	createUserController := controllers.NewCreateUserController(*createUserService)
	deleteUserController := controllers.NewDeleteUserController(*deleteUserService)
	findUserByIdController := controllers.NewFindUserByIdController(*findUserByIdService)
	findAllUsersController := controllers.NewFindAllUsersController(*findAllUsersService)
	updateUserController := controllers.NewUpdateUserController(*updateUserService)

	UsersRoutes(router, UsersHandlers{
		Create:   createUserController,
		Delete:   deleteUserController,
		FindById: findUserByIdController,
		FindAll:  findAllUsersController,
		Update:   updateUserController,
	})
}
