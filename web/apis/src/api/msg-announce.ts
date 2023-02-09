import request from '../request';
import type { AddAnnounceType } from '../types/msg-announce-type';

//获取公告
export const getAnnounceAPI = (page: number, page_size: number) => {
    return request.get(`v1/message/announce/get?page=${page}&page_size=${page_size}`);
}

//获取最新重要公告
export const getImportantAnnounceAPI = () => {
    return request.get('v1/message/announce/important/get');
}

//添加公告
export const addAnnounceAPI = (addAnnounce: AddAnnounceType) => {
    return request.post('v1/message/announce/add', addAnnounce);
}

//删除公告
export const deleteAnnounceAPI = (id: number) => {
    return request.post('v1/message/announce/delete', { id });
}
