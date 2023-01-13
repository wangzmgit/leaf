import request from '../request';

//获取回复通知
export const getReplyMessageAPI = (page: number, page_size: number) => {
    return request.get(`v1/message/reply/get?page=${page}&page_size=${page_size}`);
}
