package initialize

import (
	"fmt"
	"standard/internal/global"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func InitOrm() {
	var err error
	dsn := dsn()
	global.Orm, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		global.Logger.Panicf("数据库连接失败:%v", err)
	}
	global.Logger.Debug("数据库连接成功")
}

func dsn() string {
	username := global.Conf.Mysql.Username
	passord := global.Conf.Mysql.Password
	host := global.Conf.Mysql.Host
	port := global.Conf.Mysql.Port
	query := global.Conf.Mysql.Query
	dbname := global.Conf.Mysql.Database
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		username, passord, host, port, dbname, query)
}
