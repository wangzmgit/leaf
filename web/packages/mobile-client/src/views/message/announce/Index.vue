<template>
    <div v-title data-title="系统公告">
        <header-bar></header-bar>
        <div v-if="loading" class="loading">
            <n-spin type="loading" />
        </div>
        <div v-else class="announce-box">
            <div class="announce-item" v-for="(item, index) in announceList" :key="index">
                <div class="title">
                    <p class="item-title">{{ item.title }}</p>
                    <n-time class="item-time" :time="new Date(item.created_at)"></n-time>
                </div>
                <span>{{ item.content }}</span>
                <span class="link" v-if="item.url">
                    <span @click="linkTo(item.url)">网页链接</span>
                </span>
            </div>
            <div class="more">
                <n-button v-if="!noMore" @click="getAnnounceList(++page)">加载更多</n-button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onBeforeMount, ref } from "vue";
import { NSpin, NTime, NButton,useMessage } from 'naive-ui';
import { getAnnounceAPI } from "@leaf/apis";
import type { AnnounceType } from "@leaf/apis";
import HeaderBar from "@/components/header-bar/Index.vue";
import { statusCode } from "@leaf/utils";

const message = useMessage();

let page = 1;
const pageSize = 10;
let noMore = false;
const loading = ref(false);
const announceList = ref<Array<AnnounceType>>([]);

const getAnnounceList = (page: number) => {
    loading.value = true;
    getAnnounceAPI(page, pageSize).then((res) => {
        if (res.data.code === statusCode.OK) {
            if (res.data.data.announces) {
                announceList.value.push(...res.data.data.announces);
                if (res.data.data.announces.length < pageSize) {
                    noMore = true;
                    message.info("没有更多了");
                }
            } else {
                noMore = true;
                message.info("没有更多了");
            }
        }
        
        loading.value = false;
    });
}

const linkTo = (url: string) => {
    window.open(url, "_blank");
}

onBeforeMount(() => {
    getAnnounceList(page);
})
</script>

<style lang="less" scoped>
.loading {
    font-size: 30px;
    text-align: center;
}

.announce-box {
    min-height: 600px;

    .announce-item {
        padding: 8px 6px;
        border-bottom: 1px solid #e1e2e3;

        .title {
            display: flex;
            align-items: center;
            justify-content: space-between;
            margin: 5px 0;

            .item-title {
                font-size: 16px;
                font-weight: 500;
                width: calc(100% - 160px);
                margin: 0;
            }

            .item-time {
                color: #999;
            }
        }

        .link {
            margin-left: 10px;
            color: #1890ff;
            cursor: pointer;
        }
    }

    .more {
        text-align: center;
        margin-top: 10px;
    }
}
</style>