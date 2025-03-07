package application

import (
	"mi-tienda-online/src/payments/domain"
	"mi-tienda-online/src/payments/domain/entities"
)

type CreatePaymentUseCase struct {
	PaymentRepository domain.PaymentRepository
}

func NewCreatePaymentUseCase(PaymentRepository domain.PaymentRepository) *CreatePaymentUseCase {
	return &CreatePaymentUseCase{PaymentRepository: PaymentRepository}
}

func (useCase *CreatePaymentUseCase) Execute(payment *entities.Payment) error {
	err := useCase.PaymentRepository.Save(payment)
	if err != nil {
		return err
	}
	return nil
}
