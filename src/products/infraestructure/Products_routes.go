package infraestructure

import (
	"mi-tienda-online/src/products/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

type ProductsHandlers struct {
	Create   *controllers.CreateProductController
	Delete   *controllers.DeleteProductController
	FindById *controllers.FindProductByIdController
	FindAll  *controllers.FindAllProductsController
	Update   *controllers.UpdateProductController
}

func ProductsRoutes (router *gin.Engine, handlers ProductsHandlers) {
	productsGroup := router.Group("/products")
	{
		productsGroup.POST("/", handlers.Create.Execute)
		productsGroup.DELETE("/:id", handlers.Delete.Execute)
		productsGroup.GET("/:id", handlers.FindById.Execute)
		productsGroup.GET("/", handlers.FindAll.Execute)
		productsGroup.PUT("/:id", handlers.Update.Execute)
	}
}