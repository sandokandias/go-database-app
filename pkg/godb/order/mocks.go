package order

import (
	"context"

	"github.com/sandokandias/go-database-app/pkg/godb/customer"
	"github.com/sandokandias/go-database-app/pkg/godb/db"
)

type MockTxManager struct {
}

func (MockTxManager) Exec(ctx context.Context, fn func(txc db.TxContext) error) error {
	return nil
}

type MockOrderStorage struct {
}

func (MockOrderStorage) Order(ctx context.Context, id string) (Order, error) {
	return Order{}, nil
}

func (MockOrderStorage) SaveOrder(tcx db.TxContext, order Order) error {
	return nil
}

func (MockOrderStorage) DeleteOrder(tcx db.TxContext, id string) error {
	return nil
}

type MockCustomerStorage struct {
}

func (MockCustomerStorage) Customer(ctx context.Context, document string) (customer.Customer, error) {
	return customer.Customer{}, nil
}

func (MockCustomerStorage) SaveCustomer(tcx db.TxContext, customer customer.Customer) error {
	return nil
}
