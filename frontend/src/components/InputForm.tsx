import { useState, useContext } from 'react'
import { WORK_ON_PROGRESS } from '../constants'
import { client } from '../libs/axios'
import { TodoContext } from '../provider/TodoProvider'

export const InputForm = () => {
  const [todoName, setTodoName] = useState('')
  const { setTodos } = useContext(TodoContext)
  const onChangeTodoName: React.ChangeEventHandler<HTMLInputElement> = (
    event
  ) => {
    setTodoName(event.target.value)
  }

  const addTodo = async () => {
    const body = {
      name: todoName,
      status: WORK_ON_PROGRESS,
    }
    await client.post('todo/add-todo', body)
    client.get('todo/fetch-todos').then(({ data }) => {
      setTodos(data)
    })
  }

  return (
    <div className="w-auto h-30 mb-4 p-4 border border-gray-200 rounded shadow-lg">
      <p className="mb-2 font-bold">新しいタスクを追加する</p>
      <input
        placeholder="買い物"
        className="mr-4 border shadow-md border-teal-500 rounded"
        onChange={onChangeTodoName}
      />
      <button
        onClick={addTodo}
        className="px-2 h-7 border border-white rounded bg-teal-400 shadow-md text-white"
      >
        追加
      </button>
    </div>
  )
}
