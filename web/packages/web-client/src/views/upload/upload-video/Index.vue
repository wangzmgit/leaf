<template>
    <div class="upload-video">
        <n-steps class="step" :current="current" :status="currentStatus">
            <n-step title="视频信息" />
            <n-step title="上传视频" />
            <n-step title="审核" />
            <n-step title="上传成功" />
        </n-steps>
        <div class="upload-center">
            <upload-video-info v-if="current === 1" :info="videoInfo" @finish="infoFinish"></upload-video-info>
            <!--<upload-video :vid="videoInfo.vid" v-else-if="current === 2"></upload-video> -->
        </div>

    </div>
</template>

<script setup lang="ts">
import { onBeforeMount, ref } from 'vue';
import { NStep, NSteps } from "naive-ui";
import { useRoute } from 'vue-router';
import { useNotification } from 'naive-ui';
import UploadVideoInfo from './component/UploadVideoInfo.vue';
import { statusCode, reviewCode } from '@leaf/utils';
import { getVideoStatusAPI } from '@leaf/apis';
import type { VideoStatusType } from '@leaf/apis';
// import UploadVideo from './components/UploadVideo.vue';

const route = useRoute();
const notification = useNotification();//通知

// 当前步骤
const current = ref(1);
const currentStatus = ref<"wait" | "error" | "finish" | "process">("process");

const videoInfo = ref<VideoStatusType>({
    vid: 0,
    status: 0,
    title: "",
    cover: "",
    desc: "",
    copyright: false,
    partition: 0
});


const infoFinish = (vid: number) => {
    videoInfo.value.vid = vid;
    current.value = 2;
}

onBeforeMount(() => {
    const vid = Number(route.params.vid);
    const modify = (route.query.modify || "").toString();
    if (vid) {
        getVideoStatusAPI(vid).then((res) => {
            if (res.data.code === statusCode.OK) {
                videoInfo.value = res.data.data.video;
                switch (videoInfo.value.status) {
                    case reviewCode.CREATED_VIDEO:
                        current.value = 2;
                        break;
                    case reviewCode.AUDIT_APPROVED:
                        if (modify === "info") {
                            current.value = 1;
                        } else if (modify === "video") {
                            current.value = 2;
                        }
                        break;
                    default:
                        current.value = 4; //默认结果页
                        currentStatus.value = "finish";
                        break;
                }
            } else {
                notification.error({
                    title: '获取失败',
                    content: res.data.msg,
                    duration: 5000,
                })
            }
        }).catch(() => {
            notification.error({
                title: '获取失败',
                duration: 5000,
            })
        });
    }
})
</script>

<style lang="less" scoped>
.step {
    width: calc(100% - 100px);
    margin-left: 100px;
    padding-top: 30px;
}
</style>
