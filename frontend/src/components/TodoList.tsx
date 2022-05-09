import { useEffect, useContext } from 'react'
import { client } from '../libs/axios'
import { TodoContext } from '../provider/TodoProvider'

const omitText = (text: string): string => {
  if (text.length > 10) {
    return text.substring(0, 7) + '...'
  }
  return text
}

export const TodoList = () => {
  const { todos, setTodos } = useContext(TodoContext)

  useEffect(() => {
    client.get('fetch-todos').then(({ data }) => {
      setTodos(data)
    })
  }, [])

  const changeTodo = async (id: string, status: string) => {
    const body = new URLSearchParams({
      id,status
    })
    await client.post('change-todo', body)
    client.get('fetch-todos').then(({ data }) => {
      setTodos(data)
    })
  }

  const deleteTodo = async (id: string) => {
    const body = new URLSearchParams({
      id,
    })
    await client.post('delete-todo', body)
    client.get('fetch-todos').then(({ data }) => {
      setTodos(data)
    })
  }

  return (
    <div className="p-4 border border-gray-200 rounded shadow-lg">
      <p className="font-bold mb-2">タスク一覧</p>
      <table className="border-collapse table-auto">
        <thead>
          <tr>
            <th className="py-1">番号</th>
            <th className="p-1">タスク名</th>
            <th className="p-1">ステータス</th>
            <th className="p-1"></th>
            <th className="p-1"></th>
          </tr>
        </thead>
        <tbody>
          {todos.map((todo, index) => {
            return (
              <tr key={index}>
                <td className="p-1">{index + 1}</td>
                <td className="p-1">{omitText(todo.name)}</td>
                <td className="p-1">{todo.status}</td>
                <td className="p-1">
                  <button
                    className="px-2 h-7 border border-white rounded bg-teal-400 shadow-md text-white"
                    onClick={() => {
                      changeTodo(todo.id, todo.status)
                    }}
                  >
                    変更する
                  </button>
                </td>
                <td className="p-1">
                  <button
                    className="px-2 h-7 border border-white rounded bg-teal-400 shadow-md text-white"
                    onClick={() => {
                      deleteTodo(todo.id)
                    }}
                  >
                    削除する
                  </button>
                </td>
              </tr>
            )
          })}
        </tbody>
      </table>
    </div>
  )
}
