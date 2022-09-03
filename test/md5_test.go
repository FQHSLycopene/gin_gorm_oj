package test

import (
	"gin_gorm_oj/utils"
	"testing"
)

func TestMd5(t *testing.T) {
	md5 := utils.GetMd5("password2")
	println(md5)
}
