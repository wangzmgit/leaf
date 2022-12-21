<template>
    <div class="comment-box">
        <!--头像-->
        <common-avatar :url="userInfo.avatar"></common-avatar>
        <!--输入框-->
        <n-input class="comment-input" v-model:value="commentForm.content" placeholder="在这里写点什么吧~" maxlength="255"
            show-count type="textarea" :autosize="descSize" />
        <n-button v-if="!userInfo.uid" disabled>未登录</n-button>
        <n-button v-else type="primary" @click="addComment">发表</n-button>
    </div>
    <div class="comment-item" v-for="(item, index) in commentList" :key="index">
        <!--头像和昵称-->
        <div class="comment-user">
            <common-avatar :url="item.author.avatar"></common-avatar>
            <div class="user-name">
                <span @click="goUserSpace(item.author.uid)">{{ item.author.name }}</span>
            </div>
            <n-time class="comment-time" type="relative" :time="new Date(item.created_at)"></n-time>
            <div class="comment-btn">
                <n-button v-if="userInfo" text @click="showReply(index, '')">回复</n-button>
                <n-button v-if="item.author.uid === userInfo.uid" text
                    @click="deleteClick(item.id, null, index)">删除</n-button>
            </div>
        </div>
        <!--评论内容-->
        <div class="comment-content">
            <div v-if="item.content.indexOf('@') === -1">
                <span>{{ item.content }}</span>
            </div>
            <div class="content-text" v-else v-for="content in handleMention(item.content)">
                <span :class="content.class" @click="goMention(content.key)">{{ content.value }}</span>
            </div>
        </div>
        <!--动态回复框-->
        <div v-if="showReplyFlag[index]">
            <div class="reply-box">
                <common-avatar :url="userInfo.avatar"></common-avatar>
                <!-- 输入框 -->
                <n-input class="reply-input" v-model:value="replyForm.content" :placeholder="replyTip" maxlength="200"
                    show-count type="textarea" :autosize="descSize" />
                <n-button type="primary" @click="submitReply(item)">回复</n-button>
            </div>
        </div>
        <!--回复-->
        <div class="reply-list" v-if="item.reply">
            <div class="reply-item" v-for="(reply, i) in item.reply" :key="i">
                <!-- 头像和昵称 -->
                <div class="reply-user">
                    <div class="info">
                        <common-avatar :url="reply.author.avatar" :size="26" :iconsize="18"></common-avatar>
                        <div class="reply-user-name">
                            <span @click="goUserSpace(reply.author.uid)">{{ reply.author.name }}</span>
                        </div>
                        <n-time class="reply-time" type="relative" :time="new Date(reply.created_at)"></n-time>
                    </div>
                    <div class="reply-btn">
                        <n-button v-if="userInfo" text @click="showReply(index, reply.author.name)">回复</n-button>
                        <n-button v-if="reply.author.uid === userInfo.uid" text
                            @click="deleteClick(item.id, reply.id, index, i)">删除
                        </n-button>
                    </div>
                </div>
                <!-- 二级评论内容 -->
                <div class="reply-content">
                    <div v-if="reply.content.indexOf('@') === -1">
                        <span>{{ reply.content }}</span>
                    </div>
                    <div class="content-text" v-else v-for="content in handleMention(reply.content)">
                        <span :class="content.class" @click="goMention(content.key)">{{ content.value }}</span>
                    </div>
                </div>
            </div>
            <div class="more">
                <n-button :disabled="item.noMore" text @click="getMoreReply(item.id, index)">加载更多</n-button>
            </div>
        </div>
    </div>
    <!-- 加载更多 -->
    <div v-if="!noMore" class="more-btn">
        <n-button @click="getMore">加载更多</n-button>
    </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { onBeforeMount, reactive, ref } from "vue";
import useComment from '@/hooks/comment-hooks';
import useMention from '@/hooks/mention-hooks';
import { statusCode, storageData } from '@leaf/utils';
import { CommonAvatar } from '@leaf/components';
import { NButton, NInput, NTime, useNotification } from "naive-ui";
import { addCommentAPI, type ReplyType, } from '@leaf/apis';
import type { AddCommentType, CommentType, UserInfoType } from '@leaf/apis';

const props = defineProps<{
    vid: number
}>();

const replyTip = ref('留下条评论吧');
//输入框大小
const descSize = {
    minRows: 3,
    maxRows: 3
}
const commentForm = reactive<AddCommentType>({
    vid: props.vid,
    content: "",
    parentId: "",
})

const replyUser = ref("");
const replyForm = reactive({
    vid: props.vid,
    content: "",
    parentId: "",
})

const userInfo = ref<UserInfoType>({
    uid: 0,
    name: "",
    avatar: ""
});

const notification = useNotification();//通知
const { noMore, commentList, getCommentList, getReplyListSync, deleteCommentSync } = useComment();

//评论回复
const replyCount = 2;//获取评论时查询的回复数
const addComment = () => {
    addCommentAPI(commentForm).then((res) => {
        if (res.data.code === statusCode.OK) {
            //加载评论
            if (noMore.value) {
                commentList.value.push({
                    id: res.data.data.id,
                    author: {
                        uid: userInfo.value.uid,
                        name: userInfo.value.name,
                        avatar: userInfo.value.avatar,
                    },
                    content: commentForm.content,
                    created_at: new Date().getTime(),
                    reply: [],
                    page: 1,
                    noMore: true
                })
            }
            notification.success({
                title: '发布成功',
                duration: 5000,
            });
            commentForm.content = "";
        } else {
            notification.error({
                title: '发布失败',
                description: res.data.msg,
                duration: 5000,
            })
        }
    })
}

