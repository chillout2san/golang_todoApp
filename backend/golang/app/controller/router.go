package controller

import (
	"net/http"
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
	
}
