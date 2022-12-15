package vo

import "kuukaa.fun/leaf/domain/model"

type ResourceVo struct {
	ID uint `json:"id"`
	// 分P使用的标题
	Title string `json:"title"`
	// 不同分辨率
	Url string `json:"url"`
	// 时长
	Duration float64 `json:"duration"`
	// 审核状态
	Status int `json:"status"`
	// 资源最大分辨率
	Quality int `json:"quality"`
}

func ResourceListToVoList(resources []model.Resource) []ResourceVo {
	length := len(resources)
	newResources := make([]ResourceVo, length)
	for i := 0; i < length; i++ {
		newResources[i].ID = resources[i].ID
		newResources[i].Title = resources[i].Title
		newResources[i].Url = resources[i].Url
		newResources[i].Duration = resources[i].Duration
		newResources[i].Status = resources[i].Status
		newResources[i].Quality = resources[i].Quality
	}
	return newResources
}
