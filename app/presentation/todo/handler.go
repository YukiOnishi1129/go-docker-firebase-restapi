package todo

import (
	"github.com/YukiOnishi1129/go-docker-firebase-restapi/application/todo"
	"github.com/labstack/echo/v4"
	"net/http"
)

type handler struct {
	saveTodoUseCase     *todo.SaveTodoUseCase
	findByIdTodoUseCase *todo.FindByIdTodoUseCase
	fetchTodoUseCase    *todo.FetchTodoUseCase
}

func NewHandler(
	saveTodoUseCase *todo.SaveTodoUseCase,
	findByIdTodoUseCase *todo.FindByIdTodoUseCase,
	fetchTodoUseCase *todo.FetchTodoUseCase,
) *handler {
	return &handler{
		saveTodoUseCase:     saveTodoUseCase,
		findByIdTodoUseCase: findByIdTodoUseCase,
		fetchTodoUseCase:    fetchTodoUseCase,
	}
}

func (h handler) PostTodo(ctx echo.Context) error {
	var params PostTodosParams
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(400, err)
		return err
	}

	//	TODO: バリデーション

	input := todo.SaveTodoUseCaseInputDTO{
		Title:       params.Title,
		Description: params.Description,
	}

	dto, err := h.saveTodoUseCase.Run(ctx.Request().Context(), input)
	if err != nil {
		ctx.JSON(500, err)
		return err
	}

	response := postTodoResponse{
		todoResponseModel{
			Id:          dto.ID,
			Title:       dto.Title,
			Description: dto.Description,
		},
	}
	err = ctx.JSON(http.StatusCreated, response)
	if err != nil {
		return err
	}
	return nil
}

// GetTodoByID godoc
func (h *handler) GetTodoByID(ctx echo.Context) error {
	//	TODO: バリデーション

	id := ctx.Param("id")
	dto, err := h.findByIdTodoUseCase.Run(ctx.Request().Context(), id)
	if err != nil {
		ctx.JSON(500, err)
		return err
	}

	response := getTodoResponse{
		&todoResponseModel{
			Id:          dto.ID,
			Title:       dto.Title,
			Description: dto.Description,
		},
	}
	err = ctx.JSON(http.StatusOK, response)
	if err != nil {
		return err
	}

	return nil
}

// GetTodos godoc
func (h handler) GetTodos(ctx echo.Context) error {
	dtos, err := h.fetchTodoUseCase.Run(ctx.Request().Context())
	if err != nil {
		ctx.JSON(500, err)
		return err
	}

	var response []getTodoResponse
	for _, dto := range dtos {
		response = append(response, getTodoResponse{
			&todoResponseModel{
				Id:          dto.ID,
				Title:       dto.Title,
				Description: dto.Description,
			},
		})
	}

	err = ctx.JSON(http.StatusOK, response)
	if err != nil {
		return err
	}
	return nil
}

func (h handler) PutTodoByID(ctx echo.Context) error {
	var params PutTodosParams
	// id := ctx.Param("id")
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(400, err)
		return err
	}
	return nil
}
