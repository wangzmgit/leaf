<template>
    <div class="header-bar">
        <div class="header-content">
            <div class="header-left">
            </div>
            <div class="header-right">
                <!-- 用户头像 -->
                <div v-if="userInfo" class="avatar-box">
                    <common-avatar :url="userInfo.avatar" :size="40"></common-avatar>
                    <div class="header-menu">
                        <div class="menu-item">
                            <span class="btn">{{ userInfo.name }}</span>
                        </div>
                        <div class="menu-item">
                            <span class="btn" @click="logout">退出登录</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useRouter } from 'vue-router';
import { storageData } from '@leaf/utils';
import { CommonAvatar } from '@leaf/components';

const router = useRouter();

const logout = () => {
    storageData.remove("token");
    storageData.remove('user_info');
    router.push({ name: 'Login' });
}

const userInfo = computed(() => {
    return storageData.get("user_info");
})
</script>

<style lang="less" scoped>
.header-bar {
    display: flex;
    align-items: center;
    width: 100%;
    height: 60px;
    min-width: 1000px;
    background-color: #fff;
    box-shadow: 0 1px 4px rgba(0, 21, 41, .3);

    .header-content {
        margin: 0 auto;
        width: 1200px;
        height: 50px;
        display: flex;
        justify-content: space-between;

        .header-left {
            display: flex;
            align-items: center;
            padding-top: 5px;
        }

        .header-right {
            display: flex;
            align-items: center;

            .avatar-box {
                padding-top: 5px;
                margin-right: 6px;
            }

            .upload-btn {
                margin-left: 6px;
            }
        }

    }
}

.avatar-box {
    position: relative;
    cursor: pointer;
    margin: 0 10px;

    &:hover {
        .header-menu {
            display: block;
        }
    }

    .header-menu {
        display: none;
        width: 180px;
        height: 100px;
        top: 40px;
        left: -80px;
        position: absolute;
        z-index: 999;
        background-color: #fff;
        border-radius: 10px;
        animation: menu .3s ease-in;
        box-shadow: 0 0 30px rgba(0, 0, 0, .1);
    }

    .menu-item {
        margin-top: 7px;
        width: 150px;
        height: 36px;
        margin-left: 16px;
        -webkit-user-select: none;
        -moz-user-select: none;
        -ms-user-select: none;
        user-select: none;

        .btn {
            display: block;
            color: #18191b;
            line-height: 36px;
            text-align: left;
            border-radius: 6px;
            padding-left: 6px;

            &:hover {
                background-color: #c9ccd0;
            }
        }

    }
}

@keyframes menu {
    0% {
        opacity: 0;
        transform: translateY(-10px);
    }

    100% {
        opacity: 1;
    }
}
</style>