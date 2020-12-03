package order

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/hashicorp/go-multierror"

	"github.com/sandokandias/go-database-app/pkg/godb"
	"github.com/sandokandias/go-database-app/pkg/godb/validators"
)

// Service type that implements OrderService interface
type Service struct {
	storage godb.OrderStorage
}

// NewService creates a new order service with storage dependency
func NewService(storage godb.OrderStorage) Service {
	return Service{storage: storage}
}

// Order validate id field and gets from storage
func (s Service) Order(ctx context.Context, id string) (godb.Order, error) {
	if err := validators.RequiredString("id", id); err != nil {
		return godb.Order{}, err
	}

	return s.storage.Order(ctx, id)
}

// CreateOrder validates required fields and stores in storage
func (s Service) CreateOrder(ctx context.Context, o godb.CreateOrder) error {
	var result error

	if err := validators.RequiredString("id", o.ID); err != nil {
		result = multierror.Append(result, err)
	}

	if err := validators.Int64GreaterZero("amount", o.Amount); err != nil {
		result = multierror.Append(result, err)
	}

	if len(o.Items) == 0 {
		err := godb.ErrRequiredField("items")
		result = multierror.Append(result, err)
	}

	for i, it := range o.Items {
		if err := validators.RequiredString(fmt.Sprintf("items[%d].name", i), it.Name); err != nil {
			result = multierror.Append(result, err)
		}

		if err := validators.Int64GreaterZero(fmt.Sprintf("items[%d].price", i), it.Price); err != nil {
			result = multierror.Append(result, err)
		}

		if err := validators.IntGreaterZero(fmt.Sprintf("items[%d].quantity", i), it.Quantity); err != nil {
			result = multierror.Append(result, err)
		}
	}

	if result != nil {
		return result
	}

	items := godb.Items{}
	for _, i := range o.Items {
		item := godb.Item{
			ID:       uuid.New().String(),
			Name:     i.Name,
			Price:    i.Price,
			Quantity: i.Quantity,
		}
		items = append(items, item)
	}

	order := godb.Order{
		ID:        o.ID,
		Amount:    o.Amount,
		CreatedAt: time.Now(),
		Items:     items,
	}

	if err := s.storage.SaveOrder(ctx, order); err != nil {
		return err
	}

	return nil
}

// DeleteOrder validates id field and removes order from storage
func (s Service) DeleteOrder(ctx context.Context, id string) error {
	if err := validators.RequiredString("id", id); err != nil {
		return err
	}

	return s.storage.DeleteOrder(ctx, id)
}
