package order

import (
	"context"
	"time"

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
	if err := validators.RequiredString("id", o.ID); err != nil {
		return err
	}

	order := godb.Order{
		ID:        o.ID,
		Amount:    o.Amount,
		CreatedAt: time.Now(),
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