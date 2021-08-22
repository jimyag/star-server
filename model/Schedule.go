package model

import "github.com/jinzhu/gorm"

type Schedule struct {
	gorm.Model
	SectorName  string `gorm:"type:varchar(50)" json:"sector_name"`
	StudentId   string `gorm:"type:varchar(15)" json:"student_id"`
	DayOfWeek   int    `gorm:"type:int" json:"day_of_week"`
	CourseIndex int    `gorm:"type:int" json:"course_index"`
	Address     string `gorm:"type:" json:"address"`
}
