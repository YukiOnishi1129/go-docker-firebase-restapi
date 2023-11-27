package repository

import (
	"context"
	"github.com/YukiOnishi1129/go-docker-firebase-restapi/domain/todo"
)

type todoRepository struct {
}

func NewTodoRepository() todo.TodoRepository {
	return &todoRepository{}
}

func (r *todoRepository) Save(ctx context.Context, todo *todo.Todo) error {
	return nil
}

func (r *todoRepository) FindByID(ctx context.Context, id string) (*todo.Todo, error) {
	return nil, nil
}

func (r *todoRepository) FindAll(ctx context.Context) ([]*todo.Todo, error) {
	return nil, nil
}
