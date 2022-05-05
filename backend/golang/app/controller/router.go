package controller

import (
	"net/http"
	"os"
)

type Router interface {
	FetchTodos(w http.ResponseWriter, r *http.Request)
}

type router struct {
	tc TodoController
}

func CreateRouter(tc TodoController) Router {
	return &router{tc}
}

func (ro *router) FetchTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ORIGIN"))
	ro.tc.FetchTodos(w, r)
}
