<template>
    <div class="announce-container" :style="themeStyle" @scroll="lazyLoading">
        <div class="announce-item" v-for="(item, index) in announceList" :key="index">
            <div class="title">
                <p class="item-title">{{ item.title }}</p>
                <n-time class="item-time" :time="new Date(item.created_at)"></n-time>
            </div>
            <span>{{ item.content }}</span>
            <span class="link" v-if="item.url">
                <span @click="linkTo(item.url)">网页链接</span>
            </span>
        </div>
    </div>
</template>

<script setup lang="ts">
import { computed, onBeforeMount, ref } from 'vue';
import { NTime } from 'naive-ui';
import { statusCode } from '@leaf/utils';
import type { AnnounceType } from '@leaf/apis';
import { getAnnounceAPI } from '@leaf/apis';
import { getTheme } from '@/theme';

const themeStyle = computed(() => {
    const theme = getTheme();
    return {
        "--link-color": theme.primaryColor
    }
})

let page = 1;
let pageSize = 10;
let noMore = false;
let loading = false;
const announceList = ref<Array<AnnounceType>>([]);
const getAnnounce = () => {
    loading = true;
    getAnnounceAPI(page, pageSize).then((res) => {
        if (res.data.code === statusCode.OK) {
            if (res.data.data.announces.length < pageSize) {
                noMore = true;
            }
            announceList.value = announceList.value.concat(res.data.data.announces);
        }
        loading = false;
    }).catch(()=>{
        loading = false;
    })
}

const lazyLoading = (() => {
    var timer: number | null = null;
    return (e: Event) => {
        if (timer) return;
        const target = e.target as HTMLElement;
        if (target.scrollTop + target.clientHeight >= (target.scrollHeight - 10)) {
            if (!noMore && !loading) {
                page++;
                getAnnounce();
            }
        }

        timer = setTimeout(() => {
            timer = null;
        }, 100);
    }
})();

const linkTo = (url: string) => {
    window.open(url, "_blank");
}


onBeforeMount(() => {
    getAnnounce();
})
</script>

<style lang="less" scoped>
.announce-container {
    height: 620px;
    overflow-y: auto;
    margin: 16px 20px 30px;

    &::-webkit-scrollbar {
        width: 6px;
    }

    &::-webkit-scrollbar-thumb {
        background-color: #e4e6eb;
        outline: none;
        border-radius: 2px;
    }

    &::-webkit-scrollbar-track {
        box-shadow: none;
        border-radius: 2px;
    }
}

.announce-item {
    padding: 8px 0 16px;
    margin-right: 10px;
    border-bottom: 1px solid #e1e2e3;

    .title {
        display: flex;
        align-items: center;
        justify-content: space-between;
        margin: 5px 0;

        .item-title {
            font-size: 16px;
            font-weight: 600;
            color: #333;
            width: calc(100% - 160px);
            margin: 0;
        }

        .item-time {
            color: #999;
        }
    }

    .link {
        margin-left: 10px;
        color: var(--link-color);
        cursor: pointer;
    }
}
</style>