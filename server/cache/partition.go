package cache

import (
	"encoding/json"

	"go.uber.org/zap"
	"kuukaa.fun/leaf/domain/vo"
)

func GetPartition() (partitions []vo.PartitionVo) {
	jsonStr := Get(PARTITION_KEY)
	// 反序列化
	if err := json.Unmarshal([]byte(jsonStr), &partitions); err != nil {
		zap.L().Error("分区反序列化失败: " + err.Error())
	}
	return
}

func SetPartition(partitions []vo.PartitionVo) {
	//先序列化
	pb, err := json.Marshal(partitions)
	if err != nil {
		zap.L().Error("分区序列化失败: " + err.Error())
		return
	}
	Set(PARTITION_KEY, pb, PARTITION_EXPRIRATION_TIME)
}

func DelPartition() {
	Del(PARTITION_KEY)
}
