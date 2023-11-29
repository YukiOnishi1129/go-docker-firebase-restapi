package todo

type PostTodosParams struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type PutTodosParams struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}
