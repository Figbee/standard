package cmd

import (
	"standard/app/model"
	"standard/internal/global"
	"standard/pkg/tools"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "数据迁移命令",
	Run: func(cmd *cobra.Command, args []string) {
		initTable()
	},
}

//数据库表迁移命令
func initTable() {
	err := global.Orm.AutoMigrate(
		new(model.SysUser),
		new(model.SysRole),
		new(model.SysMenu),
	)
	if err != nil {
		global.Logger.Panicf("数据库迁移错误:%v", err)
	}
	//初始化数据
	initDb()
}

//初始化数据
func initDb() {
	global.Orm.Create(&model.SysRole{
		Name:    "admin",
		Keyword: "admin",
		Desc:    "管理员",
	})
	pass := tools.GenPwd("admin")
	global.Orm.Create(&model.SysUser{Username: "admin", Password: pass, RoleId: 1})
}
