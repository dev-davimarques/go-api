package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main(){
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	// Camada usecase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
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