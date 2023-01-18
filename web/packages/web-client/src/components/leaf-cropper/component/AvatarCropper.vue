<template>
    <div class="avatar-cropper">
        <div class="card-left">
            <pic-cropper ref="cropperRef" :key="props.file?.name" :file="props.file" :minHeight="60" :minWidth="60"
                @change="imgChange"></pic-cropper>
        </div>
        <div class="card-right">
            <span class="title">头像预览</span>
            <img :src="previewURL" alt="图像预览" />
            <n-button type="primary" @click="uploadAvatar">裁剪并上传</n-button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import PicCropper from './PicCropper.vue';
import { NButton } from 'naive-ui';
import { uploadFileAPI } from "@leaf/apis";

const props = defineProps<{
    file?: File
}>();

const cropperRef = ref<InstanceType<typeof PicCropper> | null>(null);

const previewURL = ref("");
const imgChange = (url: string) => {
    previewURL.value = url;
}

const emits = defineEmits(["stateChange"])
const uploadAvatar = async () => {
    if (cropperRef.value) {
        const file = await cropperRef.value.getFile();
        await uploadFileAPI({
            name: "image",
            action: "v1/upload/image",
            file: file,
            onProgress: () => { },
            onError: () => {
                emits("stateChange", "error")
            },
            onFinish: (data?: any) => {
                emits("stateChange", "success", data)
            },
        })
    }
}

</script>

<style lang="less" scoped>
.avatar-cropper {
    box-sizing: border-box;
    padding: 10px;
    display: flex;
    height: 280px;

    .card-left {
        width: 260px;
        text-align: center
    }

    .card-right {
        width: 200px;
        display: flex;
        align-items: center;
        justify-content: space-between;
        flex-direction: column;

        img {
            width: 120px;
            height: 120px;
        }
    }
}
</style>