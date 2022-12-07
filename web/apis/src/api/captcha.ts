import request from '../request';

// 获取滑块
export const getSliderAPI = (email: string) => {
    return request.get(`v1/captcha/get?email=${email}`);
}

// 验证滑块
export const validateSliderAPI = (email: string, x: number) => {
    return request.post("v1/captcha/validate", { email, x });
}