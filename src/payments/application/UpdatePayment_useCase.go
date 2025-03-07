package application

import (
	"mi-tienda-online/src/payments/domain"
	"mi-tienda-online/src/payments/domain/entities"
)

type UpdatePaymentUseCase struct {
	PaymentRepository domain.PaymentRepository
}

func NewUpdatePaymentUseCase(PaymentRepository domain.PaymentRepository) *UpdatePaymentUseCase {
	return &UpdatePaymentUseCase{PaymentRepository: PaymentRepository}
}

func (useCase *UpdatePaymentUseCase) Execute(payment *entities.Payment) error {
	err := useCase.PaymentRepository.Update(payment)
	if err != nil {
		return err
	}
	return nil
}
