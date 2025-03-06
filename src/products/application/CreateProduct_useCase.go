package application

import (
	"mi-tienda-online/src/products/domain"
	"mi-tienda-online/src/products/domain/entities"
)

type CreateProductUseCase struct {
	ProductRepository domain.ProductRepository
}
func NewCreateProductUseCase(ProductRepository domain.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{ProductRepository: ProductRepository}
}

func (useCase *CreateProductUseCase) Execute(product *entities.Product) error {
	err:= useCase.ProductRepository.Save(product)
	if err != nil {
		return err
	}
	return nil
}