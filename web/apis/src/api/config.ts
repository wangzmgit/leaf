import request from '../request';
import type { EmailConfigType, OtherConfigType, StorageConfigType } from '../types/config-type';

//获取其他配置
export const getOtherConfigAPI = () => {
    return request.get('v1/config/other/get');
}

//修改其他配置
export const setOtherConfigAPI = (config: OtherConfigType) => {
    return request.post('v1/config/other/set', config);
}

//获取存储配置
export const getStorageConfigAPI = () => {
    return request.get('v1/config/storage/get');
}

//修改存储配置
export const setStorageConfigAPI = (config: StorageConfigType) => {
    return request.post('v1/config/storage/set', config);
}


//获取邮箱配置
export const getEmailConfigAPI = () => {
    return request.get('v1/config/email/get');
}

//修改邮箱配置
export const setEmailConfigAPI = (config: EmailConfigType) => {
    return request.post('v1/config/email/set', config);
}
