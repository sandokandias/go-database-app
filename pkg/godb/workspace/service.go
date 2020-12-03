package workspace

import (
	"context"

	"github.com/sandokandias/go-database-app/pkg/godb"
	"github.com/sandokandias/go-database-app/pkg/godb/validators"
)

// Service type that implements WorkspaceService interface
type Service struct {
	storage godb.WorkspaceStorage
}

func NewService(storage godb.WorkspaceStorage) Service {
	return Service{storage: storage}
}

// Workspace validates if the name is not empty and gets workspace by name from database
func (s Service) Workspace(ctx context.Context, name string) (godb.Workspace, error) {
	if err := validators.RequiredString("name", name); err != nil {
		return godb.Workspace{}, err
	}

	return s.storage.Workspace(ctx, name)
}

// CreateWorkspace validates if the name is not empty and stores in database
func (s Service) CreateWorkspace(ctx context.Context, w godb.CreateWorkspace) error {
	if err := validators.RequiredString("name", w.Name); err != nil {
		return err
	}

	entity := godb.Workspace{
		Name: w.Name,
	}

	if err := s.storage.SaveWorkspace(ctx, entity); err != nil {
		return err
	}

	return nil
}

// DeleteWorkspace validates if the name is not empty and removes workspace from database
func (s Service) DeleteWorkspace(ctx context.Context, name string) error {
	if err := validators.RequiredString("name", name); err != nil {
		return err
	}

	return s.storage.DeleteWorkspace(ctx, name)
}
