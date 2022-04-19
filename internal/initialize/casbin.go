package initialize

import (
	"fmt"
	"standard/internal/global"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

func InitCasbin() {
	var err error
	a, err := gormadapter.NewAdapterByDB(global.Orm)
	if err != nil {
		global.Logger.Panicf("权限模块载入失败:%v", err)
		panic(fmt.Sprintf("初始化Casbin失败：%v", err))
	}
	global.Casbin, err = casbin.NewEnforcer("conf/rbac_model.conf", a)
	if err != nil {
		global.Logger.Panicf("权限模块载入失败:%v", err)
		panic(fmt.Sprintf("初始化Casbin失败：%v", err))
	}
	err = global.Casbin.LoadPolicy()
	if err != nil {
		global.Logger.Panicf("loadpolicy error is :%v", err)
		panic(fmt.Sprintf("初始化Casbin失败：%v", err))
	}
	global.Logger.Info("初始化casbin完成")
}
