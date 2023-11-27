package todo

type PostTodosParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type PutTodosParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
