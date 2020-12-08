package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TxManager struct {
	DB *pgxpool.Pool
}

func NewTxManager(db *pgxpool.Pool) *TxManager {
	return &TxManager{DB: db}
}

func (t *TxManager) Exec(ctx context.Context, fn func(ctx context.Context, tx pgx.Tx) error) error {
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
