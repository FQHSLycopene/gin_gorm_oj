package test

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

//
//type User struct {
//	ID          uint `gorm:"primaryKey" json:"id"`
//	CreatedAt   time.Time
//	UpdatedAt   time.Time
//	DeletedAt   gorm.DeletedAt `gorm:"index;"`
//	Identity    string         `gorm:"index;type:varchar(36)"`
//	Name        string         `gorm:"type:varchar(255)"`
//	CreditCards []*CreditCard  `gorm:"foreignKey:UserIdentity;references:Identity"`
//}

//type CreditCard struct {
//	ID           uint `gorm:"primaryKey" json:"id"`
//	CreatedAt    time.Time
//	UpdatedAt    time.Time
//	DeletedAt    gorm.DeletedAt `gorm:"index"`
//	Identity     string         `gorm:"index;type:varchar(36)"`
//	Name         string         `gorm:"type:varchar(255)"`
//	UserIdentity string         `gorm:"type:varchar(36)"`
//}

type User struct {
	gorm.Model
	Identity  string     `gorm:"index;type:varchar(36)"`
	Languages []Language `gorm:"many2many:user_languages;foreignKey:Identity;joinForeignKey:UserIdentity;References:Identity;joinReferences:LanguageIdentity"`
}

type Language struct {
	gorm.Model
	Identity string `gorm:"index;type:varchar(36)"`
	Name     string
}

func TestGormTest(t *testing.T) {
	dsn := "Lycopene:MiMaJiuShi123321!@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//db.AutoMigrate(&ProblemBasic{}, &CategoryBasic{}, &TestCase{})
	db.AutoMigrate(&User{}, &Language{})
	data := User{
		Identity: "User_1",
		Languages: []Language{
			{Identity: "lang_1", Name: "英语"},
			{Identity: "lang_2", Name: "数学"},
		},
	}

	//var creditCards []*CreditCard
	//for i := 0; i < 3; i++ {
	//	creditCards = append(creditCards, &CreditCard{
	//		Identity: utils.GetUUID(),
	//		Name:     strconv.Itoa(i),
	//	})
	//}
	//
	//data := User{
	//	Identity:    utils.GetUUID(),
	//	Name:        "wjw",
	//	CreditCards: creditCards,
	//}
	db.Create(&data)
}

