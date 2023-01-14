package vo

import (
	"time"

	"kuukaa.fun/leaf/domain/model"
)

type CarouselVO struct {
	ID        uint      `json:"id"`
	Img       string    `json:"img"`
	Url       string    `json:"url"`
	Title     string    `json:"title"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
}

func ToCarouselVoList(carousels []model.Carousel) []CarouselVO {
	length := len(carousels)
	newCarousels := make([]CarouselVO, length)
	for i := 0; i < length; i++ {
		newCarousels[i].ID = carousels[i].ID
		newCarousels[i].Img = carousels[i].Img
		newCarousels[i].Url = carousels[i].Url
		newCarousels[i].Title = carousels[i].Title
		newCarousels[i].Color = carousels[i].Color
		newCarousels[i].CreatedAt = carousels[i].CreatedAt
	}

	return newCarousels
}
