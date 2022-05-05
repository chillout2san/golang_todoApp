package controller

import (
	"crud/model"
	"encoding/json"
	"fmt"
	"net/http"
)

type TodoController interface {
	FetchTodos(w http.ResponseWriter, r *http.Request)
	AddTodo(w http.ResponseWriter, r *http.Request)
	ChangeTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
}

type todoController struct {
	tm model.TodoModel
}

func CreateTodoController(tm model.TodoModel) TodoController {
	return &todoController{tm}
}

func (tc *todoController) FetchTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := tc.tm.FetchTodos()

	if err != nil {
		// TODO: ログ出力方法を検討
		fmt.Println(err)
	}

	json, err := json.Marshal(todos)

	if err != nil {
		// TODO: ログ出力方法を検討
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, string(json))
}

func (tc *todoController) AddTodo(w http.ResponseWriter, r *http.Request) {

}

func (tc *todoController) ChangeTodo(w http.ResponseWriter, r *http.Request) {

}

func (tc *todoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {

}
