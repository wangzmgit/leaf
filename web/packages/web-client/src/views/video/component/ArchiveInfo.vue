<template>
    <div class="archive-data" :style="initTheme()">
        <!--点赞收藏-->
        <div class="archive-item">
            <n-icon :class="[likeAnimation, archive.is_like ? 'active' : 'icon']" @click="likeClick">
                <like></like>
            </n-icon>
            <p>0</p>
        </div>
        <div class="archive-item">
            <n-icon :class="archive.is_collect ? 'active' : 'icon'" @click="showCollect = true">
                <collect></collect>
            </n-icon>
            <p>0</p>
        </div>
    </div>
    <!-- <collection-list v-if="showCollect" :vid="vid" @close="closeCollection"></collection-list> -->
</template>

<script setup lang="ts">
import { NIcon } from 'naive-ui';
import { getTheme } from "@/theme"
import { Like, Collect } from "@leaf/icons";
import { reactive, ref, onBeforeMount } from 'vue';

const props = defineProps<{
    vid: Number
}>()

const initTheme = () => {
    const theme = getTheme();

    return {
        "--primary-color": theme.primaryColor
    }
}
// const stat = ref(props.stat);//点赞收藏数据
//点赞收藏数据
const archive = reactive({
    is_collect: false,
    is_like: false
})
const showCollect = ref(false);
const likeAnimation = ref('');

//获取点赞收藏关注信息
const getArchiveInfo = () => {
    // getArchiveAPI(props.vid).then((res) => {
    //     if (res.data.code === 2000) {
    //         const resArchive = res.data.data.archive;
    //         archive.is_like = resArchive.is_like;
    //         archive.is_collect = resArchive.is_collect;
    //     }
    // })
}

//点赞点赞按钮
const likeClick = () => {
    // if (!archive.is_like) {
    //     //调用点赞接口
    //     likeAPI(props.vid);
    //     likeAnimation.value = 'like-active';
    //     stat.value.like++;
    // } else {
    //     dislikeAPI(props.vid);
    //     stat.value.like--;
    // }
    // archive.is_like = !archive.is_like;
}

//关闭收藏夹回调
const closeCollection = (val: number) => {
    // if (val === 1) {
    //     stat.value.collect++;
    //     archive.is_collect = true;
    // } else if (val === -1) {
    //     stat.value.collect--;
    //     archive.is_collect = false;
    // }
    // showCollect.value = false;
}

onBeforeMount(() => {
    getArchiveInfo();
})

</script>

<style lang="less" scoped>
.archive-data {
    height: 30px;

    .archive-item {
        float: left;
        user-select: none;
        margin-right: 20px;

        i {
            font-size: 26px;
            line-height: 30px;
            cursor: pointer;
        }

        p {
            font-size: 16px;
            float: right;
            margin: 0 6px;
            line-height: 30px;
        }

        .icon:hover {
            color: var(--primary-color);
        }

        .active {
            color: var(--primary-color);
        }

        .like-active {
            animation: scaleDraw .3s ease-in-out;
        }
    }
}


@keyframes scaleDraw {
    0% {
        transform: scale(1);
        /*开始为原始大小*/
    }

    25% {
        transform: scale(1.2);
        /*放大1.1倍*/
    }

    100% {
        transform: scale(1);
    }
}
</style>