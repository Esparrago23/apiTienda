package application

import (
	"mi-tienda-online/src/products/domain"
)

type DeleteProductUseCase struct {
	ProductRepository domain.ProductRepository
}

func NewDeleteProductUseCase(ProductRepository domain.ProductRepository) *DeleteProductUseCase {
	return &DeleteProductUseCase{ProductRepository: ProductRepository}
}

func (useCase *DeleteProductUseCase) Execute(product_id int) error {
	err := useCase.ProductRepository.Delete(product_id)
	if err != nil {
		return err
	}
	return nil
}