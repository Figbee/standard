package dao

import (
	"errors"
	"standard/app/model"
	"standard/internal/global"
	"standard/pkg/tools"
)

func CheckAuth(username string, password string) (*model.SysUser, error) {
	var u model.SysUser
	//err := global.Orm.Preload("Role").Where("username=?", username).First(&u).Error
	err := global.Orm.Preload("Role").Where("username=?", username).First(&u).Error
	if err != nil {
		return nil, err
	}
	//校验密码
	if ok := tools.ComparePwd(password, u.Password); !ok {
		return nil, errors.New("密码错误")
	}
	return &u, nil
}

func GetUserInfo(uname string) (*model.SysUser, error) {
	var u model.SysUser
	err := global.Orm.Select("username,avatar,status").Where("username=?", uname).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}
