package models

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type SubmitBasic struct {
	gorm.Model
	Identity        string        `gorm:"column:identity;type:varchar(36)" json:"identity"` //提交的唯一表示
	ProblemIdentity string        `gorm:"column:problem_identity;type:varchar(36)" json:"problem_identity"`
	ProblemBasic    *ProblemBasic `gorm:"foreignKey:identity;references:problem_identity"`
	UserIdentity    string        `gorm:"column:user_identity;type:varchar(36)" json:"user_identity"`
	UserBasic       *UserBasic    `gorm:"foreignKey:identity;references:user_identity"`
	Path            string        `gorm:"column:path;type:varchar(255)" json:"path"` //代码存放路劲
	Status          int           `gorm:"column:status;type:tinyint" json:"status"`
}

func (table *SubmitBasic) TableName() string {
	return "submit_basic"
}

func GetSubmitList(pageStr, sizeStr, problemIdentity, userIdentity string) (interface{}, error) {
	data := make([]*SubmitBasic, 0)
	var total int64
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return nil, errors.New("page 不是数字")
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return nil, errors.New("size 不是数字")
	}
	tx := DB.Model(&data).Preload("ProblemBasic", func(db *gorm.DB) *gorm.DB {
		return db.Omit("content")
	})
	if problemIdentity != "" {
		tx = tx.Where("problem_identity = (SELECT identity from problem_basic pb WHERE pb.identity = ?)", problemIdentity)
	}
	if userIdentity != "" {
		tx = tx.Where("user_identity = (SELECT identity from user_basic ub WHERE ub.identity = ?)", userIdentity)
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
