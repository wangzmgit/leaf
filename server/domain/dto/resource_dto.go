package dto

import (
	"kuukaa.fun/leaf/common"
	"kuukaa.fun/leaf/domain/model"
)

/**
 * resDTO结构体转化为Resource结构体
 * param: vid 视频id
 * param: quality 视频最大分辨率
 * param: duration 视频时长
 * param: resDTO 不同分辨率DTO结构体
 * return: Resource结构体
 */
func ResourceDtoToResource(vid uint, quality int, duration float64, url string) model.Resource {
	return model.Resource{
		Vid:      vid,
		Url:      url,
		Duration: duration,
		Status:   common.VIDEO_PROCESSING,
		Quality:  quality,
	}
}
