package model

import "time"

type Model struct {
	Id        uint       `json:"id"  gorm:"comment:'自增编号';primary_key"`
	CreatedAt time.Time  `json:"createdAt" gorm:"comment:'创建时间';"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"comment:'更新时间';"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"comment:'删除时间'" sql:"index"`
}
