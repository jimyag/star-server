package model

import (
	"github.com/jinzhu/gorm"
	"star-server/utils/errmsg"
)

type WorkForm struct {
	gorm.Model
	Remark     string `gorm:"type:varchar(255)"json:"remark"`
	SectorName string `gorm:"type:varchar(50)"json:"sector_name"`
	StudentId  string `gorm:"type:varchar(15)"json:"student_id"`
}

func CreateForm(data *WorkForm) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func UpdateForm(id int, data *WorkForm) int {
	var form WorkForm
	var maps = make(map[string]interface{})
	maps["remark"] = data.Remark
	err = db.Model(&form).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
