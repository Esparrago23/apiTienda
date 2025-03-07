package entities

import (
	"time"
)

type Payment struct {
	PaymentID     int       `json:"payment_id" db:"payment_id"`
	OrderID       int       `json:"order_id" db:"order_id"`
	PaymentStatus string    `json:"payment_status" db:"payment_status"`
	PaymentDate   time.Time `json:"payment_date" db:"payment_date"`
}

func NewPayment(paymentID int, orderID int, paymentStatus string, paymentDate time.Time) *Payment {
	return &Payment{
		PaymentID:     paymentID,
		OrderID:       orderID,
		PaymentStatus: paymentStatus,
		PaymentDate:   paymentDate,
	}
}

func (p *Payment) GetPaymentID() int {
	return p.PaymentID
}

func (p *Payment) GetOrderID() int {
	return p.OrderID
}

func (p *Payment) GetPaymentStatus() string {
	return p.PaymentStatus
}

func (p *Payment) GetPaymentDate() time.Time {
	return p.PaymentDate
}
