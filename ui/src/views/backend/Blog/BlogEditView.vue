<template>
    <div class="breadEdit">
        <a-breadcrumb>
            <a-breadcrumb-item>博客后端</a-breadcrumb-item>
            <a-breadcrumb-item>文章管理</a-breadcrumb-item>
            <a-breadcrumb-item>文章编辑</a-breadcrumb-item>
        </a-breadcrumb>
    </div>
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
            <a-button type="primary"  @click="handleCreate">新建博客</a-button>
        </span>
    </div>
    <MdEditor v-model="blog.content" />
</template>

<script setup>
import { ref, reactive } from 'vue';
import { MdEditor } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import { CREATE_BLOG } from '../../../api/vblog/vblog';

const blog = reactive({
    title: 'title',
    author: 'author',
    summary: 'summary',
    content: '请输入内容',
})

const text = blog.content;

var handleCreate = async () => {
    try {
        await CREATE_BLOG(blog)
    } finally {
        console.log(blog)
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