package model

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	StudentName string `gorm:"type:varchar(40)" json:"student_name"`
	StudentId   string `gorm:"type:varchar(15)" json:"student_id"`
	Major       string `gorm:"type:varchar(20)" json:"major"`
	Phone       string `gorm:"type:varchar(20)" json:"phone"`
}
