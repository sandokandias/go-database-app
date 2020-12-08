package customer

import "context"

// Storage interface that defines the customer storage operations
type Storage interface {
	SaveCustomer(ctx context.Context, customer Customer) error
}
