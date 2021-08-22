package model

import "github.com/jinzhu/gorm"

type Paper struct {
	gorm.Model
	Title         string `gorm:"type:varchar(255)" json:"title"`
	ReleaseTime   string `gorm:"type:varchar(50)" json:"release_time"`
	ReleaseSector string `gorm:"type:varchar(50)" json:"release_sector"`
	Content       string `gorm:"type:text" json:"content"`
	ReleaseName   string `gorm:"type:varchar(40)" json:"release_name"`
	Phone         string `gorm:"type:varchar(20)" json:"phone"`
	Attachment    string `gorm:"type:text" json:"attachment"`
}
