import { ResourceType } from "./resource-type"
import { UserInfoType } from "./user-type"

export interface BaseVideoType {
    vid: number,
    title: string,
    cover: string,
    created_at: string
}

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
    partition: number,
    resources: ResourceType[]
}

// 视频
export interface VideoType extends BaseVideoType {
    desc: string,
    copyright: boolean,
    clicks: number
    author: UserInfoType
}

