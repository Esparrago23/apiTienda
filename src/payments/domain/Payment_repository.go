package domain

import (
	"mi-tienda-online/src/payments/domain/entities"
)

type PaymentRepository interface {
	FindAll() ([]entities.Payment, error)
	FindById(payment_id int) (entities.Payment, error)
	Save(payment *entities.Payment) error
	Update(payment *entities.Payment) error
	Delete(payment_id int) error
}
