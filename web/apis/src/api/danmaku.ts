import request from '../request';
import type { AddDanmakuType } from '../types/danmaku-type';

//发送弹幕
export const sendDanmakuAPI = (danmaku: AddDanmakuType) => {
    return request.post('v1/danmaku/send', danmaku);
}

//获取弹幕
export const getDanmakuAPI = (vid: number, part: number) => {
    return request.get(`v1/danmaku/list?vid=${vid}&part=${part}`);
}
