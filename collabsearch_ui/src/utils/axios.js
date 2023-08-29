import axios from 'axios'
import { useAccountStore } from '@/stores/account'

const instance = axios.create({
  baseURL: 'http://localhost:8080',
  timeout: 10000
})

instance.defaults.headers.post['Content-Type'] =
  'application/json;charset=UTF-8'
instance.defaults.headers.get['Content-Type'] = 'charset=UTF-8'
instance.defaults.headers.put['Content-Type'] = 'application/json;charset=UTF-8'
instance.defaults.headers.delete['Content-Type'] =
  'application/json;charset=UTF-8'

instance.interceptors.request.use(
  function (config) {
    // Before sending request...
    const store = useAccountStore()
    if (store.token) {
      config.headers.Authorization = 'Bearer ' + store.token
    }
    return config
  },
  function (error) {
    // Do something if request error.
    return Promise.reject(error)
  }
)

instance.interceptors.response.use(
  function (response) {
    // Handled if http code = 2xx. Extract token from header.
    const token = response.headers.authorization
    if (token) {
      const extractedToken = token.split(' ')[1]
      if (extractedToken) {
        const store = useAccountStore()
        store.setToken(extractedToken)
      }
    }
    // return response.data instead of response.
    return response.data
  },
  function (error) {
    // Do something if response error.
    // Only 3 types of status code.
    // http.BadRequest: 400 -- JSON body validation failed.
    // http.Unauthorized: 401 -- Token required.
    // http.UnprocessedEntity: 422 -- Failed to process data for some reason.
    var errMsg = error.response.data.Msg
    if (error.response.status === 400) {
      return Promise.reject({ Code: 400, Msg: errMsg })
    } else if (error.response.status === 401) {
      return Promise.reject({ Code: 401, Msg: errMsg })
    } else if (error.response.status === 422) {
      return Promise.reject({ Code: 422, Msg: errMsg })
    }
    return Promise.reject(error)
  }
)

const outside = axios.create({
  timeout: 10000
})

outside.defaults.headers.get['Content-Type'] = 'application/json;charset=UTF-8'

outside.interceptors.request.use(
  function (config) {
    // Before sending request...
    return config
  },
  function (error) {
    // Do something if request error.
    return Promise.reject(error)
  }
)

outside.interceptors.response.use(
  function (response) {
    return response
  },
  function (error) {
    return Promise.reject(error)
  }
)

export default instance
export { outside }
