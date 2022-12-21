package model

// 收藏
type Collect struct {
	Cid      uint   `json:"cid" bson:"cid"`
	Uid      uint   `json:"uid" bson:"uid"`
	VideoIds []uint `json:"video_ids" bson:"video_ids"`
}
