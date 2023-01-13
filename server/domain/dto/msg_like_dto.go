package dto

import "kuukaa.fun/leaf/domain/model"

/**
 * 转化为LikeMessage结构体
 * param: videoId 视频ID
 * param: userId 用户ID
 * param: fid 发送用户ID
 * return: LikeMessage结构体
 */
func ToLikeMessage(videoId, userId, fid uint) model.LikeMessage {
	return model.LikeMessage{
		Vid: videoId,
		Uid: userId,
		Fid: fid,
	}
}
