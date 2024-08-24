import client from '../client'

export var LOGOUT = (rt) => {
    console.log(rt)
    return client.delete('/vblog/api/v1/tokens', {
        headers: {
            "X-REFRESH-TOKEN": rt
        }
    })
}