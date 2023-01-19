package dto

import (
	"kuukaa.fun/leaf/common"
	"kuukaa.fun/leaf/domain/model"
)

type UploadVideoDTO struct {
	Title     string
	Cover     string
	Desc      string
	Copyright bool
	Partition uint //分区ID
}

// 修改视频信息
type ModifyVideoDTO struct {
	VID       uint
	Title     string
	Cover     string
	Desc      string
	Copyright bool
}

// 审核视频
type ReviewDTO struct {
	ID     uint
	Status int
}

/**
 * 上传视频DTO结构体转化为Video结构体
 * param: userId 用户Id
 * param: uploadVideoDTO 上传视频DTO结构体
 * return: Video结构体
 */
func UploadVideoDtoToVideo(userId uint, uploadVideoDTO UploadVideoDTO) model.Video {
	return model.Video{
		Uid:         userId,
		Title:       uploadVideoDTO.Title,
		Cover:       uploadVideoDTO.Cover,
		Desc:        uploadVideoDTO.Desc,
		Copyright:   uploadVideoDTO.Copyright,
		PartitionId: uploadVideoDTO.Partition,
		Status:      common.CREATED_VIDEO,
	}
}
