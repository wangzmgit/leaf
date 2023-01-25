import type { UserInfoType } from "./user-type";
import type { BaseVideoType } from "./video-type";

export interface LikeMessageType {
    video: BaseVideoType,
    user: UserInfoType,
    created_at: string,
}