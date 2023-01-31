package mongodb

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var mongoClient *MongoClient

func Init() {
	var err error

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?%s",
		viper.GetString("mongo.username"),
		viper.GetString("mongo.password"),
		viper.GetString("mongo.host"),
		viper.GetString("mongo.port"),
		viper.GetString("mongo.param"),
	)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		// TODO 数据库连接失败处理
		zap.L().Error("mongodb连接失败")
		return
	}

	mongoClient = &MongoClient{
		db: client.Database(viper.GetString("mongo.datasource")),
	}

	zap.L().Info("mongodb连接成功")
}

func GetMongoClient() *MongoClient {
	return mongoClient
}
