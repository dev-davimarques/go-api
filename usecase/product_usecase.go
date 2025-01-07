package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct{
	// Repository
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUsecase{
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProduct() ([]model.Product, error){
	return pu.repository.GetProduct()
}