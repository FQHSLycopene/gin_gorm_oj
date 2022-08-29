package models

import "gorm.io/gorm"

type ProblemCategory struct {
	gorm.Model
	ProblemID     int            `gorm:"column:problem_id;int(11)" json:"problem_id"`
	CategoryID    int            `gorm:"column:category_id;int(11)" json:"category_id"`
	CategoryBasic *CategoryBasic `gorm:"foreignKey:id;references:category_id"` //关联分类的基础信息表
}

func (table *ProblemCategory) TableName() string {
	return "problem_category"
}
