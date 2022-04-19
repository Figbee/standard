package model

type SysUser struct {
	Model
	Username string `json:"username"`
	//json"-",表示json序列化时不显示
	Password string  `json:"-"`
	Avatar   string  `json:"avatar"`
	Status   *bool   `json:"status" gorm:"type:tinyint(1);default:1;comment:'用户状态'"`
	RoleId   uint    `json:"roleId" gorm:"comment:'角色id外键';"`
	Role     SysRole `json:"role"  gorm:"foreignkey:RoleId;"` //将sysUser.RoleId指定为外键
}

type Login struct {
	Username string `json:"username" form:"username"  binding:"required"`
	Password string `json:"password" form:"password"  binding:"required"`
}

type Claim struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	RoleId   uint   `json:"role_id"`
	RoleName string `json:"roleName"`
}

func (u SysUser) TableName() string {
	return "sys_user"
}
