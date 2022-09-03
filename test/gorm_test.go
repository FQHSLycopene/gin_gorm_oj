package test

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

type User struct {
	ID          uint `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Identity    string         `gorm:"type:varchar(36)"`
	Name        string         `gorm:"type:varchar(255)"`
	CreditCards []CreditCard   `gorm:"foreignKey:UserIdentity;references:Identity"`
}

type CreditCard struct {
	ID           uint `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Identity     string         `gorm:"type:varchar(36)"`
	Name         string         `gorm:"type:varchar(255)"`
	UserIdentity string         `gorm:"type:varchar(36)"`
}

func TestGormTest(t *testing.T) {
	dsn := "Lycopene:MiMaJiuShi123321!@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&User{}, &CreditCard{})
}
