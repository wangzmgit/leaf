package service

import (
	"kuukaa.fun/leaf/domain/model"
)

func SelectPartition() []model.Partition {
	var partitions []model.Partition
	mysqlClient.Model(&model.Partition{}).Find(&partitions)
	return partitions
}

func InsertPartition(partition model.Partition) {
	mysqlClient.Create(&partition)
}

func DelPartitionByID(id uint) {
	mysqlClient.Where("id = ?", id).Delete(&model.Partition{})
}

// 所属分区是否存在
func IsParentPartitionExist(id uint) bool {
	var partition model.Partition
	mysqlClient.First(&partition, id)
	if partition.ID != 0 && partition.ParentId == 0 {
		return true
	}
	return false
}

// 是否为子分区
func IsSubpartition(id uint) bool {
	var partition model.Partition
	mysqlClient.First(&partition, id)
	if partition.ID != 0 && partition.ParentId != 0 {
		return true
	}
	return false
}
