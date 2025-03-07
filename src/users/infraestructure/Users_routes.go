package infraestructure

import (
	"mi-tienda-online/src/users/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

type UsersHandlers struct {
	Create   *controllers.CreateUserController
	Delete   *controllers.DeleteUserController
	FindById *controllers.FindUserByIdController
	FindAll  *controllers.FindAllUsersController
	Update   *controllers.UpdateUserController
}

func UsersRoutes(router *gin.Engine, handlers UsersHandlers) {
	usersGroup := router.Group("/users")
	{
		usersGroup.POST("/", handlers.Create.Execute)
		usersGroup.DELETE("/:id", handlers.Delete.Execute)
		usersGroup.GET("/:id", handlers.FindById.Execute)
		usersGroup.GET("/", handlers.FindAll.Execute)
		usersGroup.PUT("/:id", handlers.Update.Execute)
	}
}
