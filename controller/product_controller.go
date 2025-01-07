package controller

import (
	// "go-api/model"
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct{
	productUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController{
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProduct(ctx *gin.Context){
	products, err := p.productUseCase.GetProduct()
	if(err != nil){
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}