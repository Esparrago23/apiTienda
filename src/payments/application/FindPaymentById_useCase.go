package application

import (
	"mi-tienda-online/src/payments/domain"
	"mi-tienda-online/src/payments/domain/entities"
)

type FindPaymentByIdUseCase struct {
	PaymentRepository domain.PaymentRepository
}

func NewFindPaymentByIdUseCase(PaymentRepository domain.PaymentRepository) *FindPaymentByIdUseCase {
	return &FindPaymentByIdUseCase{PaymentRepository: PaymentRepository}
}

func (useCase *FindPaymentByIdUseCase) Execute(payment_id int) (*entities.Payment, error) {
	payment, err := useCase.PaymentRepository.FindById(payment_id)
	if err != nil {
		return nil, err
	}
	return &payment, nil
}
