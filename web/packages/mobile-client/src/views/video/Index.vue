<template>
    <div class="video" :style="initTheme()" v-title :data-title="`视频播放-${title}`">
        <header-bar class="header-bar"></header-bar>
        <div class="video-main">
            <div v-if="!loading" class="video-player">
                <video-player v-if="ios === 0" :vid="vid" :part="part" :resources="resources" :mobile="true"
                    :theme="theme.primaryColor"></video-player>
                <ios-video-player class="ios" v-else :vid="vid" :resources="resources" :theme="theme.primaryColor"
                    :part="part"></ios-video-player>
                <!-- 视频信息 -->
                <div class="video-info">
                    <div class="title-wrapper">
                        <span :class="['title', fold ? 'title-fold' : '']">{{ videoInfo?.title }}</span>
                        <div class="fold" @click="fold = !fold">
                            <n-icon size="22" color='#666'>
                                <arrow-down v-if="fold" />
                                <arrow-up v-else />
                            </n-icon>
                        </div>
                    </div>
                    <n-collapse-transition :show="!fold">
                        <p class="info-item">{{ videoInfo?.desc }}</p>
                        <!-- 视频分集 -->
                        <div v-if="resources.length > 1">
                            <part-list :resources="resources" :active="part" @change="changePart"></part-list>
                        </div>
                        <p class="info-item" v-show="videoInfo?.copyright">
                            <n-icon color='#fd6d6f'>
                                <forbid />
                            </n-icon>
                            <span>未经作者授权，禁止转载</span>
                        </p>
                    </n-collapse-transition>
                    <div class="author">
                        <div class="info-fold">
                            <common-avatar :size="30" :iconsize="16" :url="videoInfo!.author.avatar"></common-avatar>
                            <span class="name">{{ videoInfo!.author.name }}</span>
                        </div>
                        <n-time type="relative" :time="new Date(videoInfo!.created_at)"></n-time>
                    </div>
                </div>
                <!--发表评论-->
                <comment-list :vid="videoInfo!.vid"></comment-list>
            </div>
            <div v-else class="player-skeleton">
                <n-skeleton height="200px" />
                <n-skeleton text :repeat="2" />
                <n-skeleton text width=" 60%" />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { getTheme } from "@/theme";
import { useRoute, useRouter } from 'vue-router';
import { NIcon, NTime, NSkeleton, NCollapseTransition } from 'naive-ui';
import { onBeforeMount, ref } from 'vue';
import CommentList from './component/CommentList.vue';
import PartList from './component/PartList.vue';
import HeaderBar from '@/components/header-bar/Index.vue';
import { CommonAvatar, VideoPlayer, IosVideoPlayer } from '@leaf/components';
import { Forbid, ArrowDown, ArrowUp } from "@leaf/icons";
import { globalConfig, statusCode } from '@leaf/utils';
import type { VideoType } from '@leaf/apis';
import { getVideoInfoAPI } from '@leaf/apis';

const title = ref(window.$title || globalConfig.title);

const route = useRoute();
const router = useRouter();
const theme = getTheme();
const vid = parseInt(route.params.vid.toString());

const ios = ref(navigator.userAgent.includes("iPhone") ? 1 : 0);
const part = ref(1);//当前分集
const fold = ref(true);//折叠

const initTheme = () => {
    return {
        "--hover-color": theme.primaryHoverColor,
        "--primary-color": theme.primaryColor
    }
}


//获取视频信息
const loading = ref(true);
const resources = ref([]);
const videoInfo = ref<VideoType | null>(null);
const getVideoInfo = (vid: number, ios: number) => {
    getVideoInfoAPI(vid, ios).then((res) => {
        if (res.data.code === statusCode.OK) {
            videoInfo.value = res.data.data.video;
            resources.value = res.data.data.video.resources;
            //设置播放的资源
            if (!resources.value[part.value - 1]) {
                part.value = 1;
            }

            //修改网站标题
            document.title = `${res.data.data.video.title}-${window.$title || globalConfig.title}`
            loading.value = false;
        }
    })
}

const changePart = (target: number) => {
    if (resources.value[target - 1]) {
        part.value = target;
    }
    router.replace({ query: { p: part.value } });
}

onBeforeMount(() => {
    getVideoInfo(vid, ios.value);
    if (route.query.p) {
        part.value = Number(route.query.p);
    }
})
</script>

<style lang="less" scoped>
.video {
    height: 100%;
    width: 100%;
}

.header-bar {
    position: fixed;
    top: 0;
    z-index: 1000;
}

.video-main {
    width: 100%;
    background: #fff;
    margin-top: 50px;
    display: flex;
    flex-wrap: nowrap;

    .video-player {
        width: 100%;
    }

    //骨架占位
    .player-skeleton {
        margin: 0 auto;
        width: 100%;
    }
}

.video-info {
    margin: 50px 0 10px;
    padding: 0 10px 16px;
    border-bottom: 1px solid #efeff5;

    .title-wrapper {
        display: flex;
        align-items: center;
        justify-content: space-between;

        .title {
            font-size: 17px;
        }

        .fold {
            width: 50px;
            padding-top: 6px;
            text-align: right;
        }
    }

    .info-item {
        color: #666;
        margin: 4px 0 2px 0;
    }

    .author {
        margin-top: 12px;
        display: flex;
        align-items: center;
        justify-content: space-between;

        .info-fold {
            display: flex;
            align-items: center;

            .name {
                font-size: 16px;
                line-height: 30px;
                padding-left: 6px;
            }
        }
    }
}
</style>