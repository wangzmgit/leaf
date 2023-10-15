<template>
    <div class="dashboard">
        <n-scrollbar>
            <div class="card card-list">
                <div class="data-item" style="background-color: rgba(255, 156, 110, 0.2);">
                    <n-statistic label="待审核视频数">
                        <span style="color: rgb(255, 156, 110);">{{ cardData.review_video_count }}</span>
                    </n-statistic>
                </div>
                <div class="data-item" style="background-color: rgba(105, 192, 255, 0.2);">
                    <n-statistic label="今日新增用户数">
                        <span style="color: rgb(105, 192, 255);">{{ cardData.new_user_count }}</span>
                    </n-statistic>
                </div>
                <div class="data-item" style="background-color: rgba(179, 127, 235, 0.2);">
                    <n-statistic label="今日新增视频数">
                        <span style="color: rgb(179, 127, 235);">{{ cardData.new_video_count }}</span>
                    </n-statistic>
                </div>
                <div class="data-item" style="background-color: rgba(255, 133, 160, 0.2)">
                    <n-statistic label="总用户数">
                        <span style="color: rgb(255, 133, 160);">{{ cardData.user_count }}</span>
                    </n-statistic>
                </div>
                <div class="data-item" style="background-color: rgba(92, 219, 211, 0.2);">
                    <n-statistic label="总视频数">
                        <span style="color: rgb(92, 219, 211);">{{ cardData.video_count }}</span>
                    </n-statistic>
                </div>
            </div>
            <div class="card">
                <trend-chart v-if="!loadingTrendChart" :data="trendData"></trend-chart>
            </div>
            <div class="card">
                <partition-data></partition-data>
            </div>
        </n-scrollbar>
    </div>
</template>

<script setup lang="ts">
import { onBeforeMount, reactive, ref } from 'vue';
import { getCardDataAPI, getTrendDataAPI, getPartitionAPI } from '@leaf/apis';
import { statusCode } from '@leaf/utils';
import type { CardDataType, PartitionType } from '@leaf/apis';
import { NStatistic, NCard, NScrollbar, useNotification } from 'naive-ui';
import TrendChart from './components/TrendChart.vue';
import PartitionData from './components/PartitionData.vue';

const notification = useNotification();

const cardData = ref<CardDataType>({
    user_count: 0,
    video_count: 0,
    review_video_count: 0,
    new_user_count: 0,
    new_video_count: 0,
})

const getCardData = () => {
    getCardDataAPI().then((res) => {
        if (res.data.code === statusCode.OK) {
            cardData.value = res.data.data.data;
        }
    })
}

const trendData = ref([]);
const loadingTrendChart = ref(true);
const getTrendData = () => {
    getTrendDataAPI().then((res) => {
        if (res.data.code === statusCode.OK) {
            trendData.value = res.data.data.data.reverse();
            loadingTrendChart.value = false;
        }
    })
}

onBeforeMount(() => {
    getCardData();
    getTrendData();
})
</script>

<style lang="less" scoped>
.dashboard {
    width: 100%;
    height: calc(100vh - 60px);
}

.card {
    margin: 10px;
    width: calc(100% - 20px);
    border-radius: 10px;
    background-color: #fff;
}

.card-list {
    display: grid;
    grid-template-columns: repeat(5, 1fr);

    .data-item {
        margin: 16px 16px;
        padding: 12px 16px;
        border-radius: 6px;
    }
}
</style>