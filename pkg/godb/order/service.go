package order

import (
	"context"
	"strings"

	"github.com/sandokandias/go-database-app/pkg/godb/validators"
)

// DefaultService type that implements Service interface
type DefaultService struct {
	storage Storage
}

// NewService creates a new order service with storage dependency
func NewService(storage Storage) DefaultService {
	return DefaultService{storage: storage}
}

// Order validates id field and gets from storage
func (s DefaultService) Order(ctx context.Context, id string) (Order, error) {
	if strings.TrimSpace(id) == "" {
		return Order{}, nil
	}

	return s.storage.Order(ctx, id)
}

// CreateOrder validates required fields and stores in storage
func (s DefaultService) CreateOrder(ctx context.Context, o CreateOrder) error {
	order, err := NewOrder(o.ID, o.Amount, o.Items)
	if err != nil {
		return err

	}

	if err := s.storage.SaveOrder(ctx, order); err != nil {
		return err
	}

	return nil
}

// DeleteOrder validates id field and removes order from storage
func (s DefaultService) DeleteOrder(ctx context.Context, id string) error {
	if err := validators.StringRequired("id", id); err != nil {
		return err
	}

	return s.storage.DeleteOrder(ctx, id)
}
