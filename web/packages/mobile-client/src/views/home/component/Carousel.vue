<template>
    <n-carousel autoplay show-arrow>
        <div class="carousel" v-for="(item, index) in carousels" :key="index">
            <img class="carousel-img" :src="getResourceUrl(item.img)" alt="轮播图" />
        </div>
    </n-carousel>
</template>

<script setup lang="ts">
import { NCarousel } from 'naive-ui';
import { ref, onBeforeMount } from 'vue';
import { getResourceUrl, statusCode } from "@leaf/utils";
import type { CarouselType } from '@leaf/apis';
import { getCarouselAPI } from '@leaf/apis';

const carousels = ref(Array<CarouselType>());
const getCarousel = () => {
    getCarouselAPI().then((res) => {
        if (res.data.code === statusCode.OK) {
            carousels.value = res.data.data.carousels;
        }
    })
}

onBeforeMount(() => {
    getCarousel();
})
</script>

<style lang="less" scoped>
.carousel {
    width: 100%;
    height: 100%;

    .carousel-img {
        width: 100%;
        height: 100%;
        object-fit: fill;
    }
}
</style>