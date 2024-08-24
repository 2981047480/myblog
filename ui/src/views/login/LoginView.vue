<template>
    <div class="login-component-wapper">
        <a-form class="login-component" :model="form" @submit="ClickHandler">
            <div class="title">登录</div>
            <div class="login-form">
                <a-form-item field="username" label="用户名">
                    <a-input v-model="form.username" placeholder="请输入用户名">
                        <template #prefix>
                            <icon-user />
                        </template>
                    </a-input>
                </a-form-item>
                <a-form-item field="password" label="密码">
                    <a-input-password v-model="form.password" placeholder="请输入密码">
                        <template #prefix>
                            <icon-lock />
                        </template>
                    </a-input-password>
                </a-form-item>
                <a-form-item field="is_remember">
                    <a-checkbox v-model="form.is_remember" style="color: white;">
                        记住我
                    </a-checkbox>
                </a-form-item>
            </div>
            <div class="submit-buttom">
                <a-button type="primary" html-type="submit" :loading="isLoading" auto-label-width>
                    submit
                </a-button>
            </div>
        </a-form>
    </div>
</template>

<script setup>
    import { ref } from "vue";
    import { reactive } from "vue";
    import { LOGIN } from '../../api/login/login'
    import { useRouter } from 'vue-router'
    import app from '../../stores/app'
    document.body.setAttribute('arco-theme', 'dark')

    const isLoading=ref(false)

    const form = reactive({
        username: '',
        password: '',
        is_remember: false,
    });

    const router = useRouter()

    const ClickHandler = async (data) => {
        if (!data.errors) {
            console.log(data.values);
            try {
                isLoading.value = true
                const resp = await LOGIN(data.values)
                
                app.value.token = resp
                router.push({name: "backend"})
                document.body.removeAttribute('arco-theme');
                console.log(resp)
            }
            finally {
                isLoading.value = false
            }
            
        }
        
    };
</script>

<style scoped>
    .title{
        font-size: 30px;
        color: #d0d0d0;
        padding-top: 15px;
        margin: auto;
        margin-bottom: 25px;
    }

    .login-component-wapper{
        width: 100%;
        height: 100%;
        display: flex;
        background-image: url('../../../images/background.jpg');
        background-size: 100% 100%;
    }
    
    .login-component {
        backdrop-filter: blur(5px);
        height: 300px;
        width: 500px;
        margin: auto;
        border-radius: 10px;
        border: solid 1px #d0d0d0;
    }
   
    .arco-btn-primary, .arco-btn-primary[type="button"], .arco-btn-primary[type="submit"] {
        color: #fff;
        border-radius: 10px;
        background-color: #ffa163;
        backdrop-filter: blur(5px);
        height: 40px;
        width: 350px;
        border: 0px;
        padding: 0px;
        margin: 0px;
    }

    .submit-buttom {
        width: 350px;
        margin: auto;
    }

    .login-form {
        width: 350px;
        height: 130px;
        margin: auto;
        margin-bottom: 0px;
        margin-top: 0px;
    }

    .arco-input-wrapper {
        border-radius: 5px;
        border: solid 1px #d0d0d0;
    }
</style>