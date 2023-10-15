<template>
    <div v-title :data-title="`找回密码-${title}`">
        <header-bar></header-bar>
        <div class="find-password">
            <steps class="steps" :current="current" :title-list="['填写账号', '重置密码', '操作成功']"></steps>
            <div class="form" v-if="current === 1">
                <n-form :rules="rules" :model="modifyForm" label-width="auto">
                    <n-form-item label="邮箱" path="email">
                        <n-input placeholder="请输入绑定的邮箱" v-model:value="modifyForm.email"></n-input>
                    </n-form-item>
                    <n-button class="submit" type="primary" @click="checkUser">验证</n-button>
                </n-form>
            </div>
            <div v-else-if="current === 2">
                <n-form ref="formRef" :rules="rules" :model="modifyForm" label-width="auto">
                    <n-form-item label="邮箱">
                        <span>{{ desensiEmail(modifyForm.email) }}</span>
                    </n-form-item>
                    <n-form-item label="密码" path="password">
                        <n-input placeholder="请输入密码" type="password" v-model:value="modifyForm.password" />
                    </n-form-item>
                    <n-form-item label="确认密码" path="reenteredPassword">
                        <n-input placeholder="请输入确认密码" type="password" v-model:value="modifyForm.reenteredPassword" />
                    </n-form-item>
                    <n-form-item label="验证码" path="code">
                        <n-input placeholder="请输入验证码" v-model:value="modifyForm.code" />
                        <n-button :disabled="disabledSend" @click="beforeSendCode">{{ sendBtnText }}</n-button>
                    </n-form-item>
                    <n-button class="submit" type="primary" @click="submitForm">保存</n-button>
                </n-form>
            </div>
            <div v-else-if="current === 3">
                <n-result status="success" title="成功" description="重置密码成功">
                    <template #footer>
                        <n-button @click="goHome">返回首页</n-button>
                    </template>
                </n-result>
            </div>
        </div>
    </div>
    <slider-captcha v-model:show="showCaptcha" :email="modifyForm.email" @success="captchaSuccess"></slider-captcha>
</template>

<script setup lang="ts">
import { reactive, ref } from "vue";
import HeaderBar from "@/components/header-bar/Index.vue"
import { globalConfig, statusCode } from '@leaf/utils';
import type { FormInst, FormItemRule, FormRules } from "naive-ui";
import { NResult, NInput, NButton, NForm, NFormItem, useMessage } from "naive-ui";
import { resetpwdCheckAPI, mpdifyPwdAPI } from "@leaf/apis";
import { SliderCaptcha } from "@leaf/components"
import useSendCode from "@/hooks/send-code-hooks";
import { useRouter } from "vue-router";
import Steps from "@/components/steps/Index.vue"

const title = ref(window.$title || globalConfig.title);

const message = useMessage();

const current = ref(1);
let captchaUsers = "";// 人机验证使用者
const showCaptcha = ref(false);
const modifyForm = reactive({
    email: "",
    password: "",
    reenteredPassword: "",
    code: ""
})

const validatePasswordSame = (rule: FormItemRule, value: string): boolean => {
    return value === modifyForm.password
}

const rules: FormRules = {
    email: [
        { required: true, message: "请输入邮箱", trigger: ['blur', 'input'] },
        { type: "email", message: "请输入正确的邮箱地址", trigger: ['blur', 'input'] },
    ],
    password: [{ required: true, message: '请输入密码', trigger: ['input', 'blur'] }],
    reenteredPassword: [
        { required: true, message: '请再次输入密码', trigger: ['input', 'blur'] },
        { validator: validatePasswordSame, message: '两次密码输入不一致', trigger: ['blur', 'password-input'] }
    ],
    code: { required: true, message: '请输入验证码', trigger: ['blur', 'input'] },
}


const checkUser = () => {
    if (!modifyForm.email) {
        message.error("邮箱不能为空");
        return;
    };
    resetpwdCheckAPI(modifyForm.email).then((res) => {
        if (res.data.code === statusCode.CAPTCHA_REQUIRED) {
            captchaUsers = "resetpwd";
            showCaptcha.value = true;
            return;
        }
        if (res.data.code === statusCode.OK) {
            current.value = 2;
        } else {
            message.error(res.data.msg);
        }
    })
}

// 处理显示的邮箱
const desensiEmail = (email: string) => {
    return email.replace(/(.{0,3}).*@(.*)/, "$1***@$2")
}

const { disabledSend, sendBtnText, sendEmailCodeAsync } = useSendCode();
const beforeSendCode = () => {
    if (!modifyForm.email) {
        return;
    }
    sendEmailCodeAsync(modifyForm.email).then((res) => {
        if (res === statusCode.CAPTCHA_REQUIRED) {
            captchaUsers = "emailcode";
            showCaptcha.value = true;
        }
    })
}

// 提交表单
const formRef = ref<FormInst | null>(null);
const submitForm = () => {
    formRef.value?.validate((errors) => {
        if (!errors) {
            mpdifyPwdAPI(modifyForm).then((res) => {
                if (res.data.code === statusCode.OK) {
                    current.value = 3;
                    message.success("修改成功");
                } else {
                    message.error(res.data.msg);
                }
            })
        } else {
            message.error("请检查输入的数据");
        }
    })
}

const router = useRouter();
const goHome = () => {
    router.push({ name: "Home" });
}

// 人机验证通过
const captchaSuccess = () => {
    if (captchaUsers === "resetpwd") {
        checkUser();
    } else {
        beforeSendCode();
    }
}
</script>

<style lang="less" scoped>
.find-password {
    width: 800px;
    margin: 30px auto;

    .steps {
        width: 100%;
    }

    .form {
        width: 520px;
        margin-left: 160px;

        .submit {
            width: 100%;
            margin-top: 30px;
        }
    }
}
</style>