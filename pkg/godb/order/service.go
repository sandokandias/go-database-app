package order

import (
	"context"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/jackc/pgx/v4"
	"github.com/sandokandias/go-database-app/pkg/godb/customer"
	"github.com/sandokandias/go-database-app/pkg/godb/db"
	"github.com/sandokandias/go-database-app/pkg/godb/validators"
)

// DefaultService type that implements Service interface
type DefaultService struct {
	txManager       *db.TxManager
	orderStorage    Storage
	customerStorage customer.Storage
}

// NewService creates a new order service with storage dependency
func NewService(txManager *db.TxManager,
	orderStorage Storage,
	customerStorage customer.Storage) DefaultService {
	return DefaultService{
		txManager:       txManager,
		orderStorage:    orderStorage,
		customerStorage: customerStorage,
	}
}

// Order validates id field and gets from storage
func (s DefaultService) Order(ctx context.Context, id string) (Order, error) {
	if strings.TrimSpace(id) == "" {
		return Order{}, nil
	}

	return s.orderStorage.Order(ctx, id)
}

// CreateOrder validates required fields and stores in storage
func (s DefaultService) CreateOrder(ctx context.Context, o CreateOrder) error {
	var result error

	customer, err := customer.NewCustomer(o.Customer.Name, o.Customer.Document, o.Customer.Address)
	if err != nil {
		result = multierror.Append(result, err)
	}

	order, err := NewOrder(o.ID, o.Amount, o.Items, customer.Document)
	if err != nil {
		result = multierror.Append(result, err)
	}

	if result != nil {
		return result
	}

	customerExists, err := s.customerExists(ctx, customer)
	if err != nil {
		return err
	}

	err = s.txManager.Exec(ctx, func(ctx context.Context, tx pgx.Tx) error {
		if !customerExists {
			if err := s.customerStorage.SaveCustomer(ctx, tx, customer); err != nil {
				return err
			}
		}

		if err := s.orderStorage.SaveOrder(ctx, tx, order); err != nil {
			return err
		}

		return nil
	})

	return nil
}

func (s DefaultService) customerExists(ctx context.Context, customer customer.Customer) (bool, error) {
	c, err := s.customerStorage.Customer(ctx, customer.Document)
	if err != nil {
		return false, err
	}

	return c.Document != "", nil
}

// DeleteOrder validates id field and removes order from storage
func (s DefaultService) DeleteOrder(ctx context.Context, id string) error {
	if err := validators.StringRequired("id", id); err != nil {
		return err
	}

	return s.txManager.Exec(ctx, func(ctx context.Context, tx pgx.Tx) error {
		return s.orderStorage.DeleteOrder(ctx, tx, id)
	})
}
