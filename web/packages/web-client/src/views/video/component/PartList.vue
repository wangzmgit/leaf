<template>
    <div class="part-head">
        <span class="title">分段列表</span>
        <span class="part">({{ current }}/{{ resources?.length }})</span>
    </div>
    <n-scrollbar style="max-height: 300px;margin-bottom: 20px;">
        <div :class="['video-item', current - 1 === index ? 'active-part' : '']" v-for="(item, index) in resources"
            @click="changePart(index)">
            <n-ellipsis class="item-left" style="max-height: 230px">
                <span>P{{ index + 1 }} {{ item.title }}</span>
            </n-ellipsis>
            <div class="item-right">
                <span class="duration">{{ toDuration(item.duration) }}</span>
            </div>
        </div>
    </n-scrollbar>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { NScrollbar, NEllipsis } from "naive-ui";
import type { ResourceType } from "@leaf/apis";

const emits = defineEmits(['change']);
const props = withDefaults(defineProps<{
    active?: number
    resources: Array<ResourceType>
}>(), {
    active: 1
})

const current = ref(props.active);
const toDuration = (duration: number) => {
    let m: number = Math.floor(duration / 60);
    let s: number = Math.floor(duration % 60);
    const mm = m < 10 ? "0" + m : m;
    const ss = s < 10 ? "0" + s : s;
    return mm + ":" + ss;
}

const changePart = (part: number) => {
    current.value = part + 1;
    emits('change', part + 1)
}
</script>

<style lang="less" scoped>
.part-head {
    display: flex;
    align-items: flex-end;
    justify-content: space-between;
    margin-top: 12px;

    .title {
        color: #333;
        font-size: 16px;
        font-weight: 500;
    }

    .part {
        color: #666;
        float: right;
        font-size: 14px;
    }
}

.active-part {
    color: #18a058 !important;
}

.video-item {
    height: 40px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin: 6px 0;
    color: #4c4c4c;
    padding: 0 10px;
    border: 1px solid #efeff5;
    cursor: pointer;

    &:hover {
        color: #18a058;
    }

    .item-left {
        display: flex;
        align-items: center;
        justify-content: space-between;
    }
}
</style>