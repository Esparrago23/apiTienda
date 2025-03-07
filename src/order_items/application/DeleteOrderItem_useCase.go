package application

import (
	"mi-tienda-online/src/order_items/domain"
)

type DeleteOrderItemUseCase struct {
	OrderItemRepository domain.OrderItemRepository
}

func NewDeleteOrderItemUseCase(OrderItemRepository domain.OrderItemRepository) *DeleteOrderItemUseCase {
	return &DeleteOrderItemUseCase{OrderItemRepository: OrderItemRepository}
}

func (useCase *DeleteOrderItemUseCase) Execute(order_item_id int) error {
	err := useCase.OrderItemRepository.Delete(order_item_id)
	if err != nil {
		return err
	}
	return nil
}
