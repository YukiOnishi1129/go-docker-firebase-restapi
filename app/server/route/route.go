package route

import (
	todoApp "github.com/YukiOnishi1129/go-docker-firebase-restapi/application/todo"
	"github.com/YukiOnishi1129/go-docker-firebase-restapi/infrastructure/firestore/repository"
	todoPre "github.com/YukiOnishi1129/go-docker-firebase-restapi/presentation/todo"
	"github.com/labstack/echo/v4"
)

func InitRoute(api *echo.Echo) {
	vi := api.Group("/api/v1")
	{
		todoRoute(vi)
	}
}

func todoRoute(r *echo.Group) {
	todoRepository := repository.NewTodoRepository()
	h := todoPre.NewHandler(
		todoApp.NewSaveTodoUseCase(todoRepository),
		todoApp.NewFindByIdTodoUseCase(todoRepository),
		todoApp.NewFetchTodoUseCase(todoRepository),
	)
	group := r.Group("/todos")
	group.GET("", h.GetTodos)
	group.GET("/:id", h.GetTodoByID)
	group.POST("", h.PostTodo)
	group.PUT("/:id", h.PutTodoByID)
}
