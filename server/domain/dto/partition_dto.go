package dto

import "kuukaa.fun/leaf/domain/model"

type PartitionDTO struct {
	Content  string
	ParentId uint //所属分区ID
}

/**
 * 分区DTO结构体转化为Partition结构体
 * param: partitionDTO 分区DTO结构体
 * return: Partition结构体
 */
func PartitionDtoToPartition(partitionDTO PartitionDTO) model.Partition {
	return model.Partition{
		Content:  partitionDTO.Content,
		ParentId: partitionDTO.ParentId,
	}
}
