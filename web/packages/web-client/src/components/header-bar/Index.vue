<template>
    <div class="header-bar" :style="themeStyle()">
        <div class="header-content">
            <div class="header-left" @click="goPage('Home')">
                <h1 class="title">{{ globalConfig.title }}</h1>
            </div>
            <div v-show="route.name !== 'Search'" class="header-center">
                <n-input v-model:value="keywords" round placeholder="搜索" @keydown.enter="search">
                    <template #suffix>
                        <n-icon @click="search">
                            <search-icon />
                        </n-icon>
                    </template>
                </n-input>
            </div>
            <div class="header-right">
                <!-- 用户头像 -->
                <div v-if="isLogin" class="avatar-box">
                    <div @click="goPage('Space')">
                        <common-avatar :url="userInfo.avatar" :size="40" :iconSize="22"></common-avatar>
                    </div>
                    <div class="header-menu">
                        <div class="menu-item">
                            <span class="btn" @click="goPage('Space')">{{ userInfo.name }}</span>
                        </div>
                        <div class="menu-item">
                            <span class="btn" @click="logout">{{ t("common.logOut") }}</span>
                        </div>
                    </div>
                </div>
                <div v-else>
                    <n-button class="login-btn" text @click="goPage('Login')">
                        {{ t("common.login") }}/{{ t("common.register") }}
                    </n-button>
                </div>
                <!-- 图形按钮 -->
                <div class="icon-btn" @click="goPage('Message')">
                    <n-icon class="icon">
                        <message></message>
                    </n-icon>
                    <div class="icon-action">
                        <div class="icon-action-msg">{{ t("common.message") }}</div>
                    </div>
                </div>
                <div class="icon-btn" @click="goPage('Collection')">
                    <n-icon class="icon">
                        <collection></collection>
                    </n-icon>
                    <div class="icon-action">
                        <div class="icon-action-msg">{{ t("common.collect") }}</div>
                    </div>
                </div>
                <!-- 投稿按钮 -->
                <n-button type="primary" class="upload-btn" @click="goPage('Upload')">
                    <template #icon>
                        <n-icon>
                            <upload></upload>
                        </n-icon>
                    </template>
                    {{ t("common.submission") }}
                </n-button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts" >
import { useI18n } from 'vue-i18n';
import { onBeforeMount, ref } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { CommonAvatar } from '@leaf/components';
import { NInput, NIcon, NButton } from 'naive-ui';
import type { UserInfoType } from '@leaf/apis';
import { globalConfig, storageData } from "@leaf/utils"
import { Search as SearchIcon, Upload, Message, Collection } from "@leaf/icons";
import { getTheme } from '@/theme';

// i18n
const { t } = useI18n();

const router = useRouter();
const route = useRoute();

const logout = () => {
    storageData.remove("token");
    storageData.remove('user_info');
    isLogin.value = false;
}

const goPage = (name: string) => {
    router.push({ name: name });
}

//搜索功能
const keywords = ref("");
const search = () => {
    let searchUrl = router.resolve({ name: "Search", params: { keywords: keywords.value } });
    window.open(searchUrl.href, '_blank');
}

const isLogin = ref(false);
const userInfo = ref<UserInfoType>({
    uid: 0,
    name: "",
    avatar: ""
});

// 主题
const themeStyle = () => {
    const theme = getTheme();
    return {
        "--header-bar-btn-active-color": theme.primaryColor
    }
}

onBeforeMount(() => {
    const tmpUserInfo = storageData.get("user_info");
    if (!tmpUserInfo) return;
    userInfo.value = tmpUserInfo;
    isLogin.value = true;
})
</script>

<style lang="less" scoped>
.header-bar {
    display: flex;
    align-items: center;
    width: 100%;
    height: 60px;
    // min-width: 1100px;
    background-color: #fff;
    -webkit-box-shadow: 0px 0px 3px #c8c8c8;
    -moz-box-shadow: 0px 0px 3px #c8c8c8;
    box-shadow: 0px 0px 3px #c8c8c8;

    .header-content {
        margin: 0 auto;
        width: 1100px;
        height: 50px;
        display: flex;
        justify-content: space-between;

        .header-left {
            display: flex;
            align-items: center;
            width: 260px;

            .title {
                color: rgba(0, 0, 0, .85);
                margin: 0;
                font-size: 16px;
                font-weight: 500;
                vertical-align: top;
                cursor: pointer;
            }
        }


        .header-center {
            width: 300px;
            margin-top: 8px;
        }

        .header-right {
            display: flex;
            align-items: center;
            width: 310px;

            .avatar-box {
                padding-top: 5px;
                margin-right: 6px;
            }

            .login-btn {
                width: 80px;
            }

            .upload-btn {
                margin-left: 6px;
                width: 120px;
            }
        }
    }
}


.icon-btn {
    width: 50px;
    height: 50px;
    margin: 0 3px;
    position: relative;
    display: inline-flex;
    align-items: center;
    justify-content: center;

    .icon {
        cursor: pointer;
        font-size: 20px;
        transition: font .46s cubic-bezier(.4, 0, .2, 1), transform .46s cubic-bezier(.4, 0, .2, 1);
    }

    .icon-action {
        position: absolute;
        top: calc(50% + 10px);
        opacity: 0;
        transition: opacity .46s cubic-bezier(.4, 0, .2, 1), transform .46s cubic-bezier(.4, 0, .2, 1);

        .icon-action-msg {
            cursor: pointer;
            font-size: 12px;
            color: var(--header-bar-btn-active-color);
        }
    }

    &:hover {
        .icon-action {
            opacity: 1;
            transform: translateY(-10px);
        }

        .icon {
            font-size: 16px;
            color: var(--header-bar-btn-active-color);
            transform: translateY(-10px);
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
        width: 200px;
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
        width: 160px;
        height: 36px;
        margin-left: 20px;
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

@media (min-width: 1400px) {
    .header-bar {
        .header-content {
            width: 1300px;
        }
    }
}
</style>