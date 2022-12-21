package vo

import (
	"time"

	"kuukaa.fun/leaf/domain/model"
)

type UserVO struct {
	ID         uint      `json:"uid"`
	Name       string    `json:"name"`
	Sign       string    `json:"sign"`
	Avatar     string    `json:"avatar"`
	SpaceCover string    `json:"spacecover"`
	Gender     int       `json:"gender"`
	Role       int       `json:"role"`
	Birthday   time.Time `json:"birthday"`
}

func ToUserVO(user model.User) UserVO {
	return UserVO{
		ID:         user.ID,
		Name:       user.Username,
		Sign:       user.Sign,
		Avatar:     user.Avatar,
		SpaceCover: user.SpaceCover,
		Gender:     user.Gender,
		Role:       user.Role,
		Birthday:   user.Birthday,
	}
}

func ToUserVoList(users []model.User) []UserVO {
	length := len(users)
	newUsers := make([]UserVO, length)
	for i := 0; i < length; i++ {
		newUsers[i].ID = users[i].ID
		newUsers[i].Name = users[i].Username
		newUsers[i].Sign = users[i].Sign
		newUsers[i].Avatar = users[i].Avatar
		newUsers[i].SpaceCover = users[i].SpaceCover
		newUsers[i].Gender = users[i].Gender
		newUsers[i].Birthday = users[i].Birthday
	}

	return newUsers
}
