package model

import (
	"crud/model/entities"
	"database/sql"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

type TodoModel interface {
	FetchTodos() ([]*entities.Todo, error)
	AddTodo(r entities.Todo) (sql.Result, error)
	ChangeTodo(r entities.Todo) (sql.Result, error)
	DeleteTodo(r entities.Todo) (sql.Result, error)
}

type todoModel struct {
}

func CreateTodoModel() TodoModel {
	return &todoModel{}
}

func (tm *todoModel) FetchTodos() ([]*entities.Todo, error) {
	sql := `SELECT id, name, status FROM todos`

	rows, err := Db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*entities.Todo

	for rows.Next() {
		var (
			id, name, status string
		)

		if err := rows.Scan(&id, &name, &status); err != nil {
			return nil, err
		}

		todos = append(todos, &entities.Todo{
			Id:     id,
			Name:   name,
			Status: status,
		})
	}

	return todos, nil
}

func (tm *todoModel) AddTodo(r entities.Todo) (sql.Result, error) {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	req := entities.Todo{
		Id:     id.String(),
		Name:   r.Name,
		Status: r.Status,
	}

	sql := `INSERT INTO todos(id, name, status) VALUES(?, ?, ?)`

	result, err := Db.Exec(sql, req.Id, req.Name, req.Status)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (tm *todoModel) ChangeTodo(r entities.Todo) (sql.Result, error) {
	sql := `UPDATE todos SET status = ? WHERE id = ?`

	afterStatus := "作業中"
	if r.Status == "作業中" {
		afterStatus = "完了"
	}

	result, err := Db.Exec(sql, afterStatus, r.Id)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (tm *todoModel) DeleteTodo(r entities.Todo) (sql.Result, error) {
	sql := `DELETE FROM todos WHERE id = ?`

	result, err := Db.Exec(sql, r.Id)

	if err != nil {
		return result, err
	}

	return result, nil
}
