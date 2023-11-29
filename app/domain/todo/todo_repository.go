package todo

import "context"

// TodoRepository is an interface for persistence of todo domain object
type TodoRepository interface {
	Save(ctx context.Context, todo *Todo) error
	FindByID(ctx context.Context, id string) (*Todo, error)
	FindAll(ctx context.Context) ([]*Todo, error)
	Remove(ctx context.Context, id string) error
}
