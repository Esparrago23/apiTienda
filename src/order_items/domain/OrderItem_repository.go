package domain

import (
	"mi-tienda-online/src/order_items/domain/entities"
)

type OrderItemRepository interface {
	FindAll() ([]entities.OrderItem, error)
	FindById(order_item_id int) (entities.OrderItem, error)
	Save(orderItem *entities.OrderItem) error
	Update(orderItem *entities.OrderItem) error
	Delete(order_item_id int) error
}
