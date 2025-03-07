package infraestructure

import (
	"mi-tienda-online/src/orders/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

type OrdersHandlers struct {
	Create   *controllers.CreateOrderController
	Delete   *controllers.DeleteOrderController
	FindById *controllers.FindOrderByIdController
	FindAll  *controllers.FindAllOrdersController
	Update   *controllers.UpdateOrderController
}

func OrdersRoutes(router *gin.Engine, handlers OrdersHandlers) {
	ordersGroup := router.Group("/orders")
	{
		ordersGroup.POST("/", handlers.Create.Execute)
		ordersGroup.DELETE("/:id", handlers.Delete.Execute)
		ordersGroup.GET("/:id", handlers.FindById.Execute)
		ordersGroup.GET("/", handlers.FindAll.Execute)
		ordersGroup.PUT("/:id", handlers.Update.Execute)
	}
}
