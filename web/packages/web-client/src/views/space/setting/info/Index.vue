<template>
    <div class="modify-info-from">
        <n-form label-placement="left" label-width="80px">
            <div class="avatar-box">
                <p class="avatar-label">{{ t("common.avatar") }}:</p>
                <common-avatar class="avatar" :url="getResourceUrl(userInfo.avatar)" :size="60"
                    @click="avatarClick"></common-avatar>
            </div>
            <n-form-item label="UID:">
                <p class="uid form-item">{{ userInfo.uid }}</p>
            </n-form-item>
            <n-form-item :label="`${t('common.nickname')}:`">
                <n-input class="form-input name" v-model:value="userInfo.name" placeholder="请输入昵称" maxlength="20"
                    show-count />
            </n-form-item>
            <n-form-item :label="`${t('common.gender')}:`">
                <n-radio-group class="form-item" v-model:value="userInfo.gender">
                    <n-radio-button :value="0">{{ t("common.unknown") }}</n-radio-button>
                    <n-radio-button :value="1">{{ t("common.male") }}</n-radio-button>
                    <n-radio-button :value="2">{{ t("common.female") }}</n-radio-button>
                </n-radio-group>
            </n-form-item>
            <n-form-item :label="`${t('common.birthday')}:`">
                <n-date-picker class="form-item" placeholder="选择出生日期" type="date" :value="userInfoBirthday"
                    @update-value="updateBirthday" />
            </n-form-item>
            <n-form-item :label="`${t('common.sign')}:`">
                <n-input class="form-input" v-model:value="userInfo.sign" placeholder="请输入个性签名" maxlength="50"
                    show-count type="textarea" :autosize="descSize" />
            </n-form-item>
            <div class="setting-save-btn">
                <n-button type="primary" @click="modifyUserInfo">{{ t("common.modify") }}</n-button>
            </div>
        </n-form>
        <leaf-cropper ref="cropperRef">
            <template #content="fileSlot">
                <avatar-cropper :file="fileSlot.file" @state-change="changeUpload"></avatar-cropper>
            </template>
        </leaf-cropper>
    </div>
</template>
<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { storageData, statusCode, getResourceUrl } from "@leaf/utils";
import { ref, onBeforeMount } from "vue";
import {
    NForm, NFormItem, NButton, NRadioGroup, NInput,
    NRadioButton, NDatePicker, useNotification,
} from 'naive-ui';
import { modifyUserInfoAPI, getUserInfoAPI } from "@leaf/apis";
import { CommonAvatar } from "@leaf/components";
import type { UserInfoType, ModifyUserInfoType } from "@leaf/apis";
import LeafCropper from "@/components/leaf-cropper/Index.vue";
import AvatarCropper from "@/components/leaf-cropper/component/AvatarCropper.vue";

// i18n
const { t } = useI18n();

const notification = useNotification();//通知


//简介输入框大小
const descSize = {
    minRows: 3,
    maxRows: 3
}

const userInfo = ref<UserInfoType>({
    uid: 0,
    avatar: "",
    name: "",
    sign: "",
    gender: 0,
    birthday: "",
});

//用户生日信息
const userInfoBirthday = ref(0);
const updateBirthday = (value: number, formattedValue: string) => {
    userInfoBirthday.value = value;
    userInfo.value.birthday = formattedValue;
}

const cropperRef = ref<InstanceType<typeof LeafCropper> | null>(null);
const avatarClick = () => {
    cropperRef.value?.open();
}

//上传变化的回调
const changeUpload = (status: string, data: any) => {
    switch (status) {
        case "success":
            userInfo.value.avatar = data.data.url;
            notification.success({
                title: '上传完成',
                duration: 3000,
            });
            break;
        case "error":
            notification.error({
                title: '文件上传失败',
                duration: 5000,
            });
            break;
    }
    cropperRef.value?.closeCropper();
}



const modifyUserInfo = () => {
    const modifyForm: ModifyUserInfoType = {
        avatar: userInfo.value.avatar,
        name: userInfo.value.name,
        sign: userInfo.value.sign || "",
        gender: userInfo.value.gender || 0,
        birthday: userInfo.value.birthday || '1970-1-1',
    }

    if (!modifyForm.name) {
        notification.error({
            content: "请填写昵称",
            duration: 5000,
        });
        return;
    }
    modifyUserInfoAPI(modifyForm).then((res) => {
        if (res.data.code === statusCode.OK) {
            getUserInfo();//获取用户信息
            notification.success({
                title: '修改成功',
                duration: 5000,
            })
        } else {
            notification.error({
                title: '修改失败',
                content: res.data.msg,
                duration: 5000,
            })
        }
    }).catch((err) => {
        notification.error({
            title: '修改失败',
            content: err.response.data.msg,
            duration: 5000,
        })
    });
}

//获取用户信息
const getUserInfo = () => {
    getUserInfoAPI().then((res) => {
        if (res.data.code === statusCode.OK) {
            storageData.update("user_info", res.data.data.user_info);
        }
    })
}


onBeforeMount(() => {
    const info = storageData.get('user_info');
    userInfo.value = Object.assign(userInfo.value, info);
    userInfoBirthday.value = Date.parse(info.birthday);
})

</script>

<style lang="less" scoped>
.avatar-box {
    height: 60px;
    margin: 20px 0;
    display: flex;
    position: relative;

    .avatar-label {
        box-sizing: border-box;
        width: 80px;
        margin: 0;
        padding: 0 12px 0 0;
        text-align: right;
        font-size: 14px;
        line-height: 60px;
    }

    .upload-avatar {
        display: none;
    }

    .avatar {
        margin-left: 36px;

        &:hover::after {
            display: block;
        }

        //头像遮罩
        &::after {
            font-size: 10px;
            display: none;
            position: absolute;
            content: "更换头像";
            line-height: 60px;
            vertical-align: middle;
            width: 100%;
            height: 100%;
            cursor: pointer;
            background-color: rgba(0, 0, 0, 0.3);
        }
    }
}

.uid {
    margin: 0;
}

.form-item {
    padding-left: 20px;
}

.name {
    &:deep(.n-input__input-el) {
        height: auto;
    }
}

.form-input {
    margin-left: 20px;
}

.setting-save-btn {
    float: right;
    margin-top: 30px;

    button {
        width: 120px;
        height: 40px;
    }
}
</style>