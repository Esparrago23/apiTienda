package infraestructure

import (
	"mi-tienda-online/src/payments/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

type PaymentsHandlers struct {
	Create   *controllers.CreatePaymentController
	Delete   *controllers.DeletePaymentController
	FindById *controllers.FindPaymentByIdController
	FindAll  *controllers.FindAllPaymentsController
	Update   *controllers.UpdatePaymentController
}

func PaymentsRoutes(router *gin.Engine, handlers PaymentsHandlers) {
	paymentsGroup := router.Group("/payments")
	{
		paymentsGroup.POST("/", handlers.Create.Execute)
		paymentsGroup.DELETE("/:id", handlers.Delete.Execute)
		paymentsGroup.GET("/:id", handlers.FindById.Execute)
		paymentsGroup.GET("/", handlers.FindAll.Execute)
		paymentsGroup.PUT("/:id", handlers.Update.Execute)
	}
}
