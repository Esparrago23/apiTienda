package application

import (
	"mi-tienda-online/src/products/domain"
	"mi-tienda-online/src/products/domain/entities"
)

type UpdateProductUseCase struct {
	ProductRepository domain.ProductRepository
}

func NewUpdateProductUseCase(ProductRepository domain.ProductRepository) *UpdateProductUseCase {
	return &UpdateProductUseCase{ProductRepository:	ProductRepository}
}

func (useCase *UpdateProductUseCase) Execute(product *entities.Product) error {
	err:= useCase.ProductRepository.Update(product)
	if err != nil {
		return err
	}
	return nil
}