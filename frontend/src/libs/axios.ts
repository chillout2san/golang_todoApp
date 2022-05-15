import axios from 'axios'

axios.defaults.headers.common['Content-Type'] = 'application/json'

export const client = axios.create({
  baseURL: 'http://localhost:8080/',
})
