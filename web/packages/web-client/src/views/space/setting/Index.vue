<template>
    <div class="setting-title">
        <div class="content" :style="themeStyle">
            <span ref="infoRef" :class="isInfo ? 'active-text' : ''" @click="setRouter(true)">
                {{ t("space.basicInformation") }}
            </span>
            <span ref="securityRef" :class="!isInfo ? 'active-text' : ''" @click="setRouter(false)">
                {{ t("space.accountSecurity") }}
            </span>
        </div>
    </div>
    <router-view class="router" />
</template>

<script setup lang="ts">
import { getTheme } from '@/theme';
import { useI18n } from 'vue-i18n';
import { computed, onMounted, ref } from 'vue';
import { useRouter, useRoute } from 'vue-router';

const { t } = useI18n();
const route = useRoute();
const router = useRouter();

const isInfo = ref(true);
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
        min-width: 200px;
        margin: 0 auto;
        position: relative;

        span {
            line-height: 100px;
            font-size: 20px;
            user-select: none;
            cursor: pointer;

            &:nth-child(2) {
                margin-left: 30px;
            }

        }
    }
}

/**激活时 */
.active-text {
    color: var(--space-menu-active-color);
}

.router {
    width: 80%;
    margin: 0 auto;
}
</style>