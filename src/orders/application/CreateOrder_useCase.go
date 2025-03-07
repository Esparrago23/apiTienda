package application

import (
	"mi-tienda-online/src/orders/domain"
	"mi-tienda-online/src/orders/domain/entities"
)

type CreateOrderUseCase struct {
	OrderRepository domain.OrderRepository
}

func NewCreateOrderUseCase(OrderRepository domain.OrderRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{OrderRepository: OrderRepository}
}

func (useCase *CreateOrderUseCase) Execute(order *entities.Order) error {
	err := useCase.OrderRepository.Save(order)
	if err != nil {
		return err
	}
	return nil
}
