import client from "../client"

export var GET_BLOG_LIST = (params) => {
    // 这里一定要用params 不然会失败。。。 血的教训
    console.log(params)
    return client.get('/vblog/api/v1/blogs', { params })
}
// export var GET_BLOG_LIST = (params) => {
//     return client.get('/vblog/api/v1/blogs', { params })
// }

export var DELETE_BLOG = (id) => {
    return client.delete(`/vblog/api/v1/blogs/${id}`)
}