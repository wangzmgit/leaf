<template>
    <div v-title data-title="全部视频">
        <n-affix class="header" :top="0">
            <header-bar></header-bar>
        </n-affix>
        <!-- 视频列表 -->
        <div class="history-list">
            <div v-for="item in historyList">
                <span class="date-title">{{ relativeDate(item.date) }}</span>
                <div class="video-list">
                    <div class="video-item" v-for="(history, index) in item.videoList" :key="index"
                        @click="goVideo(history.video.vid, history.part)">
                        <img class="cover" :src="getResourceUrl(history.video.cover)" />
                        <div class="info">
                            <p class="title">{{ history.video.title }}</p>
                            <!--播放进度-->
                            <p class="clicks">播放到&ensp;{{ toTimeText(history.time) }}</p>
                        </div>
                    </div>
                </div>
            </div>

        </div>
    </div>
</template>

<script setup lang="ts">
import { useRouter } from "vue-router";
import { ref, onBeforeMount, onBeforeUnmount } from "vue";
import { NAffix } from "naive-ui";
import HeaderBar from "@/components/header-bar/Index.vue";
import type { HistoryVideoType } from "@leaf/apis";
import { getHistoryVideoAPI } from "@leaf/apis";
import { statusCode, timeStrToDateStr, relativeDate, getResourceUrl } from "@leaf/utils";

let page = 1;
let noMore = false;
let loading = false;
const historyMap = new Map<string, number>();
const historyList = ref<Array<{ date: string, videoList: Array<HistoryVideoType> }>>([]);
const getVideoList = () => {
    loading = true;
    getHistoryVideoAPI(page, 15).then((res) => {
        if (res.data.code === statusCode.OK) {
            if (res.data.data.history) {
                if (res.data.data.history.length < 15) {
                    noMore = true;
                }
                // 处理历史记录
                res.data.data.history.map((item: HistoryVideoType) => {
                    const date = timeStrToDateStr(item.created_at);
                    const index = historyMap.get(date);
                    if (index !== undefined) {
                        historyList.value[index].videoList.push(item);
                    } else {
                        const len = historyList.value.push({ date: date, videoList: [item] });
                        historyMap.set(date, len - 1);
                    }
                })
            } else {
                noMore = true;
            }
        }
        loading = false;
    });
}

//加载更多
const lazyLoading = () => {
    const scrollTop = document.documentElement.scrollTop || document.body.scrollTop;
    const clientHeight = document.documentElement.clientHeight;
    const scrollHeight = document.documentElement.scrollHeight;
    if (scrollTop + clientHeight >= (scrollHeight - 10)) {
        if (!noMore && !loading) {
            page++;
            getVideoList();
        }
    }
}

// 播放进度时间格式化
const toTimeText = (time: number) => {
    let m: number = Math.floor(time / 60);
    let s: number = Math.floor(time % 60);
    const mm = m < 10 ? "0" + m : m;
    const ss = s < 10 ? "0" + s : s;
    return mm + ":" + ss;
}

const router = useRouter();
const goVideo = (vid: number, part: number) => {
    router.push({ name: "Video", params: { vid: vid }, query: { part: part } });
}

onBeforeMount(() => {
    getVideoList();
    window.addEventListener('scroll', lazyLoading, true);
})

onBeforeUnmount(() => {
    window.removeEventListener('scroll', lazyLoading);
})
</script>

<style lang="less" scoped>
.header {
    z-index: 999;
    width: 100%;
}

.history-list {
    width: 90%;
    margin: 90px auto 20px;

    .date-title {
        color: #222;
        font-size: 18px;
        display: block;
        margin: 10px 0;
    }

    .video-list {
        display: grid;
        gap: 12px 16px;
        grid-template-columns: repeat(5, 1fr);
    }
}




.video-item {
    position: relative;
    width: 100%;
    height: 160px;

    .cover {
        width: 100%;
        height: 100%;
        overflow: hidden;
        border-radius: 6px;
    }

    .info {
        color: #fff;
        position: absolute;
        bottom: 0;
        width: 100%;
        border-radius: 6px;
        background: linear-gradient(0deg, rgba(0, 0, 0, 0.7), transparent);

        p {
            margin: 0;
            padding-left: 10px;
        }

        .title {
            font-size: 14px;
            line-height: 16px;
            overflow: hidden;
            text-overflow: ellipsis;
            display: -webkit-box;
            -webkit-line-clamp: 1;
            -webkit-box-orient: vertical;
        }

        .clicks {
            font-size: 10px;
        }
    }
}
</style>