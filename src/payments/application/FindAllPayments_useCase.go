package application

import (
	"mi-tienda-online/src/payments/domain"
	"mi-tienda-online/src/payments/domain/entities"
)

type FindAllPaymentsUseCase struct {
	PaymentRepository domain.PaymentRepository
}

func NewFindAllPaymentsUseCase(PaymentRepository domain.PaymentRepository) *FindAllPaymentsUseCase {
	return &FindAllPaymentsUseCase{PaymentRepository: PaymentRepository}
}

func (useCase *FindAllPaymentsUseCase) Execute() ([]entities.Payment, error) {
	payments, err := useCase.PaymentRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return payments, nil
}
