<template>
    <div class="login" :style="initTheme()" v-title :data-title="`后台登录-${globalConfig.title}`">
        <div class="login-left">
            <img class="illustration" src="@/assets/login-illustration.png" />
        </div>
        <div class="login-right">
            <div class="login-form">
                <span class="title">{{ title }}后台管理系统</span>
                <login-form @success="loginSuccess"></login-form>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import LoginForm from './component/LoginForm.vue';
import { globalConfig } from '@leaf/utils';
import { getTheme } from '@/theme';
import { ref } from 'vue';

const title = ref(window.$title || globalConfig.title);

const router = useRouter();

const initTheme = () => {
    const theme = getTheme();

    return {
        "--primary-color": theme.primaryColor,
        "--title-color": theme.primaryHoverColor
    }
}

const loginSuccess = () => {
    router.push({ name: "Dashboard" });
}
</script>

<style lang="less" scoped>
.login {
    display: flex;
    width: 100%;
    height: 100vh;

    .login-left {
        width: 50%;
        height: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
        background-color: var(--primary-color);

        .illustration {
            width: 80%;
        }
    }

    .login-right {
        width: 50%;
        height: 100%;
    }
}

.login-form {
    width: 60%;
    margin: 220px auto 0;

    .title {
        text-align: center;
        display: block;
        width: 100%;
        margin: 20px 0;
        color: var(--title-color);
        font-size: 28px;
    }
}
</style>