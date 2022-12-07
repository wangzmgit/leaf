import request from '../request';

// 发送注册验证码
export const sendRegisterCodeAPI = (email: string) => {
    return request.post("v1/user/emailcode", { email })
}