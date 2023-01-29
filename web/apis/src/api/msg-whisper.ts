import request from '../request';
import type { WhisperType } from '../types/msg-whisper-type';

//获取私信列表
export const getWhisperListAPI = () => {
    return request.get('v1/message/whisper/list')
}

//获取私信列表
export const getWhisperDetailsAPI = (fid: number, page: number, page_size: number) => {
    return request.get(`v1/message/whisper/details?fid=${fid}&page=${page}&page_size=${page_size}`)
}

//发送私信
export const sendWhisperAPI = (msg: WhisperType) => {
    return request.post('v1/message/whisper/send', msg)
}

//已读私信
export const readWhisperAPI = (id: number) => {
    return request.post('v1/message/whisper/read', { id })
}