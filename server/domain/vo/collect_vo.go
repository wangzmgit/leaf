package vo

import (
	"time"

	"kuukaa.fun/leaf/domain/model"
)

// 收藏夹
type CollectionVO struct {
	ID        uint      `json:"id"`
	Cover     string    `json:"cover"`
	Name      string    `json:"name"` //收藏夹名称
	Desc      string    `json:"desc"` //简介
	Open      bool      `json:"open"` //是否公开
	CreatedAt time.Time `json:"created_at"`
}

func CollectionToVo(collection model.Collection) CollectionVO {

	return CollectionVO{
		ID:        collection.ID,
		Cover:     collection.Cover,
		Name:      collection.Name,
		Desc:      collection.Desc,
		Open:      collection.Open,
		CreatedAt: collection.CreatedAt,
	}
}

func CollectionListToVoList(collections []model.Collection) []CollectionVO {
	length := len(collections)
	newCollections := make([]CollectionVO, length)
	for i := 0; i < length; i++ {
		newCollections[i].ID = collections[i].ID
		newCollections[i].Cover = collections[i].Cover
		newCollections[i].Name = collections[i].Name
		newCollections[i].Desc = collections[i].Desc
		newCollections[i].Open = collections[i].Open
		newCollections[i].CreatedAt = collections[i].CreatedAt
	}
	return newCollections
}
