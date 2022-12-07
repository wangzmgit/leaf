<template>
    <div class="login-form">
        <!-- 账号登录 -->
        <n-tabs default-value="account" size="large" justify-content="space-evenly">
            <n-tab-pane class="form-container" name="account" tab="邮箱注册">
                <n-form ref="emailFormRef" :rules="rules" :model="registerForm" label-placement="left" label-width="70">
                    <n-form-item label="邮箱" path="email">
                        <n-input placeholder="请输入邮箱" v-model:value="registerForm.email" />
                    </n-form-item>
                    <n-form-item label="密码" path="password">
                        <n-input placeholder="请输入密码" type="password" v-model:value="registerForm.password" />
                    </n-form-item>
                    <n-form-item label="验证码" path="emailcode">
                        <n-input placeholder="请输入验证码" v-model:value="registerForm.emailcode" />
                        <n-button :disabled="disabledSend" @click="beforeSendCode">{{ sendBtnText }}</n-button>
                    </n-form-item>
                </n-form>
            </n-tab-pane>
            <!-- 邮箱登录 -->
            <n-tab-pane name="email" tab="手机号注册" :disabled="true"></n-tab-pane>
        </n-tabs>
        <div class="login-btn">
            <n-button @click="emits('changeForm')">返回登录</n-button>
            <n-button type="primary" @click="sendRegisterRequest">注册</n-button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue';
import useSendCode from '@/hooks/send-code-hooks';

import type { UserRegisterType } from "@leaf/apis";
import { registerAPI } from "@leaf/apis";
import { isEmail } from '@leaf/utils';

import type { FormRules, FormInst } from 'naive-ui';
import { NTabs, NTabPane, NForm, NFormItem, NInput, NButton, useNotification } from 'naive-ui';

const emits = defineEmits(["changeForm"]);

//通知组件
const notification = useNotification();

//登录表单
const registerForm = reactive<UserRegisterType>({
    email: "",
    password: "",
    emailcode: ""
});

//校验规则
const rules: FormRules = {
    email: [
        { required: true, message: "请输入邮箱", trigger: ['blur', 'input'] },
        { type: "email", message: "请输入正确的邮箱地址", trigger: ['blur', 'input'] },
    ],
    password: { required: true, message: '请输入密码', trigger: ['blur', 'input'] },
    emailcode: { required: true, message: '请输入验证码', trigger: ['blur', 'input'] },
}

//发送验证码相关
const { disabledSend, sendBtnText, sendEmailCode } = useSendCode();
const beforeSendCode = () => {
    if (!registerForm.email || !isEmail(registerForm.email)) {
        return;
    }
    sendEmailCode(registerForm.email);
}


//登录相关
const emailFormRef = ref<FormInst | null>(null);
//登录
const sendRegisterRequest = () => {

    emailFormRef.value?.validate((err) => {
        if (!err) {
            registerAPI(registerForm).then((res) => {
                console.log('res', res)
            });
        } else {
            notification.error({
                title: '请检查输入的数据',
                duration: 3000,
            })
        }
    });
}
</script>

<style lang="less" scoped>
.login-form {
    .form-container {
        box-sizing: border-box;
        padding: 30px 30px 0 30px;

        .login-type-text {
            display: block;
            text-align: center;
            font-size: 22px;
            margin: 16px 0;
        }
    }

    .login-btn {
        display: flex;
        justify-content: space-between;
        margin: 6px 30px 0;

        button {
            width: 160px;
        }
    }
}
</style>