package models

import "gorm.io/gorm"

type CategoryBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36)" json:"identity"` //分类的唯一表示
	Name     string `gorm:"column:name;type:varchar(100)" json:"name"`
	Pid      string `gorm:"column:pid;type:int(11)" json:"pid"`
}

func (table *CategoryBasic) TableName() string {
	return "category_basic"
}
