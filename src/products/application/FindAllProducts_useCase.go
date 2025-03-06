package application

import (
	"mi-tienda-online/src/products/domain"
	"mi-tienda-online/src/products/domain/entities"
)

type FindAllProductsUseCase struct {
	ProductRepository domain.ProductRepository
}

func NewFindAllProductsUseCase(ProductRepository domain.ProductRepository) *FindAllProductsUseCase {
	return &FindAllProductsUseCase{ProductRepository: ProductRepository}
}

func (useCase *FindAllProductsUseCase) Execute() ([]entities.Product, error) {
	products, err := useCase.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}