//type ProblemBasic struct {
//	ID         uint `gorm:"primaryKey"`
//	CreatedAt  time.Time
//	UpdatedAt  time.Time
//	DeletedAt  gorm.DeletedAt   `gorm:"index"`
//	Identity   string           `gorm:"index;not null;column:identity;type:varchar(36)" json:"identity"` //问题唯一表示
//	Title      string           `gorm:"column:title;type:varchar(255)" json:"title"`                     //文章标题
//	Content    string           `gorm:"column:content;type:text" json:"content"`                         //
//	PassNum    int              `gorm:"column:pass_num;type:int" json:"pass_num"`
//	TotalNum   int              `gorm:"column:total_num;type:int" json:"total_num"`
//	MaxRuntime int              `gorm:"column:max_runtime;type:int(11)" json:"max_runtime"`                                                                                                          //最大运行时长
//	MaxMem     int              `gorm:"column:max_mem;type:int(11)" json:"max_mem"`                                                                                                                  //最大运行内存
//	Categories []*CategoryBasic `gorm:"many2many:problem_category;foreignKey:Identity;joinForeignKey:ProblemIdentity;References:Identity;joinReferences:CategoryIdentity" json:"problem_categories"` //关联问题分类表
//	TestCases  []*TestCase      `gorm:"foreignKey:ProblemIdentity;references:Identity" json:"test_cases"`
//}
//
//func (table *ProblemBasic) TableName() string {
//	return "problem_basic"
//}
//
//type UserBasic struct {
//	ID               uint `gorm:"primaryKey"`
//	CreatedAt        time.Time
//	UpdatedAt        time.Time
//	DeletedAt        gorm.DeletedAt `gorm:"index"`
//	Identity         string         `gorm:"index;not null;column:identity;type:varchar(36)" json:"identity"` //用户的唯一标识
//	Name             string         `gorm:"column:name;type:varchar(100)" json:"name"`
//	Password         string         `gorm:"column:password;type:varchar(32)" json:"password"`
//	Phone            string         `gorm:"column:phone;type:varchar(20)" json:"phone"`
//	Mail             string         `gorm:"column:mail;type:varchar(100)" json:"mail"`
//	FinishProblemNum int64          `gorm:"column:finish_problem_num;type:int(11)" json:"finish_problem_num"`
//	SubmitNum        int64          `gorm:"column:submit_num;type:int(11)" json:"submit_num"`
//}
//
//func (table *UserBasic) TableName() string {
//	return "user_basic"
//}
//
//type CategoryBasic struct {
//	ID        uint `gorm:"primaryKey"`
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt gorm.DeletedAt `gorm:"index"`
//	Identity  string         `gorm:"index;not null;column:identity;type:varchar(36)" json:"identity"` //分类的唯一表示
//	Name      string         `gorm:"column:name;type:varchar(100)" json:"name"`
//	Pid       string         `gorm:"column:pid;type:int(11)" json:"pid"`
//}
//
//func (table *CategoryBasic) TableName() string {
//	return "category_basic"
//}
//
//type SubmitBasic struct {
//	ID              uint `gorm:"primaryKey"`
//	CreatedAt       time.Time
//	UpdatedAt       time.Time
//	DeletedAt       gorm.DeletedAt `gorm:"index"`
//	Identity        string         `gorm:"index;not null;column:identity;type:varchar(36)" json:"identity"` //提交的唯一表示
//	ProblemIdentity string         `gorm:"column:problem_identity;type:varchar(36)" json:"problem_identity"`
//	ProblemBasic    *ProblemBasic  `gorm:"foreignKey:ProblemIdentity;references:Identity"`
//	UserIdentity    string         `gorm:"column:user_identity;type:varchar(36)" json:"user_identity"`
//	UserBasic       *UserBasic     `gorm:"foreignKey:UserIdentity;references:Identity"`
//	Path            string         `gorm:"column:path;type:varchar(255)" json:"path"` //代码存放路劲
//	Status          int            `gorm:"column:status;type:tinyint" json:"status"`
//}
//
//func (table *SubmitBasic) TableName() string {
//	return "submit_basic"
//}
//
//type ProblemCategory struct {
//	ID               uint `gorm:"primaryKey"`
//	CreatedAt        time.Time
//	UpdatedAt        time.Time
//	DeletedAt        gorm.DeletedAt `gorm:"index"`
//	ProblemIdentity  string         `gorm:"column:problem_identity;varchar(36)" json:"problem_identity"`
//	CategoryIdentity string         `gorm:"column:category_identity;varchar(36)" json:"category_identity"`
//	CategoryBasic    *CategoryBasic `gorm:"foreignKey:CategoryIdentity;references:Identity"` //关联分类的基础信息表
//}
//
//func (table *ProblemCategory) TableName() string {
//	return "problem_category"
//}
//
//type TestCase struct {
//	ID              uint `gorm:"primaryKey"`
//	CreatedAt       time.Time
//	UpdatedAt       time.Time
//	DeletedAt       gorm.DeletedAt `gorm:"index"`
//	Identity        string         `gorm:"index;not null;column:identity;type:varchar(36)" json:"identity"` //
//	ProblemIdentity string         `gorm:"column:problem_identity;type:varchar(36)" json:"problem_identity"`
//	Input           string         `gorm:"column:input;type:text" json:"input"`
//	Output          string         `gorm:"column:output;type:text" json:"output"`
//}
//
//func (table *TestCase) TableName() string {
//	return "test_case"
//}
