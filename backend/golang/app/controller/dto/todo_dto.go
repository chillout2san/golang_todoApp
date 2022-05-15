package dto

// リクエストボディがないので、Requestのdto不要
type FetchTodoResponse struct {
    Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// レスポンスはSQLの実行結果なので、Responseのdto不要
type AddTodoRequest struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

// レスポンスはSQLの実行結果なので、Responseのdto不要
type ChangeTodoRequest struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

// レスポンスはSQLの実行結果なので、Responseのdto不要
type DeleteTodoRequest struct {
	Id string `json:"id"`
}
