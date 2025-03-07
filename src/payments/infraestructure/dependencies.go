package infraestructure

import (
	"mi-tienda-online/src/payments/application"
	"mi-tienda-online/src/payments/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	ps := NewMySQL()

	createPaymentService := application.NewCreatePaymentUseCase(ps)
	deletePaymentService := application.NewDeletePaymentUseCase(ps)
	findPaymentByIdService := application.NewFindPaymentByIdUseCase(ps)
	findAllPaymentsService := application.NewFindAllPaymentsUseCase(ps)
	updatePaymentService := application.NewUpdatePaymentUseCase(ps)

	createPaymentController := controllers.NewCreatePaymentController(*createPaymentService)
	deletePaymentController := controllers.NewDeletePaymentController(*deletePaymentService)
	findPaymentByIdController := controllers.NewFindPaymentByIdController(*findPaymentByIdService)
	findAllPaymentsController := controllers.NewFindAllPaymentsController(*findAllPaymentsService)
	updatePaymentController := controllers.NewUpdatePaymentController(*updatePaymentService)

	PaymentsRoutes(router, PaymentsHandlers{
		Create:   createPaymentController,
		Delete:   deletePaymentController,
		FindById: findPaymentByIdController,
		FindAll:  findAllPaymentsController,
		Update:   updatePaymentController,
	})
}
