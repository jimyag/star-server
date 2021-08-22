package model

import "github.com/jinzhu/gorm"

type WorkForm struct {
	gorm.Model
	Remark     string `gorm:"type:varchar(255)"json:"remark"`
	SectorName string `gorm:"type:varchar(50)"json:"sector_name"`
	StudentId  string `gorm:"type:varchar(15)"json:"student_id"`
}
