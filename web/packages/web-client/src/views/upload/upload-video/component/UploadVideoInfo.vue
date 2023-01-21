<template>
    <div class="upload-video-info">
        <cover-uploader class="upload" v-if="showUpload" :cover="videoInfo.cover" @finish="finishUpload" />
        <n-form class="info-form" label-placement="left" label-width="auto">
            <n-form-item :label="t('common.title')">
                <n-input v-model:value="videoInfo.title" placeholder="请输入标题" maxlength="50" show-count />
            </n-form-item>
            <n-form-item :label="t('upload.videoDesc')">
                <n-input v-model:value="videoInfo.desc" placeholder="简单介绍一下视频~" maxlength="200" show-count
                    type="textarea" :autosize="descSize" />
            </n-form-item>
            <n-form-item :label="t('upload.copyright')">
                <n-switch v-model:value="videoInfo.copyright" />
            </n-form-item>
            <n-form-item :label="t('common.partition')">
                <n-input v-if="partitionText" disabled :value="partitionText"></n-input>
                <partition-selector v-else @selected="selectedPartition"></partition-selector>
            </n-form-item>
            <div class="upload-next-btn">
                <n-button v-if="isModify" type="primary" @click="modifyVideoInfo">{{ t('common.modify') }}</n-button>
                <n-button v-else type="primary" @click="uploadInfo">{{ t('common.next') }}</n-button>
            </div>
        </n-form>
    </div>
</template>

<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { ref, onMounted, reactive } from "vue";
import CoverUploader from "./CoverUploader.vue";
import PartitionSelector from "./PartitionSelector.vue";
import { useNotification } from 'naive-ui';
import { NForm, NFormItem, NButton, NInput, NSwitch } from 'naive-ui';//表单相关
import { getPartitionAPI, uploadVideoInfoAPI, modifyVideoInfoAPI } from "@leaf/apis";
import type { VideoStatusType, PartitionType } from '@leaf/apis';
import { statusCode } from "@leaf/utils";

const emits = defineEmits(["finish"]);
const props = defineProps<{
    info: VideoStatusType
}>();

// i18n
const { t } = useI18n();

//简介输入框大小
const descSize = {
    minRows: 3,
    maxRows: 5
}

const notification = useNotification();//通知

const videoInfo = reactive({
    vid: 0,
    title: "",
    cover: "",
    desc: "",
    copyright: true,
    partition: 0,
    created_at: ""
})

const isModify = ref(false);
const showUpload = ref(false);//显示上传
const partitionText = ref("");//分区名称


//封面上传完成
const finishUpload = (cover: string) => {
    videoInfo.cover = cover;
}

//选中分区
const selectedPartition = (value: number) => {
    videoInfo.partition = value;
}

//上传视频信息
const uploadInfo = () => {
    if (!videoInfo.cover) {
        notification.error({
            title: '上传失败',
            content: "请上传视频封面",
            duration: 5000,
        });
        return;
    }
    if (!videoInfo.title) {
        notification.error({
            title: '上传失败',
            content: "请填写视频标题",
            duration: 5000,
        });
        return;
    }
    if (!videoInfo.partition) {
        notification.error({
            title: '上传失败',
            content: "请选择视频分区",
            duration: 5000,
        });
        return;
    }
    uploadVideoInfoAPI(videoInfo).then((res) => {
        if (res.data.code === statusCode.OK) {
            emits('finish', res.data.data.vid);
        }
    })
}

//修改视频信息
const modifyVideoInfo = () => {
    if (!videoInfo.cover) {
        notification.error({
            title: '请上传视频封面',
            duration: 5000,
        });
        return;
    }
    if (!videoInfo.title) {
        notification.error({
            title: '请填写视频标题',
            duration: 5000,
        });
        return;
    }
    modifyVideoInfoAPI(videoInfo).then((res) => {
        if (res.data.code === statusCode.OK) {
            notification.success({
                title: '修改成功',
                duration: 5000,
            });
        }
    })
}

// 获取分区名
const getPartitionName = (id: number) => {
    let partitionList: Array<PartitionType> = [];
    getPartitionAPI().then((res) => {
        if (res.data.code === statusCode.OK) {
            partitionList = res.data.data.partitions;
            const subpartition = partitionList.find((item) => {
                return item.id === id;
            })

            const partition = partitionList.find((item) => {
                return item.id === subpartition?.parent_id;
            })

            partitionText.value = `${partition?.content}/${subpartition?.content}`;
        }
    })
}

onMounted(() => {
    if (props.info.vid) {
        isModify.value = true;
        videoInfo.vid = props.info.vid;
        videoInfo.title = props.info.title;
        videoInfo.desc = props.info.desc;
        videoInfo.cover = props.info.cover;
        videoInfo.copyright = props.info.copyright;
        getPartitionName(props.info.partition);
    }
    showUpload.value = true;
})
</script>

<style lang="less" scoped>
.upload {
    margin: 50px auto;
}

.info-form {
    width: calc(100% - 260px);
    margin-left: 130px;

    .upload-next-btn {
        float: right;
        margin-bottom: 20px;

        button {
            width: 160px;
            height: 40px;
        }
    }
}
</style>