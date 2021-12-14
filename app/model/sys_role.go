package model

type SysRole struct {
	Model
	Name    string `json:"name" gorm:"comment:'角色名称'"`
	Keyword string `json:"keyword" gorm:"unique;comment:'角色关键字'"`
	Desc    string `json:"desc" gorm:"comment:'角色说明';"`
	Status  *bool  `json:"status" gorm:"comment:'角色状态';type:tinyint(1);default:1"`
	Creator string `json:"creator" gorm:"comment:'创建人';"`
	//一对多关系
	Users []SysUser `json:"users"  gorm:"foreignkey:RoleId"`
	Menus []SysMenu `json:"menus" gorm:"many2many:relation_role_menu"`
}

func (m *SysRole) TableName() string {
	return "sys_role"
}
