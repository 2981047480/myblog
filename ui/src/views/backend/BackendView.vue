<template>
    <div class="warpper">
        <a-layout>
            <div class="sider">
                <h1 class="sider-header">博客后端</h1>
                <a-layout-sider :style="{width: '300px', height: 'auto'}">
                    <a-menu @menu-item-click="routerHandler"
                    :style="{width: '100%', height: '100%'}"
                    :default-open-keys="['BlogManagement']"
                    :defalut-selected-keys="['BackendBlogList']"
                    breakpoint="xl">
                        <a-sub-menu key="BlogManagement">
                            <template #title>
                                文章管理
                            </template>
                            <a-menu-item key="BackendBlogList">
                                文章列表
                            </a-menu-item>
                            <a-menu-item key="BackendTagList">
                                标签列表
                            </a-menu-item>
                        </a-sub-menu>
                        <a-sub-menu key="CommentManagement">
                            <template #title>
                                评论管理
                            </template>
                            <a-menu-item>
                                评论列表
                            </a-menu-item>
                        </a-sub-menu>
                        <a-sub-menu>
                            <template #title>
                                用户管理
                            </template>
                            <a-menu-item>
                                用户列表
                            </a-menu-item>
                        </a-sub-menu>
                        <a-sub-menu>
                            <template #title>
                                权限管理
                            </template>
                            <a-menu-item>
                                权限管理
                            </a-menu-item>
                        </a-sub-menu>
                    </a-menu>
                    <template #trigger="{ collapsed }">
                        <IconCaretRight v-if="collapsed"></IconCaretRight>
                        <IconCaretLeft v-else></IconCaretLeft>
                      </template>
                </a-layout-sider>
            </div>
            <a-layout>
                <a-layout-header style="height:60px;">
                  <a-page-header title="Header">
                    <template #extra>
                        quit
                    </template>
                  </a-page-header>
                </a-layout-header >
                <a-layout-content class="page">
                    <router-view></router-view>
                </a-layout-content>
            </a-layout>
        </a-layout>
    </div>
</template>

<script setup>
import { useRouter } from 'vue-router';

const router = useRouter()

var routerHandler = (v) => {
    return router.push({name: v})
}

</script>

<style>
    .arco-layout-sider{
        box-shadow: none;
    }
    
    .page {
        background-color: white;
        margin-left: 450px;
        padding: 50px;
        margin-top: 50px;
        margin-right: 50px;
        border-radius: 50px;
    }

    .sider {
        width: auto;
        height: auto;
        border: 1px solid var(--color-border);
        border-radius: 50px;
        position: fixed;
        background-color: white;
        padding-bottom: 50px;
        top: 100px;
        left: 100px;
    }


    .sider-header {
        height: 60px;
        margin-top: 50px;
        margin-left: 90px;
    }

    .warpper {
        width: 100%;
        height: 100%;
        background-color: var(--color-fill-2);
    }

    .arco-layout-has-sider {
        flex-direction: row;
        background-color: var(--color-fill-2);
    }
    
    .arco-menu-light .arco-menu-item, .arco-menu-light .arco-menu-group-title, .arco-menu-light .arco-menu-pop-header, .arco-menu-light .arco-menu-inline-header {
        font-size: medium;
        box-shadow: none; 
    }

    .arco-page-header {
        padding: 16px 0;
        padding-left: 430px;
    }

    .header {
        height: 100px;
        border-bottom: 1px solid var(--color-border);
    }
</style>