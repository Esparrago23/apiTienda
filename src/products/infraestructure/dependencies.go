package infraestructure

import (
	"mi-tienda-online/src/products/application"
	"mi-tienda-online/src/products/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine){
	ps := NewMySQL()

	createProductService := application.NewCreateProductUseCase(ps)
	deleteProductService := application.NewDeleteProductUseCase(ps)
	findProductByIdService := application.NewFindProductByIdUseCase(ps)
	findAllProductsService := application.NewFindAllProductsUseCase(ps)
	updateProductService := application.NewUpdateProductUseCase(ps)
	

	createProductController := controllers.NewCreateProductController(*createProductService)
	deleteProductController := controllers.NewDeleteProductController(*deleteProductService)
	findProductByIdController := controllers.NewFindProductByIdController(*findProductByIdService)
	findAllProductsController := controllers.NewFindAllProductsController(*findAllProductsService)
	updateProductController := controllers.NewUpdateProductController(*updateProductService)

	ProductsRoutes(router, ProductsHandlers{
		Create: createProductController,
		Delete: deleteProductController,
		FindById: findProductByIdController,
		FindAll: findAllProductsController,
		Update: updateProductController,
	})
}