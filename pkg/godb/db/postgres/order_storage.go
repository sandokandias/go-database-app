package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sandokandias/go-database-app/pkg/godb/db"
	"github.com/sandokandias/go-database-app/pkg/godb/order"
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
func (s OrderStorage) Order(ctx context.Context, id string) (order.Order, error) {
	SQL := `SELECT order_id, amount, created_at FROM orders WHERE id = $1`
	var o order.Order

	if err := s.db.QueryRow(ctx, SQL, id).Scan(&o.ID, &o.Amount, &o.CreatedAt); err != nil {
		if err == pgx.ErrNoRows {
			return o, nil
		}
		return o, fmt.Errorf("query order by %v: %w", id, err)
	}
	return o, nil
}

// SaveOrder persists entity order in postgresql
func (s OrderStorage) SaveOrder(tcx db.TxContext, o order.Order) error {
	orderSQL := `INSERT INTO orders(order_id, amount, created_at, customer_id) VALUES($1, $2, $3, $4)`
	itemsSQL := `INSERT INTO items(item_id, order_id, name, price, quantity) VALUES($1, $2, $3, $4, $5)`

	ctx := tcx.Context()
	tx := tcx.Tx()

	batch := &pgx.Batch{}
	batch.Queue(orderSQL, o.ID, o.Amount, o.CreatedAt, o.CustomerID)
	for _, i := range o.Items {
		batch.Queue(itemsSQL, i.ID, o.ID, i.Name, i.Price, i.Quantity)
	}

	br := tx.SendBatch(ctx, batch)

	for i := 0; i < 1+len(o.Items); i++ {
		_, err := br.Exec()
		if err != nil {
			return fmt.Errorf("exec insert order %v: %w", o, err)
		}
	}

	err := br.Close()
	if err != nil {
		return fmt.Errorf("close batch insert order %v: %w", o, err)
	}

	return nil
}

// DeleteOrder removes entity order from postgresql
func (s OrderStorage) DeleteOrder(tcx db.TxContext, id string) error {
	SQL := `DELETE FROM orders WHERE id = $1`

	ctx := tcx.Context()
	tx := tcx.Tx()

	_, err := tx.Exec(ctx, SQL, id)
	if err != nil {
		return fmt.Errorf("exec delete order %v: %w", id, err)
	}

	return nil
}
