package main

import (
	"time"
	"github.com/gin-contrib/cors"
	"mi-tienda-online/src/products/infraestructure"
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
	infraestructure.Init(r)
	r.Run(":8080")
}