package vo

import (
	"time"

	"github.com/spf13/viper"
	"kuukaa.fun/leaf/domain/model"
)

type RoomVO struct {
	Number int `json:"number"`
}

type VideoVO struct {
	ID        uint         `json:"vid"`
	Title     string       `json:"title"`
	Cover     string       `json:"cover"`
	Desc      string       `json:"desc"`
	CreatedAt time.Time    `json:"created_at"`
	Copyright bool         `json:"copyright"`
	Author    BaseUserVO   `json:"author"`
	Resource  []ResourceVO `json:"resources"`
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
	Resources []ResourceVO `json:"resources"`
}

// 基础视频信息
type BaseVideoVO struct {
	ID        uint      `json:"vid"`
	Title     string    `json:"title"`
	Cover     string    `json:"cover"`
	Desc      string    `json:"desc"`
	Clicks    int64     `json:"clicks"`
	CreatedAt time.Time `json:"created_at"`
}

// 用户上传视频
type UserUploadVideoVO struct {
	ID        uint      `json:"vid"`
	Title     string    `json:"title"`
	Cover     string    `json:"cover"`
	Desc      string    `json:"desc"`
	Clicks    int64     `json:"clicks"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// 搜索视频
type SearchVideoVO struct {
	ID        uint       `json:"vid"`
	Title     string     `json:"title"`
	Cover     string     `json:"cover"`
	Desc      string     `json:"desc"`
	CreatedAt time.Time  `json:"created_at"`
	Copyright bool       `json:"copyright"`
	Author    BaseUserVO `json:"author"`
	Clicks    int64      `json:"clicks"`
	Partition uint       `json:"partition"`
}

func ToVideoStatusVO(video model.Video, resources []model.Resource) VideoStatusVO {
	resourcesVO := ToResourceVoList(resources)

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

func ToVideoVO(video model.Video, resource []model.Resource, ios uint) VideoVO {

	// 处理视频资源地址
	if ios == 1 && viper.GetBool("file.video_adaptation_ios") {
		for i := 0; i < len(resource); i++ {
			resource[i].Url = resource[i].OriginalUrl
		}
	}

	return VideoVO{
		ID:        video.ID,
		Title:     video.Title,
		Cover:     video.Cover,
		Desc:      video.Desc,
		CreatedAt: video.CreatedAt,
		Copyright: video.Copyright,
		Author:    ToBaseUserVO(video.Author),
		Resource:  ToResourceVoList(resource),
		Clicks:    video.Clicks,
	}
}

func ToBaseVideoVO(video model.Video) BaseVideoVO {
	return BaseVideoVO{
		ID:        video.ID,
		Title:     video.Title,
		Cover:     video.Cover,
		Desc:      video.Desc,
		Clicks:    video.Clicks,
		CreatedAt: video.CreatedAt,
	}
}

func ToBaseVideoVoList(videos []model.Video) []BaseVideoVO {
	length := len(videos)
	newVideos := make([]BaseVideoVO, length)
	for i := 0; i < length; i++ {
		newVideos[i].ID = videos[i].ID
		newVideos[i].Title = videos[i].Title
		newVideos[i].Cover = videos[i].Cover
		newVideos[i].Desc = videos[i].Desc
		newVideos[i].Clicks = videos[i].Clicks
		newVideos[i].CreatedAt = videos[i].CreatedAt
	}

	return newVideos
}

func ToUserUploadVideoVoList(videos []model.Video) []UserUploadVideoVO {
	length := len(videos)
	newVideos := make([]UserUploadVideoVO, length)
	for i := 0; i < length; i++ {
		newVideos[i].ID = videos[i].ID
		newVideos[i].Title = videos[i].Title
		newVideos[i].Cover = videos[i].Cover
		newVideos[i].Desc = videos[i].Desc
		newVideos[i].Clicks = videos[i].Clicks
		newVideos[i].Status = videos[i].Status
		newVideos[i].CreatedAt = videos[i].CreatedAt
	}

	return newVideos
}

func ToSearchVideoVoList(videos []model.Video) []SearchVideoVO {
	length := len(videos)
	newVideos := make([]SearchVideoVO, length)
	for i := 0; i < length; i++ {
		newVideos[i].ID = videos[i].ID
		newVideos[i].Title = videos[i].Title
		newVideos[i].Cover = videos[i].Cover
		newVideos[i].Desc = videos[i].Desc
		newVideos[i].Clicks = videos[i].Clicks
		newVideos[i].Copyright = videos[i].Copyright
		newVideos[i].CreatedAt = videos[i].CreatedAt
		newVideos[i].Author = ToBaseUserVO(videos[i].Author)
		newVideos[i].Partition = videos[i].PartitionId
	}

	return newVideos
}
