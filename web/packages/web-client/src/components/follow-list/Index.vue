<template>
    <div class="follow-list" :style="`height:${height}px`" @scroll="lazyLoading">
        <div class="follow-card" v-for="(item, index) in followList" :key="index">
            <!--头像-->
            <div class="follow-avatar">
                <common-avatar :url="item.avatar" :size="60"></common-avatar>
            </div>
            <!--昵称和个签-->
            <span class="follow-name" @click="goUserSpace(item.uid)">{{ item.name }}</span>
            <span class="follow-sign">{{ item.sign }}</span>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useRouter } from "vue-router";
import { onBeforeMount, ref, reactive } from "vue";
import { CommonAvatar } from "@leaf/components";
import { statusCode } from "@leaf/utils";
import type { UserInfoType } from '@leaf/apis';
import { getFollowingAPI, getFollowersAPI } from '@leaf/apis';
import { useNotification } from "naive-ui";


const props = withDefaults(defineProps<{
    uid: number,
    following?: boolean,
    height?: number
}>(), {
    uid: 0,
    following: false,
    height: 600
})

const pageInfo = reactive({
    current: 1,
    pageSize: 9
});
const noMore = ref(false);//没有更多
const loading = ref(false);//加载中
const followList = ref<Array<UserInfoType>>([]);
const router = useRouter();
const notification = useNotification();//通知

//进入用户空间
const goUserSpace = (uid: number) => {
    let userUrl = router.resolve({ name: "User", params: { uid: uid } });
    window.open(userUrl.href, '_blank');
}

//获取关注
const getFollowingList = () => {
    getFollowingAPI(props.uid, pageInfo.current, pageInfo.pageSize).then((res) => {
        if (res.data.code === statusCode.OK) {
            const resUser = res.data.data.users;
            if (resUser) {
                followList.value.push(...resUser);
            }
            if (!resUser || resUser.length < pageInfo.pageSize) {
                noMore.value = true;
            }
            loading.value = false;
        } else {
            notification.error({
                title: '获取失败',

                duration: 5000,
            })
        }
    })
}

//获取粉丝
const getFollowersList = () => {
    getFollowersAPI(props.uid, pageInfo.current, pageInfo.pageSize).then((res) => {
        if (res.data.code === statusCode.OK) {
            const resUser = res.data.data.users;
            if (resUser) {
                followList.value.push(...resUser);
            }
            if (!resUser || resUser.length < pageInfo.pageSize) {
                noMore.value = true;
            }
            loading.value = false;
        } else {
            notification.error({
                title: '获取失败',

                duration: 5000,
            })
        }
    })
}

const lazyLoading = (() => {
    var timer: number | null = null;
    return (e: Event) => {
        if (timer) return;
        const target = e.target as HTMLElement;
        if (target.scrollTop + target.clientHeight >= (target.scrollHeight - 10)) {
            if (!noMore.value && !loading.value) {
                pageInfo.current++;
                loading.value = true;
                if (props.following) {
                    getFollowingList();
                } else {
                    getFollowersList();
                }
            }
        }

        timer = setTimeout(() => {
            timer = null;
        }, 100);
    }
})();

onBeforeMount(() => {
    if (props.following) {
        getFollowingList();
    } else {
        getFollowersList();
    }
})
</script>

<style lang="less" scoped>
.follow-list {
    overflow-y: scroll;

    /**修改滚动条样式 */
    &::-webkit-scrollbar {
        width: 6px;
    }

    &::-webkit-scrollbar-thumb {
        /*滚动条里面小方块*/
        border-radius: 3px;
        background: #a3a3a3;
    }

    &::-webkit-scrollbar-track {
        /*滚动条里面轨道*/
        border-radius: 5px;
    }
}

.follow-card {
    height: 70px;
    position: relative;
    border-bottom: 1px solid #d1d1d1;

    /**移除最后一个的底部边框 */
    &:last-child {
        border-bottom: none;
    }
}

.follow-avatar>span {
    margin-top: 5px;
}

.follow-name {
    color: #333;
    position: absolute;
    top: 10px;
    left: 80px;
    font-weight: 600;
    cursor: pointer;
}

.follow-sign {
    position: absolute;
    top: 38px;
    left: 80px;
    font-size: 12px;
    color: #666;
}
</style>