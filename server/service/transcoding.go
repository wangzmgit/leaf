package service

import (
	"errors"
	"os"

	"strconv"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/common"
	"kuukaa.fun/leaf/util/number"
	"kuukaa.fun/leaf/util/transcoding"
)

// 预处理视频
func PreTreatmentVideo(input string) (int, float64, error) {
	videoData, err := transcoding.GetVideoInfo(input)
	if err != nil {
		return 0, 0, err
	}

	if videoData.Stream[0].CodecName != "h264" {
		return 0, 0, errors.New("not h264")
	}

	//计算最大分辨率
	width := videoData.Stream[0].Width
	height := videoData.Stream[0].Height
	quality := number.Min(getWidthRes(width), getHeigthRes(height))

	//获取视频时长
	duration, _ := strconv.ParseFloat(videoData.Stream[0].Duration, 64)

	return quality, duration, err
}

// 转码
func VideoTransCoding(resourceId uint, quality int, dirName string) {
	localDir := "./upload/video/" + dirName + "/"
	inputFile := localDir + "upload.mp4"

	// 提取音频
	audioFile, err := transcoding.ExtractingAudio(inputFile, localDir)
	if err != nil {
		zap.L().Error("音频提取失败" + err.Error())
		completeTransCoding(resourceId, common.PROCESSING_FAIL)
		return
	}

	// 生成不同分辨率的MP4
	videoFiles, err := transcoding.PressingVideo(inputFile, localDir, quality)
	if err != nil {
		zap.L().Error("分辨率处理失败" + err.Error())
		completeTransCoding(resourceId, common.PROCESSING_FAIL)
		return
	}

	// 生成dash分片
	err = transcoding.GenerateDash(videoFiles, audioFile, localDir, dirName)
	if err != nil {
		zap.L().Error("分片生成失败" + err.Error())
		completeTransCoding(resourceId, common.PROCESSING_FAIL)
		return
	}

	// 删除临时文件
	os.Remove(audioFile)
	for _, v := range videoFiles {
		os.Remove(v)
	}

	// 上传oss
	if viper.GetString("oss.type") != "local" {
		if err := UploadVideoToOss(dirName); err != nil {
			zap.L().Error("视频上传OSS失败" + err.Error())
		}
	}

	// 完成转码
	completeTransCoding(resourceId, common.WAITING_REVIEW)
}

// 获取宽度支持的最大分辨率
func getWidthRes(width int) int {
	//1920*1080
	if width >= 1920 {
		return 1080
	}
	// 1280*720
	if width >= 1280 {
		return 720
	}
	//720*480
	if width >= 720 {
		return 480
	}
	return 360
}

// 获取高度支持的最大分辨率
func getHeigthRes(height int) int {
	//1920*1080
	if height >= 1080 {
		return 1080
	}
	// 1280*720
	if height >= 720 {
		return 720
	}
	//720*480
	if height >= 480 {
		return 480
	}
	return 360
}

// 完成转码
func completeTransCoding(resourceId uint, status int) {
	// 更新资源状态
	UpadteResourceStatus(resourceId, status)
	// 获取资源信息
	resource := SelectResourceByID(resourceId)
	// 获取转码中资源的数量
	count := SelectResourceCountByStatus(resource.Vid, common.VIDEO_PROCESSING)
	// 如果没有转码中的视频，则更新视频为待审核
	if count == 0 {
		// 获取视频审核状态
		video := GetVideoInfo(resource.Vid)
		if video.Status == common.SUBMIT_REVIEW {
			UpadteVideoStatus(video.ID, common.WAITING_REVIEW)
		}
	}
}
