package main

import (
	"time"
	"github.com/gin-contrib/cors"
	products "mi-tienda-online/src/products/infraestructure"
	users "mi-tienda-online/src/users/infraestructure"
	orders "mi-tienda-online/src/orders/infraestructure"
	"github.com/gin-gonic/gin"
)

func main() {
	r:= gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	products.Init(r)
	users.Init(r)
	orders.Init(r)
	r.Run(":8080")
}