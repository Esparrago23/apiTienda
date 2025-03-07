package entities

type OrderItem struct {
	OrderItemID int     `json:"order_item_id" db:"order_item_id"`
	OrderID     int     `json:"order_id" db:"order_id"`
	ProductID   int     `json:"product_id" db:"product_id"`
	Quantity    int     `json:"quantity" db:"quantity"`
	Price       float64 `json:"price" db:"price"`
	TotalPrice  float64 `json:"total_price" db:"total_price"`
}

func NewOrderItem(orderItemID int, orderID int, productID int, quantity int, price float64, totalPrice float64) *OrderItem {
	return &OrderItem{
		OrderItemID: orderItemID,
		OrderID:     orderID,
		ProductID:   productID,
		Quantity:    quantity,
		Price:       price,
		TotalPrice:  totalPrice,
	}
}

func (oi *OrderItem) GetOrderItemID() int {
	return oi.OrderItemID
}

func (oi *OrderItem) GetOrderID() int {
	return oi.OrderID
}

func (oi *OrderItem) GetProductID() int {
	return oi.ProductID
}

func (oi *OrderItem) GetQuantity() int {
	return oi.Quantity
}

func (oi *OrderItem) GetPrice() float64 {
	return oi.Price
}

func (oi *OrderItem) GetTotalPrice() float64 {
	return oi.TotalPrice
}
