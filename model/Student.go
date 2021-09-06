package model

import (
	"github.com/jinzhu/gorm"
	"star-server/utils/errmsg"
)

type Student struct {
	gorm.Model
	StudentName string `gorm:"type:varchar(40)" json:"student_name"`
	StudentId   string `gorm:"type:char(10)" json:"student_id"`
	Major       string `gorm:"type:varchar(50)" json:"major"`
	Phone       string `gorm:"type:char(11)" json:"phone"`
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

func MatchStuNameAndId(stu_name string, stu_id string) bool {
	var stu Student
	_ = db.Where("student_id=?", stu_id).Find(&stu).Error
	if stu.StudentName == stu_name {
		return true
	}
	return false

}
