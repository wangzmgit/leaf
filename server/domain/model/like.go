package model

type Like struct {
	Vid     uint   `json:"vid" bson:"vid"`
	UserIds []uint `json:"user_ids" bson:"user_ids"`
}
