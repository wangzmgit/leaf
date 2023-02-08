<template>
    <div class="player-container">
        <w-player class="player" :key="playerKey" :options="options"></w-player>
    </div>
</template>

<script setup lang="ts">
import { onBeforeMount, ref, watch } from 'vue';
import { WPlayer } from 'vue-wplayer';
import 'vue-wplayer/dist/style.css';
import type { ResourceType, AddDanmakuType } from '@leaf/apis';
import { sendDanmakuAPI, getDanmakuAPI } from '@leaf/apis';
import type { OptionsType } from '../types/player';
import { getResourceUrl, statusCode } from '@leaf/utils';

const props = withDefaults(defineProps<{
    vid: number,
    theme: string,
    resources: Array<ResourceType>,
    part: number,
}>(), {
    part: 1,
})

const playerKey = ref(0);

const options: OptionsType = {
    resource: "",
    type: "mp4",
    mobile: true,
    theme: props.theme,
    danmaku: {
        open: true,
        data: [],
        send: (danmaku: AddDanmakuType) => {
            sendDanmaku(danmaku);
        }
    }
}

// 加载
const loadPart = async (part: number) => {
    options.resource = getResourceUrl(props.resources[part - 1].url);
    await getDanmaku(part);
    playerKey.value = Date.now();
}

watch(() => props.part, (val) => {
    loadPart(val);
});

const sendDanmaku = (danmakuForm: AddDanmakuType) => {
    danmakuForm.vid = props.vid;
    danmakuForm.part = props.part;
    sendDanmakuAPI(danmakuForm);
}

const getDanmaku = async (part: number) => {
    if (options.danmaku.data) options.danmaku.data = [];
    const res = await getDanmakuAPI(props.vid, part);
    if (res.data.code === statusCode.OK) {
        const list = res.data.data.danmaku;
        if (list) {
            options.danmaku!.data = list;
        }
    }
}

onBeforeMount(() => {
    loadPart(props.part);
});
</script>

<style lang="less" scoped>
.player-container {
    height: 0;
    width: 100%;
    padding-bottom: 56.25%;
    position: relative;
    margin-bottom: 40px;

    .player {
        width: 100%;
        height: 100%;
        position: absolute;
        background-color: black;
    }
}
</style>