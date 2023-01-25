<template>
    <div class="header-box" :style="initTheme()">
        <common-avatar :url="userInfo ? userInfo.avatar : ''" :size="36" :iconsize="20"
            @click="headerRouter('Space')" />
        <div class="search-box">
            <n-input v-model:value="keywords" round placeholder="搜索关键词~" @keydown.enter="search">
                <template #suffix>
                    <n-icon @click="search">
                        <search-icon></search-icon>
                    </n-icon>
                </template>
            </n-input>
        </div>
        <n-icon class="msg" @click="headerRouter('Message')">
            <message-icon></message-icon>
        </n-icon>
    </div>
</template>

<script setup lang="ts">
import { computed, ref } from "vue";
import { useRouter, useRoute } from "vue-router";
import { NIcon, NInput } from "naive-ui";
import { storageData } from "@leaf/utils";
import { CommonAvatar } from "@leaf/components";
import { Search as SearchIcon, Message as MessageIcon } from "@leaf/icons";
import { getTheme } from "@/theme";

const route = useRoute();
const router = useRouter();

const initTheme = () => {
    const theme = getTheme();

    return {
        "--primary-color": theme.primaryColor
    }
}

const userInfo = computed(() => {
    return storageData.get("user_info");
})

const headerRouter = (page: string) => {
    const path = route.name;
    switch (page) {
        case 'Space':
            if (path == "SpaceInfo")
                page = "Home";
            break;
        case 'Message':
            if (path == "Message")
                page = "Home";
            break;
    }
    router.push({ name: page });
}

// 搜索
const keywords = ref("");
const search = () => {
    router.push({ name: "Search", query: { keywords: keywords.value } });
}
</script>

<style lang="less" scoped>
.header-box {
    width: calc(100% - 20px);
    height: 50px;
    display: flex;
    align-items: center;
    padding: 0 10px;
    background-color: var(--primary-color);
    justify-content: space-between;
    -webkit-box-shadow: 0px 0px 3px #c8c8c8;
    -moz-box-shadow: 0px 0px 3px #c8c8c8;
    box-shadow: 0px 0px 3px #c8c8c8;

    .search-box {
        width: calc(100% - 130px);
    }

    .msg {
        color: #fff;
        font-size: 26px;
    }
}
</style>