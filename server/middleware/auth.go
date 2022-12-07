package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/cache"
	"kuukaa.fun/leaf/domain/resp"
	"kuukaa.fun/leaf/util/jwt"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 读取验证token
		tokenString := ctx.GetHeader("Authorization")
		// 验证并解析token
		_, claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			zap.L().Info("token验证失败: " + err.Error())
			resp.Response(ctx, resp.UnauthorizedError, "", nil)
			ctx.Abort()
			return
		}
		// 验证token存在 -> 判断token类型
		if claims.TokenType == 0 { // accessToken
			// 读取缓存
			accessToken := cache.GetAccessToken(claims.UserId)
			if accessToken != "" { // accessToken 未过期
				ctx.Set("userId", claims.UserId)
				ctx.Next()
			} else {
				// 刷新accessToken 和 refreshToken
				// refreshDoubleToken(ctx, claims.UserId)
				zap.L().Info("token过期")
				resp.Response(ctx, resp.TokenExpriedError, "", nil)
				return
			}
		} else if claims.TokenType == 1 { // refreshToken
			// 读取缓存
			refreshToken := cache.GetRefreshToken(claims.UserId)
			if refreshToken != "" { // refreshToken 未过期
				// 刷新accessToken 和 refreshToken
				refreshAccessToken(ctx, claims.UserId)
				return
			} else {
				zap.L().Info("token验证失败")
				resp.Response(ctx, resp.UnauthorizedError, "", nil)
				ctx.Abort()
				return
			}
		} else {
			zap.L().Info("token验证失败")
			resp.Response(ctx, resp.UnauthorizedError, "", nil)
			ctx.Abort()
			return
		}
	}
}

func refreshAccessToken(ctx *gin.Context, id uint) {
	// 生成验证token
	var err error
	var accessToken string
	if accessToken, err = jwt.GenerateAccessToken(id); err != nil {
		resp.Response(ctx, resp.Error, "验证token生成失败", nil)
		zap.L().Error("验证token生成失败")
		return
	}

	// 存入缓存
	cache.SetAccessToken(id, accessToken)

	// 返回给前端
	resp.OK(ctx, "", gin.H{"token": accessToken})

	ctx.Abort()
}
