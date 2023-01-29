<template>
    <div class="space" v-title :data-title="`${userInfo.name}的个人中心`">
        <header-bar></header-bar>
        <div class="space-container">
            <div class="space-header">
                <button class="upload-btn" @click="uploadClick">上传封面图片</button>
                <img class="cover" v-if="userInfo.spacecover" :src="getResourceUrl(userInfo.spacecover)" alt="用户封面图" />
                <div class="header-content">
                    <common-avatar :url="userInfo.avatar" :size="70" :iconsize="36"></common-avatar>
                    <div class="content-name">
                        <p class="name">{{ userInfo.name }}
                            <n-icon :color="userInfo.gender === 1 ? '#1890ff' : '#eb2f96'" size="20">
                                <male v-if="userInfo.gender === 1"></male>
                                <female v-else-if="userInfo.gender === 2"></female>
                            </n-icon>
                        </p>
                        <p class="sign">{{ userInfo.sign }}</p>
                    </div>
                    <!--关注粉丝信息-->
                    <div class="user-data">
                        <div>
                            <p class="data-title">{{ t("common.manuscript") }}</p>
                            <p>{{ videoCount }}</p>
                        </div>
                        <div>
                            <p class="data-title">{{ t("common.following") }}</p>
                            <p class="data-content" @click="goPage('Following')">{{ userData.followingCount }}</p>
                        </div>
                        <div>
                            <p class="data-title">{{ t("common.follower") }}</p>
                            <p class="data-content" @click="goPage('Follower')">{{ userData.followerCount }}</p>
                        </div>
                    </div>
                </div>
            </div>
            <div class="space-content">
                <n-menu class="space-menu" :options="menuOptions" :default-value="defaultOption" />
                <div class="space-router">
                    <router-view></router-view>
                </div>
            </div>
        </div>
        <leaf-cropper ref="cropperRef">
            <template #content="fileSlot">
                <space-cover-cropper :file="fileSlot.file" @state-change="changeUpload"></space-cover-cropper>
            </template>
        </leaf-cropper>
    </div>
</template>


<script setup lang="ts">
import { useI18n } from 'vue-i18n';
import { h, ref, onBeforeMount, reactive } from "vue";
import { RouterLink, useRoute, useRouter } from 'vue-router';
import { Video, Collection, Upload, Message, Setting, Male, Female } from "@leaf/icons";
import { getFollowDataAPI, type UserInfoType } from "@leaf/apis";
import { getResourceUrl, statusCode, storageData } from "@leaf/utils";
import { modifySpaceCoverAPI, getUserInfoAPI } from "@leaf/apis";
import useRenderIcon from "@/hooks/render-icon-hooks";
import { CommonAvatar } from "@leaf/components";
import { NIcon, NMenu, useNotification } from "naive-ui";
import HeaderBar from "@/components/header-bar/Index.vue";
import LeafCropper from "@/components/leaf-cropper/Index.vue";
import SpaceCoverCropper from "@/components/leaf-cropper/component/SpaceCoverCropper.vue";
import { storeToRefs } from "pinia";
import { useVideoCountStore } from "@/stores/video-count-store";

const { t } = useI18n();

const route = useRoute();
const router = useRouter();
const { renderIcon } = useRenderIcon();
const notification = useNotification();

const defaultOption = ref('');//默认激活菜单
const menuOptions = [
    {
        label: () =>
            h(
                RouterLink,
                {
                    to: {
                        name: "SpaceVideo",
                    }
                },
                { default: () => t("common.manuscript") }
            ),
        key: "video",
        icon: renderIcon(Video, '#609a8b'),
    },
    {
        label: () =>
            h(
                RouterLink,
                {
                    to: {
                        name: "Collection",
                    }
                },
                { default: () => t("common.collection") }
            ),
        key: "collection",
        icon: renderIcon(Collection, '#e3c0aa'),
    },
    {
        label: () =>
            h(
                RouterLink,
                {
                    to: {
                        name: "Upload",
                    }
                },
                { default: () => t("common.submission") }
            ),
        key: "upload",
        icon: renderIcon(Upload, '#7daebd'),
    },
    {
        label: () =>
            h(
                RouterLink,
                {
                    to: {
                        name: "Message",
                    }
                },
                { default: () => t("common.message") }
            ),
        key: "message",
        icon: renderIcon(Message, '#c79fa7'),
    },
    {
        label: () =>
            h(
                RouterLink,
                {
                    to: {
                        name: "Setting",
                    }
                },
                { default: () => t("common.setting") }
            ),
        key: "setting",
        icon: renderIcon(Setting, '#808080'),
    },
];

