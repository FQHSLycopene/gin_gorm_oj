package models

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	Identity   string `gorm:"column:identity;type:varchar(36)" json:"identity"` //问题唯一表示
	CategoryId string `gorm:"column:cid;type:varchar(255)" json:"cid"`          //分类ID,以逗号分隔
	Title      string `gorm:"column:title;type:varchar(255)" json:"title"`      //文章标题
	Content    string `gorm:"column:content;type:text" json:"content"`          //
	PassNum    int    `gorm:"column:pass_num;type:datetime" json:"pass_num"`
	TotalNum   int    `gorm:"column:total_num;type:datetime" json:"total_num"`
}

func (Table Problem) TableName() string {
	return "problem"
}
