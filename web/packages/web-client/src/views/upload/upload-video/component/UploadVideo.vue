<template>
    <div class="upload-video" :style="initTheme()">
        <video-uploader :vid="vid" @finish="finishUpload"></video-uploader>
        <div class="video-box">
            <n-scrollbar style="max-height: 300px;">
                <div class="video-item" v-for="(item, index) in resourceList">
                    <div class="item-left">
                        <n-icon class="item-icon">
                            <video-icon></video-icon>
                        </n-icon>
                        <span v-if="modifyIndex !== index" @click="titleClick(item, index)">
                            P{{ index + 1 }} {{ item.title }}
                        </span>
                        <n-input-group v-else>
                            <n-input ref="titleInput" v-model:value="modifyForm.title" maxlength="50" show-count
                                @blur="modifyIndex = -1" />
                            <!-- 使用mousedown触发而不是click触发，因为input的blur要早于click -->
                            <n-button style="width: 30%;" type="primary" @mousedown="modifyTitle(item)">
                                {{ t('common.modify') }}
                            </n-button>
                        </n-input-group>
                        <n-tag class="tag" :type="toTagType(item.status)">{{ toTagText(item.status) }}</n-tag>
                    </div>
                    <div class="item-right" @click="deleteResource(item.id, index)">
                        <span class="delete">
                            {{ t('common.delete') }}
                        </span>
                    </div>
                </div>
            </n-scrollbar>
        </div>
        <div class="upload-next-btn">
            <n-button type="primary" @click="submitReview">{{ t('upload.submitReview') }}</n-button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { reactive, nextTick, ref } from "vue";
import { getTheme } from "@/theme";
import VideoUploader from "./VideoUploader.vue";
import { Video as VideoIcon } from "@leaf/icons";
import { getVideoStatusAPI, modifyTitleAPI } from "@leaf/apis";
import { reviewCode, statusCode } from "@leaf/utils";
import { NScrollbar, NInput, NInputGroup, NButton, NIcon, NTag, useNotification } from "naive-ui";
import { deleteResourceAPI, submitReviewAPI, type BaseResourceType, type ResourceType } from "@leaf/apis";

const props = defineProps<{
    vid: number,
    resources: Array<ResourceType>
}>();

// i18n
const { t } = useI18n();

const notification = useNotification();// 通知


// 主题
const initTheme = () => {
    const theme = getTheme();

    return {
        "--primary-color": theme.primaryColor
    }
}

const resourceList = ref<Array<ResourceType>>(props.resources);
const finishUpload = () => {
    getVideoStatusAPI(props.vid).then((res) => {
        if (res.data.code === statusCode.OK) {
            resourceList.value = res.data.data.video.resources;
        } else {
            notification.error({
                title: '获取失败',
                content: res.data.msg,
                duration: 5000,
            })
        }
    })
}

// 获取标签类型
const toTagType = (state: number) => {
    switch (state) {
        case reviewCode.AUDIT_APPROVED:
            return "success";
        case reviewCode.VIDEO_PROCESSING: case reviewCode.WAITING_REVIEW:
            return "info";
        default:
            return "error";
    }
}

// 获取标签文本
const toTagText = (state: number) => {
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

// 提交审核
const submitReview = () => {
    if (resourceList.value.length === 0) {
        notification.error({
            title: '请先上传视频',
            duration: 5000,
        })
        return;
    }
    submitReviewAPI(props.vid).then((res) => {
        if (res.data.code === statusCode.OK) {
            notification.success({
                title: '提交成功',
                duration: 5000,
            })
        } else {
            notification.error({
                title: '提交失败',
                description: res.data.msg,
                duration: 5000,
            })
        }
    })
}

const deleteResource = (id: number, index: number) => {
    deleteResourceAPI(id).then((res) => {
        if (res.data.code === statusCode.OK) {
            resourceList.value.splice(index, 1);
        } else {
            notification.error({
                title: '删除失败',
                duration: 5000,
            })
        }
    })
}

//修改资源名
const modifyIndex = ref(-1);
const titleInput = ref<Array<InstanceType<typeof NInput>>>();
const modifyForm = reactive<BaseResourceType>({
    id: 0,
    title: '',
});

//点击标题
const titleClick = (resource: ResourceType, index: number) => {
    modifyForm.id = resource.id;
    modifyForm.title = resource.title;
    modifyIndex.value = index;
    nextTick(() => {
        if (titleInput.value) {
            titleInput.value[0].focus();
        }
    })
}

//修改标题
const modifyTitle = (resource: ResourceType) => {
    modifyTitleAPI(modifyForm).then((res) => {
        if (res.data.code === statusCode.OK) {
            resource.title = modifyForm.title;
            notification.success({
                title: '修改成功',
                duration: 5000,
            })
        } else {
            notification.error({
                title: '修改失败',
                duration: 5000,
            })
        }
    })
}
</script>

<style lang="less" scoped>
.video-box {
    width: 80%;
    margin: 0 auto;
    padding-bottom: 50px;


    .video-item {
        height: 60px;
        padding: 0 20px;
        display: flex;
        align-items: center;
        justify-content: space-between;
        border: 1px solid #efeff5;
        margin: 10px 20px;

        .item-left {
            display: flex;
            align-items: center;
            justify-content: space-between;

            .item-icon {
                font-size: 26px;
                color: var(--primary-color);
                padding-right: 10px;
            }

            .tag {
                margin-left: 10px;
            }
        }

        .item-right {
            .delete {
                cursor: pointer;
                color: #767c82;

                &:hover {
                    color: var(--primary-color);
                }
            }
        }
    }
}

.upload-next-btn {
    float: right;
    margin: 10px 20px 20px;

    button {
        width: 160px;
        height: 40px;
    }
}
</style>