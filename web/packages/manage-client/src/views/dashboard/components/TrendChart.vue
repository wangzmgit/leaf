<template>
    <div class="echarts-box">
        <div class="echart" id="trend"></div>
    </div>
</template>

<script setup lang="ts">
import theme from './theme.json';
import * as echarts from "echarts";
import { ref } from "vue";
import { onMounted, onUnmounted } from 'vue';

const props = defineProps<{
    data: Array<{
        date: string,
        user: number,
        video: number
    }>
}>();

let chart: echarts.ECharts | null = null;
const date = ref<Array<string>>([]);
const newUser = ref<Array<number>>([]);
const newVideo = ref<Array<number>>([]);

const initChart = () => {
    echarts.registerTheme('westeros', theme);//注册主题
    chart = echarts.init(document.getElementById("trend")!, "westeros");
    chart.setOption({
        legend: {
            data: ['新增用户', '新增视频']
        },
        xAxis: {
            type: "category",
            data: date.value
        },
        tooltip: {
            trigger: "axis"
        },
        yAxis: {
            type: "value"
        },
        series: [
            {
                name: '新增用户',
                data: newUser.value,
                type: "line",
                smooth: true
            },
            {
                name: '新增视频',
                data: newVideo.value,
                type: "line",
                smooth: true
            }
        ]
    });
}

const resize = () => {
    chart?.resize();
}

onMounted(() => {
    props.data.forEach(item => {
        date.value.push(item.date);
        newUser.value.push(item.user);
        newVideo.value.push(item.video);
    });
    initChart();
    window.addEventListener("resize", resize);
});

onUnmounted(() => {
    window.removeEventListener("resize", resize);
    chart?.dispose();
});
</script>

<style lang="less" scoped>
.echarts-box {
    width: 100%;
    height: 360px;
    padding-top: 30px;

    .echart {
        width: 100%;
        height: 100%;
    }
}
</style>