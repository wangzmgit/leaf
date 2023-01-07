import { UserInfoType } from "./user-type";
import { BaseVideoType } from "./video-type";

export interface AtMessageType {
    video: BaseVideoType,
    user: UserInfoType,
    created_at: string,
}