package service

import (
	"strconv"
	"time"

	"github.com/spf13/viper"
	"github.com/wangzmgit/unioss"
	"kuukaa.fun/leaf/util/random"
)

func UploadImgService(dir, fileName string) (string, error) {
	objectKey := dir + "/" + fileName
	filePath := "./upload/" + objectKey
	if viper.GetString("oss.type") != "local" {
		// 调用OSS上传
		oss, err := unioss.GetStorage()
		if err != nil {
			return "", err
		}
		oss.PutObjectFromFile(objectKey, filePath)
		//获取文件url
		return oss.GetObjectUrl(objectKey), nil
	}

	// 返回文件URL
	return GenerateUrl(objectKey), nil
}

/**
 * 随机生成一个不重复的文件名
 * return: 文件名字符串
 */
func GenerateUniqueFilename(prefix, suffix string) string {
	// 前缀 + 时间戳(36进制) + 3位随机数 + 后缀
	return prefix + strconv.FormatInt(time.Now().UnixNano(), 36) + random.GenerateNumberCode(3) + suffix
}

//生成文件url
func GenerateUrl(objectKey string) string {
	protocol := "http://"
	if viper.GetBool("file.https") {
		protocol = "https://"
	}
	if len(viper.GetString("file.domain")) == 0 {
		return "/api/" + objectKey
	} else {
		return protocol + viper.GetString("file.domain") + "/api/" + objectKey
	}
}
