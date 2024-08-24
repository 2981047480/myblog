<template>
    <div class="breadList">
        <a-breadcrumb>
            <a-breadcrumb-item>博客前端</a-breadcrumb-item>
            <a-breadcrumb-item>文章列表</a-breadcrumb-item>
        </a-breadcrumb>
    </div>
    <div class="content">
        <a-spin :loading="isLoading" tip="加载中，请稍候...">
            <template v-for="item of data.Item">
                <a-list class="myList" :style="{width: '600px'}" :hoverable="true">
                    <a-list-item :id=item.Id>
                        <a-list-item-meta>
                            <template #title>
                            {{ item.title }}
                            </template>
                        <template #description>
                            {{ item.summary }}
                        </template> 
                    </a-list-item-meta>
                    <template #actions>
                        <icon-eye @click="$router.push({ name: 'FrontendBlogEdit', query: {id: item.Id} })" />
                    </template>
                    </a-list-item>
                </a-list>
            </template>
            <a-pagination class="myPagination" :total="data.Total"
            :current="pagination.current" :pageSize="pagination.pageSize" showPageSize
            showTotal @change="changePage" @page-size-change="changePageSize"/>
        </a-spin>
    </div>
</template>

<script setup>
import { ref, onMounted, reactive } from 'vue';
import { GET_BLOG_LIST } from '../../../api/vblog/vblog';

var data = ref({ Total: 0, Item: [] })
var isLoading = ref(true)

var pagination = reactive({
    pageSize: 10,
    current: 1,
})

const Load = async () => {
    try {
        isLoading.value = true
        data = await GET_BLOG_LIST({
            "page_size": pagination.pageSize,
            "page_number": pagination.current
        })
        console.log(data.Item,data.Total)
    }
    finally {
        isLoading.value = false
    }
}

onMounted(()=>{
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
</script>

<style>
.breadList {
    float: left;
    display: flex;
}

.content {
    padding-top: 50px;
}

.arco-spin.MyLoading {
    display: flex;
}

.arco-list-wrapper.myList{
    float: left;
    padding-right: 20px;
    overflow: hidden;
}

.arco-scrollbar-container {
    border-radius: 20px;
    position: relative;
    scrollbar-width: none;
    margin-bottom: 10px;
}

.arco-pagination.myPagination {
    align-self: end;
    clear: both;
    display: flex;
    align-items: center;
    font-size: 16px;
    flex-direction: row;
    justify-content: flex-end;
}
</style>