const videoCountStore = useVideoCountStore();
const { videoCount } = storeToRefs(videoCountStore);

const userInfo = ref<UserInfoType>({
    uid: 0,
    name: "",
    avatar: "",
    spacecover: ""
});

const userData = reactive({
    followingCount: 0,
    followerCount: 0
})

//前往关注和粉丝页面
const goPage = (name: string) => {
    router.push({ name: name });
}

const cropperRef = ref<InstanceType<typeof LeafCropper> | null>(null);
const uploadClick = () => {
    cropperRef.value?.open();
}

//上传变化的回调
const changeUpload = (status: string, data: any) => {
    switch (status) {
        case "success":
            //更新封面图
            const url = data.data.url;
            modifySpaceCoverAPI(url).then((res) => {
                if (res.data.code = statusCode.OK) {
                    getUserInfo();
                }
            })
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

// 获取用户信息
const getUserInfo = () => {
    getUserInfoAPI().then((res) => {
        if (res.data.code === statusCode.OK) {
            const tmpInfo = res.data.data.user_info as UserInfoType;
            userInfo.value = tmpInfo;
            storageData.update("user_info", tmpInfo);
        }
    })
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

onBeforeMount(() => {
    userInfo.value = storageData.get("user_info");
    switch (route.name) {
        case 'SpaceVideo':
            defaultOption.value = 'video';
            break;
        case 'Collection':
            defaultOption.value = 'collection';
            break;
        case 'InfoSetting': case 'SecuritySetting':
            defaultOption.value = 'setting';
            break;
        default:
            defaultOption.value = 'home';
            break;
    }
    getFollowData(userInfo.value.uid);
})

</script>


<style lang="less">
body {
    background: #f6f7f8;
}
</style>

<style lang="less" scoped>
.space {
    width: 100%;
}

.space-container {
    width: 1100px;
    margin: 0 auto;
}

.space-header {
    position: relative;
    width: 100%;
    height: 220px;
    background-color: #999;

    .upload-btn {
        position: absolute;
        top: 12px;
        right: 20px;
        width: 120px;
        height: 32px;
        color: #d3d3d3;
        background-color: transparent;
        cursor: pointer;
        border-radius: 3px;
        border: 1px solid rgba(255, 255, 255, .2);

        &:hover {
            border-color: #d3d3d3;
        }
    }

    .cover {
        width: 100%;
        height: 100%;
        object-fit: cover;
    }

    .header-content {
        width: 100%;
        height: 100px;
        position: absolute;
        bottom: 0;
        display: flex;
        align-items: center;
        box-sizing: border-box;
        padding: 0 20px;

        .content-name {
            flex: 1;
            padding-left: 20px;

            .name {
                color: #fff;
                font-size: 18px;
                font-weight: 600;
                margin: 0 0;
            }

            .sign {
                font-size: 12px;
                color: #ccd0d7;
                margin: 6px 0 0 0;
            }
        }

        .user-data {
            width: 236px;
            display: flex;

            div {
                color: #fff;
                width: 78px;
                text-align: center;

                .data-title {
                    margin-bottom: 6px;
                    font-weight: bold;
                }

                .data-content:hover {
                    cursor: pointer;
                }
            }
        }
    }
}


.space-content {
    display: flex;
    margin: 10px auto 30px;


    .space-menu {
        width: 200px;
        height: 260px;
        background-color: #fff;
        border: 1px solid #efeff5;
    }

    .space-router {
        margin-left: 10px;
        width: calc(100% - 210px);
        min-height: 630px;
        background-color: #fff;
    }
}

@media (min-width: 1400px) {
    .space-container {
        width: 1300px;
    }

    .space-header {
        height: 260px;
    }
}
</style>