package model

import (
	"star-server/utils/errmsg"
)

type Authentication struct {
	Model
	Openid string `gorm:"type:char(128);not null" json:"openid"`
	Token  string `gorm:"type:text;not null" json:"token"`
	Uid    uint   `gorm:"type:int" json:"uid"`
}

func CreateTokens(data *Authentication) int {
	if result := db.Create(&data); result.RowsAffected == 0 {
		return errmsg.InsertError
	}
	return errmsg.SUCCESS
}

func UseUidGetAuth(uid uint) (Authentication, int) {
	var data Authentication
	if result := db.Limit(1).Where("uid=?", uid).Find(&data); result.RowsAffected == 0 {
		return Authentication{}, errmsg.SelectError
	}
	return data, errmsg.SUCCESS
}

func UseTokenGetAuth(token string) (Authentication, int) {
	var data Authentication
	if result := db.Limit(1).Where("token=?", token).Find(&data); result.RowsAffected == 0 {
		return Authentication{}, errmsg.SelectError
	}
	return data, errmsg.SUCCESS
}

func UseOpenidGetAuth(openid string) (Authentication, int) {
	var data Authentication
	if result := db.Limit(1).Where("openid=?", openid).Find(&data); result.RowsAffected == 0 {
		return Authentication{}, errmsg.SelectError
	}
	return data, errmsg.SUCCESS
}
