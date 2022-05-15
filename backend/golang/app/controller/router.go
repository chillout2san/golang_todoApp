package controller

import (
	"net/http"
	"os"
)

type Router interface {
	HandleRequest()
}

type router struct {
	tc TodoController
}

func CreateRouter(tc TodoController) Router {
	return &router{tc}
}

func (ro *router) HandleRequest() {
	http.HandleFunc("/todo/", ro.HandleTodoRequest)
}

func (ro *router) HandleTodoRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ORIGIN"))
	w.Header().Set("Content-Type", "application/json")

	// preflightでAPIが二度実行されてしまうことを防ぐ。
	if r.Method == "OPTIONS" {
		return
	}

	prefix := "/todo/"

	switch r.URL.Path {
	case prefix + "fetch-todos":
		ro.tc.FetchTodos(w, r)
	case prefix + "add-todo":
		ro.tc.AddTodo(w, r)
	case prefix + "delete-todo":
		ro.tc.DeleteTodo(w, r)
	case prefix + "change-todo":
		ro.tc.ChangeTodo(w, r)
	default:
		w.WriteHeader(405)
	}
}
