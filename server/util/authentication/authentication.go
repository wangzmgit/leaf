package authentication

import (
	"github.com/casbin/casbin/v2"
	"go.uber.org/zap"
)

var enforcer *casbin.Enforcer

func InitCasbin() {
	var err error
	enforcer, err = casbin.NewEnforcer("./static/casbin/model.conf", "./static/casbin/policy.csv")
	if err != nil {
		zap.L().Error("初始化casbin错误，err：" + err.Error())
	}
}

func Check(sub, obj, act string) bool {
	ok, _ := enforcer.Enforce(sub, obj, act)
	return ok
}
