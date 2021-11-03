package model

import (
	"star-server/utils/errmsg"
)

type StuSector struct {
	Model
	Uid        int    `gorm:"type:int;" json:"uid"`
	StudentId  string `gorm:"type:char(10)" json:"student_id"`
	SectorName string `gorm:"type:varchar(50)" json:"sector_name"`
}

func CreateStuSect(stuSector *StuSector) (code int) {
	if result := db.Create(&stuSector); result.RowsAffected == 0 {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func FindStuSectorUseSidSeName(data *StuSector) (StuSector, int) {
	var stuS StuSector
	if result := db.Limit(1).Where("student_id=? and sector_name=?", data.StudentId, data.SectorName).Find(&stuS); result.RowsAffected == 0 {
		return StuSector{}, errmsg.ERROR
	}
	return stuS, errmsg.SUCCESS
}

func FindStuSectorUseUid(data StuSector) (StuSector, int) {
	var stus StuSector
	if result := db.Limit(1).Where("uid=?", data.Uid).Find(&stus); result.RowsAffected == 0 {
		return stus, errmsg.ERROR
	}
	return stus, errmsg.SUCCESS
}
