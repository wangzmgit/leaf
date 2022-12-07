<template>
    <div class="leaf-upload">
        <input class="upload-input" ref="inputRef" type="file" @change="checkFile">
        <div class="upload-trigger" @click="onClickTrigger">
            <slot></slot>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import type { UploadOptions } from "@leaf/apis";

const props = withDefaults(defineProps<{
    name: string
    action: string
    beforeUpload?: (file: File) => Boolean
    upload?: (options: UploadOptions) => void
}>(), {
    name: "file",
    beforeUpload: () => {
        return true;
    }
});

const emits = defineEmits(["stateChange"]);

const inputRef = ref<HTMLInputElement | null>(null);
const onClickTrigger = () => {
    inputRef.value?.click();
}

const uploadStatus = ref<"uploading" | "success" | "error">("uploading")
const checkFile = (e: Event) => {
    const target = (e.target as HTMLInputElement)
    const files = target.files;
    if (files) {
        const uploadFile = Array.from(files)[0];
        const validateFormat = props.beforeUpload(uploadFile);
        if (!validateFormat) return;
        if (props.upload) {
            props.upload({
                action: props.action,
                name: props.name,
                file: uploadFile,
                onProgress: onProgress,
                onFinish: onFinish,
                onError: onError
            })
        }
    }
}

const onProgress = (progress: number) => {
    console.log('progress', progress)
    emits("stateChange", uploadStatus.value, progress)
}

const onFinish = (data?: any) => {
    uploadStatus.value = "success";
    emits("stateChange", uploadStatus.value, data)
}

const onError = () => {
    uploadStatus.value = "error";
    emits("stateChange", uploadStatus.value)
}
</script>

<style lang="less" scoped>
.leaf-upload {
    .upload-input {
        display: none;
    }
}
</style>