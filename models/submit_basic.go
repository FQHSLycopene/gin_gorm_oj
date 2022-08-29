package models

import "gorm.io/gorm"

type SubmitBasic struct {
	gorm.Model
	Identity        string `gorm:"column:identity;type:varchar(36)" json:"identity"` //提交的唯一表示
	ProblemIdentity string `gorm:"column:problem_identity;type:varchar(36)" json:"problem_identity"`
	UserIdentity    string `gorm:"column:user_identity;type:varchar(36)" json:"user_identity"`
	Path            string `gorm:"column:path;type:varchar(255)" json:"path"` //代码存放路劲
	Status          int    `gorm:"column:status;type:tinyint" json:"status"`
}

func (table *SubmitBasic) TableName() string {
	return "submit_basic"
}
