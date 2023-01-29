<template>
    <div v-title data-title="个人中心"></div>
    <header-bar></header-bar>
    <div class="user-card">
        <div class="card-avatar">
            <common-avatar :url="userInfo.avatar" :size="76" :iconsize="50"></common-avatar>
        </div>
        <!--关注粉丝信息-->
        <div class="card-info-box">
            <div class="card-info">
                <div>
                    <p class="info-title">投稿</p>
                    <p class="info-content">{{ userData.videoCount }}</p>
                </div>
                <div>
                    <p class="info-title">关注</p>
                    <p class="info-content">{{ userData.followingCount }}</p>
                </div>
                <div>
                    <p class="info-title">粉丝</p>
                    <p class="info-content">{{ userData.followerCount }}</p>
                </div>
            </div>
            <n-button type="primary" ghost @click="goPage('EditInfo')">编辑信息</n-button>
        </div>
    </div>
    <div class="card-name">
        <p class="name">{{ userInfo.name }}</p>
        <p class="sign">{{ userInfo.sign }}</p>
    </div>

    <div ref="scrollBox" class="video-list" @scroll="scrollList">
        <video-list :videos="videoList"></video-list>
        <div v-show="loading" class="loading">
            <n-spin></n-spin>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from "vue";
import { NButton, NSpin, useMessage } from 'naive-ui';
import { useRouter } from "vue-router";
import {  statusCode, storageData } from "@leaf/utils";
import type { UserInfoType, VideoType } from '@leaf/apis';
import { getUploadVideoAPI, getFollowDataAPI } from '@leaf/apis';
import { CommonAvatar } from "@leaf/components";
import HeaderBar from "@/components/header-bar/Index.vue";
import VideoList from '@/components/video-list/Index.vue'

const router = useRouter();
const message = useMessage();

const userData = reactive({
    videoCount: 0,
    followingCount: 0,
    followerCount: 0
})

const userInfo = ref<UserInfoType>({
    uid: 0,
    name: "",
    avatar: ""
});

//前往关注和粉丝页面
const goPage = (name: string) => {
    router.push({ name: name });
}

//获取关注数和粉丝数
const getFollowData = (id: number) => {
    getFollowDataAPI(id).then((res) => {
        if (res.data.code === statusCode.OK) {
            userData.followerCount = res.data.data.follower;
            userData.followingCount = res.data.data.following;
        }
    })
}

//视频相关
let page = 1;
const pageSize = 10;
let noMore = false;
const loading = ref(false);
const videoList = ref<Array<VideoType>>([]);
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
                    getMyVideo();
                }
            }
            timer = null;
        }, 500);
    }
})();

//获取我的视频
const getMyVideo = () => {
    getUploadVideoAPI(page, pageSize).then((res) => {
        if (res.data.code === statusCode.OK) {
            if (res.data.data.videos) {
                videoList.value.push(...res.data.data.videos);
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
    }).catch(() => {
        message.error('获取失败');
    });
}

onMounted(() => {
    userInfo.value = storageData.get("user_info");
    getFollowData(userInfo.value.uid);
    getMyVideo();
})
</script>

<style lang="less" scoped>
.user-card {
    display: flex;
    height: 120px;
    align-items: flex-start;

    .card-avatar {
        width: 160px;

        span {
            margin: 16px 30px;
        }
    }

    .card-info-box {
        .card-info {
            width: 210px;
            display: flex;
            justify-content: space-around;

            div {
                width: 70px;
                text-align: center;

                .info-title {
                    margin-bottom: 2px;
                    // font-weight: bold;
                }

                .info-content {
                    margin: 2px 0 6px 0;
                }
            }
        }

        button {
            width: 100%;
        }
    }
}

.card-name {
    margin: 0 10px;

    .name {
        margin: 0;
        font-size: 18px;
        font-weight: 500;
    }

    .sign {
        margin: 3px 0;
        font-size: 14px;
        color: #666;
    }
}

.video-list {
    height: calc(100vh - 200px);
    overflow: scroll;

    .loading {
        padding: 20px 0;
        text-align: center;
    }
}
</style>