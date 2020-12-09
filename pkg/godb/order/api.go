package order

import "context"

// CreateOrder type that represents a request for order creation
type CreateOrder struct {
	ID       string       `json:"id"`
	Amount   int64        `json:"amount"`
	Customer CustomerData `json:"customer"`
	Items    []ItemData   `json:"items"`
}

// CustomerData type that represents a customer of the order request
type CustomerData struct {
	Name     string `json:"name"`
	Document string `json:"document"`
	Address  string `json:"address"`
}

// ItemData type that represents a item of the order request
type ItemData struct {
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int    `json:"quantity"`
}

// ItemsData type that represents a collection of items
type ItemsData []ItemData

// Service interface that defines the order use cases
type Service interface {
	Order(ctx context.Context, id string) (Order, error)
	CreateOrder(ctx context.Context, order CreateOrder) error
	DeleteOrder(ctx context.Context, id string) error
}
