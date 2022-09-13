package models

import (
	"context"
	"github.com/go-redis/redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB = InitDB()
var RDB = InitRedisDB()
var CTX = context.Background()

func InitDB() *gorm.DB {
	dsn := "Lycopene:MiMaJiuShi123321!@tcp(127.0.0.1:3306)/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	}
	db.AutoMigrate(&ProblemBasic{}, &UserBasic{}, &CategoryBasic{}, &SubmitBasic{}, &ProblemCategory{}, &TestCase{})
	return db
}
func InitRedisDB() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}
