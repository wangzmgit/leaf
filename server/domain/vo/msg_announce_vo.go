package vo

import (
	"time"

	"kuukaa.fun/leaf/domain/model"
)

type AnnounceVO struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

func ToAnnounceVoList(announces []model.Announce) []AnnounceVO {
	length := len(announces)
	newAnnounces := make([]AnnounceVO, length)

	for i := 0; i < length; i++ {
		newAnnounces[i].ID = announces[i].ID
		newAnnounces[i].Title = announces[i].Title
		newAnnounces[i].Content = announces[i].Content
		newAnnounces[i].Url = announces[i].Url
		newAnnounces[i].CreatedAt = announces[i].CreatedAt
	}

	return newAnnounces
}