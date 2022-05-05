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
          {todos.map((datum, index) => {
            return (
              <tr key={datum.name}>
                <td className="p-1">{index + 1}</td>
                <td className="p-1">{omitText(datum.name)}</td>
                <td className="p-1">{datum.status}</td>
                <td className="p-1">
                  <button className="px-2 h-7 border border-white rounded bg-teal-400 shadow-md text-white">
                    変更する
                  </button>
                </td>
                <td className="p-1">
                  <button className="px-2 h-7 border border-white rounded bg-teal-400 shadow-md text-white">
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
