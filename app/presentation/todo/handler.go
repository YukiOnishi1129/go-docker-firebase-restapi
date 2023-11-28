package todo

import (
	"github.com/YukiOnishi1129/go-docker-firebase-restapi/application/todo"
	responseError "github.com/YukiOnishi1129/go-docker-firebase-restapi/presentation/error"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	saveTodoUseCase     *todo.SaveTodoUseCase
	findByIdTodoUseCase *todo.FindByIdTodoUseCase
	fetchTodoUseCase    *todo.FetchTodoUseCase
}

func NewHandler(
	saveTodoUseCase *todo.SaveTodoUseCase,
	findByIdTodoUseCase *todo.FindByIdTodoUseCase,
	fetchTodoUseCase *todo.FetchTodoUseCase,
) *Handler {
	return &Handler{
		saveTodoUseCase:     saveTodoUseCase,
		findByIdTodoUseCase: findByIdTodoUseCase,
		fetchTodoUseCase:    fetchTodoUseCase,
	}
}

func (h Handler) PostTodo(ctx echo.Context) error {
	var params PostTodosParams
	if err := ctx.Bind(&params); err != nil {
		err := ctx.JSON(http.StatusBadRequest, responseError.NewErrorResponse(http.StatusBadRequest, err))
		if err != nil {
			return err
		}
		return err
	}

	//	TODO: バリデーション

	input := todo.SaveTodoUseCaseInputDTO{
		Title:       params.Title,
		Description: params.Description,
	}

	dto, err := h.saveTodoUseCase.Run(ctx.Request().Context(), input)
	if err != nil {
		err := ctx.JSON(http.StatusInternalServerError, responseError.NewErrorResponse(http.StatusInternalServerError, err))
		if err != nil {
			return err
		}
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
func (h *Handler) GetTodoByID(ctx echo.Context) error {
	//	TODO: バリデーション

	id := ctx.Param("id")
	dto, err := h.findByIdTodoUseCase.Run(ctx.Request().Context(), id)
	if err != nil {
		err := ctx.JSON(http.StatusInternalServerError, responseError.NewErrorResponse(http.StatusInternalServerError, err))
		if err != nil {
			return err
		}
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
func (h Handler) GetTodos(ctx echo.Context) error {
	dtoList, err := h.fetchTodoUseCase.Run(ctx.Request().Context())
	if err != nil {
		err := ctx.JSON(http.StatusInternalServerError, responseError.NewErrorResponse(http.StatusInternalServerError, err))
		if err != nil {
			return err
		}
		return err
	}

	var response []getTodoResponse
	for _, dto := range dtoList {
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

func (h Handler) PutTodoByID(ctx echo.Context) error {
	var params PutTodosParams
	// id := ctx.Param("id")
	if err := ctx.Bind(&params); err != nil {
		err := ctx.JSON(http.StatusBadRequest, responseError.NewErrorResponse(http.StatusBadRequest, err))
		if err != nil {
			return err
		}
		return err
	}
	return nil
}
