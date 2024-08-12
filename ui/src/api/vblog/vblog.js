import client from "../client"

export var GET_BLOG_LIST = () => {
    return client.get("/vblog/api/v1/blogs")
}