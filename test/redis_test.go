package test

import (
	"context"
	"fmt"
	"gin_gorm_oj/models"
	"testing"
	"time"
)

var ctx = context.Background()

func TestRedisSet(t *testing.T) {
	err := models.RDB.Set(ctx, "name", "lycopene", time.Second*10).Err()
	models.RDB.Set(ctx, "wjw2002227@126.com", "123456", 0).Err()
	if err != nil {
		panic(err)
	}
}

func TestRedisGet(t *testing.T) {
	val, err := models.RDB.Get(ctx, "wjw2002227@126.com").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}
