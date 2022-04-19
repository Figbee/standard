package initialize

import (
	"fmt"
	"standard/internal/global"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var (
	dataSource string
)

func InitMysql() {
	mysqlConf := global.Conf.Mysql
	dataSource = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		mysqlConf.Username,
		mysqlConf.Password,
		mysqlConf.Host,
		mysqlConf.Port,
		mysqlConf.Database,
		mysqlConf.Query)
	//开启mysql日志
	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		global.Logger.DPanicf("数据库连接失败%d", err)
		panic("数据库连接失败")
	}
	if mysqlConf.LogMode {
		db.Debug()
	}
	global.Orm = db

}

