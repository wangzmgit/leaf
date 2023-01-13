<template>
    <div v-title data-title="搜索">
        <header-bar></header-bar>
        <div class="search">
            <n-input round placeholder="搜索关键词~" v-model:value="keywords" @keydown.enter="searchVideo(true)">
                <template #suffix>
                    <n-icon @click="searchVideo(true)">
                        <Search />
                    </n-icon>
                </template>
            </n-input>
        </div>
        <div class="card-list">
            <video-item v-for="item in videoList" :info="item" :keywords="keywords"></video-item>
        </div>
    </div>
</template>

<script setup lang="ts">
import { searchVideoAPI } from "@leaf/apis";
import type { VideoType } from "@leaf/apis";
import VideoItem from "@/components/video-item/Index.vue";
import HeaderBar from "@/components/header-bar/Index.vue";
import { onBeforeMount, onBeforeUnmount, ref } from "vue";
import { NInput, NIcon, useMessage } from "naive-ui";
import { useRoute } from "vue-router";
import { Search } from "@leaf/icons";
import { statusCode } from "@leaf/utils";

const route = useRoute();
const message = useMessage();


let page = 1;
let pageSize = 15;
let noMore = false;
let loading = false;
const keywords = ref("");
const videoList = ref<Array<VideoType>>([]);
const searchVideo = (init = false) => {
    loading = true;
    if (init) {
        page = 1;
        videoList.value = [];
        noMore = false;
    }
    searchVideoAPI(page, pageSize, keywords.value).then((res) => {
        if (res.data.code === statusCode.OK) {
            if (res.data.data.videos) {
                videoList.value.push(...res.data.data.videos);
                if (res.data.data.videos.length < 15) {
                    noMore = true;
                }
            } else {
                noMore = true;
                message.error("获取失败");
            }
        }
        loading = false;
    })
}

//加载更多
const lazyLoading = () => {
    const scrollTop = document.documentElement.scrollTop || document.body.scrollTop;
    const clientHeight = document.documentElement.clientHeight;
    const scrollHeight = document.documentElement.scrollHeight;
    if (scrollTop + clientHeight >= (scrollHeight - 10)) {
        if (!noMore &&!loading) {
            page++;
            searchVideo();
        }
    }
}

onBeforeMount(() => {
    keywords.value = route.params.keywords.toString();
    searchVideo();
    window.addEventListener('scroll', lazyLoading, true);
})

onBeforeUnmount(()=>{
    window.removeEventListener('scroll', lazyLoading);
})
</script>

<style lang="less" scoped>
.search {
    width: 100%;
    margin: 30px auto;
    width: 500px;
}

.card-list {
    display: grid;
    gap: 0 16px;
    grid-template-columns: repeat(5, 1fr);
    width: 90%;
    margin: 20px auto;
}
</style>