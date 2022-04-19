package cmd

import (
	"github.com/spf13/cobra"
	"standard/internal/app/model"
	"standard/internal/global"
	"standard/pkg/tools"
)

var cfgFile bool

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "数据库初始命令",

	Run: func(cmd *cobra.Command, args []string) {
		//执行数据建表操作
		migrate()
		if cfgFile {
			//	执行数据填充操作
			initDatas()
		}
	},
}

func init() {
	//定义参数,并将参数值存储到某个变量中
	migrateCmd.PersistentFlags().BoolVar(&cfgFile, "config", false, "值为0或者1")
}

func migrate() {
	err := global.Orm.AutoMigrate(&model.SysUser{}, &model.SysRole{}, &model.SysMenu{}, &model.RoleCasbin{}, &model.Api{})
	if err != nil {
		global.Logger.Panic("数据库表迁移失败:%v", err)
	}
}

func initDatas() {
	global.Orm.Create(&model.SysRole{
		Name:    "admin",
		Keyword: "admin",
		Desc:    "管理员",
	})
	password := tools.GenPwd("admin")
	global.Orm.Create(&model.SysUser{
		Username: "admin",
		Password: password,
		RoleId:   1,
	})
}
