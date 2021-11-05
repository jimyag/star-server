package model

import (
	"star-server/utils/errmsg"
)

type Record struct {
	Model
	BeginTime  int    `gorm:"type:int" json:"begin_time"`
	EndTime    int    `gorm:"type:int" json:"end_time"`
	Remark     string `gorm:"type:varchar(100)" json:"remark"`
	SectorName string `gorm:"type:varchar(50)" json:"sector_name"`
	StudentId  string `gorm:"type:char(10)" json:"student_id"`
}

func CreateForm(data *Record) int {
	if result := db.Create(&data); result.RowsAffected == 0 {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func UpdateForm(data *Record) int {
	if result := db.Model(&data).Updates(data); result.RowsAffected == 0 {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetForms(studentId string) ([]Record, int) {
	var forms []Record
	if result := db.Where("student_id=?", studentId).Find(&forms); result.RowsAffected == 0 {
		return nil, errmsg.ERROR
	}
	return forms, errmsg.SUCCESS
}

func GetFormUseId(id uint) (Record, int) {
	var w Record
	if result := db.Limit(1).Where("id=?", id).Find(&w); result.RowsAffected == 0 {
		return w, errmsg.ERROR
	}
	return w, errmsg.SUCCESS
}
