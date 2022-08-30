package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36)" json:"identity"` //用户的唯一标识
	Name     string `gorm:"column:name;type:varchar(100)" json:"name"`
	Password string `gorm:"column:password;type:varchar(32)" json:"password"`
	Phone    string `gorm:"column:phone;type:varchar(20)" json:"phone"`
	Mail     string `gorm:"column:mail;type:varchar(100)" json:"mail"`
}

func (tabel *UserBasic) TableName() string {
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
