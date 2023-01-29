<template>
    <div class="collection-list" :style="initTheme()">
        <p class="create-title">{{ t('common.collection') }}</p>
        <n-scrollbar style="max-height: 550px;">
            <div class="card-item" v-for="(item, index) in collections" :key="index">
                <div class="card-left">
                    <img v-if="item.cover" :src="getResourceUrl(item.cover)" alt="收藏夹封面">
                    <div class="no-cover" v-else>
                        <img src="@/assets/collection.png" alt="默认封面">
                    </div>
                </div>
                <div class="card-center">
                    <p class="title" @click="goCollectionDetails(item.id)">{{ item.name }}</p>
                    <span class="desc">{{ t('common.desc') }}：{{ item.desc }}</span>
                    <span class="desc">
                        {{ t('common.createdAt') }}：
                        <n-time :time="new Date(item.created_at!)" />
                    </span>
                </div>
                <div class="card-right">
                    <n-icon class="edit" size="20" @click="beforeEdit(item)">
                        <edit></edit>
                    </n-icon>
                    <n-popconfirm negative-text="取消" positive-text="确认" @positive-click="deleteClick(item.id, index)">
                        <template #trigger>
                            <n-icon class="edit" size="20">
                                <delete></delete>
                            </n-icon>
                        </template>
                        是否删除这个收藏夹
                    </n-popconfirm>
                </div>
            </div>
        </n-scrollbar>
    </div>
    <n-drawer v-model:show="active" :width="500" placement="right">
        <n-drawer-content :title="t('space.editCollection')">
            <cover-uploader :cover="collectionInfo.cover" @finish="finishUpload"></cover-uploader>
            <n-form class="info-form">
                <n-form-item :label="t('space.collectionName')">
                    <n-input v-model:value="collectionInfo.name" placeholder="请输入名称" maxlength="20" show-count />
                </n-form-item>
                <n-form-item :label="t('common.desc')">
                    <n-input v-model:value="collectionInfo.desc" placeholder="收藏夹简介~" maxlength="150" show-count
                        type="textarea" :autosize="descSize" />
                </n-form-item>
                <n-form-item :label="t('common.open')">
                    <n-switch v-model:value="collectionInfo.open" />
                </n-form-item>
                <div class="upload-next-btn">
                    <n-button type="primary" @click="modifyCollection">{{ t('common.save') }}</n-button>
                </div>
            </n-form>
        </n-drawer-content>
    </n-drawer>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n';
import { getTheme } from "@/theme";
import { useRouter } from 'vue-router';
import { ref, reactive, onBeforeMount } from 'vue';
import useCollection from '@/hooks/collection-hooks';
import type { CollectionType, ModifyCollectionType } from '@leaf/apis';
import CoverUploader from "./component/CoverUploader.vue";
import { Edit, Delete } from "@leaf/icons";
import { modifyCollectionAPI, deleteCollectionAPI } from '@leaf/apis';
import {
    NTime, NIcon, NForm, NFormItem, NButton, NInput, NSwitch,
    NScrollbar, NPopconfirm, NDrawer, NDrawerContent, useNotification
} from 'naive-ui';
import { getResourceUrl, statusCode } from '@leaf/utils';

// i18n
const { t } = useI18n();

const initTheme = () => {
    const theme = getTheme();

    return {
        "--hover-color": theme.primaryHoverColor
    }
}

//简介输入框大小
const descSize = {
    minRows: 3,
    maxRows: 3
}

const notification = useNotification();
const { collections, getCollectionList } = useCollection();

const active = ref(false);//显示编辑抽屉
const collectionInfo = reactive<ModifyCollectionType>({
    id: 0,
    cover: "",
    name: "",
    desc: "",
    open: false,
});
const beforeEdit = (item: CollectionType) => {
    collectionInfo.id = item.id;
    collectionInfo.cover = item.cover!;
    collectionInfo.name = item.name!;
    collectionInfo.desc = item.desc!;
    collectionInfo.open = item.open!;
    active.value = true;
}

//封面上传完成
const finishUpload = (cover: string) => {
    collectionInfo.cover = cover;
}

//修改收藏夹
const modifyCollection = () => {
    if (!collectionInfo.name) {
        notification.error({
            title: '请输入收藏夹名',
        });
        return;
    }
    modifyCollectionAPI(collectionInfo).then((res) => {
        if (res.data.code === statusCode.OK) {
            getCollectionList();
        } else {
            notification.error({
                title: '修改失败',
                content: res.data.msg,
                duration: 5000,
            })
        }
    })
}

//删除收藏夹
const deleteClick = (id: number, index: number) => {
    deleteCollectionAPI(id).then((res) => {
        if (res.data.code === statusCode.OK) {
            collections.value.splice(index, 1);
        } else {
            notification.error({
                title: '删除失败',
                content: res.data.msg,
                duration: 5000,
            })
        }
    })
}

//前往收藏夹详情
const router = useRouter();
const goCollectionDetails = (id: number) => {
    router.push({ name: "CollectionDetails", params: { id: id } });
}

onBeforeMount(() => {
    getCollectionList();
})
</script>

<style lang="less" scoped>
.collection-list {
    margin: 0 26px;

    .create-title {
        line-height: 30px;
        font-size: 20px;
        user-select: none;
    }

    .card-item {
        width: 100%;
        height: 80px;
        margin-bottom: 12px;
        border-bottom: 1px solid #e6e6e6;
        padding-bottom: 12px;
    }
}

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

    .no-cover {
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
</style>
