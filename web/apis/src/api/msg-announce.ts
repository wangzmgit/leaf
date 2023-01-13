import request from '../request';

//获取公告
export const getAnnounceAPI = (page: number, page_size: number) => {
    return request.get(`v1/message/announce/get?page=${page}&page_size=${page_size}`);
}
