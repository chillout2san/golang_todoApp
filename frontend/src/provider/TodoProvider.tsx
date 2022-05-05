import {
  createContext,
  ReactNode,
  useState,
  Dispatch,
  SetStateAction,
} from 'react'
import { TodoType } from '../types'

interface PropType {
  children: ReactNode
}

interface ProviderType {
  todos: TodoType[]
  setTodos: Dispatch<SetStateAction<TodoType[]>>
}

export const TodoContext = createContext<ProviderType>({} as ProviderType)

export const TodoProvider = (props: PropType) => {
  const { children } = props

  const [todos, setTodos] = useState<TodoType[]>([])

  return (
    <TodoContext.Provider value={{ todos, setTodos }}>
      {children}
    </TodoContext.Provider>
  )
}
