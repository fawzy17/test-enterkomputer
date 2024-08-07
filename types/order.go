package types

import (
	"time"
)

type OrderStore interface {
	CreateOrder(order []Order, meja string) error
	GetBill(orderId string) ([]BillResponse, error)
	GetMeja(orderId string) (*string, error)
}

type Order struct {
	ID         int       `json:"id"`
	OrderId    string    `json:"orderId"`
	ProductId  int       `json:"productId"`
	Quantity   int       `json:"quantity"`
	TotalPrice int       `json:"totalPice"`
	CreatedAt  time.Time `json:"createdAt"`
}

type OrderPayload struct {
	Meja int `json:"meja" validate:"required"`
	Products []ProductOrder `json:"products"`
}

type ProductOrder struct {
	ID       int `json:"id" validate:"required"`
	Quantity int `json:"quantity" validate:"required,min=1,max=100"`
}

type OrderResponse struct {
	Name       string `json:"name"`
	Variant    string `json:"variant"`
	Quantity   int    `json:"quantity"`
	TotalPrice int    `json:"totalPrice"`
}

type BillResponse struct {
	Quantity   int    `json:"quantity"`
	TotalPrice int    `json:"totalPrice"`
	Name       string `json:"name"`
	Variant    string `json:"variant"`
	Price      int    `json:"price"`
}
