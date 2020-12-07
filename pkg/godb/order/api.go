package order

import "context"

// CreateOrder type that represents a request for order creation
type CreateOrder struct {
	ID     string     `json:"id"`
	Amount int64      `json:"amount"`
	Items  []ItemData `json:"items"`
}

// ItemData type that represents a item of the order entity
type ItemData struct {
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int    `json:"quantity"`
}

// ItemsData type that represents a collection of items
type ItemsData []ItemData

// Storage interface that defines the order storage operations
type Storage interface {
	Order(ctx context.Context, id string) (Order, error)
	SaveOrder(ctx context.Context, order Order) error
	DeleteOrder(ctx context.Context, id string) error
}

// Service interface that defines the workspace business logic
type Service interface {
	Order(ctx context.Context, id string) (Order, error)
	CreateOrder(ctx context.Context, order CreateOrder) error
	DeleteOrder(ctx context.Context, id string) error
}
