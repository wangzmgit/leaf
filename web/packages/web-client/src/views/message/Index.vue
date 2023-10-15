<template>
    <div class="message" :style="themeStyle" v-title :data-title="`消息中心-${title}`">
        <header-bar></header-bar>
        <div class="message-container">
            <div class="message-side-bar">
                <span class="title">{{ t("message.messageCenter") }}</span>
                <ul class="list">
                    <li class="item" :class="defaultOption === item.link ? 'active' : ''" v-for="item in menuOption"
                        @click="changeMenu(item.link)">{{ item.label }}</li>
                </ul>
            </div>
            <div class="message-router">
                <router-view></router-view>
            </div>
        </div>
    </div>
</template>


<script setup lang="ts">
import i18n from '@/locale';
import { useI18n } from "vue-i18n";
import { ref, onBeforeMount, computed } from "vue";
import { useRoute, useRouter } from 'vue-router';
import HeaderBar from '@/components/header-bar/Index.vue';
import { globalConfig } from "@leaf/utils";
import { getTheme } from "@/theme";

const title = ref(window.$title || globalConfig.title);

// i18n
const { t } = useI18n();

const themeStyle = computed(() => {
    const theme = getTheme();
    return {
        "--active-color": theme.primaryColor
    }
})

const route = useRoute();
const router = useRouter();
const defaultOption = ref('');//默认激活菜单

const menuOption = [
    {
        label: i18n.global.t("message.announce"),
        link: "Announce",
    },
    {
        label: i18n.global.t("message.like"),
        link: "Like",
    },
    {
        label: i18n.global.t("message.reply"),
        link: "Reply",
    },
    {
        label: i18n.global.t("message.at"),
        link: "At",
    },
    {
        label: i18n.global.t("message.whisper"),
        link: "Whisper",
    }
]

const changeMenu = (name: string) => {
    router.push({ name: name });
    defaultOption.value = name;
}

onBeforeMount(() => {
    defaultOption.value = route.name as string;
})
</script>

<style lang="less" scoped>
.message {
    height: 100%;
    width: 100%;
    top: 0;
    bottom: 0;
    position: fixed;
    overflow-y: scroll;
    background: #e9e9e9;
}

.message-container {
    width: 1100px;
    display: flex;
    margin: 10px auto 30px;
    min-height: 630px;

    .message-router {
        margin-left: 20px;
        width: calc(100% - 220px);
        height: 100%;
        background-color: #fff;
    }
}

.message-side-bar {
    width: 200px;
    min-height: 100%;
    border: 1px solid #efeff5;
    background-color: #fff;

    .title {
        height: 40px;
        line-height: 40px;
        font-size: 16px;
        color: #333;
        text-align: left;
        font-weight: 600;
        padding: 0 16px;
    }

    .list {
        margin: 0;
        padding: 0 0 0 20px;
        list-style-type: none;

        .item {
            height: 42px;
            cursor: pointer;
            display: flex;
            align-items: center;

            &:hover {
                color: var(--active-color);
            }

            &::before {
                content: '';
                left: 16px;
                top: 50%;
                width: 8px;
                height: 8px;
                margin-right: 10px;
                border-radius: 50%;
                background: #e5e5e5;
            }
        }

        .active {
            color: var(--active-color);
        }

    }
}

@media (min-width: 1400px) {
    .message-container {
        width: 1300px;
    }
}
</style>