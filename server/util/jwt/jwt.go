package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"kuukaa.fun/leaf/cache"
)

type Claims struct {
	UserId uint
	// TokenType uint // 0:accessToken,1:refreshtToken
	jwt.RegisteredClaims
}

/**
 * 生成token
 * param: id 用户id
 * return: token字符串、错误信息
 */
func GenerateToken(id uint) (string, error) {
	refreshJwtKey := []byte(viper.GetString("security.jwt_secret"))
	// token过期时间
	expirationTime := time.Now().Add(cache.TOKEN_EXPRIRATION_TIME * time.Hour) // 14天有效

	refreshClaims := &Claims{
		UserId: id,
		// TokenType: 1,
		RegisteredClaims: jwt.RegisteredClaims{
			//发放时间等
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "leaf",
		},
	}
	return generateToken(refreshJwtKey, refreshClaims)
}

/**
 * 获取token的荷载数据
 * param: tokenString token字符串s
 * return: *Claims token的负载结构体
 * return: error 解析token的错误信息
 */
func GetTokenClaims(tokenString string) (*Claims, error) {
	claims := &Claims{}
	parser := jwt.NewParser()
	_, _, err := parser.ParseUnverified(tokenString, claims)
	return claims, err
}

/**
 * 验证token
 * param: tokenString token字符串
 * return: *jwt.Token token结构体
 * return: error 解析token的错误信息
 */
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	secret := []byte(viper.GetString("security.jwt_secret"))
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, e error) {
		return secret, nil
	})

	return token, claims, err
}

/**
 * 通过密钥和负载生成token字符串
 * param: key 密钥
 * claims: jwt负载
 * return: token字符串、错误信息
 */
func generateToken(key []byte, claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
