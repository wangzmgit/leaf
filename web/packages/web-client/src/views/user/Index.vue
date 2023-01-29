<template>
    <div class="space" v-title :data-title="`${userInfo.name}的个人中心`">
        <header-bar></header-bar>
        <div class="space-container">
            <div class="space-header">
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
                    <div class="card-btn">
                        <n-button size="small" type="default" @click="goWhisper">私信</n-button>
                        <n-button size="small" v-if="isFollow" type="primary" @click="followClick">已关注</n-button>
                        <n-button size="small" v-else type="error" @click="followClick">关注</n-button>
                    </div>
                </div>
            </div>
            <div class="user-menu">
                <n-menu class="menu" mode="horizontal" :options="menuOptions" :default-value="defaultOption" />
                <div class="menu-right">
                    <div>
                        <p class="data-title">关注</p>
                        <p class="data-content">{{ userData.followingCount }}</p>
                    </div>
                    <div>
                        <p class="data-title">粉丝</p>
                        <p class="data-content">{{ userData.followerCount }}</p>
                    </div>
                </div>
            </div>
            <div class="user-content">
                <router-view></router-view>
            </div>
        </div>
    </div>
</template>


<script setup lang="ts">
import { h, ref, onBeforeMount, reactive } from "vue";
import { RouterLink, useRoute, useRouter } from 'vue-router';
import { Male, Female } from "@leaf/icons";
import { getResourceUrl, statusCode } from "@leaf/utils";
import { CommonAvatar } from "@leaf/components";
import { getFollowDataAPI, getOtherUserInfoAPI, type UserInfoType } from "@leaf/apis";
import { NIcon, NMenu, useNotification, NButton } from "naive-ui";
import HeaderBar from "@/components/header-bar/Index.vue";
import useUserFollow from "@/hooks/user-follow-hooks";

const route = useRoute();
const router = useRouter();
const notification = useNotification();

const defaultOption = ref('');//默认激活菜单
const menuOptions = [
    {
        label: () =>
            h(
                RouterLink,
                {
                    to: {
                        name: "UserVideo",
                    }
                },
                { default: () => `${getGender()}的作品` }
            ),
        key: "video",
    },
    {
        label: () =>
            h(
                RouterLink,
                {
                    to: {
                        name: "UserFollowing",
                    }
                },
                { default: () => '关注' }
            ),
        key: "following",
    },

    {
        label: () =>
            h(
                RouterLink,
                {
                    to: {
                        name: "UserFollower",
                    }
                },
                { default: () => '粉丝' }
            ),
        key: "follower",
    },
];

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

//获取性别
const getGender = () => {
    switch (userInfo.value.gender) {
        case 1:
            return '他';
        case 2:
            return '她';
        default:
            return 'TA';
    }
}

// 获取用户信息
const getUserInfo = (uid: number) => {
    getOtherUserInfoAPI(uid).then((res) => {
        if (res.data.code === statusCode.OK) {
            userInfo.value = res.data.data.user;
            document.title = `${userInfo.value.name}的空间`;
        } else {
            notification.error({
                title: '获取用户信息失败',
                duration: 5000,
            })
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

//关注相关
const { isFollow, getFollowStatus, follow, unfollow } = useUserFollow();
const followClick = () => {
    const uid = Number(route.params.uid)
    if (isFollow.value) {
        unfollow(uid);
    } else {
        follow(uid);
    }
}

// 前往私信页面
const goWhisper = () => {
    router.push({ name: "Whisper", query: { fid: userInfo.value.uid } });
}

onBeforeMount(() => {
    const uid = Number(route.params.uid);
    getUserInfo(uid);
    getFollowData(uid);
    getFollowStatus(uid);
    switch (route.name) {
        case 'UserFollowing':
            defaultOption.value = 'following';
            break;
        case 'UserFollowers':
            defaultOption.value = 'follower';
            break;
        default:
            defaultOption.value = 'video';
            break;
    }
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

        .card-btn {
            width: 130px;
            display: flex;
            margin-right: 20px;
            justify-content: space-between;
        }
    }
}

.user-menu {
    height: 50px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    box-shadow: 0 0 0 1px #eee;
    background-color: #fff;

    .menu-right {
        width: 160px;
        display: flex;

        div {
            color: #fff;
            width: 70px;
            font-size: 12px;
            text-align: center;

            .data-title {
                margin: 0;
                color: #999;
            }

            .data-content {
                margin: 0 0 5px 0;
                color: #222;
            }
        }
    }
}

.user-content {
    margin-top: 10px;
    min-height: 300px;
    background-color: #fff;
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