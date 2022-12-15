<template>
    <div class="video">
        <header-bar></header-bar>
        <div class="video-main">
            <div class="content-left">
                <div v-if="!loading" class="video-player">
                    <video-player :vid="vid" :theme="theme.primaryColor" :part="part" :resources="resources"></video-player>
                    <div class="video-title-box">
                        <p class="video-title">{{ videoInfo!.title }}</p>
                        <p v-show="videoInfo!.copyright" class="copyright">
                            <n-icon color='#fd6d6f'>
                                <forbid />
                            </n-icon>
                            <span>未经作者授权，禁止转载</span>
                        </p>
                    </div>
                    <!-- 点赞收藏等数据 -->
                    <div class="video-toolbar">
                        <div class="toolbar-left">
                            <archive-info :stat="stat" :vid="videoInfo!.vid"></archive-info>
                        </div>
                        <!-- 日期播放和在线人数 -->
                        <div class="toolbar-right">
                            <span>{{ number }}人在看</span>
                            <span>{{ videoInfo!.clicks }}播放</span>
                            <span>上传于<n-time :time="new Date(videoInfo!.created_at)"></n-time></span>
                        </div>
                    </div>
                    <!--视频简介-->
                    <div class="desc">
                        <div :class="['desc-content', more ? 'open' : '']">{{ videoInfo!.desc }}</div>
                        <n-button text @click="more = !more">{{ more ? '收起' : '展开更多' }}</n-button>
                    </div>
                    <!--发表评论-->
                    <!-- <comment :vid="videoInfo!.vid"></comment> -->
                </div>
                <div v-else class="player-skeleton">
                    <n-skeleton height="500px" />
                    <n-skeleton text :repeat="2" />
                    <n-skeleton text width=" 60%" />
                </div>
            </div>
            <div class="content-right">
                <!--作者信息-->
                <author-card :loading="loading" :author="videoInfo?.author"></author-card>
                <!-- 视频分集 -->
                <div v-if="resources.length > 1">
                    <part-list :resources="resources" :active="part" @change="changePart"></part-list>
                </div>
                <!-- 用户视频 -->
                <!-- <author-video v-if="!loading" :uid="videoInfo!.author.uid"></author-video> -->
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { getTheme } from "@/theme"
// import config from '@/config';
// import { statType } from '@/types/archive';
import { useRoute, useRouter } from 'vue-router';
import { NIcon, NTime, NButton, NSkeleton } from 'naive-ui';
import { onBeforeMount, onBeforeUnmount, reactive, ref } from 'vue';
// import { OnlineSocketURL } from '@/utils/request';
// import Comment from './components/Comment.vue';
import PartList from './component/PartList.vue';
import HeaderBar from '@/components/header-bar/Index.vue';
import AuthorCard from './component/AuthorCard.vue';
import { VideoPlayer } from "@leaf/components";
// import AuthorVideo from './components/AuthorVideo.vue';
import ArchiveInfo from './component/ArchiveInfo.vue';

import { Forbid } from "@leaf/icons";
import type { PartListType, VideoType } from '@leaf/apis';
import { getVideoInfoAPI } from '@leaf/apis';
import { statusCode } from '@leaf/utils';

const route = useRoute();
const router = useRouter();
const theme = getTheme();
const vid = parseInt(route.params.vid.toString());
// const title = config.title;
const part = ref(1);//当前分集
const more = ref(false);//是否展开简介
const stat = ref({
    like: 0,
    collect: 0
});//点赞收藏数据
const loading = ref(true);
const resources = ref<Array<PartListType>>([]);
const videoInfo = ref<VideoType>();

//获取视频信息
const getVideoInfo = (vid: number) => {
    getVideoInfoAPI(vid).then((res) => {
        if (res.data.code === statusCode.OK) {
            videoInfo.value = res.data.data.video;
            resources.value = res.data.data.video.resources;
            //设置播放的资源
            if (!resources.value[part.value - 1]) {
                part.value = 1;
            }

            stat.value = res.data.data.stat;
            //修改网站标题
            // document.title = `${res.data.data.video.title}-${title}`
            loading.value = false;
        }
    })
}

//websocket
const number = ref(1);//在线人数
const SocketURL = ref("");
const websocket = ref<WebSocket | null>(null);
//初始化weosocket
const initWebSocket = (vid: number) => {
    // const wsProtocol = window.location.protocol === 'http:' ? 'ws://' : 'wss://';
    // if (OnlineSocketURL === "/api/v1/video/online/ws") {
    //     SocketURL.value = wsProtocol + window.location.host + "/api/v1/video/online/ws";
    // } else {
    //     //处理协议部分
    //     let reg = new RegExp('^http(s)?:')
    //     SocketURL.value = OnlineSocketURL.replace(reg, wsProtocol);
    // }
    // SocketURL.value += `?vid=${vid}`;
    // websocket.value = new WebSocket(SocketURL.value);
    // websocket.value.onmessage = websocketOnmessage;
}

//数据接收
const websocketOnmessage = (e: any) => {
    const res = JSON.parse(e.data);
    number.value = res.number;
}

const changePart = (target: number) => {
    if (resources.value[target - 1]) {
        part.value = target;
    }
    router.replace({ query: { p: part.value } });
}

onBeforeMount(() => {
    initWebSocket(Number(route.params.vid));
    getVideoInfo(Number(route.params.vid));
    if (route.query.p) {
        part.value = Number(route.query.p);
    }
})

onBeforeUnmount(() => {
    if (websocket.value) {
        websocket.value.close();
    }
})
</script>

<style lang="less" scoped>
.video {
    height: 100%;
    width: 100%;
    top: 0;
    bottom: 0;
    position: fixed;
    overflow-y: scroll;
}

.video-main {
    width: calc(100% - 100px);
    background: #fff;
    margin: 20px auto 0 auto;
    display: flex;
    flex-wrap: nowrap;
}

.content-left {
    width: calc(100% - 350px);
    min-width: 680px;

    .video-player {
        margin: 0 auto;
        width: calc(100% - 120px);
        /*16:9*/
        min-width: 680px;
        min-height: 382px;
    }

    .video-title-box {
        width: 100%;
        height: 36px;
        display: flex;

        .video-title {
            width: calc(100% - 160px);
            margin: 0;
            font-size: 20px;
            color: #303030;
            line-height: 30px;
            font-weight: 500;
            overflow: hidden;
            white-space: nowrap;
            text-overflow: ellipsis;
        }

        .copyright {
            width: 160px;
            display: flex;
            align-items: center;
            justify-content: flex-end;
            font-size: 12px;
            color: #636363;
        }
    }

    //骨架占位
    .player-skeleton {
        margin: 0 auto;
        width: calc(100% - 120px);
    }
}

.video-toolbar {
    color: #505050;
    display: flex;
    align-items: center;
    justify-content: space-between;
    border-bottom: 1px solid #e5e9f0;

    .toolbar-right {
        display: inline-block;

        span {
            margin-left: 20px;
        }
    }
}

//简介
.desc {
    .desc-content {
        max-height: 36px;
        font-size: 14px;
        color: #212121;
        line-height: 18px;
        overflow: hidden;
        white-space: pre-line;
        transition: height 0.3s;
    }

    .open {
        // height: auto;
        max-height: none;
    }
}

/**右半部分 */
.content-right {
    width: 350px;
    height: 600px;
    min-width: 350px;
}
</style>