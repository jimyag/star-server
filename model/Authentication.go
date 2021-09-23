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
		return errmsg.InsertError
	}
	return errmsg.SUCCESS
}

func UseUidGetAuth(uid uint) (Authentication, int) {
	var data Authentication
	err := db.Where("uid=?", uid).Find(&data).Error
	if err != nil {
		return Authentication{}, errmsg.SelectError
	}
	return data, errmsg.SUCCESS
}

func UseTokenGetAuth(token string) (Authentication, int) {
	var data Authentication
	err := db.Where("token=?", token).Find(&data).Error
	if err != nil {
		return Authentication{}, errmsg.SelectError
	}
	return data, errmsg.SUCCESS
}

func UseOpenidGetAuth(openid string) (Authentication, int) {
	var data Authentication
	err := db.Where("openid=?", openid).Find(&data).Error
	if err != nil {
		return Authentication{}, errmsg.SelectError
	}
	return data, errmsg.SUCCESS
}
