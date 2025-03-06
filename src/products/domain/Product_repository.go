package domain

import (
	"mi-tienda-online/src/products/domain/entities"
)

type ProductRepository interface {
	FindAll() ([]entities.Product, error)
	FindById(product_id int) (entities.Product, error)
	Save(product *entities.Product) error
	Update(product *entities.Product) error
	Delete(product_id int) error
}
