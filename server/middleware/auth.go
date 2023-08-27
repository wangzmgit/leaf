package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/cache"
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/resp"
	"kuukaa.fun/leaf/service"
	"kuukaa.fun/leaf/util/authentication"
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

		// 读取缓存
		if cache.IsTokenExist(claims.UserId, tokenString) { // token 存在
			// 验证权限
			user := service.SelectUserByID(claims.UserId)
			role := dto.GetRoleString(user.Role)
			if !authentication.Check(role, ctx.FullPath(), ctx.Request.Method) {
				zap.L().Info("权限不足")
				resp.Response(ctx, resp.UnauthorizedError, "", nil)
				ctx.Abort()
				return
			}
			ctx.Set("userId", claims.UserId)
			ctx.Next()

			return
		} else {
			zap.L().Info("token验证失败")
			resp.Response(ctx, resp.UnauthorizedError, "", nil)
			ctx.Abort()
			return
		}
	}
}

func WsAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 读取验证token
		tokenString := ctx.Query("token")
		// 验证并解析token
		_, claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			zap.L().Info("token验证失败: " + err.Error())
			resp.Response(ctx, resp.UnauthorizedError, "", nil)
			ctx.Abort()
			return
		}

		// 读取缓存
		if cache.IsTokenExist(claims.UserId, tokenString) { // token存在 未过期
			ctx.Set("userId", claims.UserId)
			ctx.Next()
		}
	}
}
