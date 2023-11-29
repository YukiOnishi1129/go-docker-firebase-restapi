package todo

import (
	"context"
	todoDomain "github.com/YukiOnishi1129/go-docker-firebase-restapi/domain/todo"
)

type DeleteTodoUseCase struct {
	todoRepo todoDomain.TodoRepository
}

func NewDeleteTodoUseCase(todoRepo todoDomain.TodoRepository) *DeleteTodoUseCase {
	return &DeleteTodoUseCase{
		todoRepo: todoRepo,
	}
}

func (uc *DeleteTodoUseCase) Run(ctx context.Context, id string) error {
	t, err := uc.todoRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	err = uc.todoRepo.Remove(ctx, t.ID())
	if err != nil {
		return err
	}
	return nil
}
