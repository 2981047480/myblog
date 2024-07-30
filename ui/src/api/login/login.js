import client from '../client'

export var LOGIN = (data) => {
    return client.post('/vblog/api/v1/tokens', data)
}