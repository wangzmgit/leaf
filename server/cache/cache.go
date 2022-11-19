package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var redisClient *redis.Client
var ctx = context.Background()

func Init() {
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	password := viper.GetString("redis.password")
	redisClient = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0, // use default DB
	})
	zap.L().Info("redis连接成功")
}

func Set(key string, value interface{}, expiration time.Duration) {
	redisClient.Set(ctx, key, value, expiration)
}

func Get(key string) string {
	return redisClient.Get(ctx, key).Val()
}

func Del(key string) {
	redisClient.Del(ctx, key)
}
