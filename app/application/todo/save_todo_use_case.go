package todo

import (
	"context"
	todoDomain "github.com/YukiOnishi1129/go-docker-firebase-restapi/domain/todo"
)

type SaveTodoUseCase struct {
	todoRepo todoDomain.TodoRepository
}

type SaveTodoUseCaseInputDTO struct {
	Title       string
	Description string
}

type SaveTodoUseCaseOutputDTO struct {
	ID          string
	Title       string
	Description string
}

func (uc *SaveTodoUseCase) Run(ctx context.Context, input SaveTodoUseCaseInputDTO) (*SaveTodoUseCaseOutputDTO, error) {
	t, err := todoDomain.NewTodo(input.Title, input.Description)
	if err != nil {
		return nil, err
	}
	err = uc.todoRepo.Save(ctx, t)
	if err != nil {
		return nil, err
	}
	return &SaveTodoUseCaseOutputDTO{}, nil
}
