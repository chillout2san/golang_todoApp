import { Routes, Route } from 'react-router-dom'
import { Dashboard } from '../pages/dashboard'

export const AppRouter = () => {
  return (
    <Routes>
      <Route path="/" element={<Dashboard />}></Route>
    </Routes>
  )
}
