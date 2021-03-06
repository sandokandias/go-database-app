package customer

import (
	"github.com/hashicorp/go-multierror"
	"github.com/sandokandias/go-database-app/pkg/godb/validators"
)

// Customer type that represents a customer entity
type Customer struct {
	Name     string
	Document string
	Address  string
}

// New validates the customer fields and if ok, creates a new customer
func New(name, document, address string) (Customer, error) {
	var result error

	if err := validators.StringRequired("customer.name", name); err != nil {
		result = multierror.Append(result, err)
	}

	if err := validators.StringRequired("customer.document", document); err != nil {
		result = multierror.Append(result, err)
	}

	if err := validators.StringRequired("customer.address", address); err != nil {
		result = multierror.Append(result, err)
	}

	if result != nil {
		return Customer{}, result
	}

	customer := Customer{
		Name:     name,
		Document: document,
		Address:  address,
	}

	return customer, nil

}
