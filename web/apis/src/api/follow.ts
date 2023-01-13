import request from '../request'

//关注
export const followAPI = (id: number) => {
    return request.post('v1/follow/add', { id });
}

//取消关注
export const unfollowAPI = (id: number) => {
    return request.post('v1/follow/cancel', { id })
}

//获取关注状态
export const getFollowStatusAPI = (fid: number) => {
    return request.get(`v1/follow/status?fid=${fid}`)
}

//获取关注数据
export const getFollowDataAPI = (uid: number) => {
    return request.get(`v1/follow/count?uid=${uid}`)
}

//获取关注列表
export const getFollowingAPI = (uid: number, page: number, page_size: number) => {
    return request.get(`v1/follow/following?uid=${uid}&page=${page}&page_size=${page_size}`)
}

//获取粉丝列表
export const getFollowersAPI = (uid: number, page: number, page_size: number) => {
    return request.get(`v1/follow/follower?uid=${uid}&page=${page}&page_size=${page_size}`)
}