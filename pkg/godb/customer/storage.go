package customer

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// Storage interface that defines the customer storage operations
type Storage interface {
	Customer(ctx context.Context, document string) (Customer, error)
	SaveCustomer(ctx context.Context, tx pgx.Tx, customer Customer) error
}
