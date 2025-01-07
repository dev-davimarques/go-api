package main

import (
	"go-api/controller"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main(){
	server := gin.Default()

	ProductUseCase := usecase.NewProductUseCase()
	// Camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)

	// Rota de ping
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message":"pong",
		})
	})

	server.GET("/products", ProductController.GetProduct)

	server.Run(":9000")
}