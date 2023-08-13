import axios from 'axios'

const instance = axios.create({
  baseURL: 'http://localhost:8080',
  timeout: 10000
})

instance.defaults.headers.post['Content-Type'] =
  'application/json;charset=UTF-8'
instance.defaults.headers.get['Content-Type'] = 'application/json;charset=UTF-8'
instance.defaults.headers.put['Content-Type'] = 'application/json;charset=UTF-8'
instance.defaults.headers.delete['Content-Type'] =
  'application/json;charset=UTF-8'

instance.interceptors.request.use(
  function (config) {
    // Before sending request...
    return config
  },
  function (error) {
    // Do something if request error.
    return Promise.reject(error)
  }
)

instance.interceptors.response.use(
  function (response) {
    // Handled if http code = 2xx.
    // Do something to response data.
    return response
  },
  function (error) {
    // Do something if response error.
    return Promise.reject(error)
  }
)

export default instance
