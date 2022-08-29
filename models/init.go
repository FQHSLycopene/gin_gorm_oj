package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB = Init()

func Init() *gorm.DB {
	dsn := "Lycopene:MiMaJiuShi123321!@tcp(127.0.0.1:3306)/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	}
	db.AutoMigrate(&ProblemBasic{}, &CategoryBasic{}, &ProblemCategory{})
	return db
}
