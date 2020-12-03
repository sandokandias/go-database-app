package godb

import (
	"context"
	"time"
)

// Order type that represents a order entity
type Order struct {
	ID        string
	Amount    int64
	CreatedAt time.Time
}

// OrderStorage interface that defines the order storage operations
type OrderStorage interface {
	Order(ctx context.Context, id string) (Order, error)
	SaveOrder(ctx context.Context, order Order) error
	DeleteOrder(ctx context.Context, id string) error
}

// OrderService interface that defines the workspace business logic
type OrderService interface {
	Order(ctx context.Context, id string) (Order, error)
	CreateOrder(ctx context.Context, order CreateOrder) error
	DeleteOrder(ctx context.Context, id string) error
}