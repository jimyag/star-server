package model

import (
	"github.com/jinzhu/gorm"
	"star-server/utils/errmsg"
)

type StuSector struct {
	gorm.Model
	Uid        int    `gorm:"type:int;" json:"uid"`
	StudentId  string `gorm:"type:char(10)" json:"student_id"`
	SectorName string `gorm:"type:varchar(50)" json:"sector_name"`
}

func CreateStuSect(stuSector *StuSector) (code int) {
	err := db.Create(&stuSector).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func FindSector(data *StuSector) (StuSector, int) {
	var stuS StuSector
	err := db.Model(&data).Find(&stuS)
	if err != nil {
		return StuSector{}, errmsg.ERROR
	}
	return stuS, errmsg.SUCCESS
}
