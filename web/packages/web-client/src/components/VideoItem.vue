<template>
    <div class="video-item">
        <div class="img"></div>
        <div class="video-info">
            <span class="title">【标题】标题标题标题标题标题标题标题标题标题</span>
            <div class="author">
                <div class="avatar"></div>
                <span class="name-date">用户名 · 刚刚</span>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">

</script>

<style lang="less" scoped>
.video-item {
    width: 100%;
    height: 260px;

    .img {
        width: 100%;
        height: 160px;
        border-radius: 10px;
        background-color: rgba(0, 0, 0, .2);
    }

    .video-info {
        margin-top: 8px;

        .title {
            height: 44px;
            color: #18191c;
            padding-right: 12px;
            font-size: 15px;
            line-height: 22px;
            overflow: hidden;
            text-overflow: ellipsis;
            display: -webkit-box;
            -webkit-line-clamp: 2;
            -webkit-box-orient: vertical;
            font-weight: 500;
        }

        .author {
            display: flex;
            align-items: center;
            font-size: 13px;
            color: #9499a0;
            margin-top: 5px;

            .avatar {
                width: 26px;
                height: 26px;
                border-radius: 50%;
                background-color: rgba(0, 0, 0, .2);
            }

            .name-date {
                margin-left: 10px;
            }

        }
    }
}
</style>