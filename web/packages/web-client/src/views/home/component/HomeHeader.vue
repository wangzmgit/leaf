<template>
    <div class="header-bar">
        <div class="sidebar-title">
            <n-icon class="menu-icon-btn" size="22" @click="foldClick">
                <menu-outline></menu-outline>
            </n-icon>
            <span>{{ globalConfig.title }}</span>
        </div>
        <div class="header-search">
            <n-input class="search-input" v-model:value="keywords" round placeholder="搜索" @keydown.enter="search">
                <template #suffix>
                    <n-icon @click="search">
                        <search-icon />
                    </n-icon>
                </template>
            </n-input>
        </div>
        <div class="header-user-info">
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
                <n-button text @click="showLogin = true">
                    {{ t("common.login") }}/{{ t("common.register") }}
                </n-button>
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
        <login-dialog v-if="showLogin" @close="loginClose"></login-dialog>
    </div>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n';
import { onBeforeMount, ref } from 'vue';
import { globalConfig, storageData } from "@leaf/utils";
import { useRouter } from 'vue-router';
import { CommonAvatar } from '@leaf/components';
import { NInput, NIcon, NButton } from 'naive-ui';
import type { UserInfoType } from "@leaf/apis";
import { useLoginStore } from '@/stores/login-store';
import { Search as SearchIcon, Upload, MenuOutline } from "@leaf/icons"
import LoginDialog from "@/components/login-dialog/Index.vue";
import { storeToRefs } from 'pinia';


const emits = defineEmits(["changeFold"]);

// i18n
const { t } = useI18n();

const isLogin = ref(false);
const router = useRouter();

const loginStore = useLoginStore();
const { showLogin } = storeToRefs(loginStore);

// 用户信息
const userInfo = ref<UserInfoType>({
    uid: 0,
    name: "",
    avatar: "",
    spacecover: ""
});

// 加载用户信息
const loadUserInfo = () => {
    const info = storageData.get('user_info');
    if (info) {
        userInfo.value = Object.assign(userInfo.value, info);
        isLogin.value = true;
    }
}

// 左侧菜单
const menuFold = ref(false);
const foldClick = () => {
    menuFold.value = !menuFold.value;
    emits("changeFold", menuFold.value);
}

// 退出登录
const logout = () => {
    storageData.remove("token");
    storageData.remove('user_info');
    isLogin.value = false;
}

// 页面跳转
const goPage = (name: string) => {
    router.push({ name: name });
}

//搜索功能
const keywords = ref("");
const search = () => {
    let searchUrl = router.resolve({ name: "Search", params: { keywords: keywords.value } });
    window.open(searchUrl.href, '_blank');
}

const loginClose = () => {
    loadUserInfo();
    loginStore.setLoginState(false);
}

onBeforeMount(() => {
    loadUserInfo();
})
</script>

<style lang="less" scoped>
.header-bar {
    display: flex;
    align-items: center;
    width: 54;
    height: 60px;
    display: flex;
    align-items: center;
    box-shadow: inset 0 -1px #f1f2f3;

    .sidebar-title {
        display: flex;
        align-items: center;
        height: 100%;
        width: 220px;

        .menu-icon-btn {
            margin: 0 6px;
            padding: 6px;
            border-radius: 50%;
            cursor: pointer;

            &:hover {
                background-color: rgba(0, 0, 0, .1);
            }
        }
    }


    .header-search {
        flex: 1;

        .search-input {
            width: 60%;
            margin-left: 60px;
        }
    }

    .header-user-info {
        width: 280px;
        display: flex;
        align-items: center;

        .avatar-box {
            padding-top: 5px;
            margin-right: 6px;
        }

        .upload-btn {
            margin: 0 36px;
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
</style>