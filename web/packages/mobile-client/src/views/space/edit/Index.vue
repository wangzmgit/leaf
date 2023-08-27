<template>
    <div class="info" v-title data-title="编辑资料">
        <header-bar></header-bar>
        <table class="info-card">
            <tr>
                <td>头像</td>
                <td>
                    <common-avatar :size="50" :src="getResourceUrl(userInfo.avatar)"></common-avatar>
                </td>
            </tr>
            <tr>
                <td>昵称</td>
                <td>{{ userInfo.name }}</td>
            </tr>
            <tr>
                <td>UID</td>
                <td>{{ userInfo.uid }}</td>
            </tr>
            <tr>
                <td>性别</td>
                <td>{{ toGender(userInfo.gender) }}</td>
            </tr>
            <tr>
                <td>出生日期</td>
                <td>
                    <n-time :time="new Date(userInfo.birthday || 0)"></n-time>
                </td>
            </tr>
            <tr>
                <td>个性签名</td>
                <td>{{ userInfo.sign }}</td>
            </tr>
        </table>
        <div class="btn-card">
            <div class="item" @click="logout()">退出登录</div>
            <div class="item" @click="goMine()">返回空间</div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { computed } from "vue";
import { useRouter } from "vue-router";
import { NTime } from "naive-ui";
import { getResourceUrl, storageData } from "@leaf/utils";
import type { UserInfoType } from "@leaf/apis";
import { CommonAvatar } from "@leaf/components";
import HeaderBar from "@/components/header-bar/Index.vue";

const router = useRouter();


const userInfo = computed(() => {
    return storageData.get('user_info') as UserInfoType;
});

const goMine = () => {
    router.push({ name: "SpaceInfo" });
}

const logout = () => {
    //清除token和用户信息并刷新页面
    storageData.remove("token");
    storageData.remove("user_info");
    router.push({ name: "Home" });
}

const toGender = (gender: number | undefined) => {
    if (gender == 1) return "男";
    else if (gender == 2) return "女";
    else return "未知";
}
</script>

<style lang="less" scoped>
.info {
    height: 100%;
    width: 100%;
    top: 0;
    bottom: 0;
    position: fixed;
    background: #f4f4f4;
}

.info-card {
    margin-top: 12px;
    width: 100%;
    background: #fff;
    border-collapse: collapse;
    font-size: 14px;

    tr {
        border-bottom: 1px solid #eee;

        td {
            padding: 10px 0;

            &:nth-child(1) {
                padding-left: 12px;
                text-align: left;
                color: #505050;
                -webkit-user-select: none;
                -moz-user-select: none;
                -ms-user-select: none;
                user-select: none;
            }

            &:nth-child(2) {
                padding-right: 12px;
                text-align: right;
                color: #999;
                max-width: 40%;
                overflow: hidden;
                text-overflow: ellipsis;
                white-space: nowrap;
            }

        }
    }
}


.btn-card {
    margin-top: 20px;

    .item {
        display: block;
        padding: 10px;
        text-align: center;
        color: #505050;
        text-decoration: none;
        background: #fff;
        font-size: 14px;
        -webkit-user-select: none;
        -moz-user-select: none;
        -ms-user-select: none;
        user-select: none;

        &:first-child {
            border-bottom: 1px solid #eee;
        }
    }
}
</style>