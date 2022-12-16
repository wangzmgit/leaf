import request from '../request';
import { ModifyVideoType, UploadVideoType } from '../types/video-type';

//上传视频信息
export const uploadVideoInfoAPI = (uploadVideo: UploadVideoType) => {
    return request.post('v1/video/info/upload', uploadVideo);
}

//获取视频状态
export const getVideoStatusAPI = (vid: number) => {
    return request.get('v1/video/status?vid=' + vid);
}

//修改视频信息
export const modifyVideoInfoAPI = (modifyVideo: ModifyVideoType) => {
    return request.post('v1/video/info/modify', modifyVideo);
}

//获取视频信息
export const getVideoInfoAPI = (vid: number) => {
    return request.get(`v1/video/get?vid=${vid}`);
}


//提交审核
export const submitReviewAPI = (id: number) => {
    return request.post('v1/video/review/submit', { id });
}