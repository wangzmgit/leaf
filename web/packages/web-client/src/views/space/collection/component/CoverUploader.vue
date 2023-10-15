<template>
    <div class="upload-cover">
        <n-upload multiple directory-dnd :show-file-list="false" @before-upload="beforeUploadCover"
            :custom-request="fileChange">
            <img v-if="currentCover" class="cover" :src="getResourceUrl(currentCover)" alt="封面" />
            <n-upload-dragger v-else>
                <div v-if="!uploading">
                    <div style="margin-bottom: 12px">
                        <n-icon size="48" :depth="3">
                            <upload />
                        </n-icon>
                    </div>
                    <n-text style="font-size: 16px">
                        点击或拖拽图片到此处上传封面
                    </n-text>
                    <n-p depth="3" style="margin: 8px 0 0 0">
                        上传文件大小需小于{{ maxImgSize }}M,仅支持.jpg .jpeg .png格式文件
                    </n-p>
                </div>
                <n-progress v-else type="circle" :percentage="percent" />
            </n-upload-dragger>
        </n-upload>
        <leaf-cropper ref="cropperRef">
            <template #content="fileSlot">
                <cover-cropper :file="fileSlot.file" @state-change="changeUpload"></cover-cropper>
            </template>
        </leaf-cropper>
    </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { Upload } from "@leaf/icons";
import { getResourceUrl, globalConfig } from "@leaf/utils";
import LeafCropper from "@/components/leaf-cropper/Index.vue";
import CoverCropper from "@/components/leaf-cropper/component/CoverCropper.vue";
import type { UploadCustomRequestOptions } from 'naive-ui';
import { NIcon, NUpload, NUploadDragger, NText, NP, NProgress, useNotification } from 'naive-ui';

const maxImgSize = window.$maxImgSize || globalConfig.maxImgSize;

const emits = defineEmits(["finish"]);
const props = defineProps<{
    cover?: string
}>()

const currentCover = ref(props.cover);

const percent = ref(0);//上传百分比
const uploading = ref(false);//是否上传中
const notification = useNotification();//通知

//上传之前的回调
const beforeUploadCover = async (options: any) => {
    const file = options.file;
    const isJpgOrPng =
        file.type === "image/jpeg" || file.type === "image/png";
    if (!isJpgOrPng) {
        notification.error({
            title: '上传失败',
            content: "文件只支持jpg/jpeg/png格式",
            duration: 5000,
        });
    }
    const isLtMaxSize = file.file.size / 1024 / 1024 < maxImgSize;
    if (!isLtMaxSize) {
        notification.error({
            title: '上传失败',
            content: `图片大小不能超过${maxImgSize}M`,
            duration: 5000,
        });
    }
    return isJpgOrPng && isLtMaxSize;
}

const cropperRef = ref<InstanceType<typeof LeafCropper> | null>(null);
const fileChange = ({ file }: UploadCustomRequestOptions) => {
    cropperRef.value?.setImgFile(file.file);
    cropperRef.value?.open();
}

//上传变化的回调
const changeUpload = (status: string, data: any) => {
    switch (status) {
        case "success":
            //更新封面图
            currentCover.value = data.data.url;
            emits("finish", currentCover.value);
            notification.success({
                title: '上传完成',
                duration: 3000,
            });
            break;
        case "error":
            notification.error({
                title: '文件上传失败',
                duration: 5000,
            });
            break;
    }
    cropperRef.value?.closeCropper();
}
</script>

<style lang="less" scoped>
.upload-cover {
    width: 350px;
    margin: 20px auto;

    .cover {
        width: 350px;
        height: 200px;
    }

}
</style>