<template>
    <div class="setting-title">
        <div class="content" :style="themeStyle">
            <span ref="infoRef" :class="isInfo ? 'active-text' : ''" @click="setRouter(true)">基本信息</span>
            <span ref="securityRef" :class="!isInfo ? 'active-text' : ''" @click="setRouter(false)">账号安全</span>
            <div class="selected" :class="isInfo ? '' : 'active-security'"></div>
        </div>
    </div>
    <router-view class="router" />
</template>

<script setup lang="ts">
import { getTheme } from '@/theme';
import { computed, onMounted, ref } from 'vue';
import { useRouter, useRoute } from 'vue-router';


const isInfo = ref(true);
const route = useRoute();
const router = useRouter();

const infoRef = ref<HTMLSpanElement | null>(null);
const securityRef = ref<HTMLSpanElement | null>(null);

const setRouter = (isInfoPage: boolean) => {
    if (isInfo.value !== isInfoPage) {
        //设置当前选中并修改页面
        isInfo.value = isInfoPage;
        router.push({ name: isInfoPage ? "InfoSetting" : "SecuritySetting" });
    }
}

const themeStyle = computed(() => {
    const theme = getTheme();
    return {
        "--space-menu-active-color": theme.primaryColor
    }
})

onMounted(() => {
    if (route.name === "SecuritySetting") {
        isInfo.value = false;
    }
})
</script>

<style lang="less" scoped>
.setting-title {

    height: 100px;
    text-align: center;

    .content {
        width: 200px;
        margin: 0 auto;
        position: relative;

        span {
            line-height: 100px;
            font-size: 20px;
            user-select: none;

            &:nth-child(2) {
                margin-left: 30px;
            }

        }
    }
}

.selected {
    position: absolute;
    top: 70px;
    left: 10px;
    width: 72px;
    height: 4px;
    border-radius: 2px;
    background-color: var(--space-menu-active-color);
    transition: left 0.3s;
    -moz-transition: left 0.3s;
    -webkit-transition: left 0.3s;
    -o-transition: left 0.3s;
}

/**激活时 */
.active-text {
    color: var(--space-menu-active-color);
}

.active-security {
    left: 120px;
}

.router {
    width: 80%;
    margin: 0 auto;
}
</style>