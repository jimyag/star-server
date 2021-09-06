package model

import (
	"github.com/jinzhu/gorm"
	"star-server/utils/errmsg"
)

type SectorKey struct {
	gorm.Model
	SectorName string `gorm:"type:varchar(50)" json:"sector_name"`
	Key        string `gorm:"type:text" json:"key"`
}

func CreateSectKey(key *SectorKey) int {
	err := db.Create(&key).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func FindSectorKey(key SectorKey) (SectorKey, int) {
	var k SectorKey
	err := db.Where("sector_name = ?", key.SectorName).Find(&k).Error
	if err != nil {
		return SectorKey{}, errmsg.ERROR
	}
	return k, errmsg.SUCCESS
}
