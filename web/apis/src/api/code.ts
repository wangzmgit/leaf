import request from '../request';

// 发送验证码
export const sendEmailCodeAPI = (email: string) => {
    return request.post("v1/user/code/email", { email })
}