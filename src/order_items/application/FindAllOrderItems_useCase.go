package application

import (
	"mi-tienda-online/src/order_items/domain"
	"mi-tienda-online/src/order_items/domain/entities"
)

type FindAllOrderItemsUseCase struct {
	OrderItemRepository domain.OrderItemRepository
}

func NewFindAllOrderItemsUseCase(OrderItemRepository domain.OrderItemRepository) *FindAllOrderItemsUseCase {
	return &FindAllOrderItemsUseCase{OrderItemRepository: OrderItemRepository}
}

func (useCase *FindAllOrderItemsUseCase) Execute() ([]entities.OrderItem, error) {
	orderItems, err := useCase.OrderItemRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return orderItems, nil
}
