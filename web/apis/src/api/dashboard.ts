import request from '../request';

// 获取仪表盘卡片数据
export const getCardDataAPI = () => {
    return request.get('v1/dashboard/card/data');
}

// 获取仪表盘趋势数据
export const getTrendDataAPI = () => {
    return request.get('v1/dashboard/trend');
}

// 获取仪表盘视频分区数
export const getPartitionDataAPI = (partitionId: number) => {
    return request.get(`v1/dashboard/partition?partition_id=${partitionId}`);
}
