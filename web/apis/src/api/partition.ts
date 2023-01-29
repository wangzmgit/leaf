import request from '../request'
import type { AddPartitionType } from '../types/partition-type';

//获取分区
export const getPartitionAPI = () => {
    return request.get(`v1/partition/get`);
}

//新增分区
export const addPartitionAPI = (addPartition: AddPartitionType) => {
    return request.post(`v1/partition/add`, addPartition);
}

//删除分区
export const deletePartitionAPI = (id: number) => {
    return request.post(`v1/partition/delete`, { id });
}
