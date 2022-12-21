package service

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"kuukaa.fun/leaf/domain/model"
)

// 插入评论
func InsertComment(comment model.Comment) (primitive.ObjectID, error) {
	_, err := mongoClient.Comment().InsertOne(context.TODO(), comment)
	return comment.ID, err
}

// 插入回复
func InsertReply(commentId primitive.ObjectID, reply model.Reply) (primitive.ObjectID, error) {
	_, err := mongoClient.Comment().UpdateOne(context.TODO(), bson.M{"_id": commentId}, bson.M{
		"$addToSet": bson.M{
			"reply": reply,
		},
	})

	return reply.ID, err
}

// 查询评论
func SelectCommentByID(commentId primitive.ObjectID) (model.Comment, error) {
	var comments []model.Comment
	cursor, err := mongoClient.Comment().Aggregate(context.TODO(), bson.A{
		bson.M{
			"$match": bson.M{
				"_id": commentId,
			},
		},
		bson.M{
			"$project": bson.M{
				"uid":        "$uid",
				"content":    "$content",
				"created_at": "$created_at",
			},
		},
		bson.M{
			"$limit": 1,
		},
	})

	if err != nil {
		return model.Comment{}, err
	}

	if err := cursor.All(context.TODO(), &comments); err != nil {
		return model.Comment{}, err
	}

	if len(comments) == 0 {
		return model.Comment{}, errors.New("获取评论失败")
	}

	return comments[0], nil
}

// 查询回复
func SelectReplyByID(commentId, replyId primitive.ObjectID) (model.Reply, error) {
	var replies []model.Reply
	cursor, err := mongoClient.Comment().Aggregate(context.TODO(), bson.A{
		bson.M{
			"$match": bson.M{
				"_id": commentId,
			},
		},
		bson.M{
			"$project": bson.M{
				"reply": bson.M{
					"$filter": bson.M{
						"input": "$reply",
						"as":    "item",
						"cond": bson.M{
							"$eq": bson.A{"$$item._id", replyId},
						},
					},
				},
			},
		},
		bson.M{
			"$unwind": "$reply",
		},
		bson.M{
			"$project": bson.M{
				"_id":        "$reply._id",
				"uid":        "$reply.uid",
				"content":    "$reply.content",
				"created_at": "$reply.created_at",
			},
		},
		bson.M{
			"$limit": 1,
		},
	})

	if err != nil {
		return model.Reply{}, err
	}

	if err := cursor.All(context.TODO(), &replies); err != nil {
		return model.Reply{}, err
	}

	if len(replies) == 0 {
		return model.Reply{}, errors.New("获取回复失败")
	}

	return replies[0], nil
}

// 查询评论
func SelectCommentList(videoId uint, page, pageSize int) ([]model.Comment, error) {
	var comments []model.Comment
	cursor, err := mongoClient.Comment().Aggregate(context.TODO(), bson.A{
		bson.M{
			"$match": bson.M{
				"vid":       videoId,
				"is_delete": false,
			},
		},
		bson.M{
			"$project": bson.M{
				"uid":        "$uid",
				"content":    "$content",
				"created_at": "$created_at",
				"reply": bson.M{
					"$filter": bson.M{
						"input": "$reply",
						"as":    "item",
						"cond": bson.M{
							"$eq": bson.A{"$$item.is_delete", false},
						},
					},
				},
			},
		},
		bson.M{
			"$project": bson.M{
				"uid":        "$uid",
				"content":    "$content",
				"created_at": "$created_at",
				"reply": bson.M{
					"$slice": bson.A{"$reply", 0, 2},
				},
			},
		},
		bson.M{
			"$skip": (page - 1) * pageSize,
		},
		bson.M{
			"$limit": pageSize,
		},
	})

	if err != nil {
		return comments, err
	}

	if err := cursor.All(context.TODO(), &comments); err != nil {
		return comments, err
	}

	return comments, nil
}

// 查询回复
func SelectReplyList(id string, page, pageSize int) ([]model.Reply, error) {
	var replies []model.Reply
	objectId, _ := primitive.ObjectIDFromHex(id)
	cursor, err := mongoClient.Comment().Aggregate(context.TODO(), bson.A{
		bson.M{
			"$match": bson.M{
				"_id":       objectId,
				"is_delete": false,
			},
		},
		bson.M{
			"$project": bson.M{
				"reply": bson.M{
					"$filter": bson.M{
						"input": "$reply",
						"as":    "item",
						"cond": bson.M{
							"$eq": bson.A{"$$item.is_delete", false},
						},
					},
				},
			},
		},
		bson.M{
			"$unwind": "$reply",
		},
		bson.M{
			"$project": bson.M{
				"_id":        "$reply._id",
				"uid":        "$reply.uid",
				"content":    "$reply.content",
				"created_at": "$reply.created_at",
			},
		},
		bson.M{
			"$skip": (page-1)*pageSize + 2,
		},
		bson.M{
			"$limit": pageSize,
		},
	})

	if err != nil {
		return replies, err
	}

	if err := cursor.All(context.TODO(), &replies); err != nil {
		return replies, err
	}

	return replies, nil
}

func DeleteComment(objectId primitive.ObjectID) error {
	_, err := mongoClient.Comment().UpdateOne(context.TODO(), bson.M{"_id": objectId}, bson.M{
		"$set": bson.M{
			"is_delete": true,
		},
	})

	return err
}

func DeleteReply(commentId, replyId primitive.ObjectID) error {
	filter := bson.M{
		"_id":       commentId,
		"reply._id": replyId,
	}

	_, err := mongoClient.Comment().UpdateOne(context.TODO(), filter, bson.M{
		"$set": bson.M{
			"reply.$.is_delete": true,
		},
	})

	return err
}
