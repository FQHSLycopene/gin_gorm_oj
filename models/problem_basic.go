package models

import (
	"errors"
	"gin_gorm_oj/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type ProblemBasic struct {
	gorm.Model
	Identity          string             `gorm:"column:identity;type:varchar(36)" json:"identity"`                          //问题唯一表示
	ProblemCategories []*ProblemCategory `gorm:"foreignKey:problem_identity;references:identity" json:"problem_categories"` //关联问题分类表
	TestCases         []*TestCase        `gorm:"foreignKey:problem_identity;references:identity" json:"test_cases"`
	Title             string             `gorm:"column:title;type:varchar(255)" json:"title"` //文章标题
	Content           string             `gorm:"column:content;type:text" json:"content"`     //
	PassNum           int                `gorm:"column:pass_num;type:int" json:"pass_num"`
	TotalNum          int                `gorm:"column:total_num;type:int" json:"total_num"`
	MaxRuntime        int                `gorm:"column:max_runtime;type:int(11)" json:"max_runtime"` //最大运行时长
	MaxMem            int                `gorm:"column:max_mem;type:int(11)" json:"max_mem"`         //最大运行内存
}

func (table *ProblemBasic) TableName() string {
	return "problem_basic"
}

func GetProblemList(pageStr, sizeStr, keyword, categoryIdentity string) (interface{}, error) {
	data := make([]*ProblemBasic, 0)
	var total int64
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return nil, errors.New("page 不是数字")
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return nil, errors.New("size 不是数字")
	}
	tx := DB.Model(data).Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").Preload("TestCases").
		Where(
			"title like ? OR content like ?",
			"%"+keyword+"%",
			"%"+keyword+"%",
		)
	if categoryIdentity != "" {
		tx = tx.Joins("RIGHT JOIN problem_category pc on pc.problem_identity = problem_basic.identity").
			Where("pc.category_identity = (SELECT cb.identity FROM category_basic cb where cb.identity = ?)", categoryIdentity)
	}
	tx.Offset((page - 1) * size).Limit(size).Count(&total).Find(&data)
	return gin.H{
		"list":  data,
		"total": total,
	}, nil
}

func AddProblem(title, content, maxRuntimeStr, maxMemoryStr string) (interface{}, error) {
	maxRuntime, err := strconv.Atoi(maxRuntimeStr)
	if err != nil {
		return nil, err
	}
	maxMemory, err := strconv.Atoi(maxMemoryStr)
	if err != nil {
		return nil, err
	}
	data := ProblemBasic{
		Identity:   utils.GetUUID(),
		Title:      title,
		Content:    content,
		MaxMem:     maxMemory,
		MaxRuntime: maxRuntime,
	}
	DB.Create(&data)
	return data.Identity, nil
}

func GetProblemDetail(identity string) (interface{}, error) {
	data := ProblemBasic{}
	err := DB.Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").Preload("TestCases").Where("identity = ?", identity).First(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
