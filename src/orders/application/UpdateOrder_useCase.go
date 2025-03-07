package application

import (
	"mi-tienda-online/src/orders/domain"
	"mi-tienda-online/src/orders/domain/entities"
)

type UpdateOrderUseCase struct {
	OrderRepository domain.OrderRepository
}

func NewUpdateOrderUseCase(OrderRepository domain.OrderRepository) *UpdateOrderUseCase {
	return &UpdateOrderUseCase{OrderRepository: OrderRepository}
}

func (useCase *UpdateOrderUseCase) Execute(order *entities.Order) error {
	err := useCase.OrderRepository.Update(order)
	if err != nil {
		return err
	}
	return nil
}
