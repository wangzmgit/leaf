<template>
    <div class="login-form">
        <!-- 账号登录 -->
        <n-tabs default-value="account" size="large" justify-content="space-evenly" @update:value="loginTypeChange">
            <n-tab-pane class="form-container" name="account" :tab="t('login.accountLogin')">
                <n-form ref="accountFormRef" :rules="rules" :model="loginForm" label-placement="left" label-width="70">
                    <n-form-item :label="t('common.email')" path="email">
                        <n-input :placeholder="t('input.email')" v-model:value="loginForm.email" />
                    </n-form-item>
                    <n-form-item :label="t('common.pwd')" path="password">
                        <n-input :placeholder="t('input.pwd')" v-model:value="loginForm.password" type="password">
                            <template #suffix>
                                <n-button type="primary" text @click="findPassword">{{ t('login.forgotPwd') }}</n-button>
                            </template>
                        </n-input>
                    </n-form-item>
                </n-form>
            </n-tab-pane>
            <!-- 邮箱登录 -->
            <n-tab-pane class="form-container" name="email" :tab="t('login.emailLogin')">
                <n-form ref="emailFormRef" :rules="rules" :model="loginForm" label-placement="left" label-width="70">
                    <n-form-item :label="t('common.email')" path="email">
                        <n-input :placeholder="t('input.email')" v-model:value="loginForm.email" />
                    </n-form-item>
                    <n-form-item :label="t('common.code')" path="code">
                        <n-input :placeholder="t('input.code')" v-model:value="loginForm.code" />
                        <n-button :disabled="disabledSend" @click="beforeSendCode">{{ sendBtnText }}</n-button>
                    </n-form-item>
                </n-form>
            </n-tab-pane>
        </n-tabs>
        <div class="login-btn">
            <n-button @click="emits('changeForm')">{{ t("login.registerNow") }}</n-button>
            <n-button type="primary" @click="sendLoginRequest">{{ t("common.login") }}</n-button>
        </div>
    </div>
    <slider-captcha v-model:show="showCaptcha" :email="loginForm.email" @success="captchaSuccess"></slider-captcha>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue';
import { useI18n } from 'vue-i18n';
import useSendCode from '@/hooks/send-code-hooks';
import type { UserLoginType } from "@leaf/apis";
import { loginAPI, emailLoginAPI, getUserInfoAPI } from "@leaf/apis";
import { storageData, statusCode } from '@leaf/utils';
import type { FormRules, FormInst } from 'naive-ui';
import { NTabs, NTabPane, NForm, NFormItem, NInput, NButton, useNotification } from 'naive-ui';
import { SliderCaptcha } from "@leaf/components";
import { useRouter } from 'vue-router';

const emits = defineEmits(["changeForm", "success"]);

const router = useRouter();
// i18n
const { t } = useI18n();

//通知组件
const notification = useNotification();

// 显示滑块验证
const showCaptcha = ref(false);

//登录表单
const loginForm = reactive<UserLoginType>({
    email: "",
    code: "",
    password: "",
});


//校验规则
const rules: FormRules = {
    email: [
        { required: true, message: "请输入邮箱", trigger: ['blur', 'input'] },
        { type: "email", message: "请输入正确的邮箱地址", trigger: ['blur', 'input'] },
    ],
    password: { required: true, message: '请输入密码', trigger: ['blur', 'input'] },
    code: { required: true, message: '请输入验证码', trigger: ['blur', 'input'] },
}

// 滑块验证通过
let captchaUsers = "";// 人机验证使用者
const captchaSuccess = () => {
    if (captchaUsers === "login") {
        sendLoginRequest();
    } else {
        beforeSendCode();
    }
}

//发送验证码相关
const { disabledSend, sendBtnText, sendEmailCodeAsync } = useSendCode();
const beforeSendCode = () => {
    if (!loginForm.email) {
        return;
    }
    sendEmailCodeAsync(loginForm.email).then((res) => {
        if (res === statusCode.CAPTCHA_REQUIRED) {
            captchaUsers = "emailcode";
            showCaptcha.value = true;
        }
    })
}

// 登录方式切换
const currentTabName = ref("account");
const loginTypeChange = (tabName: string) => {
    currentTabName.value = tabName;
}

//登录相关
const emailFormRef = ref<FormInst | null>(null);
const accountFormRef = ref<FormInst | null>(null);
//登录
const sendLoginRequest = () => {
    let currentRef: FormInst | null = null;
    switch (currentTabName.value) {
        case "account":
            emailLoginAPI
            currentRef = accountFormRef.value;
            break;
        case "email":
            currentRef = emailFormRef.value;
            break;
    }
    currentRef?.validate((err) => {
        if (!err) {
            const loginRequest = currentTabName.value === "account" ? loginAPI : emailLoginAPI;
            loginRequest(loginForm).then((res) => {
                switch (res.data.code) {
                    case statusCode.CAPTCHA_REQUIRED:
                        captchaUsers = "login";
                        showCaptcha.value = true;
                        break;
                    case statusCode.OK:
                        storageData.set("token", res.data.data.token, 14 * 24 * 60);
                        getUserInfoAPI().then((infoRes) => {
                            if (infoRes.data.code === statusCode.OK) {
                                storageData.set("user_info", infoRes.data.data.user_info, 14 * 24 * 60);
                            }
                            emits("success");
                        })
                        break;
                    default:
                        notification.error({
                            title: res.data.msg,
                            duration: 3000,
                        });
                }
            });
        } else {
            notification.error({
                title: '请检查输入的数据',
                duration: 3000,
            })
        }
    });
}

const findPassword = () => {
    let findPasswordUrl = router.resolve({ name: "FindPassword" });
    window.open(findPasswordUrl.href, '_blank');
}
</script>

<style lang="less" scoped>
.login-form {
    .form-container {
        box-sizing: border-box;
        padding: 40px 30px 0 30px;
    }

    .login-btn {
        display: flex;
        justify-content: space-between;
        margin: 20px 30px 0;

        button {
            width: 160px;
        }
    }
}
</style>