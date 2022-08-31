package models

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type SubmitBasic struct {
	gorm.Model
	Identity     string        `gorm:"column:identity;type:varchar(36)" json:"identity"` //提交的唯一表示
	ProblemId    string        `gorm:"column:problem_id;type:int" json:"problem_id"`
	ProblemBasic *ProblemBasic `gorm:"foreignKey:id;reference:problem_id"`
	UserId       string        `gorm:"column:user_id;type:int" json:"user_id"`
	UserBasic    *UserBasic    `gorm:"foreignKey:id;reference:user_id"`
	Path         string        `gorm:"column:path;type:varchar(255)" json:"path"` //代码存放路劲
	Status       int           `gorm:"column:status;type:tinyint" json:"status"`
}

func (table *SubmitBasic) TableName() string {
	return "submit_basic"
}

func GetSubmitList(pagestr, sizestr, problemID, userID string) (interface{}, error) {
	data := make([]*SubmitBasic, 0)
	var total int64
	page, err := strconv.Atoi(pagestr)
	if err != nil {
		return nil, errors.New("page 不是数字")
	}
	size, err := strconv.Atoi(sizestr)
	if err != nil {
		return nil, errors.New("size 不是数字")
	}
	tx := DB.Model(&data).Preload("ProblemBasic", func(db *gorm.DB) *gorm.DB {
		return db.Omit("content")
	})
	if problemID != "" {
		tx = tx.Where("problem_id = (SELECT id from problem_basic pb WHERE pb.id = ?)", problemID)
	}
	if userID != "" {
		tx = tx.Where("user_id = (SELECT id from user_basic ub WHERE ub.id = ?)", userID)
	}
	err = tx.Offset((page - 1) * size).Limit(size).Count(&total).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return gin.H{
		"total": total,
		"list":  data,
	}, nil
}
