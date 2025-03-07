package application

import (
	"mi-tienda-online/src/order_items/domain"
	"mi-tienda-online/src/order_items/domain/entities"
)

type UpdateOrderItemUseCase struct {
	OrderItemRepository domain.OrderItemRepository
}

func NewUpdateOrderItemUseCase(OrderItemRepository domain.OrderItemRepository) *UpdateOrderItemUseCase {
	return &UpdateOrderItemUseCase{OrderItemRepository: OrderItemRepository}
}

func (useCase *UpdateOrderItemUseCase) Execute(orderItem *entities.OrderItem) error {
	err := useCase.OrderItemRepository.Update(orderItem)
	if err != nil {
		return err
	}
	return nil
}
