package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", viper.GetString("cors.allow_origin"))
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "authorization,Authorization,DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}

	}
}
