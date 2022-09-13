package models

import (
	"errors"
	"gin_gorm_oj/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type UserBasic struct {
	ID               uint `gorm:"primaryKey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	Identity         string         `gorm:"index;not null;column:identity;type:varchar(36)" json:"identity"` //用户的唯一标识
	Name             string         `gorm:"column:name;type:varchar(100)" json:"name"`
	Password         string         `gorm:"column:password;type:varchar(32)" json:"password"`
	Phone            string         `gorm:"column:phone;type:varchar(20)" json:"phone"`
	Mail             string         `gorm:"column:mail;type:varchar(100)" json:"mail"`
	FinishProblemNum int64          `gorm:"column:finish_problem_num;type:int(11)" json:"finish_problem_num"`
	SubmitNum        int64          `gorm:"column:submit_num;type:int(11)" json:"submit_num"`
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserRank(pageStr, sizeStr string) (interface{}, error) {
	data := make([]*UserBasic, 0)
	var total int64
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return nil, errors.New("page 不是数字")
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return nil, errors.New("size 不是数字")
	}
	err = DB.Model(data).Select("identity", "name", "finish_problem_num", "submit_num").Offset((page - 1) * size).Limit(size).Count(&total).Order("finish_problem_num DESC,submit_num ASC").Find(&data).Error
	if err != nil {
		return nil, err
	}
	return gin.H{
		"total": total,
		"list":  data,
	}, nil
}

func GetUser(identity string) (interface{}, error) {
	data := UserBasic{}
	err := DB.Where("identity", identity).Omit("password").First(&data).Error
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
		Identity:         utils.GetUUID(),
		Name:             name,
		Password:         utils.GetMd5(password),
		Mail:             mail,
		Phone:            phone,
		FinishProblemNum: 0,
		SubmitNum:        0,
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
	//user := UserBasic{}
	err := DB.Select("mail").Where("mail = ?", email).Find(&UserBasic{}).Count(&total).Error
	if err != nil {
		return false, err
	}
	if total == 1 {
		return true, nil
	}
	return false, nil
}
