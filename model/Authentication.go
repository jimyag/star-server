package model

import (
	"github.com/jinzhu/gorm"
	"star-server/utils/errmsg"
)

type Authentication struct {
	gorm.Model
	Openid string `gorm:"type:char(128);not null" json:"openid"`
	Token  string `gorm:"type:text;not null" json:"token"`
	Uid    uint   `gorm:"type:int" json:"uid"`
}

func CreateTokens(data *Authentication) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS
}

func UseOpenidGetUid(openid string) Authentication {
	var data Authentication
	db.Where("openid=?", openid).Find(&data)
	return data
}
