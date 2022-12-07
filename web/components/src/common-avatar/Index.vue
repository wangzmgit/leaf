<template>
    <n-avatar v-if="props.url" round :size="props.size" :src="props.url" />
    <n-avatar v-else round :size="props.size">
        <n-icon :size="props.iconsize">
            <avatar></avatar>
        </n-icon>
    </n-avatar>
</template>

<script setup lang="ts">
import { NAvatar, NIcon } from "naive-ui";
import { Avatar } from "@leaf/icons";

const props = withDefaults(defineProps<{
    url?: string
    size?: number
    iconsize?: number
}>(), {
    url: '',
    size: 50,
    iconsize: 26
});
</script>