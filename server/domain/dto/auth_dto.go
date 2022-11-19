package dto

import (
	"strconv"
	"time"

	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"kuukaa.fun/leaf/domain/model"
	"kuukaa.fun/leaf/util/random"
)

type LoginDTO struct {
	// 邮箱
	Email string
	// 密码
	Password string
}

type RegisterDTO struct {
	// 邮箱
	Email string
	// 密码
	Password string
	// 邮箱验证码
	EmailCode string
}

/**
 * 注册DTO结构体转化为User结构体
 * param: registerDTO 注册DTO结构体
 * return: User结构体
 */
func RegisterToUser(registerDTO RegisterDTO) model.User {
	// 用户密码加密
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(registerDTO.Password), bcrypt.DefaultCost)
	return model.User{
		Username: generateUniqueUsername(), // 随机生成不重复的用户名
		Email:    registerDTO.Email,
		Password: string(hashedPassword),
	}
}

/**
 * 随机生成一个不重复的用户名
 * return: 用户名字符串
 */
func generateUniqueUsername() string {
	// 前缀 + 时间戳(36进制) + 3位随机数
	return viper.GetString("user.prefix") + strconv.FormatInt(time.Now().UnixNano(), 36) + random.GenerateNumberCode(3)
}
