import request from '../request';

//获取点赞通知
export const getLikeMessageAPI = (page: number, page_size: number) => {
    return request.get(`v1/message/like/get?page=${page}&page_size=${page_size}`);
}
