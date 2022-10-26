import axios from 'axios'

axios.defaults.withCredentials = true
axios.defaults.headers.post['X-Requested-With'] = 'XMLHttpRequest'
axios.defaults.headers.put['X-Requested-With'] = 'XMLHttpRequest'
axios.defaults.headers.delete['X-Requested-With'] = 'XMLHttpRequest'

export default axios.create({
    baseURL: `${process.env.REACT_APP_API_BASE_URL}`
})