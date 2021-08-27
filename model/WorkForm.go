package model

import (
	"github.com/jinzhu/gorm"
	"star-server/utils/errmsg"
)

type WorkForm struct {
	gorm.Model
	Remark     string `gorm:"type:varchar(100)" json:"remark"`
	SectorName string `gorm:"type:varchar(50)" json:"sector_name"`
	StudentId  string `gorm:"type:char(10)" json:"student_id"`
}

func CreateForm(data *WorkForm) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func UpdateForm(id int, data *WorkForm) int {
	//var form WorkForm
	//var maps = make(map[string]interface{})
	//maps["remark"] = data.Remark
	data.ID = uint(id)
	err = db.Model(&data).Updates(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
