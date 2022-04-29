package main

import (
	"crud/model"
	"crud/controller"
	"net/http"
)

var tm = model.CreateTodoModel()
var tc = controller.CreateTodoController(tm)
var ro = controller.CreateRouter(tc)

func main() {
	http.HandleFunc("/fetch-todos", ro.FetchTodos)
	http.ListenAndServe(":8080", nil)
}
