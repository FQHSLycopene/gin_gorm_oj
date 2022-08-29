package test

import (
	"fmt"
	"gin_gorm_oj/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGormTest(t *testing.T) {
	dsn := "Lycopene:MiMaJiuShi123321!@tcp(127.0.0.1:3306)/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

	}
	//data := models.Problem{
	//	Identity:   "problem_1",
	//	Title:      "文章标题",
	//	MaxRuntime: 3000,
	//	MaxMem:     3000,
	//	Content:    "文章内容",
	//	PassNum:    0,
	//	TotalNum:   0,
	//}
	//db.Create(&data)
	datas := make([]*models.Problem, 0)
	db.Find(&datas)
	for _, v := range datas {
		fmt.Printf("Problem ==> %v \n", v)
	}
}
