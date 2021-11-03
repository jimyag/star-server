package model

import (
	"star-server/utils/errmsg"
)

type Student struct {
	Model
	StudentName string `gorm:"type:varchar(40)" json:"student_name"`
	StudentId   string `gorm:"type:char(10)" json:"student_id"`
	Major       string `gorm:"type:varchar(50)" json:"major"`
	Phone       string `gorm:"type:char(11)" json:"phone"`
}

func CreateStudent(data *Student) int {
	if result := db.Create(&data); result.RowsAffected == 0 {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetStudent(studentId string) (Student, int) {
	var student Student
	if result := db.Limit(1).Where("student_id=?", studentId).Find(&student); result.RowsAffected == 0 {
		return Student{}, errmsg.ERROR
	}
	return student, errmsg.SUCCESS
}

func UpdateStudent(student Student) (Student, int) {
	var stu Student
	if result := db.Model(&student).Updates(&stu); result.RowsAffected == 0 {
		return Student{}, errmsg.ERROR
	}
	return stu, errmsg.SUCCESS
}

func MatchStuNameAndId(stuName string, stuId string) bool {
	var stu Student
	_ = db.Limit(1).Where("student_id=?", stuId).Find(&stu)
	if stu.StudentName == stuName {
		return true
	}
	return false

}
