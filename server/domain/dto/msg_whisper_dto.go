package dto

import "kuukaa.fun/leaf/domain/model"

type WhisperDTO struct {
	Fid     uint
	Content string
}

/**
 * WhisperDTO结构体转化为Whisper结构体
 * param: whisperDTO WhisperDTO结构体
 * return: Whisper结构体
 */
func WhisperDtoToWhisper(whisperDTO WhisperDTO, userId uint) []model.Whisper {
	return []model.Whisper{
		{
			Uid:     userId,
			Fid:     whisperDTO.Fid,
			FromId:  userId,
			ToId:    whisperDTO.Fid,
			Content: whisperDTO.Content,
			Status:  true,
		},
		{
			Uid:     whisperDTO.Fid,
			Fid:     userId,
			FromId:  userId,
			ToId:    whisperDTO.Fid,
			Content: whisperDTO.Content,
		},
	}
}
