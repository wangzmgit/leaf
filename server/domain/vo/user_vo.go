package vo

import (
	"time"

	"kuukaa.fun/leaf/domain/model"
	"kuukaa.fun/leaf/util/desensitization"
)

type BaseUserVO struct {
	ID         uint   `json:"uid"`
	Name       string `json:"name"`
	Sign       string `json:"sign"`
	Avatar     string `json:"avatar"`
	SpaceCover string `json:"spacecover"`
	Gender     int    `json:"gender"`
}

type UserVO struct {
	ID         uint      `json:"uid"`
	Name       string    `json:"name"`
	Sign       string    `json:"sign"`
	Email      string    `json:"email"`
	Avatar     string    `json:"avatar"`
	SpaceCover string    `json:"spacecover"`
	Gender     int       `json:"gender"`
	Role       int       `json:"role"`
	Birthday   time.Time `json:"birthday"`
	CreatedAt  time.Time `json:"created_at"`
}

func ToUserVO(user model.User, desensitize bool) UserVO {
	if desensitize {
		user.Email = desensitization.HideEmail(user.Email)
	}
	return UserVO{
		ID:         user.ID,
		Name:       user.Username,
		Email:      user.Email,
		Sign:       user.Sign,
		Avatar:     user.Avatar,
		SpaceCover: user.SpaceCover,
		Gender:     user.Gender,
		Role:       user.Role,
		Birthday:   user.Birthday,
		CreatedAt:  user.CreatedAt,
	}
}

func ToBaseUserVO(user model.User) BaseUserVO {
	return BaseUserVO{
		ID:         user.ID,
		Name:       user.Username,
		Sign:       user.Sign,
		Avatar:     user.Avatar,
		SpaceCover: user.SpaceCover,
		Gender:     user.Gender,
	}
}

func ToUserVoList(users []model.User) []UserVO {
	length := len(users)
	newUsers := make([]UserVO, length)
	for i := 0; i < length; i++ {
		newUsers[i].ID = users[i].ID
		newUsers[i].Name = users[i].Username
		newUsers[i].Sign = users[i].Sign
		newUsers[i].Email = users[i].Email
		newUsers[i].Avatar = users[i].Avatar
		newUsers[i].SpaceCover = users[i].SpaceCover
		newUsers[i].Gender = users[i].Gender
		newUsers[i].Birthday = users[i].Birthday
		newUsers[i].CreatedAt = users[i].CreatedAt
		newUsers[i].Role = users[i].Role
	}

	return newUsers
}

func ToBaseUserVoList(users []model.User) []BaseUserVO {
	length := len(users)
	newUsers := make([]BaseUserVO, length)
	for i := 0; i < length; i++ {
		newUsers[i].ID = users[i].ID
		newUsers[i].Name = users[i].Username
		newUsers[i].Sign = users[i].Sign
		newUsers[i].Avatar = users[i].Avatar
		newUsers[i].SpaceCover = users[i].SpaceCover
		newUsers[i].Gender = users[i].Gender
	}

	return newUsers
}
