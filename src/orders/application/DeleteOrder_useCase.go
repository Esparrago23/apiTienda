package application

import (
	"mi-tienda-online/src/orders/domain"
)

type DeleteOrderUseCase struct {
	OrderRepository domain.OrderRepository
}

func NewDeleteOrderUseCase(OrderRepository domain.OrderRepository) *DeleteOrderUseCase {
	return &DeleteOrderUseCase{OrderRepository: OrderRepository}
}

func (useCase *DeleteOrderUseCase) Execute(order_id int) error {
	err := useCase.OrderRepository.Delete(order_id)
	if err != nil {
		return err
	}
	return nil
}
