import axios from 'axios'
import { Message } from '@arco-design/web-vue'

var client = axios.create({
    // 这里vite代理了 所以不写
    baseURL: '',
    timeout: 1000
});

client.interceptors.response.use(
    (value) => {
        return value.data
    },

    (err) => {
        console.log(err)
        var msg = err.message
        var code = 0
        if (err.response && err.response.data) {
            msg = err.response.data.message
            code = err.response.data.code
        }
        Message.error(msg);
        return Promise.reject(err)
    }
)


export default client