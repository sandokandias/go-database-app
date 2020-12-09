package order

import (
	"context"

	"github.com/sandokandias/go-database-app/pkg/godb/db"
)

// Storage interface that defines the order storage operations
type Storage interface {
	Order(ctx context.Context, id string) (Order, error)
	SaveOrder(tcx db.TxContext, order Order) error
	DeleteOrder(tcx db.TxContext, id string) error
}
