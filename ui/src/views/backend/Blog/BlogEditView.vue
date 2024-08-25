<template>
    <a-page-header title="返回" @back="$router.go(-1)">
        <template #breadcrumb>
            <a-breadcrumb>
                <a-breadcrumb-item>博客后端</a-breadcrumb-item>
                <a-breadcrumb-item>文章管理</a-breadcrumb-item>
                <a-breadcrumb-item>文章编辑</a-breadcrumb-item>
            </a-breadcrumb>
        </template>
    </a-page-header>
    <a-spin class="MyLoading" :loading="isLoading" tip="加载中，请稍候...">
        <div class="input">
            <span class="title">
                标题：
                <a-input v-model="blog.title" :style="{height: '50px',width:'320px'}" placeholder="请输入标题" allow-clear />
            </span>
            <span class="author">
                作者：
                <a-input v-model="blog.author" :style="{height: '50px',width:'120px'}" placeholder="请输入作者" allow-clear />
            </span>
            <span class="summary">
                摘要： 
                <a-input v-model="blog.summary" :style="{height: '50px',width:'320px'}" placeholder="请输入摘要" allow-clear />
            </span>
            <span class="buttonNew">
                <a-button type="primary"  @click="handleCreate">保存博客</a-button>
            </span>
        </div>
    </a-spin>
    <MdEditor v-model="blog.content" />
</template>

<script setup>
import { ref, reactive, watch } from 'vue';
import { MdEditor } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import { useRouter } from 'vue-router';
import { CREATE_BLOG, DESCRIBE_BLOG, UPDATE_BLOG } from '../../../api/vblog/vblog';

const router = useRouter()

var blog = reactive({
    title: 'title',
    author: 'author',
    summary: 'summary',
    content: '请输入内容',
})

const text = blog.content;
var isEdit = ref(false)

var isLoading = ref(false)
watch(
    () => router.currentRoute.value.query,
    async (v) => {
        if (v.id) {
            isEdit.value = true
            try {
                isLoading.value = true
                var resp = await DESCRIBE_BLOG(v.id)
                console.log(resp)
                blog = resp
                console.log(blog)
            } finally {
                isLoading.value = false
            }
        }
    },
    { immediate: true }
)

var handleCreate = async () => {
    if (isEdit.value === true) {
        try {
            await UPDATE_BLOG(blog)
            router.go(-1)
        } finally {
            console.log(blog)
        } 
    } else {
        try {
            await CREATE_BLOG(blog)
            router.go(-1)
        } finally {
            console.log(blog)
        }
    }
}
</script>

<style>
    div {
        color: var(--color-text-2);
    }
    
    .input {
        height: 50px;
        padding: 10px;
    }

    .breadEdit {
        padding-bottom: 10px;
    }

    .title .author .summary .tag {
        height: 50px;
        padding-right: 10px;
    }

    .buttonNew {
        padding-bottom: 10px;
        float: right;
    }
</style>