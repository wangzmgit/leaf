package vo

import "kuukaa.fun/leaf/domain/model"

// 分区
type PartitionVo struct {
	ID       uint   `json:"id"`
	Content  string `json:"content"`
	ParentID uint   `json:"parent_id"`
}

func ToPartitionVo(partitions []model.Partition) []PartitionVo {
	length := len(partitions)
	partitionVo := make([]PartitionVo, length)

	for i := 0; i < length; i++ {
		partitionVo[i].ID = partitions[i].ID
		partitionVo[i].Content = partitions[i].Content
		partitionVo[i].ParentID = partitions[i].ParentId
	}

	return partitionVo
}
