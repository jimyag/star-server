package model

import (
	"github.com/jinzhu/gorm"
	"star-server/utils/errmsg"
)

type Student struct {
	gorm.Model
	StudentName string `gorm:"type:varchar(40)" json:"student_name"`
	StudentId   string `gorm:"type:varchar(15)" json:"student_id"`
	Major       string `gorm:"type:varchar(20)" json:"major"`
	Phone       string `gorm:"type:varchar(20)" json:"phone"`
}

func CreateStudent(data *Student) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetStudent(studentId string) (Student, int) {
	var student Student
	err := db.Where("student_id=?", studentId).Find(&student).Error
	if err != nil {
		return student, errmsg.ERROR
	}
	return student, errmsg.SUCCESS
}
