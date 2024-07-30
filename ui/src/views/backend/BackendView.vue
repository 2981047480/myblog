<template>
    <div class="warpper">
        <a-layout>
            <div class="sider">
                <div class="sider-header"></div>
                <a-layout-sider collapsible :collapsed="collapsed"  :style="{width: '300px', height: 'calc(100vh - 30px)'}">
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
                <a-layout-header style="height:60px; ">
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
    .sider {
        width: auto;
        height: 100vh;
        border-right: 1px solid var(--color-border);
    }

    .sider-header {
        height: 40px;
    }

    .warpper {
        width: 100%;
        height: 100%;
    }

    .arco-menu-light .arco-menu-item.arco-menu-selected, .arco-menu-light .arco-menu-group-title.arco-menu-selected, .arco-menu-light .arco-menu-pop-header.arco-menu-selected, .arco-menu-light .arco-menu-inline-header.arco-menu-selected {
        color: var(--color-text-2);
    }
    
    .arco-menu-light .arco-menu-inline-header.arco-menu-selected .arco-icon, .arco-menu-light .arco-menu-inline-header.arco-menu-selected .arco-menu-icon {
        color: var(--color-text-2);
    }
    
    .arco-menu-light .arco-menu-item, .arco-menu-light .arco-menu-group-title, .arco-menu-light .arco-menu-pop-header, .arco-menu-light .arco-menu-inline-header {
        font-size: medium;
        color: var(--color-text-2);
        background-color: var(--color-menu-light-bg);
    }

    .arco-page-header {
        padding: 16px 0;
        border-bottom: 1px solid var(--color-border);
    }

    .header {
        height: 100px;
        border-bottom: 1px solid var(--color-border);
    }
</style>