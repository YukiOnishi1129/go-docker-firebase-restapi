package todo

import (
	"context"
	todoDomain "github.com/YukiOnishi1129/go-docker-firebase-restapi/domain/todo"
)

type CreateTodoUseCase struct {
	todoRepo todoDomain.TodoRepository
}

func NewCreateTodoUseCase(todoRepo todoDomain.TodoRepository) *CreateTodoUseCase {
	return &CreateTodoUseCase{
		todoRepo: todoRepo,
	}
}

type CreateTodoUseCaseInputDTO struct {
	Title       string
	Description string
}

type CreateTodoUseCaseOutputDTO struct {
	ID          string
	Title       string
	Description string
}

func (uc *CreateTodoUseCase) Run(ctx context.Context, input CreateTodoUseCaseInputDTO) (*CreateTodoUseCaseOutputDTO, error) {
	t, err := todoDomain.NewTodo(input.Title, input.Description)
	if err != nil {
		return nil, err
	}
	err = uc.todoRepo.Save(ctx, t)
	if err != nil {
		return nil, err
	}
	return &CreateTodoUseCaseOutputDTO{
		ID:          t.ID(),
		Title:       t.Title(),
		Description: t.Description(),
	}, nil
}
