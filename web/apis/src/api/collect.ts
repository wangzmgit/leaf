import request from '../request';
import { CollectType, } from '../types/collect-type';

//收藏
export const collectAPI = (collect: CollectType) => {
    return request.post('v1/archive/collect', collect);
}

//获取收藏夹内的视频
export const getCollectVideoAPI = (id: number, page: number, page_size: number) => {
    return request.get(`v1/archive/collect/get?id=${id}&page=${page}&page_size=${page_size}`);
}

export const getCollectedCollection = (vid:number ) => {
    return request.get(`v1/archive/collect/collected?vid=${vid}`);
}


