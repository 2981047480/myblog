<template>
    <div class="warpper">
        <a-layout>
            <div class="sider">
                <a-layout-sider :style="{width: '300px', height: 'auto'}">
                    <div style="text-align: center;">
                        <a-avatar :size="150">
                            <img class="image_touxiang" alt="avatar" src="../../../images/touxiang.jpg"/>
                          </a-avatar>
                    </div>
                    <div class="MyName">
                        <a-typography>
                            <a-typography-title :heading="2">
                              zephyr
                            </a-typography-title>
                        </a-typography>
                    </div>
                    <div class="MyComment">
                        <a-typography>
                            <a-typography-paragraph type="secondary">
                                我欲同君游，北国至南洲。
                            </a-typography-paragraph>
                            <a-typography-paragraph type="secondary">
                                遍览九分雪，一分共白头。
                            </a-typography-paragraph>
                        </a-typography>
                    </div>
                </a-layout-sider>
            </div>
            <a-layout>
                <a-layout-header style="height:60px;">
                  <a-page-header title="" :show-back="false">
                    <template #extra>
                        <a-button type="primary" class="btnFrontend" @click="handleBackend">后端</a-button>
                        <a-button type="primary" class="btnQuit" @click="handleQuit">退出</a-button>
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
import app from '../../stores/app';
import { useRouter } from 'vue-router';
import { ref, onMounted } from 'vue';
import { LOGOUT } from '../../api/login/logout';

const router = useRouter()
var data = ref({ Total: 0, Item: [] })

var routerHandler = (v) => {
    return router.push({name: v})
}

var handleBackend = () => {
    return router.push({name: "backend"})
}

var handleQuit = () => {
    try {
        console.log(app.value.token.refresh_token)
        LOGOUT(app.value.token.refresh_token)
        router.push({name: "login"})
    } finally {
    }
}

onMounted(()=>{
    routerHandler('FrontendBlogList')
}) 

</script>

<style>
    .arco-layout-sider{
        box-shadow: none;
    }

    .btnFrontend {
        margin-right: 10px; 
    }

    .btnQuit {
        margin-right: 50px;
    } 

    .MyName {
        text-align: center;
    }

    div.arco-typography, p.arco-typography {
        padding-top: 0;
        margin: auto;
        text-align: center;
        /* align-content: center; */
        margin-bottom: 1em;
    }

    .MyComment {
        padding-top: 20px;
        padding-left: 10px;
        padding-right: 10px;
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
        padding-top: 50px;
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