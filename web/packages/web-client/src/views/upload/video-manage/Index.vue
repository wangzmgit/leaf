<template>
    <div class="video-list">
        <p class="create-title">{{ t("upload.manuscriptManagement") }}</p>
        <div class="card-item" v-for="(item, index) in videoList" :key="index">
            <div class="card-left">
                <img v-if="item.cover" :src="getResourceUrl(item.cover)" alt="视频封面">
            </div>
            <div class="card-center">
                <p class="title" @click="govideo(item.vid)">{{ item.title }}</p>
                <span class="desc">{{ `${t("common.desc")}：${item.desc}` }}</span>
                <span class="desc">
                    {{ t("common.createdAt") }}：
                    <n-time :time="new Date(item.created_at)"></n-time>
                </span>
            </div>
            <div class="card-right">
                <n-button class="edit-item" text @click="modifyVideo(item.vid, 'info')">
                    <template #icon>
                        <n-icon>
                            <edit></edit>
                        </n-icon>
                    </template>
                    编辑信息
                </n-button>
                <n-button class="edit-item" text @click="modifyVideo(item.vid, 'video')">
                    <template #icon>
                        <n-icon>
                            <edit></edit>
                        </n-icon>
                    </template>
                    编辑内容
                </n-button>
                <n-popconfirm negative-text="取消" positive-text="确认" @positive-click="deleteVideo(item.vid)">
                    <template #trigger>
                        <n-icon class="edit" size="20">
                            <delete></delete>
                        </n-icon>
                    </template>
                    是否删除这个视频
                </n-popconfirm>
            </div>
        </div>
    </div>
    <div class="pagination">
        <n-pagination v-model:page="page" :item-count="count" :page-size="pageSize" @update-page="pageChange" />
    </div>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n';
import { useRouter } from 'vue-router';
import { ref, onBeforeMount } from 'vue';
import { Edit, Delete } from '@leaf/icons';
import { getResourceUrl, statusCode } from '@leaf/utils';
import { NTime, NIcon, NPagination, NButton, NPopconfirm, useNotification } from 'naive-ui';
import { getUploadVideoAPI, deleteVideoAPI } from '@leaf/apis';
import type { UserUploadVideoType } from '@leaf/apis';

// i18n
const { t } = useI18n();

const notification = useNotification();

//获取用户上传的视频
const pageSize = 5;
const page = ref(1);
const count = ref(0);
const videoList = ref<Array<UserUploadVideoType>>([]);
const getUploadVideo = () => {
    getUploadVideoAPI(page.value, pageSize).then((res) => {
        if (res.data.code === statusCode.OK) {
            count.value = res.data.data.total;
            videoList.value = res.data.data.videos;
        }
    })
}

//删除视频
const deleteVideo = (id: number) => {
    deleteVideoAPI(id).then((res) => {
        if (res.data.code === statusCode.OK) {
            getUploadVideo();
        } else {
            notification.error({
                title: '删除失败',
                duration: 5000,
            })
        }
    })
}

//页码改变
const pageChange = (target: number) => {
    page.value = target;
    getUploadVideo();
}

//前往视频详情
const router = useRouter();
const govideo = (vid: number) => {
    let videoUrl = router.resolve({ name: "Video", params: { vid: vid } });
    window.open(videoUrl.href, '_blank');
}

//前往修改视频
const modifyVideo = (vid: number, status: string) => {
    router.push({ name: "UploadVideo", params: { vid: vid }, query: { modify: status } });
}

onBeforeMount(() => {
    getUploadVideo();
})
</script>

<style lang="less" scoped>
.video-list {
    margin: 0 30px;
    height: 580px;

    .create-title {
        line-height: 30px;
        font-size: 20px;
        user-select: none;
    }

    .card-item {
        width: 100%;
        height: 80px;
        margin-bottom: 12px;
        border-bottom: 1px solid #e6e6e6;
        padding-bottom: 12px;
    }
}

.card-left {
    float: left;
    width: 120px;
    height: 80px;
    margin-right: 10px;

    img {
        width: 100%;
        height: 100%;
        border-radius: 2px;
    }
}

.card-center {
    float: left;
    width: calc(100% - 400px);

    .title {
        font-size: 14px;
        color: #212121;
        line-height: 18px;
        margin: 0 0 26px;
        cursor: pointer;
        overflow: hidden;
        text-overflow: ellipsis;
        display: -webkit-box;
        -webkit-line-clamp: 1;
        -webkit-box-orient: vertical;

        &:hover {
            color: #36ad6a;
        }
    }

    .desc {
        font-size: 12px;
        color: #999;
        overflow: hidden;
        text-overflow: ellipsis;
        display: -webkit-box;
        -webkit-line-clamp: 1;
        -webkit-box-orient: vertical;
    }
}

.card-right {
    float: left;
    width: 260px;
    height: 100%;
    text-align: right;

    .edit-item {
        padding: 0 6px;
    }

    .edit {
        color: #999;
        cursor: pointer;
        line-height: 90px;
        margin-right: 20px;

        &:hover {
            color: #36ad6a;
        }
    }
}

.pagination {
    margin-left: 20px;
    padding-bottom: 20px;
}
</style>
