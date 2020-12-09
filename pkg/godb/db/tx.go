package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// TxManager type that represents the db transacion manager
type TxManager struct {
	DB *pgxpool.Pool
}

// NewTxManager creates a new transaction manager
func NewTxManager(db *pgxpool.Pool) TxManager {
	return TxManager{DB: db}
}

// Exec begins the transaction, call anomymous func and if everything is ok, then the transaction will be committed,
// otherwise the transaction will be rollbacked
func (t TxManager) Exec(ctx context.Context, fn func(ctx context.Context, tx pgx.Tx) error) error {
	tx, err := t.DB.Begin(ctx)
	if err != nil {
		return err
	}

	err = fn(ctx, tx)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rollback err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}
