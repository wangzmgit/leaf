<template>
    <div class="msg-container" :style="themeStyle" @scroll="lazyLoading">
        <div class="msg-item" v-for="(item, index) in atMessageList" :key="index">
            <div class="item-left">
                <common-avatar :url="item.user.avatar" :size="55"></common-avatar>
            </div>
            <div class="item-center">
                <p class="title">
                    <span class="user-name" @click="goUserSpace(item.user.uid)">{{ item.user.name }}</span>
                    <span> 在评论中提到了你</span>
                </p>
                <n-time class="msg-time" :time="new Date(item.created_at)"></n-time>
            </div>
            <div class="item-right">
                <img :src="getResourceUrl(item.video.cover)" alt="封面" :title="item.video.title"
                    @click="goVideo(item.video.vid)">
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useRouter } from "vue-router";
import { computed, onBeforeMount, ref } from 'vue';
import { NTime } from 'naive-ui';
import { CommonAvatar } from "@leaf/components";
import { statusCode } from '@leaf/utils';
import type { AtMessageType } from '@leaf/apis';
import { getAtMessageAPI } from '@leaf/apis';
import { getResourceUrl } from '@leaf/utils';
import { getTheme } from '@/theme';


const router = useRouter();

const themeStyle = computed(() => {
    const theme = getTheme();
    return {
        "--hover-color": theme.primaryColor
    }
})

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
const atMessageList = ref<Array<AtMessageType>>([]);
const getAtMessage = () => {
    loading = true;
    getAtMessageAPI(page, pageSize).then((res) => {
        if (res.data.code === statusCode.OK) {
            if (res.data.data.messages.length < pageSize) {
                noMore = true;
            }
            atMessageList.value = atMessageList.value.concat(res.data.data.messages);
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
    height: 60px;
    padding: 8px 0 16px;
    margin-right: 10px;
    border-bottom: 1px solid #e1e2e3;

    .item-left {
        float: left;
        width: 60px;
        margin-right: 10px;
    }

    .item-center {
        float: left;
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

        .msg-time {
            font-size: 12px;
            color: #999;
        }
    }

    .item-right {
        float: left;
        width: 90px;
        height: 100%;
        text-align: center;

        img {
            cursor: pointer;
            width: 100%;
            height: 100%;
            border-radius: 2px;
        }
    }
}
</style>