<template>
    <div class="upload-video">
        <video-uploader :vid="vid" @finish="finishUpload"></video-uploader>
        <!-- <div class="video-box">
            <n-scrollbar style="max-height: 300px;">
                <div class="video-item" v-for="(item, index) in uploadVideoList">
                    <div class="item-left">
                        <n-icon class="item-icon">
                            <Videocam />
                        </n-icon>
                        <span v-if="!item.modify" @click="titleClick(item, index)">P{{ index + 1 }} {{ item.title
                        }}</span>
                        <n-input-group v-else>
                            <n-input ref="titleInput" v-model:value="modifyForm.title" maxlength="50" show-count
                                @blur="item.modify = false" />
                            !-- 使用mousedown触发而不是click触发，因为input的blur要早于click --
                            <n-button style="width: 30%;" type="primary" @mousedown="modifyTitle(item)">修改</n-button>
                        </n-input-group>
                        <n-tag class="tag" :type="toTagType(item.review)">{{ toTagText(item.review) }}</n-tag>
                    </div>
                    <div class="item-right" @click="deleteResource(item.id, index)">
                        <span class="delete">删除</span>
                    </div>
                </div>
            </n-scrollbar>
        </div> -->
        <div class="upload-next-btn">
            <n-button type="primary" @click="submitReview">提交审核</n-button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onBeforeMount, reactive, nextTick, ref } from "vue";
// import { submitReviewAPI } from '@/api/review';
// import { deleteResourceAPI, modifyTitleAPI, getResourceListAPI } from "@/api/video";
import VideoUploader from "./VideoUploader.vue";
import { NScrollbar, NInput, NInputGroup, NButton, NIcon, NTag, useNotification } from "naive-ui";
// import { Videocam } from "@vicons/ionicons5";
// import { baseResourceType, reviewResourceType } from "@/types/resource";

const props = defineProps<{
    vid: number
}>();

const vid = props.vid;
// const uploadVideoList = ref<Array<reviewResourceType>>([]);
const notification = useNotification();//通知

const toTagType = (state: number) => {
    switch (state) {
        case 2000:
            return "success";
        case 200: case 1000:
            return "info";
        default:
            return "error";
    }
}

const toTagText = (state: number) => {
    switch (state) {
        case 200:
            return "视频处理中";
        case 1000:
            return "等待审核";
        case 2000:
            return "审核通过";
        case 3100:
            return "视频信息存在问题";
        case 3200:
            return "视频内容存在问题";
        case 3300:
            return "视频处理失败";
        default:
            return "未知";
    }
}

const finishUpload = () => {
    getResourceList();
}

const getResourceList = () => {
    // getResourceListAPI(vid!).then((res) => {
    //     if (res.data.code === 2000) {
    //         uploadVideoList.value = res.data.data.resources;
    //     }
    // }).catch((err) => {
    //     notification.error({
    //         title: '获取失败',
    //         content: "原因:" + err.response.data.msg,
    //         duration: 5000,
    //     })
    // });
}

const submitReview = () => {
    // submitReviewAPI(vid!).then((res) => {
    //     if (res.data.code === 2000) {
    //         notification.success({
    //             title: '提交成功',
    //             duration: 5000,
    //         })
    //     }
    // }).catch((err) => {
    //     notification.error({
    //         title: '提交失败',
    //         content: "原因:" + err.response.data.msg,
    //         duration: 5000,
    //     })
    // });
}

const deleteResource = (id: number, index: number) => {
    // deleteResourceAPI(id).then((res) => {
    //     if (res.data.code === 2000) {
    //         uploadVideoList.value.splice(index, 1);
    //     }
    // }).catch(() => {
    //     notification.error({
    //         title: '删除失败',
    //         duration: 5000,
    //     })
    // });
}

//修改资源名
// const titleInput = ref<Array<HTMLElement>>([]);
// const modifyForm = reactive<baseResourceType>({
//     id: 0,
//     title: '',
// });

//点击标题
// const titleClick = (resource: reviewResourceType, index: number) => {
//     modifyForm.id = resource.id;
//     modifyForm.title = resource.title;
//     resource.modify = true;
//     nextTick(() => {
//         titleInput.value[index].focus();
//     })
// }

//修改标题
// const modifyTitle = (resource: reviewResourceType) => {
//     modifyTitleAPI(modifyForm).then((res) => {
//         if (res.data.code === 2000) {
//             resource.title = modifyForm.title;
//             notification.success({
//                 title: '修改成功',
//                 duration: 5000,
//             })
//         }
//     }).catch(() => {
//         notification.error({
//             title: '修改失败',
//             duration: 5000,
//         })
//     });
// }

onBeforeMount(() => {
    getResourceList();
})


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
                color: #18a058;
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
                    color: #18a058;
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