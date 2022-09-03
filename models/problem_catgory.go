package models

import "gorm.io/gorm"

type ProblemCategory struct {
	gorm.Model
	ProblemIdentity  string         `gorm:"column:problem_identity;varchar(36)" json:"problem_identity"`
	CategoryIdentity string         `gorm:"column:category_identity;varchar(36)" json:"category_identity"`
	CategoryBasic    *CategoryBasic `gorm:"foreignKey:identity;references:category_identity"` //关联分类的基础信息表
}

func (table *ProblemCategory) TableName() string {
	return "problem_category"
}
