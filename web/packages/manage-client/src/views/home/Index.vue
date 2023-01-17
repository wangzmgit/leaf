<template>
    <div class="home" :style="initTheme()" v-title data-title="后台管理系统">
        <div class="home-menu">
            <span class="title">{{ globalConfig.title }}</span>
            <n-menu :options="menuOptions" />
        </div>
        <div class="home-box">
            <header-bar></header-bar>
            <div class="home-router">
                <router-view></router-view>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { h } from "vue";
import { globalConfig } from '@leaf/utils';
import { NMenu } from "naive-ui";
import { RouterLink } from 'vue-router';
import HeaderBar from '@/components/header-bar/Index.vue';
import { getTheme } from '@/theme';
// import useRenderIcon from "@/hooks/render-icon-hooks";

const initTheme = () => {
    const theme = getTheme();

    return {
        "--title-color": theme.primaryColor
    }
}

const menuOptions = [
    {
        // label: () =>
        //     h(
        //         RouterLink,
        //         {
        //             to: {
        //                 name: "Dashboard",
        //             }
        //         },
        //         { default: () => '仪表盘' }
        //     ),
        label: "仪表盘",
        key: "dashboard",
        // icon: renderIcon(BarChartOutline),
    },
    {
        // label: () =>
        //     h(
        //         RouterLink,
        //         {
        //             to: {
        //                 name: "Review",
        //             }
        //         },
        //         { default: () => '视频审核' }
        //     ),
        label: "视频审核",
        key: "review",
        // icon: renderIcon(CreateOutline),
    },
    {
        label: () =>
            h(
                RouterLink,
                {
                    to: {
                        name: "User",
                    }
                },
                { default: () => '用户管理' }
            ),
        key: "user",
        // icon: renderIcon(PersonOutline),
    },
    {
        // label: () =>
        //     h(
        //         RouterLink,
        //         {
        //             to: {
        //                 name: "Video",
        //             }
        //         },
        //         { default: () => '视频管理' }
        //     ),
        label: "视频管理",
        key: "video",
        // icon: renderIcon(VideocamOutline),
    },
    {
        // label: () =>
        //     h(
        //         RouterLink,
        //         {
        //             to: {
        //                 name: "Announce",
        //             }
        //         },
        //         { default: () => '公告管理' }
        //     ),
        label: "公告管理",
        key: "announce",
        // icon: renderIcon(CheckboxOutline),
    },
    {
        label: () =>
            h(
                RouterLink,
                {
                    to: {
                        name: "Carousel",
                    }
                },
                { default: () => '轮播图管理' }
            ),
        key: "carousel",
        // icon: renderIcon(ImageOutline),
    },
    {
        label: () =>
            h(
                RouterLink,
                {
                    to: {
                        name: "Partition",
                    }
                },
                { default: () => '分区管理' }
            ),
        key: "partition",
        // icon: renderIcon(GitMergeOutline),
    },
    {
        // label: () =>
        //     h(
        //         RouterLink,
        //         {
        //             to: {
        //                 name: "About",
        //             }
        //         },
        //         { default: () => '关于' }
        //     ),
        label: "关于",
        key: "about",
        // icon: renderIcon(InformationCircleOutline),
    },
];

</script>

<style lang="less" scoped>
.home {
    display: flex;
    background: #f7f7f7;

    .home-box {
        width: calc(100% - 220px);
    }

    .home-menu {
        .title {
            display: block;
            height: 60px;
            font-size: 18px;
            color: var(--title-color);
            line-height: 60px;
            text-align: center;
            border-bottom: 1px solid rgba(60, 60, 67, .12);
        }

        width: 220px;
        height: 100vh;
        position: relative;
        background-color: #fff;
        box-shadow: 0 1px 4px rgba(0, 21, 41, .3);
    }

    .home-router {
        width: 100%;
        min-height: 500px;
        min-width: 1000px;
    }
}
</style>