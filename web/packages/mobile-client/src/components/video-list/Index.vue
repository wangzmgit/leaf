<template>
    <div class="card-list">
        <div class="card-box" v-for="(item, index) in props.videos" :key="index">
            <div class="card" @click="govideo(item.vid)">
                <img class="cover" :src="getResourceUrl(item.cover)" />
                <div class="info">
                    <p class="title">{{ item.title }}</p>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useRouter } from "vue-router";
import type { VideoType } from '@leaf/apis';
import { getResourceUrl } from "@leaf/utils";

const props = defineProps<{
    videos: Array<VideoType>
}>();

const router = useRouter();
const govideo = (vid: number) => {
    router.push({ name: "Video", params: { vid: vid } });
}

</script>

<style lang="less" scoped>
.card-list {
    width: 100%;
    display: grid;
    grid-template-columns: repeat(2, 50%);

    .card-box {
        margin: 0 auto;
        height: 100%;
        width: calc(100% - 10px);

        .card {
            position: relative;
            width: 100%;
            height: 150px;
            margin-top: 15px;
            border-radius: 6px;

            .cover {
                width: 100%;
                height: 100%;
                border-radius: 6px;
            }

            .info {
                color: #fff;
                overflow: hidden;
                position: absolute;
                bottom: 0;
                width: 100%;
                border-bottom-left-radius: inherit;
                border-bottom-right-radius: inherit;
                background: linear-gradient(0deg, rgba(0, 0, 0, 0.7), transparent);

                .title {
                    padding-left: 10px;
                    font-size: 14px;
                    line-height: 16px;
                    overflow: hidden;
                    text-overflow: ellipsis;
                    display: -webkit-box;
                    -webkit-line-clamp: 1;
                    -webkit-box-orient: vertical;
                }
            }
        }
    }
}
</style>