package service

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"kuukaa.fun/leaf/domain/model"
)

func InsertLike(id uint) error {
	_, err := mongoClient.Like().InsertOne(context.TODO(), model.Like{
		Vid:     id,
		UserIds: []uint{},
	})

	return err
}

// 点赞
func Like(videoId, userId uint) error {
	_, err := mongoClient.Like().UpdateOne(context.TODO(), bson.M{"vid": videoId}, bson.M{
		"$addToSet": bson.M{
			"user_ids": userId,
		},
	})

	return err
}

// 取消赞
func CancelLike(videoId, userId uint) error {
	_, err := mongoClient.Like().UpdateOne(context.TODO(), bson.M{"vid": videoId}, bson.M{
		"$pull": bson.M{
			"user_ids": userId,
		},
	})

	return err
}

// 是否点赞
func IsLike(videoId, userId uint) (bool, error) {
	c, err := mongoClient.Like().CountDocuments(context.TODO(), bson.M{
		"vid": videoId,
		"user_ids": bson.M{
			"$elemMatch": bson.M{"$eq": userId},
		},
	})

	if err != nil {
		return false, err
	}

	return c > 0, nil
}

// 获取点赞数
func SelectLikeCount(videoId uint) (int32, error) {
	cursor, err := mongoClient.Like().Aggregate(context.TODO(), bson.A{
		bson.M{
			"$match": bson.M{
				"vid": videoId,
			},
		},
		bson.M{
			"$project": bson.M{
				"count": bson.M{
					"$size": "$user_ids",
				},
			},
		},
		bson.M{
			"$project": bson.M{
				"_id": 0,
			},
		},
	})
	if err != nil {
		return 0, err
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return 0, err
	}

	if len(results) > 0 {
		return results[0]["count"].(int32), nil
	}

	return 0, errors.New("点赞数获取失败")
}
