package service

import (
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/model"
)

func InsertCollection(collection model.Collection) (uint, error) {
	if err := mysqlClient.Create(&collection).Error; err != nil {
		return 0, err
	}

	// 插入收藏内容数据
	if err := InsertCollect(collection.ID, collection.Uid); err != nil {
		// 删除MySQL中的收藏夹
		DeleteCollection(collection.ID)
		return 0, err
	}

	return collection.ID, nil
}

func SelectCollectionByID(collectionId uint) (collection model.Collection) {
	mysqlClient.First(&collection, collectionId)
	return
}

func ModifyCollection(modifyDTO dto.ModifyCollectionDTO) error {
	if err := mysqlClient.Model(&model.Collection{}).Where("id = ?", modifyDTO.ID).Updates(
		map[string]interface{}{
			"name":  modifyDTO.Name,
			"cover": modifyDTO.Cover,
			"desc":  modifyDTO.Desc,
			"open":  modifyDTO.Open,
		},
	).Error; err != nil {
		return err
	}
	return nil
}

// 通过用户ID查询收藏夹
func SelectCollectionListByUid(userId uint) (collections []model.Collection) {
	mysqlClient.Model(&model.Collection{}).Where("uid = ?", userId).Find(&collections)
	return
}

// 删除收藏夹
func DeleteCollection(id uint) {
	mysqlClient.Where("id = ?", id).Delete(&model.Collection{})
}

// 收藏夹是否属于用户
func IsCollectionBelongUser(id, userId uint) bool {
	var collection model.Collection
	mysqlClient.Where("id = ? and uid = ?", id, userId).First(&collection)

	return collection.ID != 0
}
