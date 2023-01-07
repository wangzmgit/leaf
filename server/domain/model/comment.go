package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Vid       uint               `json:"vid" bson:"vid"` //视频ID
	CreatedAt int64              `json:"created_at" bson:"created_at"`
	Content   string             `json:"content" bson:"content"` //内容
	Uid       uint               `json:"uid" bson:"uid"`         //用户ID
	Author    User               `bson:"-"`
	Reply     []Reply            `json:"reply" bson:"reply"`
	At        []uint             `json:"at" bson:"at"`
	IsDelete  bool               `json:"is_delete" bson:"is_delete"`
}

type Reply struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	CreatedAt int64              `json:"created_at" bson:"created_at"`
	Content   string             `json:"content" bson:"content"` //内容
	Uid       uint               `json:"uid" bson:"uid"`         //用户ID
	Author    User               `bson:"-"`
	At        []uint             `json:"at" bson:"at"`
	IsDelete  bool               `json:"is_delete" bson:"is_delete"`
}
