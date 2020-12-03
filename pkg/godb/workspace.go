package godb

import (
	"context"
)

// Workspace type that represents a workspace entity
type Workspace struct {
	Name string
}

// WorkspaceStorage interface that defines the workspace storage
type WorkspaceStorage interface {
	Workspace(ctx context.Context, name string) (Workspace, error)
	SaveWorkspace(ctx context.Context, workspace Workspace) error
	DeleteWorkspace(ctx context.Context, name string) error
}

// WorkspaceService interface that defines the workspace service
type WorkspaceService interface {
	Workspace(ctx context.Context, name string) (Workspace, error)
	CreateWorkspace(ctx context.Context, workspace CreateWorkspace) error
	DeleteWorkspace(ctx context.Context, name string) error
}
