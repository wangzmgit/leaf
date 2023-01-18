<template>
    <div class="select-partition">
        <n-select class="select" :key="key" placeholder="选择分区" label-field="content" value-field="id" remote
            :options="partitions" @update:value="partitionChange" />
        <n-select class="select" :key="key" label-field="content" value-field="id" remote placeholder="选择二级分区"
            :options="subpartitions" @update:value="selectedPartition" />

        <n-button @click="reset">重置</n-button>
    </div>
</template>

<script setup lang="ts">
import { onBeforeMount, ref } from "vue";
import type { PartitionType } from '@leaf/apis';
import { getPartitionAPI, } from '@leaf/apis';
import { NSelect, NButton, useNotification } from 'naive-ui';
import { statusCode } from "@leaf/utils";

const emits = defineEmits(["selected"]);
const notification = useNotification();

const key = ref(Date.now());
const partitionList = ref<Array<PartitionType>>([]);//所有分区
const partitions = ref<Array<PartitionType>>([]);//分区
const subpartitions = ref<Array<PartitionType>>([]);//二级分区

//改变分区
const partitionChange = (parentId: number) => {
    subpartitions.value = partitionList.value.filter((item) => {
        return item.parent_id === parentId;
    })
    emits("selected", parentId);
}

const selectedPartition = (value: number) => {
    emits("selected", value);
}

const reset = () => {
    key.value = Date.now();
    emits("selected", 0);
}

onBeforeMount(() => {
    // 获取分区列表
    getPartitionAPI().then((res) => {
        if (res.data.code === statusCode.OK) {
            partitionList.value = res.data.data.partitions;
            partitions.value = partitionList.value.filter((item) => {
                return item.parent_id === 0;
            })

        }
    }).catch((err) => {
        notification.error({
            title: '获取分区失败',
            content: "原因:" + err.response.data.msg,
            duration: 5000,
        })
    });
})

defineExpose({
    reset
})
</script>

<style lang="less" scoped>
.select-partition {
    width: 380px;
    display: flex;
    justify-content: space-between;

    .select {
        width: 150px;
    }
}
</style>