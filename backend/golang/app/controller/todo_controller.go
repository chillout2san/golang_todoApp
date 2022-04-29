package controller

import (
	"net/http"
	"crud/model"
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

}

func (tc *todoController) AddTodo(w http.ResponseWriter, r *http.Request) {

}

func (tc *todoController) ChangeTodo(w http.ResponseWriter, r *http.Request) {

}

func (tc *todoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {

}
