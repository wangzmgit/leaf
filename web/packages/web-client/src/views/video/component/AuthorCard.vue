<template>
    <div v-if="loading" class="video-author">
        <div class="video-author-box">
            <!--头像-->
            <div class="author-avatar">
                <common-avatar :size="50"></common-avatar>
            </div>
            <div class="author-info">
                <n-skeleton height="20px" width="70px" />
                <n-skeleton height="16px" width="80%" />
            </div>
        </div>
    </div>
    <div v-else class="video-author">
        <div class="video-author-box">
            <!--头像-->
            <div class="author-avatar">
                <common-avatar :url="authorInfo.avatar" :size="50"></common-avatar>
            </div>
            <!--昵称和个签-->
            <div class="author-info">
                <p @click="goUserSpace(authorInfo.uid)">{{ authorInfo.name }}</p>
                <p>{{ authorInfo.sign }}</p>
            </div>
            <div class="follow-btn">
                <n-button size="small" v-if="isFollow" type="primary" @click="unfollow(author!.uid)">
                    {{ t("follow.followed") }}
                </n-button>
                <n-button size="small" v-else :disabled="!isLogin" type="error" @click="follow(author!.uid)">
                    {{ t("follow.follow") }}
                </n-button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n';
import { useRouter } from 'vue-router';
import useUserFollow from '@/hooks/user-follow-hooks';
import { NButton, NSkeleton } from 'naive-ui';
import type { UserInfoType } from '@leaf/apis';
import { CommonAvatar } from '@leaf/components';
import { onBeforeMount, ref, watch } from 'vue';
import { storageData } from '@leaf/utils';

const props = defineProps<{
    author?: UserInfoType,
    loading: Boolean
}>()

// i18n
const { t } = useI18n();

const router = useRouter();

const isLogin = ref(false);
const authorInfo = ref(Object.assign({}, props.author));
const { isFollow, getFollowStatus, follow, unfollow } = useUserFollow();

const goUserSpace = (uid: number) => {
    let userUrl = router.resolve({ name: "User", params: { uid: uid } });
    window.open(userUrl.href, '_blank');
}

watch(() => props.loading, (val) => {
    if (!val && props.author) {
        authorInfo.value = props.author;
        getFollowStatus(authorInfo.value.uid);
    }
})

onBeforeMount(() => {
    if (storageData.get("user_info")?.uid) {
        isLogin.value = true;
    }
})

</script>

<style lang="less" scoped>
.video-author {
    width: 100%;
    height: 90px;
    border-radius: 6px;
    background-color: #f1f2f3;

    .video-author-box {
        display: flex;
        height: 100%;
        align-items: center;

        .author-avatar {
            width: 50px;
            height: 50px;
            padding-left: 10px;
        }

        .author-info {
            height: 50px;
            display: flex;
            flex-direction: column;
            justify-content: space-around;
            width: calc(100% - 160px);
            padding-left: 16px;

            p {
                width: 100%;
                margin: 0;
                overflow: hidden;
                text-overflow: ellipsis;
                display: -webkit-box;
                -webkit-line-clamp: 1;
                -webkit-box-orient: vertical;

                &:nth-child(1) {
                    color: #222;
                    font-size: 14px;
                    cursor: pointer;
                    font-weight: 500;
                }

                &:nth-child(2) {
                    color: #999;
                    font-size: 12px;
                }
            }
        }

        .follow-btn {
            width: 80px;
            text-align: center;
        }
    }
}
</style>