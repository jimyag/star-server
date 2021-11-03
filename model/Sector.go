package model

import (
	"star-server/utils/errmsg"
)

type Sector struct {
	Model
	SectorName        string `gorm:"type:varchar(50)" json:"sector_name"`
	SectorInformation string `gorm:"type:text" json:"sector_information"`
	Supervisor        string `gorm:"type:varchar(40)" json:"supervisor"`
	Address           string `gorm:"type:varchar(100)" json:"address"`
	Phone             string `gorm:"type:char(11)" json:"phone"`
}

func CreateSector(data *Sector) int {
	if result := db.Create(&data); result.RowsAffected == 0 {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetSector(pageSize int, pageIndex int) ([]Sector, int) {
	var sectors []Sector
	err := db.Find(&sectors).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return sectors, errmsg.SUCCESS
}

func UseNameGetSector(sectorName string) (Sector, int) {
	var sector Sector
	if result := db.Limit(1).Where("sector_name=?", sectorName).Find(&sector); result.RowsAffected == 0 {
		return sector, errmsg.ERROR
	}
	return sector, errmsg.SUCCESS

}
