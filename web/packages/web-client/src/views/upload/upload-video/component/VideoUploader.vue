<template>
    <div class="video-uploader">
        <n-upload :key="uploaderKey" multiple directory-dnd :show-file-list="false" accept="video/mp4"
            @before-upload="beforeUploadVideo" :custom-request="handleChange">
            <n-upload-dragger>
                <div v-if="!uploading">
                    <div style="margin-bottom: 12px">
                        <n-icon size="48" :depth="3">
                            <upload />
                        </n-icon>
                    </div>
                    <n-text style="font-size: 16px">
                        点击或拖拽视频到此处上传视频
                    </n-text>
                    <n-p depth="3" style="margin: 8px 0 0 0">
                        上传文件大小需小于{{ globalConfig.maxVideoSize }}M,仅支持.mp4格式文件
                    </n-p>
                </div>
                <n-progress v-else type="circle" :percentage="percent" />
            </n-upload-dragger>
        </n-upload>
    </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { Upload } from "@leaf/icons";
import { globalConfig } from "@leaf/utils";
import type { UploadCustomRequestOptions } from "naive-ui";
import { NIcon, NUpload, NUploadDragger, NText, NP, NProgress, useNotification } from 'naive-ui';
import { uploadFileAPI } from "@leaf/apis";

const emits = defineEmits(["finish"]);
const props = defineProps<{
    vid: number
}>();


const percent = ref(0);//上传百分比
const uploading = ref(false);//是否上传中
const notification = useNotification();//通知

//上传之前的回调
const beforeUploadVideo = async (options: any) => {
    const file = options.file;
    const isJpgOrPng = file.type === "video/mp4";
    if (!isJpgOrPng) {
        notification.error({
            title: '上传失败',
            content: "文件只支持mp4格式",
            duration: 5000,
        });
    }
    const isLtMaxSize = file.file.size / 1024 / 1024 < globalConfig.maxVideoSize;
    if (!isLtMaxSize) {
        notification.error({
            title: '上传失败',
            content: `视频大小不能超过${globalConfig.maxVideoSize}M`,
            duration: 5000,
        });
    }
    return isJpgOrPng && isLtMaxSize;
}

//上传变化的回调
const handleChange = ({ file }: UploadCustomRequestOptions) => {
    if (!file.file) return;
    uploading.value = true;
    uploadFileAPI({
        name: "video",
        action: `v1/upload/video/${props.vid}`,
        file: file.file,
        onProgress: (val: any) => {
            changeUpload("uploading", val);
        },
        onError: () => {
            changeUpload("error");
            uploading.value = false;
        },
        onFinish: (data?: any) => {
            changeUpload("success", data);
            uploading.value = false;
        },
    })
}

//上传变化的回调
const uploaderKey = ref(0);
const changeUpload = (status: string, data?: any) => {
    uploaderKey.value = Date.now();
    switch (status) {
        case "success":
            emits("finish");
            notification.success({
                title: '上传完成',
                duration: 3000,
            });
            break;
        case "uploading":
            percent.value = data;
            break;
        case "error":
            notification.error({
                title: '文件上传失败',
                duration: 5000,
            });
            break;
    }
}
</script>

<style lang="less" scoped>
.video-uploader {
    width: 350px;
    margin: 50px auto;
}
</style>