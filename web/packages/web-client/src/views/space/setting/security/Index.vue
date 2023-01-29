<template>
    <div class="security-info">
        <div>
            <p>{{ t("common.email") }}:&nbsp;&nbsp;{{ userInfo.email }}</p>
            <n-button @click="modify">修改邮箱</n-button>
        </div>
        <div>
            <p>{{ t("common.pwd") }}:&nbsp;&nbsp;********</p>
            <n-button @click="modifyPwd">修改密码</n-button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onBeforeMount, ref } from 'vue';
import { useI18n } from 'vue-i18n';
import { useRouter } from 'vue-router';
import { useNotification, NButton } from 'naive-ui';
import type { UserInfoType } from '@leaf/apis';
import { storageData } from '@leaf/utils';

const { t } = useI18n();

const router = useRouter();

const userInfo = ref<UserInfoType>({
    uid: 0,
    name: "",
    avatar: ""
});

const notification = useNotification();//通知

const modify = () => {
    notification.info({
        title: '当前版本暂不支持修改',
        duration: 3000,
    });
}

const modifyPwd = () => {
    let findPasswordUrl = router.resolve({ name: "FindPassword" });
    window.open(findPasswordUrl.href, '_blank');
}

onBeforeMount(() => {
    userInfo.value = storageData.get('user_info');
})
</script>

<style lang="less" scoped>
.security-info {
    div {
        height: 60px;
        width: 300px;
        display: flex;
        margin-left: 30px;
        justify-content: space-between;
        align-items: center;

        p {
            margin: 0;
        }

    }
}
</style>