package model

import (
	"database/sql"
	"math/rand"
	"net/http"
	"time"

	"github.com/oklog/ulid"
)

type TodoModel interface {
	FetchTodos() ([]*Todo, error)
	AddTodo(r *http.Request) (sql.Result, error)
	ChangeTodo(r *http.Request) (sql.Result, error)
	DeleteTodo(r *http.Request) (sql.Result, error)
}

type todoModel struct {
}

type Todo struct {
	Id     string `json:"id"`
	Name   string    `json:"name"`
	Status string    `json:"status"`
}

func CreateTodoModel() TodoModel {
	return &todoModel{}
}

func (tm *todoModel) FetchTodos() ([]*Todo, error) {
	sql := `SELECT id, name, status FROM todos`

	rows, err := Db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*Todo

	for rows.Next() {
		var (
			id, name, status string
		)

		if err := rows.Scan(&id, &name, &status); err != nil {
			return nil, err
		}

		todos = append(todos, &Todo{
			Id: id,
			Name:   name,
			Status: status,
		})
	}

	return todos, nil
}

func (tm *todoModel) AddTodo(r *http.Request) (sql.Result, error) {
	err := r.ParseForm()

	if err != nil {
		return nil, nil
	}

	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	req := Todo{
		Id:     id.String(),
		Name:   r.FormValue("name"),
		Status: r.FormValue("status"),
	}

	sql := `INSERT INTO todos(id, name, status) VALUES(?, ?, ?)`

	result, err := Db.Exec(sql, req.Id, req.Name, req.Status)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (tm *todoModel) ChangeTodo(r *http.Request) (sql.Result, error) {
	err := r.ParseForm()

	if err != nil {
		return nil, nil
	}

	sql := `UPDATE todos SET status = ? WHERE id = ?`

	afterStatus := "作業中"
	if r.FormValue("status") == "作業中" {
		afterStatus = "完了"
	}

	result, err := Db.Exec(sql, afterStatus, r.FormValue("id"))

	if err != nil {
		return result, err
	}

	return result, nil
}

func (tm *todoModel) DeleteTodo(r *http.Request) (sql.Result, error) {
	err := r.ParseForm()

	if err != nil {
		return nil, nil
	}

	sql := `DELETE FROM todos WHERE id = ?`

	result, err := Db.Exec(sql, r.FormValue("id"))

	if err != nil {
		return result, err
	}

	return result, nil
}
