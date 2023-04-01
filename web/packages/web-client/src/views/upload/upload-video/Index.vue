<template>
    <div class="upload-video">
        <n-steps class="step" :current="current" :status="currentStatus">
            <n-step :title="t('upload.uploadVideoInfo')" />
            <n-step :title="t('upload.uploadVideoFile')" />
            <n-step :title="t('upload.videoReview')" />
            <n-step :title="t('upload.completeUpload')" />
        </n-steps>
        <div class="upload-center">
            <upload-video-info v-if="current === 1" :info="videoInfo" @finish="infoFinish" />
            <upload-video :vid="videoInfo.vid" :resources="videoInfo.resources" v-else-if="current === 2" />
            <n-result v-else class="result" :title="toStatusText(videoInfo.status)"
                :status="toStatusIcon(videoInfo.status)">
                <template #footer v-if="videoInfo.status === reviewCode.WRONG_VIDEO_CONTENT ||
                videoInfo.status === reviewCode.WRONG_VIDEO_INFO">
                    <n-button @click="modify">{{ t("common.modify") }}</n-button>
                </template>
            </n-result>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n';
import { onBeforeMount, ref } from 'vue';
import { NStep, NSteps } from "naive-ui";
import { useRoute } from 'vue-router';
import { NResult, NButton, useNotification } from 'naive-ui';
import UploadVideoInfo from './component/UploadVideoInfo.vue';
import { statusCode, reviewCode } from '@leaf/utils';
import { getVideoStatusAPI } from '@leaf/apis';
import type { VideoStatusType } from '@leaf/apis';
import UploadVideo from './component/UploadVideo.vue';

// i18n
const { t } = useI18n();

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
    partition: 0,
    resources: [],
    created_at: ""
});


const infoFinish = (vid: number) => {
    videoInfo.value.vid = vid;
    current.value = 2;
}

// 获取状态文本
const toStatusText = (state: number) => {
    switch (state) {
        case reviewCode.VIDEO_PROCESSING:
            return t('review.videoProcessing');
        case reviewCode.WAITING_REVIEW:
            return t('review.waitingReview');
        case reviewCode.AUDIT_APPROVED:
            return t('review.approved');
        case reviewCode.WRONG_VIDEO_INFO:
            return t('review.wrongVideoInfo');
        case reviewCode.WRONG_VIDEO_CONTENT:
            return t('review.wrongVideoContent');
        case reviewCode.PROCESSING_FAIL:
            return t('review.processingFail');
        default:
            return t('common.unknown');
    }
}

// 状态图标
const toStatusIcon = (state: number) => {
    switch (state) {
        case reviewCode.VIDEO_PROCESSING:
        case reviewCode.WAITING_REVIEW:
            return "info";
        case reviewCode.AUDIT_APPROVED:
            return "success";
        case reviewCode.WRONG_VIDEO_INFO:
        case reviewCode.WRONG_VIDEO_CONTENT:
        case reviewCode.PROCESSING_FAIL:
            return "error";
        default:
            return "info";
    }
}

// 修改
const modify = () => {
    if (videoInfo.value.status === reviewCode.WRONG_VIDEO_CONTENT) {
        current.value = 2;
    } else if (videoInfo.value.status === reviewCode.WRONG_VIDEO_INFO) {
        current.value = 1;
    }
}

onBeforeMount(() => {
    const vid = Number(route.params.vid);
    const modify = (route.query.modify || "").toString();
    if (vid) {
        current.value = 0;
        getVideoStatusAPI(vid).then((res) => {
            if (res.data.code === statusCode.OK) {
                videoInfo.value = res.data.data.video;
                switch (videoInfo.value.status) {
                    case reviewCode.CREATED_VIDEO:
                        current.value = 2;
                        break;
                    case reviewCode.AUDIT_APPROVED:
                        videoInfo.value = res.data.data.video;
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

.result {
    margin-top: 50px;
}
</style>
