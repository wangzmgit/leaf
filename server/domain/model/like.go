package model

type Like struct {
	Vid     uint   `json:"vid" bson:"vid"`
	UserIds []uint `json:"userids" bson:"userids"`
}
