package controller

import (
	"net/http"
	"os"
)

type Router interface {
	FetchTodos(w http.ResponseWriter, r *http.Request)
	AddTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
}

type router struct {
	tc TodoController
}

func CreateRouter(tc TodoController) Router {
	return &router{tc}
}

func (ro *router) FetchTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ORIGIN"))
	w.Header().Set("Content-Type", "application/json")
	ro.tc.FetchTodos(w, r)
}

func (ro *router) AddTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ORIGIN"))
	w.Header().Set("Content-Type", "application/json")

	// preflightでAPIが二度実行されてしまうことを防ぐ。
	if r.Method == "OPTIONS" {
		return
	}

	ro.tc.AddTodo(w, r)
}

func (ro *router) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ORIGIN"))
	w.Header().Set("Content-Type", "application/json")

	// preflightでAPIが二度実行されてしまうことを防ぐ。
	if r.Method == "OPTIONS" {
		return
	}
	
	ro.tc.DeleteTodo(w, r)
}
