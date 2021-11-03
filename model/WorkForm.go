package model

import (
	"star-server/utils/errmsg"
)

type WorkForm struct {
	Model
	BeginTime  int    `gorm:"type:int" json:"begin_time"`
	EndTime    int    `gorm:"type:int" json:"end_time"`
	Remark     string `gorm:"type:varchar(100)" json:"remark"`
	SectorName string `gorm:"type:varchar(50)" json:"sector_name"`
	StudentId  string `gorm:"type:char(10)" json:"student_id"`
}

func CreateForm(data *WorkForm) int {
	if result := db.Create(&data); result.RowsAffected == 0 {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func UpdateForm(id int, data *WorkForm) int {
	var form WorkForm
	var maps = make(map[string]interface{})
	maps["remark"] = data.Remark
	maps["end_time"] = data.EndTime
	if result := db.Model(&form).Where("id=?", id).Updates(maps); result.RowsAffected == 0 {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetForm(studentId string) ([]WorkForm, int) {
	var forms []WorkForm
	if result := db.Where("student_id=?", studentId).Find(&forms); result.RowsAffected == 0 {
		return nil, errmsg.ERROR
	}
	return forms, errmsg.SUCCESS
}
