package todo

import (
	"github.com/YukiOnishi1129/go-docker-firebase-restapi/application/todo"
	responseError "github.com/YukiOnishi1129/go-docker-firebase-restapi/presentation/error"
	"github.com/YukiOnishi1129/go-docker-firebase-restapi/util/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	fetchTodoUseCase    *todo.FetchTodoUseCase
	findByIdTodoUseCase *todo.FindByIdTodoUseCase
	createTodoUseCase   *todo.CreateTodoUseCase
	updateTodoUseCase   *todo.UpdateTodoUseCase
	deleteTodoUseCase   *todo.DeleteTodoUseCase
}

func NewHandler(
	fetchTodoUseCase *todo.FetchTodoUseCase,
	findByIdTodoUseCase *todo.FindByIdTodoUseCase,
	createTodoUseCase *todo.CreateTodoUseCase,
	updateTodoUseCase *todo.UpdateTodoUseCase,
	deleteTodoUseCase *todo.DeleteTodoUseCase,
) *Handler {
	return &Handler{
		fetchTodoUseCase:    fetchTodoUseCase,
		findByIdTodoUseCase: findByIdTodoUseCase,
		createTodoUseCase:   createTodoUseCase,
		updateTodoUseCase:   updateTodoUseCase,
		deleteTodoUseCase:   deleteTodoUseCase,
	}
}

// GetTodos godoc
// @Summary Todo一覧を取得する
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {array} getTodoResponse
// @Router /v1/todos [get]
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

// GetTodoByID godoc
// @Summary idに紐づくTodoを取得する
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} getTodoResponse
// @Router /v1/todos/{id} [get]
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

// PostTodo godoc
// @Summary Todoを作成する
// @Tags todos
// @Accept json
// @Produce json
// @Param request body PostTodosParams true "作成するTodo"
// @Success 201 {object} postTodoResponse
// @Router /v1/todos [post]
func (h Handler) PostTodo(ctx echo.Context) error {
	var params PostTodosParams
	if err := ctx.Bind(&params); err != nil {
		err := ctx.JSON(http.StatusBadRequest, responseError.NewErrorResponse(http.StatusBadRequest, err))
		if err != nil {
			return err
		}
		return err
	}

	validate := validator.GetValidator()
	if err := validate.Struct(params); err != nil {
		err := ctx.JSON(http.StatusBadRequest, responseError.NewErrorResponse(http.StatusBadRequest, err))
		if err != nil {
			return err
		}
		return err
	}

	input := todo.CreateTodoUseCaseInputDTO{
		Title:       params.Title,
		Description: params.Description,
	}

	dto, err := h.createTodoUseCase.Run(ctx.Request().Context(), input)
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

// PutTodoByID godoc
// @Summary Todoを更新する
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param request body PutTodosParams true "更新するTodo"
// @Success 201 {object} postTodoResponse
// @Router /v1/todos/{id} [put]
func (h Handler) PutTodoByID(ctx echo.Context) error {
	var params PutTodosParams
	id := ctx.Param("id")
	if err := ctx.Bind(&params); err != nil {
		err := ctx.JSON(http.StatusBadRequest, responseError.NewErrorResponse(http.StatusBadRequest, err))
		if err != nil {
			return err
		}
		return err
	}

	validate := validator.GetValidator()
	if err := validate.Struct(params); err != nil {
		err := ctx.JSON(http.StatusBadRequest, responseError.NewErrorResponse(http.StatusBadRequest, err))
		if err != nil {
			return err
		}
		return err
	}

	input := todo.UpdateTodoUseCaseInputDTO{
		ID:          id,
		Title:       params.Title,
		Description: params.Description,
	}

	dto, err := h.updateTodoUseCase.Run(ctx.Request().Context(), input)
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

// DeleteTodoByID godoc
// @Summary Todoを削除する
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 204
// @Router /v1/todos/{id} [delete]
func (h Handler) DeleteTodoByID(ctx echo.Context) error {
	id := ctx.Param("id")
	err := h.deleteTodoUseCase.Run(ctx.Request().Context(), id)
	if err != nil {
		err := ctx.JSON(http.StatusInternalServerError, responseError.NewErrorResponse(http.StatusInternalServerError, err))
		if err != nil {
			return err
		}
		return err
	}
	err = ctx.JSON(http.StatusNoContent, nil)
	if err != nil {
		return err
	}
	return nil
}
