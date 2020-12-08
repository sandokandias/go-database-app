package order

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// Storage interface that defines the order storage operations
type Storage interface {
	Order(ctx context.Context, id string) (Order, error)
	SaveOrder(ctx context.Context, tx pgx.Tx, order Order) error
	DeleteOrder(ctx context.Context, tx pgx.Tx, id string) error
}
