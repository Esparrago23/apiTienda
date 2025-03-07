package infraestructure

import (
	"mi-tienda-online/src/order_items/application"
	"mi-tienda-online/src/order_items/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	os := NewMySQL()

	createOrderItemService := application.NewCreateOrderItemUseCase(os)
	deleteOrderItemService := application.NewDeleteOrderItemUseCase(os)
	findOrderItemByIdService := application.NewFindOrderItemByIdUseCase(os)
	findAllOrderItemsService := application.NewFindAllOrderItemsUseCase(os)
	updateOrderItemService := application.NewUpdateOrderItemUseCase(os)

	createOrderItemController := controllers.NewCreateOrderItemController(*createOrderItemService)
	deleteOrderItemController := controllers.NewDeleteOrderItemController(*deleteOrderItemService)
	findOrderItemByIdController := controllers.NewFindOrderItemByIdController(*findOrderItemByIdService)
	findAllOrderItemsController := controllers.NewFindAllOrderItemsController(*findAllOrderItemsService)
	updateOrderItemController := controllers.NewUpdateOrderItemController(*updateOrderItemService)

	OrderItemsRoutes(router, OrderItemsHandlers{
		Create:   createOrderItemController,
		Delete:   deleteOrderItemController,
		FindById: findOrderItemByIdController,
		FindAll:  findAllOrderItemsController,
		Update:   updateOrderItemController,
	})
}
