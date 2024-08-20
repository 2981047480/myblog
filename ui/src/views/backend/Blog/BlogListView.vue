<template>
    <div>
        <div>
            <div class="bread">
                <a-breadcrumb>
                    <a-breadcrumb-item>博客后端</a-breadcrumb-item>
                    <a-breadcrumb-item>文章管理</a-breadcrumb-item>
                    <a-breadcrumb-item>文章列表</a-breadcrumb-item>
                </a-breadcrumb>
            </div>
            <div class="search">
                <a-input-search @search="handleSearch" :style="{width:'420px'}" placeholder="请输入要搜索的文章的关键字" button-text="Search" search-button/>       
            </div>
        </div>
        <a-table :columns="columns" :data="data.Item" :loading="isLoading" show="true" :pagination="pagination" 
        align="center" @page-change="changePage" @page-size-change="changePageSize">
            <template #options="{ record }">
                <a-button size="small" type="text" @click="$router.push({ name: 'BackendBlogEdit' })">编辑</a-button>
                <a-popconfirm content="确定要删除这篇博客？"  @ok="handleDelete(record)"><a-button size="small" type="text">删除</a-button></a-popconfirm>
            </template>
        </a-table>
    </div>
    <div>
    </div>
</template>

<script setup>
    import { ref, reactive, onMounted } from 'vue'
    import { GET_BLOG_LIST, DELETE_BLOG } from '../../../api/vblog/vblog';
    const columns = [
        {
            title: "id",
            dataIndex: "Id",
        },
        {
            title: "title",
            dataIndex: "title",
        },
        {
            title: "author",
            dataIndex: "author",
        },
        {
            title: "summary",
            dataIndex: "summary",
        },
        {
            title: "tags",
            dataIndex: "tags",
        },
        {
            title: "操作",
            slotName: "options",
        },
    ];
    
    var isLoading = ref(false)
    var data = ref({ Total: 0, Item: [] })
    
    const params = reactive({
        keywords: ""
    })

    const pagination = reactive({
        pageSize: 10,
        showTotal: true,
        total: data.Total,
        current: 1,
        showPageSize: true,
    })

    const Load = async () => {
        try {
            isLoading.value = true
            console.log(pagination)
            data = await GET_BLOG_LIST({
                page_size: pagination.pageSize,
                page_number: pagination.current,
                keywords: params.keywords,
            })
            pagination.total = data.Total
            console.log(data.Item,data.Total)
        }
        finally {
            isLoading.value = false
        }
    }
    
    onMounted(() => {
        Load()
    })

    const changePage = async (v) => {
        pagination.current = v
        Load()
    }

    const changePageSize = async (v) => {
        pagination.pageSize = v
        Load()
    }

    const handleSearch = async (v) => {
        params.keywords = v
        Load()
    }

    const handleDelete = async (v) => {
        try {
            console.log(v)
            DELETE_BLOG(v.Id)
        } finally {
            Load()
        }
    }
</script>

<style>
    .search {
        display: flex;
        padding-bottom: 20px;
        justify-content: flex-end;
    }

    .bread {
        float: left;
        display: flex;
    }
</style>