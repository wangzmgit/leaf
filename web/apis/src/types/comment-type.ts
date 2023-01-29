import type { UserInfoType } from "./user-type"

export interface AddCommentType {
    vid: number,
    content: string,
    parentId?: string,
    replyUserId?: number,
    replyContent?: string,
    at: Array<string>
}


export interface CommentType {
    id: string,
    content: string,
    author: UserInfoType,
    reply: Array<ReplyType>,
    created_at: number,
    page: number,//回复页码
    noMore: boolean
}

export interface ReplyType {
    id: string,
    content: string,
    author: UserInfoType,
    created_at: number,
}