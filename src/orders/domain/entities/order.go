package entities

import (
	"time"
)

type Order struct {
	OrderID     int       `json:"order_id" db:"order_id"`
	UserID      int       `json:"user_id" db:"user_id"`
	TotalAmount float64   `json:"total_amount" db:"total_amount"`
	Status      string    `json:"status" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func NewOrder(orderID int, userID int, totalAmount float64, status string, createdAt time.Time, updatedAt time.Time) *Order {
	return &Order{
		OrderID:     orderID,
		UserID:      userID,
		TotalAmount: totalAmount,
		Status:      status,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

func (o *Order) GetOrderID() int {
	return o.OrderID
}

func (o *Order) GetUserID() int {
	return o.UserID
}

func (o *Order) GetTotalAmount() float64 {
	return o.TotalAmount
}

func (o *Order) GetStatus() string {
	return o.Status
}

func (o *Order) GetCreatedAt() time.Time {
	return o.CreatedAt
}

func (o *Order) GetUpdatedAt() time.Time {
	return o.UpdatedAt
}
