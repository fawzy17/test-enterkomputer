package types

import "time"

type ProductStore interface {
	GetAllProducts() ([]Product, error)
	GetProductById(id int) (*Product, error)
	GetPromo() ([]Product, error)
}

type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Variant   string    `json:"variant"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}

type ProductPayload struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Variant string `json:"variant"`
	Price   int    `json:"price"`
}
type ProductResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Variant  string `json:"variant"`
	Quantity int    `json:"quantity"`
}
