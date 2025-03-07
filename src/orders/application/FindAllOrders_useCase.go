package application

import (
	"mi-tienda-online/src/orders/domain"
	"mi-tienda-online/src/orders/domain/entities"
)

type FindAllOrdersUseCase struct {
	OrderRepository domain.OrderRepository
}

func NewFindAllOrdersUseCase(OrderRepository domain.OrderRepository) *FindAllOrdersUseCase {
	return &FindAllOrdersUseCase{OrderRepository: OrderRepository}
}

func (useCase *FindAllOrdersUseCase) Execute() ([]entities.Order, error) {
	orders, err := useCase.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return orders, nil
}
