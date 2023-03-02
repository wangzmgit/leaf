<template>
    <div class="collection" :style="initTheme()" v-title data-title="收藏夹详情">
        <header-bar></header-bar>
        <div class="collection-main">
            <div class="content-left">
                <div class="card-item" v-for="(item, index) in videoList" :key="index">
                    <div class="card-left">
                        <img v-if="item.cover" :src="getResourceUrl(item.cover)" alt="收藏夹封面">
                        <div class="card-no-cover" v-else>
                            <img src="@/assets/collection.png" alt="默认封面">
                        </div>
                    </div>
                    <div class="card-center">
                        <p class="title" @click="goVideo(item.vid)">{{ item.title }}</p>
                        <span class="desc">{{ t("common.clicks") }}：{{ item.clicks }}</span>
                        <span class="desc">{{ t("common.desc") }}：{{ item.desc }}</span>
                    </div>
                    <div class="card-right" v-if="uid === author?.uid">
                        <n-icon class="edit" size="20" @click="removeVideo(item.vid)">
                            <delete></delete>
                        </n-icon>
                    </div>
                </div>
                <div v-if="loadingContent" class="page-box">
                    <n-pagination v-model:page="page" :item-count="count" :page-size="8" @update-page="pageChange" />
                </div>
            </div>
            <div v-if="loadingInfo" class="content-right">
                <div class="cover">
                    <img v-if="collection?.cover" :src="getResourceUrl(collection.cover)" alt="视频封面">
                    <div class="no-cover" v-else>
                        <img src="@/assets/collection.png" alt="默认封面">
                    </div>
                </div>
                <div class="info">
                    <span class="title">{{ collection?.name }}</span>
                    <span class="desc">{{ t("common.desc") }}：{{ collection?.desc }}</span>
                    <div class="desc">
                        <n-time type="date" :time="new Date(collection?.created_at || 0)"></n-time>
                        <span>・</span>
                        <span class="open">{{ collection?.open? t("common.open"): t("common.private") }}</span>
                        <span class="author" @click="goSpace(author?.uid || 0)">{{ author?.name }}</span>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n';
