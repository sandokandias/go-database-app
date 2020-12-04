package order

import (
	"context"
	"fmt"
	"strings"
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

// Order validates id field and gets from storage
func (s Service) Order(ctx context.Context, id string) (godb.Order, error) {
	if strings.TrimSpace(id) == "" {
		return godb.Order{}, nil
	}

	return s.storage.Order(ctx, id)
}

// CreateOrder validates required fields and stores in storage
func (s Service) CreateOrder(ctx context.Context, o godb.CreateOrder) error {
	if err := s.validateOrder(o); err != nil {
		return err

	}

	order := godb.Order{
		ID:        o.ID,
		Amount:    o.Amount,
		CreatedAt: time.Now(),
		Items:     godb.Items{},
	}

	for _, it := range o.Items {
		item := godb.Item{
			ID:       uuid.New().String(),
			Name:     it.Name,
			Price:    it.Price,
			Quantity: it.Quantity,
		}
		order.Items = append(order.Items, item)
	}

	if err := s.storage.SaveOrder(ctx, order); err != nil {
		return err
	}

	return nil
}

// validateOrder validates all required fields and checks it's values
func (s Service) validateOrder(o godb.CreateOrder) error {
	var result error

	if err := validators.StringRequired("id", o.ID); err != nil {
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
	}

	return result
}

// DeleteOrder validates id field and removes order from storage
func (s Service) DeleteOrder(ctx context.Context, id string) error {
	if err := validators.StringRequired("id", id); err != nil {
		return err
	}

	return s.storage.DeleteOrder(ctx, id)
}
