package dto

import "kuukaa.fun/leaf/domain/model"

type AnnounceDTO struct {
	Title     string
	Content   string
	Url       string
	Important bool
}

/**
 * AnnounceDTO结构体转化为Announce结构体
 * param: announceDTO AnnounceDTO结构体
 * return: Announce结构体
 */
func AnnounceDtoToAnnounce(announceDTO AnnounceDTO) model.Announce {
	return model.Announce{
		Title:     announceDTO.Title,
		Content:   announceDTO.Content,
		Url:       announceDTO.Url,
		Important: announceDTO.Important,
	}
}
