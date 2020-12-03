package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sandokandias/go-database-app/pkg/godb"
)

// OrderStorage type that implements OrderStorage interface for postgresql
type OrderStorage struct {
	db *pgxpool.Pool
}

// NewOrderStorage creates a new order storage with db dependency
func NewOrderStorage(db *pgxpool.Pool) OrderStorage {
	return OrderStorage{db: db}
}

// Order returns order by id from postgresql
func (s OrderStorage) Order(ctx context.Context, id string) (godb.Order, error) {
	SQL := `SELECT order_id, amount, created_at FROM roders WHERE id = ?`
	var o godb.Order

	if err := s.db.QueryRow(ctx, SQL, id).Scan(&o.ID, &o.Amount, &o.CreatedAt); err != nil {
		return o, fmt.Errorf("query order by %v: %w", id, err)
	}
	return o, nil

}

// SaveOrder persists entity order in postgresql
func (s OrderStorage) SaveOrder(ctx context.Context, o godb.Order) error {
	SQL := `INSERT INTO orders(order_id, amount, created_at) VALUES($1,$2,$3)`

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin tx insert order %v: %w", o, err)
	}

	defer func() {
		err = tx.Rollback(ctx)
		if err != nil {
			log.Printf("rollback insert order %v: %v", o, err)
		}
	}()

	_, err = tx.Exec(ctx, SQL, o.ID, o.Amount, o.CreatedAt)
	if err != nil {
		return fmt.Errorf("exec insert order %v: %w", o, err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("commit insert order %v: %w", o, err)
	}

	return nil
}

// DeleteOrder removes entity order from postgresql
func (s OrderStorage) DeleteOrder(ctx context.Context, id string) error {
	SQL := `DELETE FROM orders WHERE id = ?`

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin tx delete order %v: %w", id, err)
	}

	defer func() {
		err = tx.Rollback(ctx)
		if err != nil {
			log.Printf("rollback delete order %s: %v", id, err)
		}
	}()

	_, err = tx.Exec(ctx, SQL, id)
	if err != nil {
		return fmt.Errorf("exec delete order %v: %w", id, err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("commit delete order %v: %w", id, err)
	}

	return nil

}
