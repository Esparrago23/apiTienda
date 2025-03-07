package entities

import (
	"time"
)

type Product struct {
	ProductID   int       `json:"product_id"`
	Name        string    `json:"name" validate:"required,min=3,max=100"`
	Description string    `json:"description" validate:"required,min=10,max=500"`
	Price       float64   `json:"price" validate:"required,gt=0"`
	Stock       int       `json:"stock" validate:"required,gte=0"`
	CreatedAt   time.Time `json:"created_at" `
}

func NewProduct(product_id int, name string, description string, price float64, stock int, created_at time.Time) *Product {
	return &Product{
		ProductID:   product_id,
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
		CreatedAt:   created_at,
	}
}

func (p *Product) GetProductID() int {
	return p.ProductID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetDescription() string {
	return p.Description
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func (p *Product) GetStock() int {
	return p.Stock
}
