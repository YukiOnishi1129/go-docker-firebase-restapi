package todo

import (
	"context"
	todoDomain "github.com/YukiOnishi1129/go-docker-firebase-restapi/domain/todo"
)

type FetchTodoUseCase struct {
	todoRepo todoDomain.TodoRepository
}

func NewFetchTodoUseCase(todoRepo todoDomain.TodoRepository) *FetchTodoUseCase {
	return &FetchTodoUseCase{
		todoRepo: todoRepo,
	}
}

type FetchTodoUseCaseDto struct {
	ID          string
	Title       string
	Description string
}

func (uc *FetchTodoUseCase) Run(ctx context.Context) ([]*FetchTodoUseCaseDto, error) {
	todos, err := uc.todoRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var output []*FetchTodoUseCaseDto

	for _, t := range todos {
		output = append(output, &FetchTodoUseCaseDto{
			ID:          t.ID(),
			Title:       t.Title(),
			Description: t.Description(),
		})
	}
	return output, nil
}
