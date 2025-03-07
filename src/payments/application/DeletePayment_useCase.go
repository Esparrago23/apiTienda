package application

import (
	"mi-tienda-online/src/payments/domain"
)

type DeletePaymentUseCase struct {
	PaymentRepository domain.PaymentRepository
}

func NewDeletePaymentUseCase(PaymentRepository domain.PaymentRepository) *DeletePaymentUseCase {
	return &DeletePaymentUseCase{PaymentRepository: PaymentRepository}
}

func (useCase *DeletePaymentUseCase) Execute(payment_id int) error {
	err := useCase.PaymentRepository.Delete(payment_id)
	if err != nil {
		return err
	}
	return nil
}
