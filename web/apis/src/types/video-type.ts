import { UserInfoType } from "./user-type"

// 上传视频
export interface UploadVideoType {
    title: string,
    cover: string,
    desc: string,
    copyright: boolean,
    partition: number
}

// 修改视频信息
export interface ModifyVideoType {
    vid: number,
    title: string,
    cover: string,
    desc: string,
    copyright: boolean
}

// 视频状态
export interface VideoStatusType {
    vid: number,
    status: number
    title: string,
    cover: string,
    desc: string,
    copyright: boolean,
    partition: number
}


export interface BaseVideoType {
    vid: number,
    title: string,
    cover: string,
    created_at: string
}

export interface VideoType extends BaseVideoType {
    desc: string,
    copyright: boolean,
    clicks: number
    author: UserInfoType
}

// 分P
export interface PartListType {
    id: number
    title: string,
    url: string,
    duration: number,
    status: number,
    quality: number,
}
