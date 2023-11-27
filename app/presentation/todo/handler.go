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

func (h *handler) PostTodos(ctx echo.Context) {
	var params PostTodosParams
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(400, err)
		return
	}

	//	TODO: バリデーション

	input := todo.SaveTodoUseCaseInputDTO{
		Title:       params.Title,
		Description: params.Description,
	}

	dto, err := h.saveTodoUseCase.Run(ctx.Request().Context(), input)
	if err != nil {
		ctx.JSON(500, err)
		return
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
		return
	}
}

// GetTodoById godoc
func (h *handler) GetTodoById(ctx echo.Context, id string) {
	//	TODO: バリデーション

	dto, err := h.findByIdTodoUseCase.Run(ctx.Request().Context(), id)
	if err != nil {
		ctx.JSON(500, err)
		return
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
		return
	}
}

// GetTodos godoc
func (h *handler) GetTodos(ctx echo.Context) {
	dtos, err := h.fetchTodoUseCase.Run(ctx.Request().Context())
	if err != nil {
		ctx.JSON(500, err)
		return
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
		return
	}
}

func (h *handler) PutTodosById(ctx echo.Context, id string) {
	var params PutTodosParams
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(400, err)
		return
	}
}
