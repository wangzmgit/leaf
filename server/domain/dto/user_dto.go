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

//管理员修改用户信息
type AdminModifyUserInfoDTO struct {
	ID    uint
	Email string
	Name  string
	Sign  string
}

//管理员修改用户权限
type ModifyRoleDTO struct {
	ID   uint
	Role int
}
