package dto

//修改用户信息
type ModifyUserInfoDTO struct {
	Avatar   string
	Name     string
	Gender   int
	Birthday string
	Sign     string
}

type ModifySpaceCoverDTO struct {
	SpaceCover string
}

//修改用户信息
type ModifyPwdDTO struct {
	Email    string
	Password string
	Code     string
}
