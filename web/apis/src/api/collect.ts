import request from '../request';
import type { CollectType } from '../types/collect-type';

//收藏
export const collectAPI = (collect: CollectType) => {
    return request.post('v1/archive/collect', collect);
}

// 获取该收藏视频的收藏夹
export const getCollectedCollection = (vid: number) => {
    return request.get(`v1/archive/collect/collected?vid=${vid}`);
}


