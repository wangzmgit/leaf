import request from '../request';
import type { AddHistoryType } from '../types/history-type';

// 上传历史记录
export const addHistoryAPI = (addHistory: AddHistoryType) => {
    return request.post('v1/history/add', addHistory);
}

// 获取播放进度
export const getHistoryProgressAPI = (vid: number) => {
    return request.get(`v1/history/progress/get?vid=${vid}`);
}

// 获取历史记录
export const getHistoryVideoAPI = (page: number,pageSize: number) => {
    return request.get(`v1/history/video/get?page=${page}&page_size=${pageSize}`);
}

