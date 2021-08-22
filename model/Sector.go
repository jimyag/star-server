package model

import "github.com/jinzhu/gorm"

type Sector struct {
	gorm.Model
	SectorName        string `gorm:"type:varchar(50)"json:"sector_name"`
	SectorInformation string `gorm:"type:text"json:"sector_information"`
	Supervisor        string `gorm:"type:varchar(40)"json:"supervisor"`
	Address           string `gorm:"type:varchar(255)"json:"address"`
	Phone             string `gorm:"type:varchar(20)"json:"phone"`
}
