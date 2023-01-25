import request from '../request';
import type { BaseResourceType } from '../types/resource-type';

//修改资源标题
export const modifyTitleAPI = (resourceTitle: BaseResourceType) => {
    return request.post('v1/resource/title/modify', resourceTitle);
}

//删除资源
export const deleteResourceAPI = (id: number) => {
    return request.post('v1/resource/delete', { id });
}