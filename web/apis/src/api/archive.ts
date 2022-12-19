import request from '../request';

// 获取点赞收藏数据
export const getArchiveStatAPI = (vid: number) => {
    return request.get('v1/archive/stat?vid=' + vid);
}

// 是否点赞
export const getLikeStatusAPI = (vid: number) => {
    return request.get('v1/archive/has/like?vid=' + vid);
}

// 是否收藏
export const getCollectStatusAPI = (vid: number) => {
    return request.get('v1/archive/has/collect?vid=' + vid);
}

//点赞
export const likeAPI = (id: number) => {
    return request.post('v1/archive/like', { id });
}

//取消赞
export const cancelLikeAPI = (id: number) => {
    return request.post('v1/archive/cancel/like', { id })
}
