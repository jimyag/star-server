package model

import (
	"star-server/utils/errmsg"
)

type Schedule struct {
	Model
	SectorName  string `gorm:"type:varchar(50)" json:"sector_name"`
	StudentId   string `gorm:"type:char(10)" json:"student_id"`
	StudentName string `gorm:"type:varchar(40);" json:"student_name"`
	DayOfWeek   int    `gorm:"type:int" json:"day_of_week"`
	CourseIndex int    `gorm:"type:int" json:"course_index"`
	Address     string `gorm:"type:varchar(100);" json:"address"`
}

func GetSectorSchedule(sectorName string) ([]Schedule, int) {
	_, code := UseNameGetSector(sectorName)
	if code == errmsg.ERROR {
		return nil, errmsg.SectorNotExist
	}
	var schedules []Schedule
	if result := db.Where("sector_name=?", sectorName).Find(&schedules); result.RowsAffected == 0 {
		return nil, errmsg.ERROR
	}
	return schedules, errmsg.SUCCESS
}

func CreateSchedule(data *Schedule) int {
	if result := db.Create(&data); result.RowsAffected == 0 {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func ScheduleRecordEquals(record Schedule) bool {
	var s Schedule
	if result := db.Where("sector_name=? and student_id=?and day_of_week=? and course_index=? and address= ?", record.SectorName, record.StudentId, record.DayOfWeek, record.CourseIndex, record.Address).Find(&s); result.RowsAffected == 0 {
		return true
	}
	return false
}
