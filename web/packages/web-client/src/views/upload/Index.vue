<template>
    <div class="upload">
        <header-bar></header-bar>
        <div class="upload-container">
            <n-menu class="upload-menu" :options="menuOptions" :default-value="defaultOption" />
            <div class="upload-router">
                <router-view></router-view>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { h, ref, onBeforeMount } from "vue";
import { NMenu } from "naive-ui";
import { RouterLink, useRoute } from 'vue-router';
import HeaderBar from '@/components/header-bar/Index.vue';
import useRenderIcon from "@/hooks/render-icon-hooks";
import { Upload, Video, Comment } from "@leaf/icons";

const route = useRoute();
const defaultOption = ref('');//默认激活菜单

const { renderIcon } = useRenderIcon();
const menuOptions = [
    {
        label: () =>
            h(
                RouterLink,
                {
                    to: {
                        name: "UploadVideo",
                    }
                },
                { default: () => '发布视频' }
            ),
        key: "upload",
        icon: renderIcon(Upload),
    },
    {
        // label: () =>
        //     h(
        //         RouterLink,
        //         {
        //             to: {
        //                 name: "VideoManage",
        //             }
        //         },
        //         { default: () => '稿件管理' }
        //     ),
        label: "稿件管理",
        key: "content",
        icon: renderIcon(Video),
    },
    {
        // label: () =>
        //     h(
        //         RouterLink,
        //         {
        //             to: {
        //                 name: "CommentManage",
        //             }
        //         },
        //         { default: () => '评论管理' }
        //     ),
        label: "评论管理",
        key: "comment",
        icon: renderIcon(Comment),
    },
];



onBeforeMount(() => {
    switch (route.name) {
        case 'UploadVideo':
            defaultOption.value = 'upload';
            break;
        // case 'VideoManage':
        //     defaultOption.value = 'content';
        //     break;
        // case 'CommentManage':
        //     defaultOption.value = 'comment';
        //     break;
        default:
            defaultOption.value = 'upload';
            break;
    }
})
</script>

<style lang="less" scoped>
.upload {
    top: 0;
    bottom: 0;
    width: 100%;
    height: 100%;
    position: fixed;
    overflow-y: scroll;
    background: #f7f7f7;

    .upload-container {
        display: flex;
        width: 1100px;
        margin: 20px auto;
    }

    .upload-menu {
        width: 220px;
        height: 500px;
        min-width: 200px;
        background-color: #fff;
    }

    .upload-router {
        flex: 1;
        min-width: 700px;
        min-height: 500px;
        margin-left: 20px;
        background-color: #fff;
    }
}


@media (min-width: 1400px) {
    .upload {
        .upload-container {
            width: 1300px;
        }
    }
}
</style>