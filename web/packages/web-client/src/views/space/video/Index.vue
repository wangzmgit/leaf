<template>
    <div class="video-box">
        <p class="video-title">{{ t('space.myVideos') }}</p>
        <div class="card-list">
            <div class="v-card" v-for="(item, index) in videoList" :key="index">
                <div class="card-item" @click="goVideo(item.status, item.vid)">
                    <img class="cover" :src="getResourceUrl(item.cover)" />
                    <div class="info">
                        <p class="title">{{ item.title }}</p>
                        <!--播放量-->
                        <p class="clicks">{{ t('common.clicks') }}&ensp;{{ item.clicks }}</p>
                    </div>
                </div>
                <div class="my-upload-card-btn">
                    <n-button text v-if="item.status !== reviewCode.AUDIT_APPROVED" @click="viewStatus(item.vid)">
                        <template #icon>
                            <n-icon>
                                <edit></edit>
                            </n-icon>
                        </template>
                        {{ t('space.viewStatus') }}
                    </n-button>
                    <n-button text v-else @mouseover="showMenu(index, true)" @mouseleave="showMenu(index, false)">
                        <template #icon>
                            <n-icon>
                                <edit></edit>
                            </n-icon>
                        </template>
                        {{ t('space.modifyContent') }}
                    </n-button>
                    <n-button text @click="deleteVideo(item.vid)">
                        <template #icon>
                            <n-icon>
                                <delete></delete>
                            </n-icon>
                        </template>
                        {{ t('common.delete') }}
                    </n-button>
                    <!--修改视频-->
                    <div v-show="modifyMenu[index]" class="modify-menu" @mouseleave="showMenu(index, false)">
                        <div class="menu-item">
                            <span class="modify-btn" @click="modifyVideo(item.vid, 'video')">
                                {{ t('space.modifyVideo') }}
                            </span>
                        </div>
                        <div class="menu-item">
                            <span class="modify-btn" @click="modifyVideo(item.vid, 'info')">
                                {{ t('space.modifyInfo') }}
                            </span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <n-pagination v-model:page="page" :item-count="count" :page-size="pageSize" @update-page="pageChange" />
    </div>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n';
import { useRouter } from 'vue-router';
import { ref, onBeforeMount } from 'vue';
import { NButton, NIcon, NPagination, useNotification } from 'naive-ui';
import { Delete, Edit } from '@leaf/icons';
import { deleteVideoAPI, getUploadVideoAPI } from '@leaf/apis';
import type { UserUploadVideoType } from '@leaf/apis';
import { getResourceUrl, reviewCode, statusCode } from '@leaf/utils';
import { useVideoCountStore } from '@/stores/video-count-store';

// i18n
const { t } = useI18n();

const router = useRouter();
const notification = useNotification();//通知

//获取我的视频
const pageSize = 8;
const page = ref(1);
const count = ref(0);
const videoCountStore = useVideoCountStore();
const videoList = ref<Array<UserUploadVideoType>>([]);
const getMyVideo = () => {
    getUploadVideoAPI(page.value, pageSize).then((res) => {
        if (res.data.code === statusCode.OK) {
            count.value = res.data.data.total;
            videoList.value = res.data.data.videos;
            videoCountStore.setVideoCountState(count.value);
        }
    })
}

//删除视频
const deleteVideo = (id: number) => {
    deleteVideoAPI(id).then((res) => {
        if (res.data.code === statusCode.OK) {
            getMyVideo();
        } else {
            notification.error({
                title: '删除失败',
                duration: 5000,
            })
        }
    })
}

//修改菜单
const modifyMenu = ref<Array<boolean>>([]);
const showMenu = (index: number, show: boolean) => {
    if (modifyMenu.value[index] && !show) {
        modifyMenu.value[index] = false;
        return;
    }
    for (let i = 0; i < videoList.value.length; i++) {
        modifyMenu.value[i] = false;
    }
    modifyMenu.value[index] = true;
}

//查看状态
const viewStatus = (vid: number) => {
    router.push({ name: "UploadVideo", params: { vid: vid } });
}

const modifyVideo = (vid: number, status: string) => {
    router.push({ name: "UploadVideo", params: { vid: vid }, query: { modify: status } });
}

const goVideo = (review: number, vid: number) => {
    if (review === reviewCode.AUDIT_APPROVED) {
        router.push({ name: "Video", params: { vid: vid } });
    }
}

//页码改变
const pageChange = (target: number) => {
    page.value = target;
    getMyVideo();
}

onBeforeMount(() => {
    getMyVideo();
})
</script>

<style lang="less" scoped>
.video-box {
    padding: 0 20px;

    .video-title {
        font-size: 18px;
        margin-top: 20px;
    }
}

.card-list {
    width: 100%;
    height: 360px;
    display: grid;
    grid-template-rows: repeat(2, 50%);
    grid-template-columns: repeat(4, 25%);

    .v-card {
        padding: 5px;
        height: 170px;
        width: calc(100% - 10px);

        .card-item {
            position: relative;
            width: 100%;
            height: 130px;

            .cover {
                width: 100%;
                height: 100%;
                border-top-left-radius: 6px;
                border-top-right-radius: 6px;
            }

            .info {
                color: #fff;
                position: absolute;
                bottom: 0;
                width: 100%;
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
    }
}

.my-upload-card-btn {
    width: calc(100% - 2px);
    border: 1px solid #d0d0d0;
    border-top: 0;
    border-bottom-left-radius: 6px;
    border-bottom-right-radius: 6px;

    button {
        width: 50%;
        color: #6f6f6f;
    }

    .modify-menu {
        width: 180px;
        height: 92px;
        z-index: 999;
        position: absolute;
        background-color: #fff;
        border-radius: 10px;
        box-shadow: 0 0 30px rgba(18, 18, 18, 0.2);

        .menu-item {
            margin-top: 7px;
            width: 144px;
            height: 36px;
            margin-left: 20px;
            -webkit-user-select: none;
            -moz-user-select: none;
            -ms-user-select: none;
            user-select: none;

            .modify-btn {
                display: block;
                color: #18191b;
                line-height: 36px;
                text-align: left;
                border-radius: 6px;
                padding-left: 6px;

                &:hover {

                    background-color: #c9ccd0;
                }
            }
        }
    }
}

@media (min-width: 1400px) {
    .video-box {
        .card-list {
            height: 400px;

            .v-card {
                height: 200px;

                .card-item {
                    height: 150px;
                }
            }
        }
    }
}
</style>