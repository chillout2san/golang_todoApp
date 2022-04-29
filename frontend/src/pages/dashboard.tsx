import { InputForm } from '../components/InputForm'
import { TodoList } from '../components/TodoList'

export const Dashboard = () => {
  return (
    <div className="w-128 mx-auto my-20">
      <InputForm />
      <TodoList />
    </div>
  )
}
