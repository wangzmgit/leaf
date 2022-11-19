package service

import (
	"kuukaa.fun/leaf/db/mysql"
	"kuukaa.fun/leaf/domain/model"
)

func GetSysMenuList() []model.SysMenu {
	mysql := mysql.GetMysqlClient()
	mysql.AutoMigrate(&model.SysMenu{})
	var sysMenuList []model.SysMenu
	mysql.Find(&sysMenuList)
	return sysMenuList
}
