package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sandokandias/go-database-app/pkg/godb"
)

type WorkspaceStorage struct {
	db *pgxpool.Pool
}

func NewWorkspaceStorage(db *pgxpool.Pool) WorkspaceStorage {
	return WorkspaceStorage{db: db}
}

func (s WorkspaceStorage) Workspace(ctx context.Context, name string) (godb.Workspace, error) {
	SQL := `SELECT name FROM workspaces WHERE name = ?`
	var w godb.Workspace

	if err := s.db.QueryRow(ctx, SQL, name).Scan(&w.Name); err != nil {
		return w, fmt.Errorf("query workspace by %v: %w", name, err)
	}
	return w, nil

}

func (s WorkspaceStorage) SaveWorkspace(ctx context.Context, w godb.Workspace) error {
	SQL := `INSERT INTO workspaces(name) VALUES($1)`

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin tx insert workspace %v: %w", w.Name, err)
	}

	defer func() {
		err = tx.Rollback(ctx)
		if err != nil {
			log.Printf("rollback insert workspace %s: %v", w, err)
		}
	}()

	_, err = tx.Exec(ctx, SQL, w.Name)
	if err != nil {
		return fmt.Errorf("exec insert workspace %v: %w", w, err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("commit insert workspace %v: %w", w.Name, err)
	}

	return nil
}

func (s WorkspaceStorage) DeleteWorkspace(ctx context.Context, name string) error {
	SQL := `DELETE FROM workspaces WHERE name = ?`

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin tx delete workspace %v: %w", name, err)
	}

	defer func() {
		err = tx.Rollback(ctx)
		if err != nil {
			log.Printf("rollback delete workspace %s: %v", name, err)
		}
	}()

	_, err = tx.Exec(ctx, SQL, name)
	if err != nil {
		return fmt.Errorf("exec delete workspace %v: %w", name, err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("commit delete workspace %v: %w", name, err)
	}

	return nil

}
