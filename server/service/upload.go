package service

import (
	"os"
	"strconv"
	"time"

	"github.com/spf13/viper"
	"github.com/wangzmgit/unioss"
	"kuukaa.fun/leaf/util/random"
)

// 上传图片到OSS
func UploadImgToOss(fileName string) error {
	objectKey := "image/" + fileName
	filePath := "./upload/" + objectKey
	oss, err := unioss.GetStorage()
	if err != nil {
		return err
	}
	return oss.PutObjectFromFile(objectKey, filePath)
}

// 上传视频到OSS
func UploadVideoToOss(dirName string) error {
	oss, err := unioss.GetStorage()
	if err != nil {
		return err
	}

	files, err := os.ReadDir("./upload/video/" + dirName)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.Name() == "upload.mp4" && !viper.GetBool("file.video_adaptation_ios") {
			continue
		}

		objectKey := "video/" + dirName + "/" + f.Name()
		filePath := "./upload/" + objectKey
		oss.PutObjectFromFile(objectKey, filePath)
	}

	return nil
}

/**
 * 随机生成图片文件名
 * return: 文件名字符串
 */
func GenerateImgFilename(suffix string) string {
	// 前缀 + 时间戳(36进制) + 3位随机数 + 后缀
	return "img_" + strconv.FormatInt(time.Now().UnixNano(), 36) + random.GenerateNumberCode(3) + suffix
}

/**
 * 随机生成一个视频文件名
 * return: 文件名字符串
 */
func GenerateVideoFilename() string {
	// 前缀 + 时间戳(36进制) + 3位随机数
	return "v_" + strconv.FormatInt(time.Now().UnixNano(), 36) + random.GenerateNumberCode(3)
}

// 生成文件url
func GenerateFileUrl(objectKey string) (string, error) {
	if viper.GetString("oss.type") != "local" {
		oss, err := unioss.GetStorage()
		if err != nil {
			return "", err
		}

		return oss.GetObjectUrl(objectKey), nil
	}

	return "/api/" + objectKey, nil
}
