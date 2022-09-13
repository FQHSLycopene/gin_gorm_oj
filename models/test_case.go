package models

import (
	"gin_gorm_oj/utils"
	"gorm.io/gorm"
	"time"
)

type TestCase struct {
	ID              uint `gorm:"primaryKey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Identity        string         `gorm:"index;not null;column:identity;type:varchar(36)" json:"identity"` //
	ProblemIdentity string         `gorm:"column:problem_identity;type:varchar(36)" json:"problem_identity"`
	Input           string         `gorm:"column:input;type:text" json:"input"`
	Output          string         `gorm:"column:output;type:text" json:"output"`
}

func (table *TestCase) TableName() string {
	return "test_case"
}

func AddTestCase(problemIdentity, inputStr, outputStr string) (interface{}, error) {
	data := TestCase{}
	data.Identity = utils.GetUUID()
	data.ProblemIdentity = problemIdentity
	data.Input = inputStr
	data.Output = outputStr
	err := DB.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
