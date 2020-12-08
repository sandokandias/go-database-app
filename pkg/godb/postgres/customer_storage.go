package postgres

import (
	"context"
	"fmt"

	"github.com/sandokandias/go-database-app/pkg/godb/customer"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// CustomerStorage type that implements customer.Storage interface for postgresql
type CustomerStorage struct {
	db *pgxpool.Pool
}

// NewCustomerStorage creates a new customer storage with db dependency
func NewCustomerStorage(db *pgxpool.Pool) CustomerStorage {
	return CustomerStorage{db: db}
}

// Customer returns customer by document from postgresql
func (s CustomerStorage) Customer(ctx context.Context, document string) (customer.Customer, error) {
	SQL := `SELECT name, customer_id, address FROM customers WHERE customer_id = $1`
	var c customer.Customer

	if err := s.db.QueryRow(ctx, SQL, document).Scan(&c.Name, &c.Document, &c.Address); err != nil {
		if err == pgx.ErrNoRows {
			return c, nil
		}
		return c, fmt.Errorf("query customer by %v: %w", document, err)
	}
	return c, nil
}

// SaveCustomer persists entity customer in postgresql
func (s CustomerStorage) SaveCustomer(ctx context.Context, tx pgx.Tx, customer customer.Customer) error {
	SQL := `INSERT INTO customers(name, customer_id, address) VALUES($1, $2, $3)`

	_, err := tx.Exec(ctx, SQL, customer.Name, customer.Document, customer.Address)
	if err != nil {
		return fmt.Errorf("exec insert customer %v: %w", customer.Document, err)
	}

	return nil
}
