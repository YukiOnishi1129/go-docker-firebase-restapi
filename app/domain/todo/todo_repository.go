package todo

import "context"

// todoのドメインオブジェクトの永続化に関するインターフェース
type TodoRepository interface {
	Save(ctx context.Context, todo *Todo) error
	FindByID(ctx context.Context, id string) (*Todo, error)
	FindAll(ctx context.Context) ([]*Todo, error)
}
