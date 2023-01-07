import { UserInfoType } from "./user-type";
import { BaseVideoType } from "./video-type";

export interface ReplyMessageType {
    video: BaseVideoType,
    user: UserInfoType,
    created_at: string,
    content: string,
    target_reply_content: string,
    root_content: string,
    comment_id: string
}
