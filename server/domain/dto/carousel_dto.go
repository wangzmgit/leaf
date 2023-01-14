package dto

import "kuukaa.fun/leaf/domain/model"

type CarouselDTO struct {
	Img   string
	Url   string
	Title string
	Color string
}

/**
 * CarouselDTO结构体转化为Carousel结构体
 * param: carouselDTO CarouselDTO
 * return: Carousel结构体
 */
func CarouselDtoToCarousel(carouselDTO CarouselDTO) model.Carousel {
	return model.Carousel{
		Img:   carouselDTO.Img,
		Url:   carouselDTO.Url,
		Title: carouselDTO.Title,
		Color: carouselDTO.Color,
	}
}
