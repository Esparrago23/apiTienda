package application

import (
	"mi-tienda-online/src/products/domain"
	"mi-tienda-online/src/products/domain/entities"
)

type FindProductByIdUseCase struct {
	ProductRepository domain.ProductRepository
}

func NewFindProductByIdUseCase(ProductRepository domain.ProductRepository) *FindProductByIdUseCase {
	return &FindProductByIdUseCase{ProductRepository: ProductRepository}
}

func (useCase *FindProductByIdUseCase) Execute(product_id int) (entities.Product, error) {
	product, err := useCase.ProductRepository.FindById(product_id)
	if err != nil {
		return product, err
	}
	return product, nil
}