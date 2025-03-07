package infraestructure

import (
	"mi-tienda-online/src/order_items/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

type OrderItemsHandlers struct {
	Create   *controllers.CreateOrderItemController
	Delete   *controllers.DeleteOrderItemController
	FindById *controllers.FindOrderItemByIdController
	FindAll  *controllers.FindAllOrderItemsController
	Update   *controllers.UpdateOrderItemController
}

func OrderItemsRoutes(router *gin.Engine, handlers OrderItemsHandlers) {
	orderItemsGroup := router.Group("/order_items")
	{
		orderItemsGroup.POST("/", handlers.Create.Execute)
		orderItemsGroup.DELETE("/:id", handlers.Delete.Execute)
		orderItemsGroup.GET("/:id", handlers.FindById.Execute)
		orderItemsGroup.GET("/", handlers.FindAll.Execute)
		orderItemsGroup.PUT("/:id", handlers.Update.Execute)
	}
}
