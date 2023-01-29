<template>
    <n-card class="card">
        <div class="info-header">
            <partition-selector ref="partitionRef" @selected="selectedPartition"></partition-selector>
            <search-box class="search" :keyword="keyword" @search="searchVideo"></search-box>
        </div>
        <n-table class="table" striped>
            <thead class="table-head">
                <tr>
                    <th>VID</th>
                    <th>封面</th>
                    <th>标题</th>
                    <th>简介</th>
                    <th>作者</th>
                    <th>上传时间</th>
                    <th>操作</th>
                </tr>
            </thead>
            <tbody class="table-body">
                <tr v-for="item in videoList">
                    <td>{{ item.vid }}</td>
                    <td>
                        <img class="cover" :src="getResourceUrl(item.cover)" alt="视频封面">
                    </td>
                    <td>{{ item.title }}</td>
                    <td>{{ item.desc }}</td>
                    <td>{{ item.author.name }}</td>
                    <td>
                        <n-time :time="new Date(item.created_at)" />
                    </td>
                    <td class="btn-list">
                        <n-button text @click="deleteVideo(item.vid)">删除</n-button>
                    </td>
                </tr>
            </tbody>
        </n-table>
        <div class="pagination">
            <n-pagination v-model:page="page" :item-count="count" :page-size="6" @update-page="pageChange" />
        </div>
    </n-card>
</template>

<script setup lang="ts">

import { onBeforeMount, ref } from 'vue';
import { SearchBox } from '@leaf/components';
import PartitionSelector from "./component/PartitionSelector.vue";
import type { VideoType } from "@leaf/apis";
import { adminGetVideoListAPI, adminSearchVideoAPI, adminDeleteVideoAPI } from "@leaf/apis";
import { NTable, NButton, NCard, NTime, NPagination, useNotification } from 'naive-ui';
import { getResourceUrl, statusCode } from '@leaf/utils';


const notification = useNotification();//通知
const partitionRef = ref<InstanceType<typeof PartitionSelector> | null>(null)

const page = ref(1);
const count = ref(0);
let partitionId = 0;
const videoList = ref<Array<VideoType>>([]);
const getVideoList = () => {
    adminGetVideoListAPI(page.value, 5,partitionId).then((res) => {
        if (res.data.code === statusCode.OK) {
            count.value = res.data.data.total;
            videoList.value = res.data.data.videos;
        }
    }).catch((err) => {
        notification.error({
            title: '加载视频失败',
            duration: 5000,
        })
    });
}

// 选择分区
const selectedPartition = (partition: number) => {
    partitionId = partition;
    getVideoList();
}

//页码改变
const pageChange = (target: number) => {
    page.value = target;
    getVideoList();
}

//搜索
const keyword = ref('');
const searchVideo = (keyword: string) => {
    page.value = 1;
    count.value = 0;
    partitionRef.value?.reset();
    if (!keyword) {
        getVideoList();
        return;
    }
    adminSearchVideoAPI(page.value, 5, keyword).then((res) => {
        if (res.data.code === statusCode.OK) {
            count.value = res.data.data.total;
            videoList.value = res.data.data.videos;
        }
    })
}

//删除
const deleteVideo = (id: number) => {
    adminDeleteVideoAPI(id).then((res) => {
        if (res.data.code === statusCode.OK) {
            getVideoList();
        }
    }).catch((err) => {
        notification.error({
            title: '删除失败',
            duration: 5000,
        })
    });
}

onBeforeMount(() => {
    getVideoList();
})
</script>

<style lang="less" scoped>
.card {
    margin: 10px;
    width: calc(100% - 20px);
    border-radius: 10px;

    .info-header {
        display: flex;
        height: 30px;
        align-items: center;
        justify-content: space-between;
        margin-bottom: 20px;

        .search {
            width: 300px;
        }
    }

    .table {
        .table-head {
            text-align: center;
        }

        .table-body {
            text-align: center;

            .cover {
                height: 60px;
                width: 100px;
            }

            .btn-list {
                button {
                    margin: 0 6px;
                }
            }
        }
    }
}

.pagination {
    margin-top: 20px;
}
</style>