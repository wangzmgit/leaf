<template>
    <div class="home" :style="initTheme()">
        <home-header class="home-header" @change-fold="changeMenuFold"></home-header>
        <div class="home-content">
            <div class="home-left" :class="menuFold ? 'home-left-fold' : ''">
                <home-sidebar class="home-sidebar" :fold="menuFold"></home-sidebar>
            </div>
            <div class="home-right">
                <div class="home-recommended" :class="menuFold ? 'recommended-fold' : ''">
                    <div class="recommended-carousel">
                        <HomeCarousel></HomeCarousel>
                    </div>
                    <VideoItem></VideoItem>
                    <VideoItem></VideoItem>
                    <VideoItem></VideoItem>
                    <VideoItem></VideoItem>
                    <VideoItem></VideoItem>
                    <VideoItem></VideoItem>
                    <VideoItem></VideoItem>
                    <VideoItem></VideoItem>
                    <VideoItem></VideoItem>
                    <VideoItem></VideoItem>
                    <VideoItem></VideoItem>
                    <VideoItem></VideoItem>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { getTheme } from '@/theme';
import HomeSidebar from './component/HomeSidebar.vue';
import HomeHeader from "./component/HomeHeader.vue";
import HomeCarousel from './component/HomeCarousel.vue';
import VideoItem from '@/components/VideoItem.vue';


const menuFold = ref(false);
const changeMenuFold = (val: boolean) => {
    menuFold.value = val;
    console.log('val', val)
}

const initTheme = () => {
    const theme = getTheme();

    return {
        "--primary-color": theme.primaryColor
    }
}
</script>

<style lang="less" scoped>
.home {
    width: 100%;
    min-width: 1200px;
    height: 100vh;

    .home-header {
        position: fixed;
        top: 0;
        width: 100%;
        z-index: 999;
        background-color: #fff;
    }
}

.home-content {
    display: flex;
    margin-top: 60px;
    .home-left {
        width: 220px;
        transition: width .25s;

        .home-sidebar {
            position: fixed;
        }
    }

    .home-left-fold {
        width: 50px;
    }

    .home-right {
        flex: 1;
    }
}

.home-recommended {
    display: grid;
    margin-left: 20px;
    width: calc(100% - 50px);
    gap: 0 16px;
    grid-template-columns: repeat(4, 1fr);
    height: 1040px;
    overflow: hidden;

    .recommended-carousel {
        height: 420px;
        grid-row: 1/ span 2;
        grid-column: 1/ span 2;
    }
}

.recommended-fold {
    height: 780px;
    grid-template-columns: repeat(5, 1fr);
}
</style>