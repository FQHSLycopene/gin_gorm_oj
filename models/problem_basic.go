package models

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type ProblemBasic struct {
	gorm.Model
	Identity          string             `gorm:"column:identity;type:varchar(36)" json:"identity"` //问题唯一表示
	ProblemCategories []*ProblemCategory `gorm:"foreignKey:problem_id;reference:id"`               //关联问题分类表
	Title             string             `gorm:"column:title;type:varchar(255)" json:"title"`      //文章标题
	Content           string             `gorm:"column:content;type:text" json:"content"`          //
	PassNum           int                `gorm:"column:pass_num;type:datetime" json:"pass_num"`
	TotalNum          int                `gorm:"column:total_num;type:datetime" json:"total_num"`
	MaxRuntime        int                `gorm:"column:max_runtime;type:int(11)" json:"max_runtime"` //最大运行时长
	MaxMem            int                `gorm:"column:max_mem;type:int(11)" json:"max_mem"`         //最大运行内存
}

func (table *ProblemBasic) TableName() string {
	return "problem_basic"
}

func GetProblemList(pagestr, sizestr, keyword, categoryId string) (interface{}, error) {
	data := make([]*ProblemBasic, 0)
	var total int64
	page, err := strconv.Atoi(pagestr)
	if err != nil {
		return nil, errors.New("page 不是数字")
	}
	size, err := strconv.Atoi(sizestr)
	if err != nil {
		return nil, errors.New("size 不是数字")
	}
	tx := DB.Model(data).Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").
		Where(
			"title like ? OR content like ?",
			"%"+keyword+"%",
			"%"+keyword+"%",
		)
	if categoryId != "" {
		tx = tx.Joins("RIGHT JOIN problem_category pc on pc.problem_id = problem_basic.id").
			Where("pc.category_id = (SELECT cb.id FROM category_basic cb where cb.id = ?)", categoryId)
	}
	tx.Offset((page - 1) * size).Limit(size).Count(&total).Find(&data)
	return gin.H{
		"list":  data,
		"total": total,
	}, nil
}

func AddProblem(problem *ProblemBasic) *ProblemBasic {
	DB.Create(problem)
	return problem
}

func GetProblemDetail(id int) (interface{}, error) {
	data := ProblemBasic{}
	DB.Model(data).Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").First(&data, id)
	return data, nil
}
