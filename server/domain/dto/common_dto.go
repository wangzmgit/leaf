package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type EmailDTO struct {
	// 邮箱
	Email string
}

type IdDTO struct {
	ID uint
}

type ObjectIdDTO struct {
	ID primitive.ObjectID
}
