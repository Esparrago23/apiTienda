package application

import (
	"mi-tienda-online/src/order_items/domain"
	"mi-tienda-online/src/order_items/domain/entities"
)

type CreateOrderItemUseCase struct {
	OrderItemRepository domain.OrderItemRepository
}

func NewCreateOrderItemUseCase(OrderItemRepository domain.OrderItemRepository) *CreateOrderItemUseCase {
	return &CreateOrderItemUseCase{OrderItemRepository: OrderItemRepository}
}

func (useCase *CreateOrderItemUseCase) Execute(orderItem *entities.OrderItem) error {
	err := useCase.OrderItemRepository.Save(orderItem)
	if err != nil {
		return err
	}
	return nil
}
