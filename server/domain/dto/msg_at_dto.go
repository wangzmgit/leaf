package dto

import "kuukaa.fun/leaf/domain/model"

/**
 * 将用户ID转化为AtMessage结构体切片
 * param: videoId 视频ID
 * param: userId 用户ID
 * param: fid 发送用户ID
 * return: AtMessage结构体切片
 */
func UserIdsToAtMessage(userIds []uint, videoId, fid uint) []model.AtMessage {
	length := len(userIds)
	newAtMsg := make([]model.AtMessage, length)
	for i := 0; i < length; i++ {
		newAtMsg[i].Uid = userIds[i]
		newAtMsg[i].Vid = videoId
		newAtMsg[i].Fid = fid

	}
	return newAtMsg
}
