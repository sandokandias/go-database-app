package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// TxContent type that presents a db transaction context
type TxContext struct {
	ctx context.Context
	tx  pgx.Tx
}

// WithTx creates a transaction context by a parent context and a db transaction
func WithTx(ctx context.Context, tx pgx.Tx) TxContext {
	return TxContext{
		ctx: ctx,
		tx:  tx,
	}
}

// Context returns the parente context
func (t TxContext) Context() context.Context {
	return t.ctx
}

// Tx returns the db transaction
func (t TxContext) Tx() pgx.Tx {
	return t.tx
}

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
func (t TxManager) Exec(ctx context.Context, fn func(txc TxContext) error) error {
	tx, err := t.DB.Begin(ctx)
	if err != nil {
		return err
	}

	txc := WithTx(ctx, tx)

	err = fn(txc)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rollback err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}
