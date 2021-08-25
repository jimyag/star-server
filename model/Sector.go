package model

import (
	"github.com/jinzhu/gorm"
	"star-server/utils/errmsg"
)

type Sector struct {
	gorm.Model
	SectorName        string `gorm:"type:varchar(50)" json:"sector_name"`
	SectorInformation string `gorm:"type:text" json:"sector_information"`
	Supervisor        string `gorm:"type:varchar(40)" json:"supervisor"`
	Address           string `gorm:"type:varchar(255)" json:"address"`
	Phone             string `gorm:"type:varchar(20)" json:"phone"`
}

func CreateSector(data *Sector) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetSector(pageSize int, pageIndex int) ([]Sector, int) {
	var sectors []Sector
	err := db.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Find(&sectors).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return sectors, errmsg.SUCCESS
}

func UseNameGetSector(sectorName string) (Sector, int) {
	var sector Sector
	err := db.Where("sector_name=?", sectorName).Find(&sector).Error
	if err != nil {
		return sector, errmsg.ERROR
	}
	return sector, errmsg.SUCCESS

}
