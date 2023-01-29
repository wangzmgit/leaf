<template>
    <div class="video-item" :style="initTheme()">
        <div class="img" @click="goVideo(info.vid)">
            <img :src="getResourceUrl(info.cover)" alt="封面" />
        </div>
        <div class="video-info">
            <span v-if="!props.keywords" class="title" @click="goVideo(info.vid)">{{ info.title }}</span>
            <span v-else class="title" @click="goVideo(info.vid)" v-dompurify-html="keyHighlight(info.title)"></span>
            <div class="author">
                <div class="avatar">
                    <common-avatar :url="info.author.avatar" :size="26" :iconsize="16"></common-avatar>
                </div>
                <div class="name-date">
                    <span class="name" @click="goUser(info.author.uid)">{{ info.author.name }}</span>
                    <span> · </span>
                    <n-time :time="new Date(info.created_at)" type="relative" />
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import type { VideoType } from '@leaf/apis';
import { NTime } from 'naive-ui';
import { CommonAvatar } from '@leaf/components';
import { getTheme } from '@/theme';
import { useRouter } from 'vue-router';
import { getResourceUrl } from '@leaf/utils';

const props = defineProps<{
    info: VideoType,
    keywords?: string,
}>()

const theme = getTheme();
const router = useRouter();

const initTheme = () => {

    return {
        "--primary-color": theme.primaryColor
    }
}

const goUser = (uid: number) => {
    let videoUrl = router.resolve({ name: "User", params: { uid: uid } });
    window.open(videoUrl.href, '_blank');
}

//前往视频详情
const goVideo = (vid: number) => {
    let videoUrl = router.resolve({ name: "Video", params: { vid: vid } });
    window.open(videoUrl.href, '_blank');
}

//关键词高亮
const keyHighlight = (title: string) => {
    let res = '';
    let indexArr: Array<number> = []; // 需要标红的字的下标数组
    const keywordsArray = props.keywords!.split(" ");
    const getReplaceStr = (str: string) => `<font color="${theme.primaryColor}">${str}</font>`;
    keywordsArray.forEach((keyword) => {
        let filterStr = title;
        let stopFlag = false;
        while (!stopFlag && filterStr && keyword) {
            const index = filterStr.indexOf(keyword); // 返回匹配的第一个字符的下标
            if (index === -1) stopFlag = true;
            else {
                keyword.split("").forEach((s, i) => {
                    indexArr.push(index + Number(i));
                });
                filterStr =
                    filterStr.substring(0, index) +
                    " " +
                    filterStr.substring(index + 1);
            }
        }
    });
    indexArr = Array.from(new Set(indexArr)); // 去重
    title.split("").forEach((char, charIndex) => {
        res += indexArr.includes(charIndex) ? getReplaceStr(char) : char;
    });
    return res;
}
</script>

<style lang="less" scoped>
.video-item {
    width: 100%;
    height: 260px;

    .img {
        width: 100%;
        height: 160px;
        border-radius: 10px;
        overflow: hidden;
        background-color: rgba(0, 0, 0, .2);

        img {
            width: 100%;
            height: 100%;
        }
    }

    .video-info {
        margin-top: 8px;

        .title {
            height: 44px;
            color: #18191c;
            padding-right: 12px;
            font-size: 15px;
            line-height: 22px;
            overflow: hidden;
            text-overflow: ellipsis;
            display: -webkit-box;
            -webkit-line-clamp: 2;
            -webkit-box-orient: vertical;
            font-weight: 500;
            cursor: pointer;

            &:hover {
                color: var(--primary-color);
            }
        }

        .author {
            display: flex;
            align-items: center;
            font-size: 13px;
            color: #9499a0;
            margin-top: 5px;

            .avatar {
                width: 26px;
                height: 26px;
                border-radius: 50%;
                background-color: rgba(0, 0, 0, .2);
            }

            .name-date {
                margin-left: 10px;

                .name {
                    cursor: pointer;

                    &:hover {
                        color: var(--primary-color);
                    }
                }
            }

        }
    }
}
</style>