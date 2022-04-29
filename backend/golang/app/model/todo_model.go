package model

type TodoModel interface {
	FetchTodos() ([]*Todo, error)
	AddTodo()
	ChangeTodo()
	DeleteTodo()
}

type todoModel struct {
}

type Todo struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

func CreateTodoModel() TodoModel {
	return &todoModel{}
}

func (tm *todoModel) FetchTodos() ([]*Todo, error) {
	sql := `SELECT name, status FROM todos`

	rows, err := Db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*Todo

	for rows.Next() {
		var (
			name, status string
		)

		if err := rows.Scan(&name, &status); err != nil {
			return nil, err
		}

		todos = append(todos, &Todo{
			Name:   name,
			Status: status,
		})
	}

	return todos, nil
}

func (tm *todoModel) AddTodo() {

}
func (tm *todoModel) ChangeTodo() {

}
func (tm *todoModel) DeleteTodo() {

}
