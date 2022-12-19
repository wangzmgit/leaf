package dto

import "kuukaa.fun/leaf/domain/model"

type CollectionDTO struct {
	Name string //收藏夹名称
}

type ModifyCollectionDTO struct {
	ID    uint
	Cover string //封面图
	Name  string //收藏夹名称
	Desc  string //简介
	Open  bool   //是否公开
}

/**
 * CollectionDTO结构体转化为Collection结构体
 * param: uid 用户id
 * param: collectionDTO CollectionDTO结构体
 * return: Collection结构体
 */
func CollectionDtoToCollection(uid uint, collectionDTO CollectionDTO) model.Collection {
	return model.Collection{
		Uid:  uid,
		Name: collectionDTO.Name,
	}
}
