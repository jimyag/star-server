package model

import (
	"gorm.io/gorm"
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

func UpdateStudent(student Student) (Student, int) {
	var stu Student
	err := db.Model(&student).Updates(&stu).Error
	if err != nil {
		return stu, errmsg.ERROR
	}
	return stu, errmsg.SUCCESS
}

func MatchStuNameAndId(stuName string, stuId string) bool {
	var stu Student
	_ = db.Where("student_id=?", stuId).Find(&stu).Error
	if stu.StudentName == stuName {
		return true
	}
	return false

}
