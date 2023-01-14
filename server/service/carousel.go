package service

import "kuukaa.fun/leaf/domain/model"

func InsertCarousel(carousel model.Carousel) error {
	return mysqlClient.Create(&carousel).Error
}

func SelectCarousel() (carousel []model.Carousel) {
	mysqlClient.Find(&carousel)
	return
}

func DeleteCarousel(id uint) {
	mysqlClient.Where("id = ?", id).Delete(&model.Carousel{})
}
