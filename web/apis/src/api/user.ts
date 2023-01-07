import request from '../request';
import type { ModifyUserInfoType, UserLoginType, UserRegisterType } from '../types/user-type'

// 登录
export const loginAPI = (login: UserLoginType) => {
    return request.post('v1/user/login', login);
}

// 注册
export const registerAPI = (register: UserRegisterType) => {
    return request.post('v1/user/register', register);
}

// 获取用户信息
export const getUserInfoAPI = () => {
    return request.get('v1/user/info/get');
}

// 修改用户信息
export const modifyUserInfoAPI = (modify: ModifyUserInfoType) => {
    return request.post('v1/user/info/modify', modify);
}

// 修改用户封面图
export const modifySpaceCoverAPI = (url: string) => {
    return request.post('v1/user/cover/modify', { spaceCover: url });
}

//通过用户ID获取用户信息
export const getOtherUserInfoAPI = (uid: number) => {
    return request.get(`v1/user/info/other?uid=${uid}`);
}

//通过用户名获取用户ID
export const getUserIdAPI = (name: string) => {
    return request.get(`v1/user/uid?name=${name}`);
}