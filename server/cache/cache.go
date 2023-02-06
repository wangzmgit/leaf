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

func Incr(key string) {
	redisClient.Incr(ctx, key)
}

func Keys(key string) []string {
	return redisClient.Keys(ctx, key).Val()
}

func TTL(key string) time.Duration {
	return redisClient.TTL(ctx, key).Val()
}

// 向有序集合插入数据
func ZAdd(key string, score float64, member interface{}) {
	redisClient.ZAdd(ctx, key, redis.Z{Score: score, Member: member})
}

// 有序集合中的数量
func ZCard(key string) int64 {
	return redisClient.ZCard(ctx, key).Val()
}

// 有序集中成员的分数值
func ZScore(key string, member string) float64 {
	return redisClient.ZScore(ctx, key, member).Val()
}

// 移除有序集中指定排名区间内的所有成员
func ZRemRangeByRank(key string, start, stop int64) {
	redisClient.ZRemRangeByRank(ctx, key, start, stop)
}
