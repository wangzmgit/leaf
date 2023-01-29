import type { UserInfoType } from "./user-type";

export interface WhisperType {
    fid: number,
    content: string,
}

export interface WhisperListType {
    user: UserInfoType,
    created_at: string | Date,
    status: boolean
}

export interface WhisperDetailsType {
    fid: number,
    from_id: number,
    content: string,
    created_at: string
}