import { onBeforeMount, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import HeaderBar from '@/components/header-bar/Index.vue';
import { Delete } from '@leaf/icons';
import { getResourceUrl, statusCode, storageData } from '@leaf/utils';
import type { UserInfoType, CollectionInfoType, VideoType } from '@leaf/apis';
import { getCollectionInfoAPI, getCollectVideoAPI, collectAPI } from '@leaf/apis';
import { NTime, NIcon, NPagination, useNotification } from 'naive-ui';
import { getTheme } from '@/theme';

// i18n
const { t } = useI18n();

const id = ref(0);
const uid = ref(0);//用户ID
const route = useRoute();
const notification = useNotification();

const initTheme = () => {
    const theme = getTheme();

    return {
        "--hover-color": theme.primaryColor
    }
}

// 加载中
const loadingInfo = ref(false);
const loadingContent = ref(false);

const author = ref<UserInfoType>();
const collection = ref<CollectionInfoType>();
const getCollectionInfo = (id: number) => {
    getCollectionInfoAPI(id).then((res) => {
        if (res.data.code === statusCode.OK) {
            author.value = res.data.data.author;
            collection.value = res.data.data.collection;
            loadingInfo.value = true;
        }
    })
}

const page = ref(1);
const count = ref(0);
const videoList = ref<Array<VideoType>>([]);
//获取收藏视频内容
const getCollectVideo = () => {
    getCollectVideoAPI(id.value, page.value, 8).then((res) => {
        if (res.data.code === statusCode.OK) {
            count.value = res.data.data.total;
            videoList.value = res.data.data.videos;
            if (count.value > 0) {
                loadingContent.value = true;
            }
        }
    })
}

//页码改变
const pageChange = (target: number) => {
    page.value = target;
    getCollectVideo();
}

//移除视频
const removeVideo = (vid: number) => {
    collectAPI({ vid, addList: [], cancelList: [id.value] }).then((res) => {
        if (res.data.code === statusCode.OK) {
            getCollectVideo();
        } else {
            notification.error({
                title: '移除失败',
                duration: 5000,
            })
        }
    })
}

//页面跳转
const router = useRouter();
const goSpace = (uid: number) => {
    let userUrl = router.resolve({ name: "User", params: { uid: uid } });
    window.open(userUrl.href, '_blank');
}

const goVideo = (vid: number) => {
    let videoUrl = router.resolve({ name: "Video", params: { vid: vid } });
    window.open(videoUrl.href, '_blank');
}

onBeforeMount(() => {
    if (storageData.get('user_info') && storageData.get('user_info').uid) {
        uid.value = storageData.get('user_info').uid;
    }
    id.value = Number(route.params.id);
    getCollectionInfo(id.value);
    getCollectVideo();
})
</script>

<style lang="less" scoped>
.collection {
    height: 100%;
    width: 100%;
    top: 0;
    bottom: 0;
    position: fixed;
    overflow-y: scroll;
    background-color: #f5f6f7;
}

.collection-main {
    // padding: 20px 20px 16px;
    width: calc(100% - 200px);
    box-sizing: border-box;
    margin: 20px auto 0 auto;
    display: flex;
    flex-wrap: nowrap;
    justify-content: space-between;
}

.content-left {
    padding-top: 20px;
    width: calc(100% - 380px);
    min-width: 680px;
    background: #fff;
}

.card-item {
    width: 100%;
    height: 80px;
    margin-bottom: 12px;
    border-bottom: 1px solid #e6e6e6;
    padding-bottom: 12px;

    .card-left {
        float: left;
        width: 120px;
        height: 80px;
        margin-right: 10px;

        img {
            width: 100%;
            height: 100%;
            border-radius: 2px;
        }

        .card-no-cover {
            width: 100%;
            height: 100%;
            background-color: #e6e6e6;
            border-radius: 2px;

            img {
                position: absolute;
                width: 50px;
                height: 50px;
                margin: 15px 35px;
                filter: grayscale(100%);
                opacity: 50%;
            }
        }
    }

    .card-center {
        float: left;
        width: calc(100% - 230px);

        .title {
            font-size: 14px;
            color: #212121;
            line-height: 18px;
            margin: 0 0 26px;
            cursor: pointer;
            overflow: hidden;
            text-overflow: ellipsis;
            display: -webkit-box;
            -webkit-line-clamp: 1;
            -webkit-box-orient: vertical;

            &:hover {
                color: var(--hover-color);
            }
        }

        .desc {
            font-size: 12px;
            color: #999;
            overflow: hidden;
            text-overflow: ellipsis;
            display: -webkit-box;
            -webkit-line-clamp: 1;
            -webkit-box-orient: vertical;
        }
    }

    .card-right {
        float: left;
        width: 90px;
        height: 100%;
        text-align: center;


        .edit {
            color: #999;
            cursor: pointer;
            line-height: 90px;
            margin-right: 20px;

            &:hover {
                color: var(--hover-color);
            }
        }
    }
}



/**右半部分 */
.content-right {
    width: 350px;
    height: 600px;
    min-width: 350px;
    padding: 10px;
    background: #fff;

    .cover {
        width: 100%;
        height: 200px;
        margin-bottom: 12px;

        img {
            width: 100%;
            height: 100%;
            border-radius: 2px;
        }

        .no-cover {
            width: 100%;
            height: 100%;
            background-color: #e6e6e6;
            border-radius: 2px;

            img {
                position: absolute;
                width: 100px;
                height: 100px;
                margin: 50px 125px;
                filter: grayscale(100%);
                opacity: 50%;
            }
        }
    }

    .info {
        .title {
            display: block;
            height: auto;
            font-size: 16px;
            color: #212121;
            line-height: 18px;
            margin-bottom: 20px;
        }

        .desc {
            display: block;
            font-size: 12px;
            font-size: 14px;
            color: #666;
            margin: 2px 0;

            .author {
                float: right;
                cursor: pointer;

                &:hover {
                    color: var(--hover-color);
                }
            }
        }


    }
}
</style>