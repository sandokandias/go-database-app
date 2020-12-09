package customer

import (
	"context"

	"github.com/sandokandias/go-database-app/pkg/godb/db"
)

// Storage interface that defines the customer storage operations
type Storage interface {
	Customer(ctx context.Context, document string) (Customer, error)
	SaveCustomer(tcx db.TxContext, customer Customer) error
}
