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
		fmt.Fprint(w, err)
		return
	}

	json, err := json.Marshal(todos)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}

func (tc *todoController) AddTodo(w http.ResponseWriter, r *http.Request) {
	result, err := tc.tm.AddTodo(r)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	json, err := json.Marshal(result)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}

func (tc *todoController) ChangeTodo(w http.ResponseWriter, r *http.Request) {
	result, err := tc.tm.ChangeTodo(r)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	json, err := json.Marshal(result)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}

func (tc *todoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	result, err := tc.tm.DeleteTodo(r)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	json, err := json.Marshal(result)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}
