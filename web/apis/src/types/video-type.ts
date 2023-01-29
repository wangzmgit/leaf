import type { ResourceType } from "./resource-type"
import type { UserInfoType } from "./user-type"

export interface BaseVideoType {
    vid: number,
    title: string,
    cover: string,
    desc: string,
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
export interface ModifyVideoType extends BaseVideoType {
    copyright: boolean,
}

// 视频状态
export interface VideoStatusType extends BaseVideoType {
    status: number,
    copyright: boolean,
    partition: number,
    resources: ResourceType[]
}

// 视频
export interface VideoType extends BaseVideoType {
    copyright: boolean,
    clicks: number,
    author: UserInfoType,
    partition?: number
}

// 视频
export interface UserUploadVideoType extends BaseVideoType {
    clicks: number,
    status: number
}

