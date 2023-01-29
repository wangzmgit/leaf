<template>
    <div class="msg-container" :style="themeStyle()" @scroll="lazyLoading">
        <div class="msg-item" v-for="(item, index) in replyMessageList" :key="index">
            <div class="item-left">
                <common-avatar :url="item.user.avatar" :size="55"></common-avatar>
            </div>
            <div class="item-center">
                <p class="title">
                    <span class="user-name" @click="goUserSpace(item.user.uid)">{{ item.user.name }}</span>
                    <span v-if="item.root_content"> 回复了你的评论</span>
                    <span v-else> 对你的视频发表评论</span>
                </p>
                <p class="content">{{ item.content }}</p>
                <p class="target-content" v-if="item.target_reply_content">{{ item.target_reply_content }}</p>
                <n-time class="msg-time" :time="new Date(item.created_at)"></n-time>
            </div>
            <div class="item-right">
                <div v-if="item.root_content">{{ item.root_content }}</div>
                <img v-else :src="getResourceUrl(item.video.cover)" alt="封面" :title="item.video.title"
                    @click="goVideo(item.video.vid)">
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useRouter } from "vue-router";
import { onBeforeMount, ref } from 'vue';
import { NTime } from 'naive-ui';
import { CommonAvatar } from "@leaf/components";
import { getResourceUrl, statusCode } from '@leaf/utils';
import type { ReplyMessageType } from '@leaf/apis';
import { getReplyMessageAPI } from '@leaf/apis';
import { getTheme } from '@/theme';


const router = useRouter();

const themeStyle = () => {
    const theme = getTheme();
    return {
        "--hover-color": theme.primaryColor
    }
}

//进入用户空间
const goUserSpace = (uid: number) => {
    let userUrl = router.resolve({ name: "User", params: { uid: uid } });
    window.open(userUrl.href, '_blank');
}

//进入视频
const goVideo = (vid: number) => {
    let userUrl = router.resolve({ name: "Video", params: { vid: vid } });
    window.open(userUrl.href, '_blank');
}



let page = 1;
let pageSize = 10;
let noMore = false;
let loading = false;
const replyMessageList = ref<Array<ReplyMessageType>>([]);
const getAtMessage = () => {
    loading = true;
    getReplyMessageAPI(page, pageSize).then((res) => {
        if (res.data.code === statusCode.OK) {
            if (res.data.data.messages.length < pageSize) {
                noMore = true;
            }
            replyMessageList.value = replyMessageList.value.concat(res.data.data.messages);
        }
        loading = false;
    }).catch(() => {
        loading = false;
    })
}

const lazyLoading = (() => {
    var timer: number | null = null;
    return (e: Event) => {
        if (timer) return;
        const target = e.target as HTMLElement;
        if (target.scrollTop + target.clientHeight >= (target.scrollHeight - 10)) {
            if (!noMore && !loading) {
                page++;
                getAtMessage();
            }
        }

        timer = setTimeout(() => {
            timer = null;
        }, 100);
    }
})();

onBeforeMount(() => {
    getAtMessage();
})
</script>

<style lang="less" scoped>
.msg-container {
    height: 620px;
    overflow-y: auto;
    margin: 16px 20px 30px;

    &::-webkit-scrollbar {
        width: 6px;
    }

    &::-webkit-scrollbar-thumb {
        background-color: #e4e6eb;
        outline: none;
        border-radius: 2px;
    }

    &::-webkit-scrollbar-track {
        box-shadow: none;
        border-radius: 2px;
    }
}

.msg-item {
    display: flex;
    min-height: 60px;
    padding: 8px 0 16px;
    margin-right: 10px;
    border-bottom: 1px solid #e1e2e3;

    .item-left {
        width: 60px;
        margin-right: 10px;
    }

    .item-center {
        flex: 1;
        width: calc(100% - 170px);

        .title {
            font-size: 14px;
            color: #212121;
            line-height: 18px;
            margin: 2px 0 12px;

            .user-name {
                cursor: pointer;
                font-weight: 600;

                &:hover {
                    color: var(--hover-color);
                }
            }
        }

        .content {
            color: #222;
            font-size: 14px;
            margin: 10px 0 6px;
        }

        .target-content {
            color: #999;
            font-size: 12px;
            margin: 5px 0;
            padding-left: 6px;
            border-left: 2px solid #e7e7e7;
        }

        .msg-time {
            font-size: 12px;
            color: #999;
        }
    }

    .item-right {
        width: 90px;
        height: 100%;
        color: #666;
        overflow: hidden;
        text-overflow: ellipsis;
        display: -webkit-box;
        -webkit-line-clamp: 3;
        -webkit-box-orient: vertical;

        img {
            cursor: pointer;
            width: 100%;
            height: 100%;
            border-radius: 2px;
        }
    }
}
</style>