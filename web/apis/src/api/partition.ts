import request from '../request'

//获取分区
export const getPartitionAPI = () => {
    return request.get(`v1/partition/get`);
}
