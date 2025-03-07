package application

import (
	"mi-tienda-online/src/order_items/domain"
	"mi-tienda-online/src/order_items/domain/entities"
)

type FindOrderItemByIdUseCase struct {
	OrderItemRepository domain.OrderItemRepository
}

func NewFindOrderItemByIdUseCase(OrderItemRepository domain.OrderItemRepository) *FindOrderItemByIdUseCase {
	return &FindOrderItemByIdUseCase{OrderItemRepository: OrderItemRepository}
}

func (useCase *FindOrderItemByIdUseCase) Execute(order_item_id int) (*entities.OrderItem, error) {
	orderItem, err := useCase.OrderItemRepository.FindById(order_item_id)
	if err != nil {
		return nil, err
	}
	return &orderItem, nil
}
