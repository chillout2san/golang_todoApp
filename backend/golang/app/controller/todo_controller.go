package controller

import (
	"crud/controller/dto"
	"crud/model"
	"crud/model/entities"
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
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	var response []dto.FetchTodoResponse

	for _, t := range todos {
		response = append(response, dto.FetchTodoResponse{
			Id:     t.Id,
			Name:   t.Name,
			Status: t.Status,
		})
	}

	json, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}

func (tc *todoController) AddTodo(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var addTodoRequest dto.AddTodoRequest
	err := json.Unmarshal(body, &addTodoRequest)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	result, err := tc.tm.AddTodo(entities.Todo{
		Name:   addTodoRequest.Name,
		Status: addTodoRequest.Status,
	})

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	json, err := json.Marshal(result)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}

func (tc *todoController) ChangeTodo(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var changeTodoRequest dto.ChangeTodoRequest
	err := json.Unmarshal(body, &changeTodoRequest)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	result, err := tc.tm.ChangeTodo(entities.Todo{
		Id:     changeTodoRequest.Id,
		Status: changeTodoRequest.Status,
	})

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	json, err := json.Marshal(result)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}

func (tc *todoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var deleteTodoRequest dto.DeleteTodoRequest
	err := json.Unmarshal(body, &deleteTodoRequest)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	result, err := tc.tm.DeleteTodo(entities.Todo{
		Id: deleteTodoRequest.Id,
	})

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	json, err := json.Marshal(result)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}