//提交回复
const submitReply = (comment: CommentType) => {
    replyForm.parentId = comment.id;
    if (replyUser.value) {
        replyForm.content = `回复 @${replyUser.value} :${replyForm.content}`;
    }
    if (comment.reply.length < replyCount) {
        comment.noMore = true;
    }

    addCommentAPI(replyForm).then((res) => {
        if (res.data.code === statusCode.OK) {
            notification.success({
                title: '发布成功',
                duration: 5000,
            });

            if (comment.noMore) {
                const newReply: ReplyType = {
                    id: res.data.data.id,
                    author: {
                        uid: userInfo.value.uid,
                        name: userInfo.value.name,
                        avatar: userInfo.value.avatar,
                    },
                    content: replyForm.content,
                    created_at: new Date().getTime(),
                }
                comment.reply.push(newReply);
            }
            replyForm.content = "";
            //关闭动态回复框
            showReplyFlag.value.forEach((item, index) => {
                if (item) showReplyFlag.value[index] = false;
            });
        } else {
            notification.error({
                title: '发布失败',
                description: res.data.msg,
                duration: 5000,
            })
        }
    })
}

//页面跳转
const router = useRouter();
const goUserSpace = (uid: number) => {
    let userUrl = router.resolve({ name: "User", params: { uid: uid } });
    window.open(userUrl.href, '_blank');
}

//显示隐藏动态回复
const showReplyFlag = ref<Array<boolean>>([]);
const showReply = (index: number, name: string) => {
    if (showReplyFlag.value[index]) {
        showReplyFlag.value[index] = false;
        return;
    }
    for (let i = 0; i < commentList.value.length; i++) {
        showReplyFlag.value[i] = false;
    }
    if (name) {
        replyTip.value = `回复 @${name}: `;
        replyUser.value = name;
    }
    showReplyFlag.value[index] = true;
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
            commentList.value[index].noMore = true;
            return;
        }

        commentList.value[index].page += 1;
        commentList.value[index].reply.push(...res);
    })
}

//加载更多
const getMore = () => {
    page.value++;
    getCommentList(props.vid, page.value, 8);
}


//删除评论回复
const deleteClick = (id: string, replyId: string | null, index: number, replyIndex: number | null = null) => {
    deleteCommentSync(id, replyId).then((res) => {
        if (res) {
            if (replyIndex !== null) {
                (commentList.value[index].reply)?.splice(replyIndex, 1);
            } else {
                commentList.value.splice(index, 1);
            }
        }
    })
}

//处理@
const { handleMention } = useMention();

//前往@的用户
const goMention = (name: string | null) => {
    // if (name) {
    //     let userUrl = router.resolve({ name: 'MentionUser', params: { name: name } });
    //     window.open(userUrl.href, '_blank');
    // }
}

onBeforeMount(() => {
    userInfo.value = storageData.get("user_info");
    getCommentList(props.vid, page.value, 8);
})
</script>

<style lang="less" scoped>
.comment-box {
    margin-top: 10px;
    padding-bottom: 20px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    border-bottom: 1px solid #e3e4e5;

    .comment-input {
        width: calc(100% - 200px);
    }

    button {
        width: 90px;
    }

}

.comment-item {
    margin-top: 10px;

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

        .comment-btn {
            float: right;

            button {
                font-size: 10px;
                color: #969996;
                margin-left: 6px;

                &:hover {
                    color: #36ad6a;
                }
            }
        }
    }

    .comment-content {
        margin: 6px 0 6px 60px;
        font-size: 14px;
        line-height: 16px;
        padding-bottom: 10px;
        border-bottom: 1px solid #e3e4e5;
    }
}

/**动态回复框 */
.reply-box {
    display: flex;
    align-items: center;
    justify-content: space-around;
    margin: 10px 0 10px 60px;
    padding-bottom: 10px;
    border-bottom: 1px solid #e3e4e5;

    .reply-input {
        width: calc(100% - 180px);
    }

    button {
        width: 80px;
    }
}

.reply-list {
    border-bottom: 1px solid #e3e4e5;

    .reply-item {
        .reply-user {
            height: 30px;
            display: flex;
            align-items: center;
            justify-content: space-between;
            margin-left: 60px;

            .info {
                display: inline-flex;
                align-items: center;

                .reply-user-name {
                    cursor: pointer;
                    color: #666;
                    font-size: 12px;
                    padding: 0 10px;
                }

                .reply-time {
                    color: #969996;
                    font-size: 10px;
                    line-height: 30px;
                }
            }

            .reply-btn {
                float: right;

                button {
                    font-size: 10px;
                    color: #969996;
                    margin-left: 6px;

                    &:hover {
                        color: #36ad6a;
                    }
                }
            }
        }

        .reply-content {
            color: #222;
            display: block;
            margin: 6px 0 0 96px;
            padding-bottom: 10px;
        }
    }

    .more {
        button {
            font-size: 12px;
            color: #969996;
            margin: 10px 60px;

            &:hover {
                color: #36ad6a;
            }
        }
    }
}

.content-text {
    display: inline-block;

    .mention-user {
        color: #36ad6a;
        cursor: pointer;
        padding: 0 2px;
    }
}

.more-btn {
    text-align: center;
}
</style>