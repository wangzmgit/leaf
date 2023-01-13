<template>
    <div v-title data-title="全部视频">
        <n-affix class="header" :top="0">
            <header-bar></header-bar>
        </n-affix>
        <!-- 分区选择 -->
        <div class="content">
            <div class="partitions">
                <p :class="partitionInfo.selectedPartition === 0 ? 'selected' : ''" @click="selectPartition(0)">
                    <span>全部</span>
                </p>
                <p v-for="item in partitionInfo.partitions" :key="item.id" @click="selectPartition(item.id)">
                    <span :class="partitionInfo.selectedPartition === item.id ? 'selected' : ''">
                        {{ item.content }}
                    </span>
                </p>
            </div>
            <div class="partitions">
                <p :class="partitionInfo.selectedSubpartition == 0 ? 'selected' : ''" @click="selectSubartition(0)">
                    <span>全部</span>
                </p>
                <p v-for="item in partitionInfo.subpartition" :key="item.id" @click="selectSubartition(item.id)">
                    <span :class="partitionInfo.selectedSubpartition == item.id ? 'selected' : ''">
                        {{ item.content }}
                    </span>
                </p>
            </div>
        </div>
        <!-- 视频列表 -->
        <div class="card-list">
            <div class="video-item" v-for="(item, index) in videoList" :key="index">
                <video-item :info="item"></video-item>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useRoute, useRouter } from "vue-router";
import { reactive, ref, onBeforeMount, onBeforeUnmount } from "vue";
import { statusCode } from "@leaf/utils";
import { NAffix, useNotification } from "naive-ui";
import type { VideoType, PartitionType } from "@leaf/apis";
import { getVideoListAPI, getPartitionAPI } from "@leaf/apis";
import HeaderBar from "@/components/header-bar/Index.vue";
import VideoItem from '@/components/video-item/Index.vue';

const partitionInfo = reactive({
    allPartition: [] as PartitionType[],
    partitions: [] as PartitionType[], //分区内容
    subpartition: [] as PartitionType[], //子分区内容
    selectedPartition: 0,//选中分区
    selectedSubpartition: 0,//选中子分区
});
const route = useRoute();
const router = useRouter();
const notification = useNotification();

// 获取所有分区
const getAllPartition = () => {
    getPartitionAPI().then((res) => {
        if (res.data.code === statusCode.OK) {
            partitionInfo.allPartition = res.data.data.partitions;
            partitionInfo.partitions = partitionInfo.allPartition.filter((item) => {
                return item.parent_id === 0;
            });
            if (partitionInfo.selectedPartition !== 0) {
                getPartitionList(partitionInfo.selectedPartition);
            }
        } else {
            notification.error({
                title: "分区加载失败",
                duration: 5000
            })
        }
    })
}


// 获取分区列表
const getPartitionList = (fid: number) => {
    partitionInfo.subpartition = partitionInfo.allPartition.filter((item) => {
        return item.parent_id === fid;
    });
}

//设置分区
const selectPartition = (id: number) => {
    if (partitionInfo.selectedPartition === id) return;
    let newQuery = JSON.parse(JSON.stringify(route.query));
    newQuery.partition = id;
    partitionInfo.selectedPartition = id;
    partitionInfo.selectedSubpartition = 0;
    router.replace({ query: newQuery });
    if (id === 0) {
        partitionInfo.subpartition = [];
    } else {
        getPartitionList(id);
    }
    getVideoList(true);
}

//设置子分区
const selectSubartition = (id: number) => {
    if (partitionInfo.selectedSubpartition === id) return;
    let newQuery = JSON.parse(JSON.stringify(route.query));
    newQuery.subpartition = id;
    partitionInfo.selectedSubpartition = id;
    router.replace({ query: newQuery });

    getVideoList(true);
}

let page = 1;
let noMore = false;
let loading = false;
const videoList = ref<Array<VideoType>>([]);
const getVideoList = (init = false) => {
    loading = true;
    if (init) {
        page = 1;
        videoList.value = [];
        noMore = false;
    }
    const current = partitionInfo.selectedSubpartition || partitionInfo.selectedPartition;
    getVideoListAPI(page, 15, current).then((res) => {
        if (res.data.code === statusCode.OK) {
            if (res.data.data.videos) {
                videoList.value.push(...res.data.data.videos);
                if (res.data.data.videos.length < 15) {
                    noMore = true;
                }
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
        if (!noMore &&!loading) {
            page++;
            getVideoList();
        }
    }
}

onBeforeMount(() => {
    partitionInfo.selectedPartition = Number(route.query.partition);
    partitionInfo.selectedSubpartition = Number(route.query.subpartition) || 0;
    getAllPartition();
    getVideoList();
    window.addEventListener('scroll', lazyLoading, true);
})

onBeforeUnmount(()=>{
    window.removeEventListener('scroll', lazyLoading);
})
</script>

<style lang="less" scoped>
.header {
    z-index: 999;
    width: 100%;
}

.content {
    width: 90%;
    min-width: 700px;
    margin: 90px auto 10px auto;
}

.partitions {
    display: flex;
    margin-top: 20px;

    p {
        cursor: pointer;
        margin: 0 20px;

        span {
            padding: 3px 6px;
        }
    }

    .selected {
        color: #fff;
        border-radius: 6px;
        background-color: rgba(24, 143, 255, 0.5);
    }
}

.card-list {
    display: grid;
    gap: 0 16px;
    grid-template-columns: repeat(5, 1fr);
    width: 90%;
    margin: 20px auto;
}
</style>