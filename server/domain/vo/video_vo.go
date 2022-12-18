package vo

import (
	"time"

	"kuukaa.fun/leaf/domain/model"
)

type VideoVo struct {
	ID        uint         `json:"vid"`
	Title     string       `json:"title"`
	Cover     string       `json:"cover"`
	Desc      string       `json:"desc"`
	CreatedAt time.Time    `json:"created_at"`
	Copyright bool         `json:"copyright"`
	Author    UserVo       `json:"author"`
	Resource  []ResourceVo `json:"resources"`
	Clicks    int64        `json:"clicks"`
}

// 视频状态VO
type VideoStatusVO struct {
	ID        uint         `json:"vid"`
	Title     string       `json:"title"`
	Cover     string       `json:"cover"`
	Desc      string       `json:"desc"`
	Status    int          `json:"status"`
	Partition uint         `json:"partition"`
	Copyright bool         `json:"copyright"`
	Resources []ResourceVo `json:"resources"`
}

func ToVideoStatusVO(video model.Video, resources []model.Resource) VideoStatusVO {
	resourcesVO := ResourceListToVoList(resources)

	return VideoStatusVO{
		ID:        video.ID,
		Title:     video.Title,
		Cover:     video.Cover,
		Desc:      video.Desc,
		Status:    video.Status,
		Partition: video.PartitionId,
		Copyright: video.Copyright,
		Resources: resourcesVO,
	}
}

func ToVideoVO(video model.Video, author model.User, clicks int64, resource []model.Resource) VideoVo {

	return VideoVo{
		ID:        video.ID,
		Title:     video.Title,
		Cover:     video.Cover,
		Desc:      video.Desc,
		CreatedAt: video.CreatedAt,
		Copyright: video.Copyright,
		Author: UserVo{
			ID:     author.ID,
			Name:   author.Username,
			Sign:   author.Sign,
			Avatar: author.Avatar,
		},
		Resource: ResourceListToVoList(resource),
		Clicks:   clicks,
	}
}
