package todo

import (
	"context"
	todoDomain "github.com/YukiOnishi1129/go-docker-firebase-restapi/domain/todo"
)

type UpdateTodoUseCase struct {
	todoRepo todoDomain.TodoRepository
}

func NewUpdateTodoUseCase(todoRepo todoDomain.TodoRepository) *UpdateTodoUseCase {
	return &UpdateTodoUseCase{
		todoRepo: todoRepo,
	}
}

type UpdateTodoUseCaseInputDTO struct {
	ID          string
	Title       string
	Description string
}

type UpdateTodoUseCaseOutputDTO struct {
	ID          string
	Title       string
	Description string
}

func (uc *UpdateTodoUseCase) Run(ctx context.Context, input UpdateTodoUseCaseInputDTO) (*UpdateTodoUseCaseOutputDTO, error) {
	t, err := todoDomain.UpdateTodo(input.ID, input.Title, input.Description)
	if err != nil {
		return nil, err
	}

	err = uc.todoRepo.Save(ctx, t)
	if err != nil {
		return nil, err
	}

	return &UpdateTodoUseCaseOutputDTO{
		ID:          t.ID(),
		Title:       t.Title(),
		Description: t.Description(),
	}, nil
}
