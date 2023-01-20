<template>
    <n-input v-model:value="keyword" round placeholder="搜索" @keydown.enter="search">
        <template #suffix>
            <n-icon @click="search">
                <Search />
            </n-icon>
        </template>
    </n-input>
</template>

<script setup lang="ts">
import { defineComponent, ref } from 'vue';
import { NInput, NIcon } from 'naive-ui';
import { Search } from '@leaf/icons';

const emits = defineEmits(['search']);
const props = defineProps<{
    keyword: string
}>();

//搜索功能
const keyword = ref(props.keyword);
const search = () => {
    emits('search', keyword.value);
}
</script>
