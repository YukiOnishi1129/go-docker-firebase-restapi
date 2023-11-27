package todo

type todoResponseModel struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type getTodoResponse struct {
	Todo *todoResponseModel `json:"todos"`
}

type postTodoResponse struct {
	Todo todoResponseModel `json:"todo"`
}

type putTodoResponse struct {
	Todo todoResponseModel `json:"todo"`
}
