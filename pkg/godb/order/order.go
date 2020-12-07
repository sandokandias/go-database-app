package order

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
	"github.com/sandokandias/go-database-app/pkg/godb"
	"github.com/sandokandias/go-database-app/pkg/godb/validators"
)

// Order type that represents a order entity
type Order struct {
	ID        string
	Amount    int64
	CreatedAt time.Time
	Items     Items
}

// NewOrder validates the order fields and if ok, creates a new order
func NewOrder(ID string, amount int64, items ItemsData) (Order, error) {
	var result error

	if err := validators.StringRequired("id", ID); err != nil {
		result = multierror.Append(result, err)
	}

	if err := validators.Int64GreaterZero("amount", amount); err != nil {
		result = multierror.Append(result, err)
	}

	if len(items) == 0 {
		err := godb.ErrRequiredField("items")
		result = multierror.Append(result, err)
	}

	ii := Items{}
	for i, it := range items {
		var err error
		if err = validators.StringRequired(fmt.Sprintf("items[%d].name", i), it.Name); err != nil {
			result = multierror.Append(result, err)
		}

		if err = validators.Int64GreaterZero(fmt.Sprintf("items[%d].price", i), it.Price); err != nil {
			result = multierror.Append(result, err)
		}

		if err = validators.IntGreaterZero(fmt.Sprintf("items[%d].quantity", i), it.Quantity); err != nil {
			result = multierror.Append(result, err)
		}

		if err == nil {
			item := Item{
				ID:       uuid.New().String(),
				Name:     it.Name,
				Price:    it.Price,
				Quantity: it.Quantity,
			}
			ii = append(ii, item)
		}

	}

	if result != nil {
		return Order{}, result
	}

	order := Order{
		ID:        ID,
		Amount:    amount,
		CreatedAt: time.Now(),
		Items:     ii,
	}

	return order, nil

}
