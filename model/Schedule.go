package model

import (
	"github.com/jinzhu/gorm"
	"star-server/utils/errmsg"
)

type Schedule struct {
	gorm.Model
	SectorName  string `gorm:"type:varchar(50)" json:"sector_name"`
	StudentId   string `gorm:"type:varchar(15)" json:"student_id"`
	DayOfWeek   int    `gorm:"type:int" json:"day_of_week"`
	CourseIndex int    `gorm:"type:int" json:"course_index"`
	Address     string `gorm:"type:" json:"address"`
}

func GetSectorSchedule(sectorName string) ([]Schedule, int) {
	_, code := UseNameGetSector(sectorName)
	if code == errmsg.ERROR {
		return nil, errmsg.SectorNotExist
	}
	var schedules []Schedule
	err := db.Where("sector_name=?", sectorName).Find(&schedules).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return schedules, errmsg.SUCCESS
}

func CreateSchedule(data *Schedule) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
