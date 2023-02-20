<template>
    <n-card class="card">
        <n-table striped>
            <thead class="table-head">
                <tr>
                    <th>VID</th>
                    <th>封面</th>
                    <th>标题</th>
                    <th>分区</th>
                    <th>简介</th>
                    <th>上传时间</th>
                    <th>允许转载</th>
                    <th>操作</th>
                </tr>
            </thead>
            <tbody class="table-body">
                <tr v-for="item in videos">
                    <td>{{ item.vid }}</td>
                    <td>
                        <img class="cover" :src="getResourceUrl(item.cover)" alt="视频封面">
                    </td>
                    <td>{{ item.title }}</td>
                    <td>{{ getPartitionName(item.partition) }}</td>
                    <td>{{ item.desc }}</td>
                    <td>
                        <n-time :time="new Date(item.created_at)" />
                    </td>
                    <td>{{ toCopyright(item.copyright) }}</td>
                    <td>
                        <div class="btn-list">
                            <n-button text @click="getResourceList(item.vid)">查看视频</n-button>
                            <n-button text @click="reviewVideo(item.vid, reviewCode.AUDIT_APPROVED)">通过</n-button>
                            <n-button text @click="reviewVideo(item.vid, reviewCode.WRONG_VIDEO_INFO)">不通过</n-button>
                        </div>
                    </td>
                </tr>
            </tbody>
        </n-table>
        <div class="pagination">
            <n-pagination v-model:page="pagination.current" :item-count="pagination.count"
                :page-size="pagination.pageSize" @update-page="pageChange" />
        </div>
    </n-card>
    <n-drawer v-model:show="activeDrawer" :width="502">
        <n-drawer-content title="视频列表">
            <video-list :list="resources"></video-list>
        </n-drawer-content>
    </n-drawer>
</template>

<script setup lang="ts">
import { onBeforeMount, reactive, ref } from 'vue';
import VideoList from './component/VideoList.vue';
import { statusCode, reviewCode, getResourceUrl } from '@leaf/utils';
import type { VideoType, PartitionType } from '@leaf/apis';
import { getReviewListAPI, getResourceListAPI, reviewVideoAPI, getPartitionAPI } from '@leaf/apis';
import { NTable, NButton, NCard, NTime, NDrawer, NPagination, NDrawerContent, useMessage } from 'naive-ui';

const pagination = reactive({
    count: 0,
    current: 1,
    pageSize: 8,
})

const videos = ref<Array<VideoType>>([]);
const message = useMessage();//通知

const getVideoList = () => {
    getReviewListAPI(pagination.current, pagination.pageSize).then((res) => {
        if (res.data.code === statusCode.OK) {
            pagination.count = res.data.data.total;
            videos.value = res.data.data.videos;
        }
    }).catch(() => {
        message.error('获取视频列表失败');
    });
}

//页码改变
const pageChange = (target: number) => {
    pagination.current = target;
    getVideoList();
}

// 获取分区列表
const partitionList = ref<Array<PartitionType>>([]);//所有分区
const getAllPartition = () => {
    getPartitionAPI().then((res) => {
        if (res.data.code === statusCode.OK) {
            partitionList.value = res.data.data.partitions;
        }
        // 先获取分区后获取视频
        getVideoList();
    })
}

const getPartitionName = (partitionId?: number) => {
    const partition = partitionList.value.find((item) => {
        return item.id === partitionId;
    })

    return partition?.content || "未知";
}

const resources = ref([]);//视频资源
const activeDrawer = ref(false);//是否打开审核抽屉
//获取审核资源
const getResourceList = (vid: number) => {
    getResourceListAPI(vid).then((res) => {
        if (res.data.code === statusCode.OK) {
            resources.value = res.data.data.resources;
            activeDrawer.value = true;
        }
    }).catch(() => {
        message.error('获取资源失败');
    });
}

//审核视频
const reviewVideo = (vid: number, status: number) => {
    reviewVideoAPI(vid, status).then((res) => {
        if (res.data.code === statusCode.OK) {
            getVideoList();
        } else {
            message.error(res.data.msg);
        }
    }).catch(() => {
        message.error('审核调用失败');
    });
}

//允许转载
const toCopyright = (copyright: boolean) => {
    if (copyright) return "否";
    else return "是";
}

onBeforeMount(() => {
    getAllPartition();
})
</script>

<style lang="less" scoped>
.card {
    margin: 10px;
    width: calc(100% - 20px);
    border-radius: 10px;
}

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

.pagination {
    margin-top: 20px;
}
</style>