package todo

import (
	"context"
	todoDomain "github.com/YukiOnishi1129/go-docker-firebase-restapi/domain/todo"
)

type FindByIdTodoUseCase struct {
	todoRepo todoDomain.TodoRepository
}

func NewFindByIdTodoUseCase(todoRepo todoDomain.TodoRepository) *FindByIdTodoUseCase {
	return &FindByIdTodoUseCase{
		todoRepo: todoRepo,
	}
}

type FindByIdTodoUseCaseOutputDTO struct {
	ID          string
	Title       string
	Description string
}

func (uc *FindByIdTodoUseCase) Run(ctx context.Context, id string) (*FindByIdTodoUseCaseOutputDTO, error) {
	t, err := uc.todoRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &FindByIdTodoUseCaseOutputDTO{
		ID:          t.ID(),
		Title:       t.Title(),
		Description: t.Description(),
	}, nil
}
