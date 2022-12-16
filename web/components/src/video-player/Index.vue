<template>
    <div class="player-container">
        <w-player class="player" :key="playerKey" :options="options"></w-player>
    </div>
</template>

<script setup lang="ts">
import { onBeforeMount, onMounted, ref, watch } from 'vue';
import dashjs from "dashjs";
import { WPlayer } from 'vue-wplayer';
import 'vue-wplayer/dist/style.css';
import { ResourceType } from '@leaf/apis';
import { OptionsType, QualityType } from '../types/player';

const props = withDefaults(defineProps<{
    vid: number,
    theme: string,
    resources: Array<ResourceType>,
    part: number
}>(), {
    part: 1
})

const playerKey = ref(0);
let dash: dashjs.MediaPlayerClass;

const options: OptionsType = {
    resource: [],
    type: "dash",
    customType: (player: HTMLVideoElement, url: string) => {
        dash = dashjs.MediaPlayer().create();
        dash.initialize(player, url, false);
    },
    customQualityChange: (quality: string) => {
        const trackIndex = getTrackIndex(Number(quality));
        const tracks = dash.getTracksFor("video");
        if (trackIndex >= tracks.length) {
            dash.setCurrentTrack(tracks[tracks.length - 1])
        } else {
            dash.setCurrentTrack(tracks[trackIndex])
        }
    },
    theme: props.theme,
    danmaku: {
        open: true,
        data: [],
    }
}

const getTrackIndex = (quality: number) => {
    switch (quality) {
        case 360:
            return 0;
        case 480:
            return 1;
        case 720:
            return 2;
        case 1080:
            return 3;
        default:
            return Number.MAX_SAFE_INTEGER
    }
}


const loadPart = async (part: number) => {
    loadResource(part);

    playerKey.value = Date.now();
}

const loadResource = (part: number) => {
    let tmpResource: QualityType = {};

    switch (props.resources[part - 1].quality) {
        case 1080:
            tmpResource[1080] = {
                name: "1080P",
                url: props.resources[part - 1].url,
            }
        case 720:
            tmpResource[720] = {
                name: "720P",
                url: props.resources[part - 1].url,
            }
        case 480:
            tmpResource[480] = {
                name: "480P",
                url: props.resources[part - 1].url,
            }
        case 360:
            tmpResource[360] = {
                name: "360P",
                url: props.resources[part - 1].url,
            }
    }

    options.resource = tmpResource;
}

watch(() => props.part, (val) => {
    loadPart(val);
});

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