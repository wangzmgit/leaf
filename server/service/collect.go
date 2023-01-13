package service

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"kuukaa.fun/leaf/domain/model"
)

func InsertCollect(id, userId uint) error {
	_, err := mongoClient.Collect().InsertOne(context.TODO(), model.Collect{
		Cid:      id,
		Uid:      userId,
		VideoIds: []uint{},
	})

	return err
}

// 收藏
func Collect(videoId, userId uint, collectionIds []uint) error {
	filter := bson.M{
		"uid": userId,
		"cid": bson.M{
			"$in": collectionIds,
		},
	}
	_, err := mongoClient.Collect().UpdateMany(context.TODO(), filter, bson.M{
		"$addToSet": bson.M{
			"video_ids": videoId,
		},
	})

	return err
}

// 取消收藏
func CancelCollect(videoId, userId uint, collectionIds []uint) error {
	filter := bson.M{
		"uid": userId,
		"cid": bson.M{
			"$in": collectionIds,
		},
	}
	_, err := mongoClient.Collect().UpdateMany(context.TODO(), filter, bson.M{
		"$pull": bson.M{
			"video_ids": videoId,
		},
	})

	return err
}

// 是否收藏
func IsCollect(videoId, userId uint) (bool, error) {
	c, err := mongoClient.Collect().CountDocuments(context.TODO(), bson.M{
		"uid": userId,
		"video_ids": bson.M{
			"$elemMatch": bson.M{"$eq": videoId},
		},
	})

	if err != nil {
		return false, err
	}

	return c > 0, nil
}

// 获取已收藏的收藏夹
func SelectCollectedInfo(userId, videoId uint) ([]uint, error) {
	var collectionIds []uint
	cursor, err := mongoClient.Collect().Aggregate(context.TODO(), bson.A{
		bson.M{
			"$match": bson.M{
				"uid": userId,
				"video_ids": bson.M{
					"$elemMatch": bson.M{
						"$eq": videoId,
					},
				},
			},
		},
		bson.M{
			"$project": bson.M{
				"_id":       0,
				"uid":       0,
				"video_ids": 0,
			},
		},
	})
	if err != nil {
		return collectionIds, err
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return collectionIds, err
	}

	if len(results) > 0 {
		for _, v := range results {
			collectionIds = append(collectionIds, uint(v["cid"].(int64)))
		}

	}
	return collectionIds, nil
}

// 获取收藏数
func SelectCollectCount(videoId uint) (int32, error) {
	cursor, err := mongoClient.Collect().Aggregate(context.TODO(), bson.A{
		bson.M{
			"$match": bson.M{
				"video_ids": bson.M{
					"$elemMatch": bson.M{
						"$eq": videoId,
					},
				},
			},
		},
		bson.M{
			"$group": bson.M{
				"_id": "$uid",
			},
		},
		bson.M{
			"$group": bson.M{
				"_id": 0,
				"count": bson.M{
					"$sum": 1,
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

	return 0, errors.New("收藏数获取失败")
}

// 获取收藏的视频
func SelectCollectVideo(collectionId uint, page, pageSize int) ([]model.Video, int32, error) {
	var results []bson.M
	var videos []model.Video
	cursor, err := mongoClient.Collect().Aggregate(context.TODO(), bson.A{
		bson.M{
			"$match": bson.M{
				"cid": collectionId,
			},
		},
		bson.M{
			"$project": bson.M{
				"_id": 0,
				"count": bson.M{
					"$size": "$video_ids",
				},
				"data": bson.M{
					"$slice": bson.A{
						"$video_ids",
						(page - 1) * pageSize,
						pageSize,
					},
				},
			},
		},
	})
	if err != nil {
		return videos, 0, err
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		return videos, 0, err
	}

	if len(results) > 0 {
		mysqlClient.Where("id in ?", results[0]["data"]).Find(&videos)
		// 更新播放量数据
		for i := 0; i < len(videos); i++ {
			videos[i].Clicks = GetVideoClicks(videos[i].ID)
		}

		return videos, results[0]["count"].(int32), nil
	}

	return videos, 0, errors.New("收藏视频获取失败")
}
