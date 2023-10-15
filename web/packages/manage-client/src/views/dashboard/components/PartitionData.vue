<template>
    <div class="partition-box">
        <div class="partitions">
            <p v-for="item in partitionInfo.partitions" :key="item.id" @click="selectPartition(item.id)">
                <span :class="partitionInfo.selectedPartition === item.id ? 'selected' : ''">
                    {{ item.content }}
                </span>
            </p>
        </div>
        <div class="echart" id="partitionChart"></div>
    </div>
</template>

<script setup lang="ts">
import { onBeforeMount, onUnmounted, onMounted, reactive, ref } from 'vue';
import theme from './theme.json';
import * as echarts from "echarts";
import { getPartitionAPI, getPartitionDataAPI } from '@leaf/apis';
import { statusCode } from '@leaf/utils';
import type { PartitionType } from '@leaf/apis';
import { useNotification } from 'naive-ui';

const notification = useNotification();

const chartOption = {
    xAxis: {
        type: 'category',
        data: [] as string[]
    },
    yAxis: {
        type: 'value'
    },
    series: [
        {
            data: [] as number[],
            type: 'bar'
        }
    ]
}

const partitionInfo = reactive({
    partitions: [] as PartitionType[], //分区内容
    selectedPartition: 0,//选中分区
});

// 获取所有分区
const getAllPartition = () => {
    getPartitionAPI().then((res) => {
        if (res.data.code === statusCode.OK) {
            const partitions = res.data.data.partitions as PartitionType[];
            partitionInfo.partitions = partitions.filter((item) => {
                return item.parent_id === 0;
            });

            selectPartition(partitionInfo.partitions[0].id);
        } else {
            notification.error({
                title: "分区加载失败",
                duration: 5000
            })
        }
    })
}

// 选中分区
const selectPartition = (partitionId: number) => {
    partitionInfo.selectedPartition = partitionId;
    getPartitionData(partitionInfo.selectedPartition);
}

const partitionVideoCount = ref<Array<{ content: string, video_count: number }>>([]);
const getPartitionData = (partitionId: number) => {
    getPartitionDataAPI(partitionId).then((res) => {
        if (res.data.code === statusCode.OK) {
            partitionVideoCount.value = res.data.data.data;
            chartOption.xAxis.data = partitionVideoCount.value.map(item => item.content);
            chartOption.series[0].data = partitionVideoCount.value.map(item => item.video_count);
            initChart();
        }
    })
}



let chart: echarts.ECharts | null = null;
const initChart = () => {
    if (!chart) {
        echarts.registerTheme('westeros', theme);//注册主题
        chart = echarts.init(document.getElementById("partitionChart")!, "westeros");
    }

    chart.setOption(chartOption);
}

const resize = () => {
    chart?.resize();
    window.addEventListener("resize", resize);
}

onBeforeMount(() => {
    getAllPartition();
})

onMounted(() => {
    window.addEventListener("resize", resize);
})

onUnmounted(() => {
    window.removeEventListener("resize", resize);
    chart?.dispose();
});
</script>

<style lang="less" scoped>
.partition-box {
    width: 100%;
    height: 360px;
    padding-top: 30px;

    .partitions {
        display: flex;

        p {
            cursor: pointer;
            margin: 0 20px;

            span {
                padding: 3px 6px;
            }
        }

        .selected {
            color: #fff;
            border-radius: 6px;
            background-color: rgba(24, 143, 255, 0.5);
        }
    }

    .echart {
        width: 100%;
        height: 100%;
    }
}
</style>