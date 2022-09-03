package models

import (
	"errors"
	"gin_gorm_oj/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36)" json:"identity"` //用户的唯一标识
	Name     string `gorm:"column:name;type:varchar(100)" json:"name"`
	Password string `gorm:"column:password;type:varchar(32)" json:"password"`
	Phone    string `gorm:"column:phone;type:varchar(20)" json:"phone"`
	Mail     string `gorm:"column:mail;type:varchar(100)" json:"mail"`
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUser(id int) (interface{}, error) {
	data := UserBasic{}
	err := DB.Omit("password").First(&data, id).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Login(username, password string) (interface{}, error) {
	data := UserBasic{}
	err := DB.Where("name = ? AND password = ?", username, password).First(&data).Error
	if err != nil {
		return nil, err
	}
	token, err := utils.GenerateToken(data.Identity, data.Name)
	if err != nil {
		return nil, err
	}
	return gin.H{
		"token": token,
	}, nil
}

func Register(name, password, phone, mail string) (interface{}, error) {
	if name == "" || password == "" || phone == "" || mail == "" {
		return nil, errors.New("参数不正确")
	}
	data := UserBasic{
		Identity: utils.GetUUID(),
		Name:     name,
		Password: utils.GetMd5(password),
		Mail:     mail,
		Phone:    phone,
	}
	var total int64 = 0
	DB.Where("name = ?", name).Count(&total)
	if total != 0 {
		return nil, errors.New("用户名已注册")
	}
	err := DB.Create(&data).Error
	if err != nil {
		return nil, err
	}
	token, err := utils.GenerateToken(data.Identity, data.Name)
	if err != nil {
		return nil, err
	}
	return gin.H{
		"token": token,
	}, nil
}

func GetCode(email string) (string, error) {
	code := utils.GetCode()
	err := RDB.Set(CTX, email, code, time.Minute*5).Err()
	if err != nil {
		return "", err
	}
	return code, nil
}
func EmailIsExist(email string) (bool, error) {
	var total int64 = 0
	err := DB.Where("mail = ?", email).Count(&total).First(&UserBasic{}).Error()
	if err != nil {
		return false, err
	}
	if total == 1 {
		return true, nil
	}
	return false, nil
}
