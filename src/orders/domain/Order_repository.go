package domain

import (
	"mi-tienda-online/src/orders/domain/entities"
)

type OrderRepository interface {
	FindAll() ([]entities.Order, error)
	FindById(order_id int) (entities.Order, error)
	Save(order *entities.Order) error
	Update(order *entities.Order) error
	Delete(order_id int) error
}
