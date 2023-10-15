<template>
    <div v-title :data-title="`${title} - 搜索`"></div>
    <header-bar class="header"></header-bar>
    <div ref="scrollBox" class="content" @scroll="scrollList">
        <video-list :videos="videoDataList"></video-list>
        <div v-show="loading" class="loading">
            <n-spin></n-spin>
        </div>
        <footer-bar></footer-bar>
    </div>
</template>

<script setup lang="ts">
import { ref, onBeforeMount, watch } from 'vue';
import { useRoute } from "vue-router";
import HeaderBar from '@/components/header-bar/Index.vue';
import FooterBar from '@/components/footer-bar/Index.vue';
import VideoList from '@/components/video-list/Index.vue';
import type { VideoType } from "@leaf/apis";
import { searchVideoAPI } from "@leaf/apis";
import { NSpin, useMessage } from 'naive-ui';
import { statusCode, globalConfig } from '@leaf/utils';

const title = window.$title || globalConfig.title;
const message = useMessage();
const route = useRoute();

const keywords = ref("");

//视频列表
let page = 1;
const pageSize = 10;
let noMore = false;
const loading = ref(false);
const videoDataList = ref(Array<VideoType>());
const searchVideo = (init:boolean = false) => {
    if (init) videoDataList.value = [];
    loading.value = true;
    searchVideoAPI(page, pageSize, keywords.value).then((res) => {
        if (res.data.code === statusCode.OK) {
            if (res.data.data.videos) {
                videoDataList.value.push(...res.data.data.videos);
                if (res.data.data.videos.length < pageSize) {
                    noMore = true;
                    message.info("没有更多了");
                }
            } else {
                noMore = true;
                message.info("没有更多了");
            }
            loading.value = false;
        }
    })
}

const scrollBox = ref<HTMLElement | null>(null);
const scrollList = (() => {
    //节流
    var timer: number | null = null;
    return () => {
        if (timer) {
            return
        }

        timer = setTimeout(() => {
            const scrollTop = scrollBox.value?.scrollTop || 0;
            const clientHeight = scrollBox.value?.clientHeight || 0;
            const scrollHeight = scrollBox.value?.scrollHeight || 0;
            if (scrollTop + clientHeight >= scrollHeight - 50) {
                if (!noMore && !loading.value) {
                    page++;
                    searchVideo();
                }
            }
            timer = null;
        }, 500);
    }
})();

watch(() => route.query.keywords, (newValue) => {
    keywords.value = (newValue || "").toString();
    searchVideo(true);
})

onBeforeMount(() => {
    keywords.value = (route.query.keywords || "").toString();
    searchVideo();
})
</script>

<style lang="less" scoped>
.header {
    top: 0;
    z-index: 999;
    position: fixed;
    height: 50px;
}

.content {
    overflow-y: scroll;
    margin-top: 50px;
    height: calc(100vh - 50px);

    .carousel {
        height: 220px;
    }

    .loading {
        padding: 20px 0;
        text-align: center;
    }
}
</style>