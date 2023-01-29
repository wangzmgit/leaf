<template>
    <div class="comment-item" v-for="(item, index) in commentList" :key="index">
        <!--头像和昵称-->
        <div class="comment-user">
            <common-avatar :size="44" :url="item.author.avatar"></common-avatar>
            <div class="user-name">
                <span @click="goUserSpace(item.author.uid)">{{ item.author.name }}</span>
            </div>
            <n-time class="comment-time" type="relative" :time="new Date(item.created_at)"></n-time>
        </div>
        <!--评论内容-->
        <div class="comment-content">
            <div v-if="item.content.indexOf('@') === -1">
                <span>{{ item.content }}</span>
            </div>
            <div class="content-text" v-else v-for="content in handleMention(item.content)">
                <span :class="content.class">{{ content.value }}</span>
            </div>
            <!--回复-->
            <div v-show="item.reply.length" class="reply-list">
                <div class="reply-item" v-for="(reply, i) in item.reply" :key="i">
                    <span class="reply-user-name" @click="goUserSpace(reply.author.uid)">
                        {{ reply.author.name }}
                    </span>
                    <span>: </span>
                    <span v-if="reply.content.indexOf('@') === -1">
                        {{ reply.content }}
                    </span>
                    <span v-else :class="content.class" v-for="content in handleMention(reply.content)">
                        {{ content.value }}
                    </span>
                </div>
                <div class="more">
                    <n-button :disabled="item.noMore" text @click="getMoreReply(item.id, index)">加载更多</n-button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { onBeforeMount, onBeforeUnmount, ref } from "vue";
import useComment from '@/hooks/comment-hooks';
import useMention from '@/hooks/mention-hooks';
import { CommonAvatar } from '@leaf/components';
import { NButton, NTime, useMessage } from "naive-ui";

const props = defineProps<{
    vid: number
}>();

//处理@
const { handleMention } = useMention();

// 获取评论时查询的回复数
const replyCount = 2;

const message = useMessage();

const { noMore, loadingComment, commentList, getCommentList, getReplyListSync } = useComment();

//页面跳转
const router = useRouter();
const goUserSpace = (uid: number) => {
    let userUrl = router.resolve({ name: "User", params: { uid: uid } });
    window.open(userUrl.href, '_blank');
}

//获取评论回复
const page = ref(1);
//加载更多回复
const getMoreReply = (cid: string, index: number) => {
    const pageSize = 2;
    if (!commentList.value[index].page) {
        commentList.value[index].page = 1;
    }
    //获取回复内容
    getReplyListSync(cid, commentList.value[index].page, pageSize).then((res) => {
        if (res === null || res.length < replyCount) {
            message.info("没有更多了");
            commentList.value[index].noMore = true;
            return;
        }

        commentList.value[index].page += 1;
        commentList.value[index].reply.push(...res);
    })
}

// 加载更多评论
const lazyLoading = () => {
    const scrollTop = document.documentElement.scrollTop || document.body.scrollTop;
    const clientHeight = document.documentElement.clientHeight;
    const scrollHeight = document.documentElement.scrollHeight;
    if (scrollTop + clientHeight >= (scrollHeight - 10)) {
        if (!noMore.value &&!loadingComment.value) {
            page.value++;
            getCommentList(props.vid, page.value, 8);
        }
    }
}

onBeforeMount(() => {
    getCommentList(props.vid, page.value, 8);
    window.addEventListener("scroll", lazyLoading, true);
})

onBeforeUnmount(() => {
    window.removeEventListener("scroll", lazyLoading);
})
</script>

<style lang="less" scoped>
.comment-item {
    margin: 0 10px;

    .comment-user {
        position: relative;

        .user-name {
            position: absolute;
            top: 0;
            left: 60px;
            font-weight: 500;
            cursor: pointer;
            color: #6d757a;
            font-size: 14px;
        }

        .comment-time {
            position: absolute;
            color: #969996;
            top: 26px;
            left: 60px;
            font-size: 12px;
        }

    }

    .comment-content {
        margin: 6px 0 30px 60px;
        font-size: 14px;
        line-height: 16px;
    }
}

.reply-list {
    margin-top: 10px;
    padding: 12px 10px;
    box-sizing: border-box;
    background-color: #f8f8f8;

    .reply-item {
        line-height: 20px;

        .reply-user-name {
            color: var(--primary-color);
            font-size: 12px;
        }
    }

    .more {
        button {
            font-size: 12px;
            color: #969996;
            margin-top: 8px;
        }
    }
}

.content-text {
    display: inline-block;
}

.mention-user {
    color: var(--primary-color);
    cursor: pointer;
    padding: 0 2px;
}
</style>