package model

import (
	"star-server/utils/errmsg"
)

type SectorKey struct {
	Model
	SectorName string `gorm:"type:varchar(50)" json:"sector_name"`
	Key        string `gorm:"type:text" json:"key"`
}

func CreateSectKey(key *SectorKey) int {
	if result := db.Create(&key); result.RowsAffected == 0 {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func FindSectorKey(key SectorKey) (SectorKey, int) {
	var k SectorKey
	if result := db.Limit(1).Where("sector_name = ?", key.SectorName).Find(&k); result.RowsAffected == 0 {
		return SectorKey{}, errmsg.ERROR
	}
	return k, errmsg.SUCCESS
}

func UseSectorNameFindSectorKey(sectorName string) (key string, code int) {
	var k SectorKey
	if result := db.Limit(1).Where("sector_name=?", sectorName).Find(&k); result.RowsAffected == 0 {
		return "", errmsg.ERROR
	}
	return k.Key, errmsg.SUCCESS
}
