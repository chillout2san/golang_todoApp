package model

type TodoModel interface {
	FetchTodos()
	AddTodo()
	ChangeTodo()
	DeleteTodo()
}

type todoModel struct {
}

func CreateTodoModel() TodoModel {
	return &todoModel{}
}

func (tm *todoModel) FetchTodos() {
	
}

func (tm *todoModel) AddTodo() {

}
func (tm *todoModel) ChangeTodo() {

}
func (tm *todoModel) DeleteTodo() {

}
